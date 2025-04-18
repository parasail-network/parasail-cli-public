package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"math/big"
	"os"
	"parasail-cli/contract"
	"sort"
	"strconv"
	"sync"

	"io"
	"net/http"
	"time"
)

type ServiceConfig struct {
	ChainEndpoint string     `toml:"chain_endpoint"`
	AvsConfig     *AvsConfig `toml:"avs_config"`
}

type AvsConfig struct {
	DepinAvsAddr        string `toml:"depin_avs_addr"`
	LocalPort           int    `toml:"local_port"`
	VaultManagerAddress string `toml:"vault_manager_contract_address"`
}

type Service struct {
	*ServiceConfig
	consensusEngine        ConsensusEngine
	HttpServer             *gin.Engine
	serviceMonitorContract *contract.ServiceMonitorContract
	vaultManagerContract   *contract.VaultManagerContract
	checkIntervalMinutes   int64
	ethClient              *ethclient.Client
	localDB                *LocalDB
}

func parseServiceConfig(confPath string) *ServiceConfig {
	c := &ServiceConfig{
		AvsConfig: &AvsConfig{},
	}
	_, err := toml.DecodeFile(confPath, c)
	if err != nil {
		panic(err)
	}
	localPort := ""
	envs := map[string]*string{
		"CHAIN_ENDPOINT_ADDR":   &c.ChainEndpoint,
		"LOCAL_PORT":            &localPort,
		"DEPIN_AVS_ADDR":        &c.AvsConfig.DepinAvsAddr,
		"VAULT_MANAGER_ADDRESS": &c.AvsConfig.VaultManagerAddress,
	}
	for name, ptr := range envs {
		v := os.Getenv(name)
		if v != "" {
			*ptr = v
		}
	}

	if localPort != "" {
		c.AvsConfig.LocalPort, err = strconv.Atoi(localPort)
		if err != nil {
			panic(err)
		}
	}
	return c
}

func RunService(confPath string) {
	c := parseServiceConfig(confPath)
	ethCli, err := ethclient.Dial(c.ChainEndpoint)
	if err != nil {
		log.Errorf("[Service]Unable to connect to chain endpoint. error: %v", err)
		panic(err)
	}
	//sc, err := contract.NewServiceMonitorContract(common.HexToAddress(c.AvsConfig.ServiceMonitorAddress), ethCli)
	//if err != nil {
	//	return err
	//}
	vmc, err := contract.NewVaultManagerContract(common.HexToAddress(c.AvsConfig.VaultManagerAddress), ethCli)
	if err != nil {
		panic(err)
	}
	serviceMonitorAddr, err := vmc.ServiceMonitor(nil)
	if err != nil {
		log.Errorf("[Service]Unable to get ServiceMonitor's address. error: %v", err)
		panic(err)
	}
	sc, err := contract.NewServiceMonitorContract(serviceMonitorAddr, ethCli)
	if err != nil {
		panic(err)
	}

	localDB, err := NewLocalDB("./badger")
	if err != nil {
		log.Errorf("[Service]Unable to open new local db")
		panic(err)
	}

	s := &Service{
		ServiceConfig:          c,
		consensusEngine:        &EmptyEngine{},
		HttpServer:             gin.Default(),
		serviceMonitorContract: sc,
		vaultManagerContract:   vmc,
		localDB:                localDB,
		ethClient:              ethCli,
	}

	err = s.consensusEngine.Init(context.Background())
	if err != nil {
		log.Errorf("[Service]Failed to init consensu engine. error: %v", err)
		panic(err)
	}

	s.consensusEngine.RegisterValidateHandler(s.validateCheckResult)

	checkInterval, err := s.serviceMonitorContract.GetCheckInterval(nil)
	if err != nil {
		panic(err)
	}
	s.checkIntervalMinutes = checkInterval.Int64()
	// start service
	go s.startHttpServer()
	s.periodicallyCheckSLA()
}

func (s *Service) periodicallyCheckSLA() {
	var lastTimeChecked int64
	ticker := time.NewTicker(time.Second * 5)
	for range ticker.C {
		cm := currentMinute()
		if cm-lastTimeChecked < s.checkIntervalMinutes {
			continue
		}
		vaults, err := s.ObtainAllVaults()
		if err != nil {
			continue
		}
		log.Infof("[Service]Checking vaults at a new round. vaults: %v", vaults)
		for _, vaultAddr := range vaults {
			vault, err := s.ObtainVaultDeals(vaultAddr)
			if err != nil {
				continue
			}
			round := vault.calculateCurrentRound()
			// 1. report the result of last round
			go func() {
				res, _ := s.localDB.Get(vault.genDBKey(round - 1))
				if res == nil {
					return
				}
				v := &Vault{}
				err := json.Unmarshal([]byte(*res), v)
				if err != nil {
					log.Errorf("[Service]Stored data of vault is malformed. data: %v", *res)
					return
				}
				s.ReportSLAStatus(v, round-1)
			}()

			// 2. check the SLA of current round and store it into local db
			go func() {
				result, _ := s.localDB.Get(vault.genDBKey(round))
				if result != nil {
					log.Infof("[Service]Already checked the vault %s at round %v", vaultAddr, round)
					return
				}
				err := s.CheckSLAOfVault(vault)
				if err != nil {
					log.Errorf("Error occured when calling SLA check api from %v, error: %v", s.AvsConfig.DepinAvsAddr, err)
					return
				}
				data, _ := json.Marshal(vault)
				err = s.localDB.Set(vault.genDBKey(round), string(data))
				if err != nil {
					log.Errorf("[Service]Failed to put check result to local db. error: %v", err)
				}
			}()

			// 3. delete earlier key
			go s.localDB.Delete(vault.genDBKey(round - 2))
		}

		lastTimeChecked = cm
	}
}

func currentMinute() int64 {
	return time.Now().Unix() / 60
}

func (s *Service) startHttpServer() {
	err := s.HttpServer.Run(fmt.Sprintf(":%v", s.AvsConfig.LocalPort))
	if err != nil {
		panic(err)
	}
}

type Vault struct {
	Address       common.Address
	Deals         []*Deal
	CheckInterval int64
	StartAt       int64
}

type Deal struct {
	VaultAddr   common.Address
	ID          string
	StartAt     int64
	CheckResult byte
}

func (p *Vault) genDBKey(round int64) string {
	return fmt.Sprintf("RESULTS-%v-%v", p.Address.String(), round)
}

func (pd *Vault) calculateCurrentRound() int64 {
	return ((time.Now().Unix() - pd.StartAt) / 60) / pd.CheckInterval
}

func (s *Service) ObtainAllVaults() ([]common.Address, error) {
	totalCnt, err := s.vaultManagerContract.GetVaultsCount(nil)
	if err != nil {
		panic(err)
	}
	cnt := totalCnt.Int64()

	addrs := []common.Address{}
	for i := int64(0); ; i += 10 {
		upper := i + 10
		if upper > cnt-1 {
			upper = cnt - 1
		}
		vaults, _ := s.vaultManagerContract.GetVaults(nil, big.NewInt(i), big.NewInt(upper))
		addrs = append(addrs, vaults...)
		if len(vaults) < 10 {
			break
		}

	}
	return addrs, nil
}

func (s *Service) ObtainVaultDeals(vault common.Address) (*Vault, error) {
	deals := []*Deal{}
	vaultContract, err := contract.NewVaultContract(vault, s.ethClient)
	if err != nil {
		log.Errorf("[Service]Failed to initiate VaultContract. error: %v", err)
		return nil, err
	}
	dealsCnt, err := vaultContract.GetActiveDealsCount(nil)
	if err != nil {
		log.Errorf("[Service]Failed to get active deals count. error: %v", err)
		return nil, err
	}
	cnt := dealsCnt.Int64()
	for i := int64(0); ; i += 10 {
		upper := i + 10
		if upper > cnt-1 {
			upper = cnt - 1
		}
		contractDeals, err := s.vaultManagerContract.GetVaultActiveDeals(nil, vault, big.NewInt(i), big.NewInt(upper))
		if err != nil {
			log.Errorf("[Service]Failed to get active deals of vault. error: %v", err)
			return nil, err
		}
		for _, deal := range contractDeals {
			deals = append(deals, &Deal{
				VaultAddr: vault,
				ID:        deal.Id,
				StartAt:   deal.StartAt.Int64(),
			})
		}
		if len(contractDeals) < 10 {
			break
		}
	}
	return &Vault{
		Address:       vault,
		Deals:         deals,
		CheckInterval: s.checkIntervalMinutes,
	}, nil
}

type CheckResult struct {
	Status int64 `json:"status"`
}

func (s *Service) CheckSLAOfVault(vault *Vault) error {
	wg := &sync.WaitGroup{}
	errChan := make(chan error, 1)
	for i := range vault.Deals {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			result, err := s.CheckSLAOfId(vault.Deals[i].ID)
			if err != nil {
				select {
				case errChan <- err:
				default:

				}
				return
			}
			vault.Deals[i].CheckResult = byte(result)
		}(i)
	}
	wg.Wait()
	select {
	case err := <-errChan:
		return err
	default:
		return nil
	}
}

func (s *Service) CheckSLAOfId(paymentID string) (int64, error) {
	log.Infof("[Service]Calling avs's api to check SLA. paymentID: %v", paymentID)
	d, _ := json.Marshal(map[string]string{
		"payment_id": paymentID,
	})

	res, err := http.Post(s.AvsConfig.DepinAvsAddr+"/check", "application/json", bytes.NewBuffer(d))
	if err != nil {
		return 0, errors.New("Failed to send check request: " + err.Error())
	}
	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return 0, errors.New("Failed to read response: " + err.Error())
	}
	resp := &CheckResult{}
	err = json.Unmarshal(body, resp)
	if err != nil {
		return 0, errors.New("Unexpected check response format: " + string(body))
	}
	log.Infof("[Service]Check result from avs's api. paymentID: %v, result: %v", paymentID, resp)
	return resp.Status, nil
}

func (s *Service) validateCheckResult(ctx context.Context, checkResult *SLACheckResult) (bool, error) {
	log.Infof("[Service]Validating check result. vault=%v, round=%v", checkResult.VaultAddress, checkResult.CheckRound)
	for _, r := range checkResult.Results {
		log.Infof("[Service]Validating check result. id=%v, result=%v", r.ID, r.Result)
	}

	payment := &Vault{
		Address: checkResult.VaultAddress,
	}
	res, _ := s.localDB.Get(payment.genDBKey(checkResult.CheckRound))
	if res == nil {
		return false, nil
	}
	err := json.Unmarshal([]byte(*res), payment)
	if err != nil {
		return false, err
	}
	results := map[string]byte{}
	for _, r := range checkResult.Results {
		results[r.ID] = r.Result
	}

	for _, deal := range payment.Deals {
		if results[deal.ID] != deal.CheckResult {
			return false, nil
		}
	}
	if len(results) != len(payment.Deals) {
		return false, err
	}

	return true, nil
}

func (s *Service) ReportSLAStatus(vault *Vault, round int64) error {
	r := &SLACheckResult{
		VaultAddress: vault.Address,
		CheckRound:   round,
	}
	results := []*DealCheckResult{}
	for _, p := range vault.Deals {
		results = append(results, &DealCheckResult{
			ID:     p.ID,
			Result: p.CheckResult,
		})
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].ID < results[j].ID
	})
	r.Results = results

	return s.consensusEngine.Propose(context.Background(), r)
}

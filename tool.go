package main

import (
	"context"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"math/big"
	"os"
	"parasail-cli/contract"
	"strconv"
	"strings"
	"time"
)

//const CHAIN_ENDPOINT = "https://indulgent-polished-cloud.matic-amoy.quiknode.pro/442e133ab1add6c7acfadf6280f04bc8d2f0d65e"

func NewTool(confPath string, flags *pflag.FlagSet) *Tool {
	conf := parseToolConfig(confPath, flags)
	ethClient, err := ethclient.Dial(conf.ChainEndpoint)
	if err != nil {
		panic(err)
	}
	return &Tool{
		ethClient: ethClient,
		conf:      conf,
	}
}

type Tool struct {
	ethClient *ethclient.Client
	conf      *ToolConfig
}

type ToolConfig struct {
	ChainEndpoint               string `toml:"chain_endpoint"`
	VaultManagerContractAddress string `toml:"vault_manager_contract_address"`
	PrivateKey                  string `toml:"private_key"`
	ProviderSpaceID             int64  `toml:"provider_space_id"`
	ProviderSpaceName           string `toml:"provider_name"`
	ProviderSpaceURI            string `toml:"provider_space_uri"`
	TokenAddress                string `toml:"token_address"`
	VaultAddress                string `toml:"vault_address"`
	DealID                      string `toml:"deal_id"`
	DealPrice                   int64  `toml:"deal_price"`
	DealQuantity                int64  `toml:"deal_quantity"`
	DealDuration                int64  `toml:"deal_duration"`
	DealClient                  string `toml:"deal_client"`
	StakeAmount                 string `toml:"stake_amount"`
}

func parseToolConfig(confPath string, flags *pflag.FlagSet) *ToolConfig {
	c := &ToolConfig{}

	_, err := toml.DecodeFile(confPath, c)
	if err != nil {
		//panic(err)
	}
	strEnvs := map[string]*string{
		"CHAIN_ENDPOINT":                 &c.ChainEndpoint,
		"VAULT_MANAGER_CONTRACT_ADDRESS": &c.VaultManagerContractAddress,
		"PRIVATE_KEY":                    &c.PrivateKey,
		"PROVIDER_SPACE_NAME":            &c.ProviderSpaceName,
		"PROVIDER_SPACE_URI":             &c.ProviderSpaceURI,
		"TOKEN_ADDRESS":                  &c.TokenAddress,
		"VAULT_ADDRESS":                  &c.VaultAddress,
		"DEAL_ID":                        &c.DealID,
		"DEAL_CLIENT":                    &c.DealClient,
		"STAKE_AMOUNT":                   &c.StakeAmount,
	}
	for name, ptr := range strEnvs {
		v := os.Getenv(name)
		if v != "" {
			*ptr = v
		}
		vv, _ := flags.GetString(strings.ToLower(name))
		if vv != "" {
			*ptr = vv
		}
	}

	intEnvs := map[string]*int64{
		"PROVIDER_SPACE_ID": &c.ProviderSpaceID,
		"DEAL_PRICE":        &c.DealPrice,
		"DEAL_QUANTITY":     &c.DealQuantity,
		"DEAL_DURATION":     &c.DealDuration,
	}
	for name, ptr := range intEnvs {
		v := os.Getenv(name)
		if v != "" {
			vv, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				log.Errorf("env %v should be an integer", name)
				panic("invalid env")
			}
			*ptr = vv
		}
		vv, _ := flags.GetInt64(strings.ToLower(name))
		if vv != 0 {
			*ptr = vv
		}
	}

	return c
}

func (t *Tool) ListContracts() {
	log.Infof("[Tool]VaultManagerContractAddress: %v", t.conf.VaultManagerContractAddress)
	vaultManager, err := contract.NewVaultManagerContract(common.HexToAddress(t.conf.VaultManagerContractAddress), t.ethClient)
	if err != nil {
		panic(err)
	}
	vault, err := vaultManager.GetVaultImplementation(nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Vault Implementation address: %v\n", vault)

	sm, err := vaultManager.ServiceMonitor(nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ServiceMonitor contract address: %v\n", sm)

	sl, err := vaultManager.Slasher(nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Slasher contract address: %v\n", sl)

	smc, err := contract.NewServiceMonitorContract(sm, t.ethClient)
	if err != nil {
		panic(err)
	}
	ce, err := smc.ConsensusEngine(nil)
	if err != nil {
		panic(err)
	}
	fmt.Printf("ConsensusEngine address: %v\n", ce)
}

func (t *Tool) SetSlasher() {

}

func (t *Tool) SetDepinHook() {

}

func (t *Tool) ListVaults() {
	log.Infof("[Tool]VaultManagerContractAddress: %v", t.conf.VaultManagerContractAddress)
	vaultManager, err := contract.NewVaultManagerContract(common.HexToAddress(t.conf.VaultManagerContractAddress), t.ethClient)
	if err != nil {
		panic(err)
	}
	fmt.Println("All Vaults: ")
	totalCnt, err := vaultManager.GetVaultsCount(nil)
	if err != nil {
		panic(err)
	}
	cnt := totalCnt.Int64()

	for i := int64(0); ; i += 10 {
		upper := i + 10
		if upper > cnt-1 {
			upper = cnt - 1
		}
		vaults, _ := vaultManager.GetVaults(nil, big.NewInt(i), big.NewInt(upper))
		for _, v := range vaults {
			fmt.Println(v)
		}
		if len(vaults) < 10 {
			break
		}
	}
}

func (t *Tool) CreateVault() {
	vaultManager, err := contract.NewVaultManagerContract(common.HexToAddress(t.conf.VaultManagerContractAddress), t.ethClient)
	if err != nil {
		panic(err)
	}

	auth := t.getContractAuth()
	tx, err := vaultManager.CreateVault(auth, common.HexToAddress(t.conf.TokenAddress), false)
	if err != nil {
		panic(err)
	}
	fmt.Println("Vault created! Transaction: ", tx.Hash())
}

func (t *Tool) StakeToVault() {
	vaultManagerAddr := common.HexToAddress(t.conf.VaultManagerContractAddress)
	vaultManager, err := contract.NewVaultManagerContract(vaultManagerAddr, t.ethClient)
	if err != nil {
		panic(err)
	}

	auth := t.getContractAuth()
	stakeAmount, ok := (&big.Int{}).SetString(t.conf.StakeAmount, 10)
	if !ok {
		log.Errorf("stake_amount is invalid")
		return
	}
	vaultAddr := common.HexToAddress(t.conf.VaultAddress)
	vaultContract, err := contract.NewVaultContract(vaultAddr, t.ethClient)
	if err != nil {
		panic(err)
	}
	tokenAddr, err := vaultContract.PaymentToken(nil)
	if err != nil {
		log.Errorf("Failed to get payment token from Vault contract. error: %v", err)
		panic(err)
	}
	token, err := contract.NewToken(tokenAddr, t.ethClient)
	if err != nil {
		panic(err)
	}
	_, err = token.Approve(auth, vaultManagerAddr, stakeAmount)
	if err != nil {
		log.Errorf("Failed to approve for paying. error: %v", err)
		panic(err)
	}
	log.Infof("Approved for paying. Wait for 30 seconds.")
	time.Sleep(time.Second * 30)
	log.Infof("Start to stake now.")
	tx, err := vaultManager.DelegateToVault(auth, common.HexToAddress(t.conf.VaultAddress), stakeAmount)
	if err != nil {
		panic(err)
	}
	fmt.Println("Staked! Transaction: ", tx.Hash())
}

func (t *Tool) VaultInfo() {
	vault, err := contract.NewVaultContract(common.HexToAddress(t.conf.VaultAddress), t.ethClient)
	if err != nil {
		panic(err)
	}
	owner, err := vault.Owner(nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Owner: ", owner.String())

	dt, err := vault.DealTotal(nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Deal Total: ", dt.String())
	dc, err := vault.DealCount(nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Deal Count: ", dc.String())
	dealsCnt, err := vault.GetActiveDealsCount(nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Active Deal Count: ", dealsCnt.Int64())
	tc, err := vault.TotalCollateral(nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Total Collateral: ", tc.String())
	pt, err := vault.PaymentToken(nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Payment Token: ", pt.String())
}

func (t *Tool) CreateDeal() {
	vault, err := contract.NewVaultContract(common.HexToAddress(t.conf.VaultAddress), t.ethClient)
	if err != nil {
		panic(err)
	}
	auth := t.getContractAuth()
	client := common.Address{}
	if t.conf.DealClient != "" {
		client = common.HexToAddress(t.conf.DealClient)
	}
	tx, err := vault.CreateDeal(auth, t.conf.DealID, big.NewInt(t.conf.DealPrice), big.NewInt(t.conf.DealQuantity), big.NewInt(t.conf.DealDuration), client)
	if err != nil {
		panic(err)
	}
	fmt.Println("Deal created! Transaction: ", tx.Hash())
}

func (t *Tool) GetDeal() {
	vault, err := contract.NewVaultContract(common.HexToAddress(t.conf.VaultAddress), t.ethClient)
	if err != nil {
		panic(err)
	}
	deal, err := vault.Deals(nil, t.conf.DealID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Deal info: ")
	fmt.Println("Id: ", deal.Id)
	fmt.Println("Price: ", deal.Price.String())
	fmt.Println("Quantity: ", deal.Quantity.String())
	fmt.Println("Duration: ", deal.Duration.String())
	fmt.Println("Client: ", deal.Client.String())
	fmt.Println("StartAt: ", deal.StartAt.String())
	fmt.Println("Status: ", deal.Status)
}

func (t *Tool) PayDeal() {
	vault, err := contract.NewVaultContract(common.HexToAddress(t.conf.VaultAddress), t.ethClient)
	if err != nil {
		panic(err)
	}

	auth := t.getContractAuth()

	deal, err := vault.GetDeal(nil, t.conf.DealID)
	if err != nil {
		panic(err)
	}
	totalAmount := big.NewInt(1).Mul(big.NewInt(1).Mul(deal.Price, deal.Quantity), deal.Duration)
	log.Infof("Amount to pay: %v", totalAmount.String())

	tokenAddr, err := vault.PaymentToken(nil)
	if err != nil {
		panic(err)
	}
	token, err := contract.NewToken(tokenAddr, t.ethClient)
	if err != nil {
		panic(err)
	}

	balance, err := token.BalanceOf(nil, auth.From)
	if err != nil {
		panic(err)
	}
	log.Infof("User's balance: %v", balance)
	_, err = token.Approve(auth, common.HexToAddress(t.conf.VaultAddress), totalAmount)
	if err != nil {
		panic(err)
	}
	log.Infof("Approved for paying. Wait for 30 seconds.")

	time.Sleep(time.Second * 30)

	tx, err := vault.PayDeal(auth, t.conf.DealID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Payed deal! Amount: ", totalAmount.String(), ", Transaction: ", tx.Hash())
}

func (t *Tool) ClaimReward() {
	vault, err := contract.NewVaultContract(common.HexToAddress(t.conf.VaultAddress), t.ethClient)
	if err != nil {
		panic(err)
	}

	auth := t.getContractAuth()
	tx, err := vault.ClaimReward(auth, t.conf.DealID)
	if err != nil {
		panic(err)
	}
	fmt.Println("Claimed! tx: ", tx)
}

func getAddress(privateKey string) common.Address {
	pk, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		panic(err)
	}
	return crypto.PubkeyToAddress(pk.PublicKey)
}

func (t *Tool) getContractAuth() *bind.TransactOpts {
	deployer_private_key, err := crypto.HexToECDSA(t.conf.PrivateKey)
	if err != nil {
		log.Errorf("Failed to parse private key. err: %v", err)
		panic(err)
	}

	chainID, err := t.ethClient.ChainID(context.Background())
	if err != nil {
		log.Errorf("Unable to get the chainid from %v. error: %v", t.conf.ChainEndpoint, err)
		panic(err)
	}

	fromAddress := crypto.PubkeyToAddress(deployer_private_key.PublicKey)
	signer := types.LatestSignerForChainID(chainID)
	return &bind.TransactOpts{
		From: fromAddress,
		Signer: func(addr common.Address, transaction *types.Transaction) (*types.Transaction, error) {
			return types.SignTx(transaction, signer, deployer_private_key)
		},
	}
}

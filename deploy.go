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
	"parasail-cli/contract"
)

type Deployer struct {
	conf         *DeployConfig
	ethCli       *ethclient.Client
	deployerAuth *bind.TransactOpts
}

type DeployConfig struct {
	ChainEndpoint  string          `toml:"chain_endpoint"`
	ContractConfig *ContractConfig `toml:"contract"`
}

type ContractConfig struct {
	OwnerPrivateKey         string `toml:"owner_private_key"`
	BlockTime               int64  `toml:"block_time"`
	DealPaymentDeadline     int64  `toml:"deal_payment_deadline"`
	CollateralFactor        int64  `toml:"collateral_factor"`
	MinCollateral           int64  `toml:"min_collateral"`
	SLACheckIntervalMinutes int64  `toml:"sla_check_interval_minutes"`
	SlashBurnRatio          int64  `toml:"slash_burn_ratio"`
	SlashTreasuryRatio      int64  `toml:"slash_treasury_ratio"`
	BurnAddress             string `toml:"burn_address"`
	TreasuryAddress         string `toml:"treasury_address"`
	MaxFailures             int64  `toml:"max_failures"`
	TerminateFeeFactor      int64  `toml:"terminate_fee_factor"`
	WithdrawDelay           int64  `toml:"withdraw_delay"`
	RewardTokenAddress      string `toml:"reward_token_address"`

	ConsensusEngineAddress      string `toml:"consensus_engine_address"`
	VaultManagerContractAddress string `toml:"vault_manager_contract_address"`
	SlasherContractAddress      string `toml:"slasher_contract_address"`
}

func parseFlagConfs(conf *DeployConfig, flags *pflag.FlagSet) {
	vmca, _ := flags.GetString("vault_manager_contract_address")
	if vmca != "" {
		conf.ContractConfig.VaultManagerContractAddress = vmca
	}
	rta, _ := flags.GetString("reward_token_address")
	if rta != "" {
		conf.ContractConfig.RewardTokenAddress = rta
	}
	sca, _ := flags.GetString("slasher_contract_address")
	if sca != "" {
		conf.ContractConfig.SlasherContractAddress = sca
	}
	cea, _ := flags.GetString("consensus_engine_address")
	if cea != "" {
		conf.ContractConfig.ConsensusEngineAddress = cea
	}
}

func NewDeployer(confPath string, flags *pflag.FlagSet) *Deployer {
	conf := &DeployConfig{}
	toml.DecodeFile(confPath, conf)
	parseFlagConfs(conf, flags)
	conf.ChainEndpoint = "https://rpc-amoy.polygon.technology/"
	deployer := &Deployer{
		conf: conf,
	}
	var err error
	deployer.ethCli, err = ethclient.Dial(deployer.conf.ChainEndpoint)
	if err != nil {
		log.Errorf("Unable to connect to endpoint. error: %v", err)
		panic(err)
	}

	deployer.deployerAuth, err = deployer.getContractAuth(deployer.ethCli)
	if err != nil {
		panic(err)
	}
	return deployer
}

func (d *Deployer) getContractAuth(ethCli *ethclient.Client) (*bind.TransactOpts, error) {
	deployer_private_key, err := crypto.HexToECDSA(d.conf.ContractConfig.OwnerPrivateKey)
	if err != nil {
		log.Errorf("Failed to parse private key. err: %v", err)
		return nil, err
	}

	chainID, err := ethCli.ChainID(context.Background())
	if err != nil {
		log.Errorf("Unable to get the chainid from %v. error: %v", d.conf.ChainEndpoint, err)
		return nil, err
	}

	fromAddress := crypto.PubkeyToAddress(deployer_private_key.PublicKey)
	signer := types.LatestSignerForChainID(chainID)
	return &bind.TransactOpts{
		From: fromAddress,
		Signer: func(addr common.Address, transaction *types.Transaction) (*types.Transaction, error) {
			return types.SignTx(transaction, signer, deployer_private_key)
		},
	}, nil
}

func (d *Deployer) deploySlasherContracts() error {
	vaultManager, err := contract.NewVaultManagerContract(common.HexToAddress(d.conf.ContractConfig.VaultManagerContractAddress), d.ethCli)
	if err != nil {
		log.Errorf("[Deploy]Failed to initiate VaultManager contract. error: %v", err)
		return err
	}
	serviceMonitorAddr, err := vaultManager.ServiceMonitor(nil)
	if err != nil {
		log.Errorf("[Deploy]Failed to get ServiceMonitor contract address. error: %v", err)
		return err
	}

	var slasherAddr = common.Address{}
	if d.conf.ContractConfig.SlasherContractAddress == "" {
		// deploy the default SlasherContract
		slasherAddr, _, _, err = contract.DeploySlasher(
			d.deployerAuth,
			d.ethCli,
			common.HexToAddress(d.conf.ContractConfig.VaultManagerContractAddress),
			serviceMonitorAddr,
		)
		if err != nil {
			log.Errorf("[Deploy]Failed to deploy Slasher contract. error: %v", err)
			return err
		}
		fmt.Println("Deployed Slasher contract. Address:", slasherAddr)
	} else {
		slasherAddr = common.HexToAddress(d.conf.ContractConfig.SlasherContractAddress)
	}
	_, err = vaultManager.SetSlasher(d.deployerAuth, slasherAddr)
	if err != nil {
		log.Errorf("[Deploy]Failed to set slasher. vaultManager: %v, slasher: %v, error: %v",
			d.conf.ContractConfig.VaultManagerContractAddress, slasherAddr, err)
		return err
	}
	fmt.Println("Succeeded in setting Slasher in VaultManager contract.")

	return nil
}

func (d *Deployer) deployRewardsContracts() error {
	rewardsAddr, _, _, err := contract.DeployRewards(
		d.deployerAuth,
		d.ethCli,
		common.HexToAddress(d.conf.ContractConfig.RewardTokenAddress),
		common.HexToAddress(d.conf.ContractConfig.VaultManagerContractAddress),
	)
	if err != nil {
		log.Errorf("[Deploy]Failed to deploy Rewards contract. error: %v", err)
		return err
	}
	fmt.Println("Deployed Rewards contract. Address:", rewardsAddr)
	return nil
}

func (d *Deployer) SetConsensusEngineAddress() error {
	vaultManager, err := contract.NewVaultManagerContract(common.HexToAddress(d.conf.ContractConfig.VaultManagerContractAddress), d.ethCli)
	if err != nil {
		log.Errorf("[Deploy]Failed to initiate VaultManager. error: %v", err)
		return err
	}
	serviceMonitorAddr, err := vaultManager.ServiceMonitor(nil)
	if err != nil {
		log.Errorf("[Deploy]Failed to get ServiceMonitor address from VaultManager. error: %v", err)
		return err
	}

	serviceMonitorContract, err := contract.NewServiceMonitorContract(serviceMonitorAddr, d.ethCli)
	if err != nil {
		log.Errorf("[Deploy]Failed to inittiate ServiceMonitor contract. error: %v", err)
		return err
	}

	consensusEngineAddress := common.HexToAddress(d.conf.ContractConfig.ConsensusEngineAddress)
	_, err = serviceMonitorContract.SetConsensusEngine(d.deployerAuth, consensusEngineAddress)
	if err != nil {
		log.Errorf("[Deploy]Failed to register ConsensusEngine contract in ServiceMonitor contract. error: %v", err)
		return err
	}
	fmt.Println("Succeeded in registering ConsensusEngine in ServiceMonitor contract.")

	return nil
}

// deploy Vault, VaultManager, ServiceMonitor contracts
func (d *Deployer) DeployParasailBaseContracts() error {
	vaultAddr, _, _, err := contract.DeployVaultContract(d.deployerAuth, d.ethCli)
	if err != nil {
		log.Errorf("[Deploy]Failed to deploy vault contract. error: %v", err)
		return err
	}
	fmt.Println("Deployed Vault contract. Address:", vaultAddr)

	vaultManagerConfig := contract.VaultManagerConfig{
		MinCollateral:      big.NewInt(d.conf.ContractConfig.MinCollateral),
		CollateralFactor:   big.NewInt(d.conf.ContractConfig.CollateralFactor),
		SlashBurnRatio:     big.NewInt(d.conf.ContractConfig.SlashBurnRatio),
		SlashTreasuryRatio: big.NewInt(d.conf.ContractConfig.SlashTreasuryRatio),
		MaxFailures:        big.NewInt(d.conf.ContractConfig.MaxFailures),
		TerminateFeeFactor: big.NewInt(d.conf.ContractConfig.TerminateFeeFactor),
		BurnAddress:        common.HexToAddress(d.conf.ContractConfig.BurnAddress),
		Treasury:           common.HexToAddress(d.conf.ContractConfig.TreasuryAddress),
		WithdrawDelay:      big.NewInt(d.conf.ContractConfig.WithdrawDelay),
	}
	vaultManagerAddr, _, vaultManager, err := contract.DeployVaultManagerContract(
		d.deployerAuth,
		d.ethCli,
		d.deployerAuth.From,
		vaultAddr,
		vaultManagerConfig,
	)
	if err != nil {
		log.Errorf("[Deploy]Failed to deploy vault manager contract. error: %v", err)
		return err
	}
	fmt.Println("Deployed VaultManager contract. Address:", vaultManagerAddr)

	serviceMonitorAddr, _, _, err := contract.DeployServiceMonitorContract(d.deployerAuth, d.ethCli, vaultManagerAddr, common.Address{}, big.NewInt(d.conf.ContractConfig.SLACheckIntervalMinutes))
	if err != nil {
		log.Errorf("[Deploy]Failed to deploy ServiceMonitor contract. error: %v", err)
		return err
	}
	fmt.Println("Deployed ServiceMonitor contract. Address:", serviceMonitorAddr)

	_, err = vaultManager.SetServiceMonitor(d.deployerAuth, serviceMonitorAddr)
	if err != nil {
		log.Errorf("[Deploy]Failed to register ServiceMonitor contract in VaultManager contract. error: %v", err)
		return err
	}
	fmt.Println("Succeeded in registering ServiceMonitor contract in VaultManager.")
	return nil
}

func (d *Deployer) DeployDemoToken(name, symbol string, supply int) error {
	tokenAddress, tx, _, err := contract.DeployToken(d.deployerAuth, d.ethCli, name, symbol, 18, big.NewInt(int64(supply)))
	if err != nil {
		panic(err)
	}
	log.Infof("Submitted transaction: %f, tokenAddress:", tx.Hash().String(), tokenAddress.String())
	return err
}

package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const VERSION = "v0.1.0"

var (
	ConfPath string
)

func main() {
	log.SetReportCaller(true)
	log.SetFormatter(&log.TextFormatter{
		ForceColors:      true,
		DisableTimestamp: false,
		FullTimestamp:    true,
		TimestampFormat:  time.RFC3339,
		CallerPrettyfier: func(frame *runtime.Frame) (function string, file string) {
			ss := strings.Split(frame.File, "/")
			return "", fmt.Sprintf("%s:%d", ss[len(ss)-1], frame.Line)
		},
	})

	cmd := &cobra.Command{
		Use:   "parasail-cli",
		Short: "parasail-cli is a tool to generator code for staking tech stack",
	}

	version := &cobra.Command{
		Use:   "version",
		Short: "show current version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(VERSION)
		},
	}

	generator := &cobra.Command{
		Use:   "new",
		Short: "generate a new template project",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				panic("please provide the project name")
			}
			GenerateRepo(args[0])
		},
	}
	serviceCmd := &cobra.Command{
		Use:   "service",
		Short: "run the service",
		Run: func(cmd *cobra.Command, args []string) {
			log.Infof("[main]ConfPath: %v", ConfPath)
			RunService(ConfPath)
		},
	}

	cmd.AddCommand(version, generator, serviceCmd, createDeployCmd(), createToolCmd())
	cmd.PersistentFlags().StringVar(&ConfPath, "conf", "config.toml", "the path of the config(default to config.toml)")
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("Failed to execute command. error: %v", err)
	}
}

func createDeployCmd() *cobra.Command {
	deployCmd := &cobra.Command{
		Use:   "deploy",
		Short: "deploy and manage the contracts",
	}

	deployBasicCmd := &cobra.Command{
		Use:   "base",
		Short: "deploy the basic three contracts: Vault, VaultManager and ServiceMonitor",
		Run: func(cmd *cobra.Command, args []string) {
			err := NewDeployer(ConfPath, deployCmd.Flags()).DeployParasailBaseContracts()
			if err != nil {
				panic(err)
			}
		},
	}
	deploySlasherCmd := &cobra.Command{
		Use:   "slasher",
		Short: "deploy the Slasher contract",
		Run: func(cmd *cobra.Command, args []string) {
			err := NewDeployer(ConfPath, deployCmd.Flags()).deploySlasherContracts()
			if err != nil {
				panic(err)
			}
		},
	}
	deployRewardsCmd := &cobra.Command{
		Use:   "rewards",
		Short: "deploy the Rewards contract",
		Run: func(cmd *cobra.Command, args []string) {
			err := NewDeployer(ConfPath, deployCmd.Flags()).deployRewardsContracts()
			if err != nil {
				panic(err)
			}
		},
	}
	setConsensusEngineCmd := &cobra.Command{
		Use:   "set-consensus-engine",
		Short: "set the address of consensus engine",
		Run: func(cmd *cobra.Command, args []string) {
			err := NewDeployer(ConfPath, deployCmd.Flags()).SetConsensusEngineAddress()
			if err != nil {
				panic(err)
			}
		},
	}
	deployDemoToken := &cobra.Command{
		Use:   "demo-token",
		Short: "Deploy a demo token for testing",
		Run: func(cmd *cobra.Command, args []string) {
			name, symbol, supply := "DemoToken", "DT", 10000
			if len(args) >= 1 {
				name = args[0]
			}
			if len(args) >= 2 {
				symbol = args[1]
			}
			if len(args) >= 3 {
				sp, err := strconv.Atoi(args[2])
				if err != nil {
					log.Errorf("supply should be a valid number")
					panic(err)
				}
				supply = sp
			}
			err := NewDeployer(ConfPath, deployCmd.Flags()).DeployDemoToken(name, symbol, supply)
			if err != nil {
				panic(err)
			}
		},
	}

	deployCmd.AddCommand(deployBasicCmd, deploySlasherCmd, deployRewardsCmd, setConsensusEngineCmd, deployDemoToken)
	deployCmd.PersistentFlags().String("vault_manager_contract_address", "", "address of VaultManager contract")
	deployCmd.PersistentFlags().String("reward_token_address", "", "address of reward_token when deploying Rewards contract")

	return deployCmd
}

func createToolCmd() *cobra.Command {
	toolCmd := &cobra.Command{
		Use:   "tool",
		Short: "tools",
	}

	listContractsCmd := &cobra.Command{
		Use:   "list-contracts",
		Short: "show all the contracts",
		Run: func(cmd *cobra.Command, args []string) {
			NewTool(ConfPath, cmd.Flags()).ListContracts()
		},
	}

	listVaultsCmd := &cobra.Command{
		Use:   "list-vaults",
		Short: "list all the vaults",
		Run: func(cmd *cobra.Command, args []string) {
			NewTool(ConfPath, cmd.Flags()).ListVaults()
		},
	}

	createVaultCmd := &cobra.Command{
		Use:   "create-vault",
		Short: "create a new vault",
		Run: func(cmd *cobra.Command, args []string) {
			NewTool(ConfPath, cmd.Flags()).CreateVault()
		},
	}
	delegateCmd := &cobra.Command{
		Use:   "stake",
		Short: "stake to a vault",
		Run: func(cmd *cobra.Command, args []string) {
			NewTool(ConfPath, cmd.Flags()).StakeToVault()
		},
	}
	getVaultCmd := &cobra.Command{
		Use:   "get-vault",
		Short: "get a vault's detailed information",
		Run: func(cmd *cobra.Command, args []string) {
			NewTool(ConfPath, cmd.Flags()).VaultInfo()
		},
	}
	createDealCmd := &cobra.Command{
		Use:   "create-deal",
		Short: "create a new deal",
		Run: func(cmd *cobra.Command, args []string) {
			NewTool(ConfPath, cmd.Flags()).CreateDeal()
		},
	}
	payDealCmd := &cobra.Command{
		Use:   "pay-deal",
		Short: "pay a deal",
		Run: func(cmd *cobra.Command, args []string) {
			NewTool(ConfPath, cmd.Flags()).PayDeal()
		},
	}
	getDealCmd := &cobra.Command{
		Use:   "get-deal",
		Short: "get a deal's detailed information",
		Run: func(cmd *cobra.Command, args []string) {
			NewTool(ConfPath, cmd.Flags()).GetDeal()
		},
	}

	claimReward := &cobra.Command{
		Use:   "claim-reward",
		Short: "claim reward of a deal",
		Run: func(cmd *cobra.Command, args []string) {
			NewTool(ConfPath, cmd.Flags()).ClaimReward()
		},
	}

	toolCmd.AddCommand(claimReward, listContractsCmd, getDealCmd, payDealCmd, createDealCmd, getVaultCmd, delegateCmd, createVaultCmd, listVaultsCmd)
	toolCmd.PersistentFlags().String("vault_manager_contract_address", "", "address of VaultManagerContract")
	toolCmd.PersistentFlags().String("chain_endpoint", "", "chain endpoint address")
	toolCmd.PersistentFlags().String("private_key", "", "private key of the user that will commit transactions")
	toolCmd.PersistentFlags().String("provider_space_name", "", "name of the new provider space")
	toolCmd.PersistentFlags().String("provider_space_uri", "", "uri of the new provider space")
	toolCmd.PersistentFlags().String("token_address", "", "address of the ERC20 token")
	toolCmd.PersistentFlags().String("vault_address", "", "address of the vault")
	toolCmd.PersistentFlags().String("deal_id", "", "deal's id")
	toolCmd.PersistentFlags().String("deal_client", "", "client's address of the deal")
	toolCmd.PersistentFlags().Int64("provider_space_id", 0, "id of provider space")
	toolCmd.PersistentFlags().Int64("deal_price", 0, "price for the deal")
	toolCmd.PersistentFlags().Int64("deal_quantity", 0, "quantity of the deal")
	toolCmd.PersistentFlags().Int64("deal_duration", 0, "duration of the deal")
	toolCmd.PersistentFlags().String("stake_amount", "0", "amount to stake")
	return toolCmd
}

package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"parasail-cli/generate"
)

var repo = map[string]interface{}{
	"avs":                generate.AvsFiles,
	"contract":           generate.ContractFiles,
	"deploy.toml":        parasailDeployConfig,
	"service.toml":       parasailServiceConfig,
	"docker-compose.yml": dockerComposeFile,
	"README.md":          "",
	"LICENSE":            "",
}

func GenerateRepo(projectName string) {
	currentDir, err := os.Getwd()
	if err != nil {
		log.Error("Can't get current directory. error: ", err)
		return
	}
	projectDir := currentDir + "/" + projectName
	for n, vals := range repo {
		if !processRepoItem(projectDir, n, vals) {
			panic("Failed to create the repo")
		}
	}
}

func processRepoItem(dirName, fileName string, v interface{}) bool {
	_, err := os.Stat(dirName)
	if os.IsNotExist(err) {
		err = os.Mkdir(dirName, 0755)
		if err != nil {
			log.Error("Can't create avs directory. error: ", err)
			return false
		}
	}

	if c, ok := v.(string); ok {
		return writeFile(dirName, fileName, c)
	}
	if c, ok := v.(map[string]string); ok {
		for fn, content := range c {
			if !writeFile(dirName+"/"+fileName, fn, content) {
				return false
			}
		}
	}
	if c, ok := v.(map[string]interface{}); ok {
		for fn, content := range c {
			if !processRepoItem(dirName+"/"+fileName, fn, content) {
				return false
			}
		}
	}
	return true
}

func writeFile(baseDir string, fileName string, content string) bool {
	_, err := os.Stat(baseDir)
	if os.IsNotExist(err) {
		err = os.Mkdir(baseDir, 0755)
		if err != nil {
			log.Error("Can't create avs directory. error: ", err)
			return false
		}
	}
	path := baseDir + "/" + fileName
	err = os.WriteFile(path, []byte(content), 0644)
	if err != nil {
		log.Error("Can't write ts scripts to avs file. error: ", err)
		return false
	}
	return true
}

const parasailDeployConfig = `chain_endpoint = "https://rpc-amoy.polygon.technology/"

[contract]
owner_private_key = ""
block_time = 30
deal_payment_deadline = 10
collateral_factor = 120
min_collateral = 0
sla_check_interval_minutes = 1440
slash_burn_ratio = 40
burn_address = ""
slash_treasury_ratio = 60
treasury_address = ""
max_failures = 14
terminate_fee_factor = 70
withdraw_delay = 100

`

const parasailServiceConfig = `chain_endpoint = "https://rpc-amoy.polygon.technology/"

[avs_config]
depin_avs_addr = ""
local_port = 4002
vault_manager_contract_address = ""
`

const dockerComposeFile = `version: '3.7'
services:
  avs:
    build:
      context: ./avs
      dockerfile: ./Dockerfile
    environment:
      - PORT=4003
    networks:
      operator_network:
        ipv4_address: 10.8.0.4

  parasail-cli:
    image: parasail/parasail-cli:latest
    volumes:
      - ./service.toml:/app/service.toml
    command: [
      "parasail-cli",
      "service",
      "run",
      "--conf",
      "/app/service.toml"
    ]
    environment:
      - DEPIN_AVS_ADDR=http://10.8.0.4:4003
    networks:
      operator_network:
        ipv4_address: 10.8.0.3


networks:
  operator_network:
    driver: bridge
    ipam:
      config:
        - subnet: 10.8.0.0/16
          gateway: 10.8.0.1`

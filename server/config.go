package main

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type cfg struct {
	SecretKey 						string `envconfig:"AWS_SECRET_KEY"`
	AccessKey 						string `envconfig:"AWS_ACCESS_KEY"`
	Bucket    						string `envconfig:"AWS_BUCKET"`
	Region    						string `envconfig:"AWS_REGION"`
	EthDeployWalletPk 				string `envconfig:"ETH_DEPLOY_WALLET_PK"`
	EthDeployWalletAddress 			string `envconfig:"ETH_DEPLOY_WALLET_ADDRESS"`
	EthEndpointUrl 					string `envconfig:"ETH_ENDPOINT_URL"`
	EthNftContractAddress 			string `envconfig:"ETH_NFT_CONTRACT_ADDRESS"`
	EthNftMarketContractAddress 	string `envconfig:"ETH_NFT_MARKET_CONTRACT_ADDRESS"`
	HecoDeployWalletPk 				string `envconfig:"HECO_DEPLOY_WALLET_PK"`
	HecoEndpointUrl 				string `envconfig:"HECO_ENDPOINT_URL"`
	HecoNftContractAddress 			string `envconfig:"HECO_NFT_CONTRACT_ADDRESS"`
	HecoNftMarketContractAddress 	string `envconfig:"HECO_NFT_MARKET_CONTRACT_ADDRESS"`
	AuroraDeployWalletPk 			string `envconfig:"AURORA_DEPLOY_WALLET_PK"`
	AuroraEndpointUrl 				string `envconfig:"AURORA_ENDPOINT_URL"`
	AuroraNftContractAddress 		string `envconfig:"AURORA_NFT_CONTRACT_ADDRESS"`
	AuroraNftMarketContractAddress 	string `envconfig:"AURORA_NFT_MARKET_CONTRACT_ADDRESS"`
	PolygonDeployWalletPk 			string `envconfig:"POLYGON_DEPLOY_WALLET_PK"`
	PolygonEndpointUrl 				string `envconfig:"POLYGON_ENDPOINT_URL"`
	PolygonNftContractAddress 		string `envconfig:"POLYGON_NFT_CONTRACT_ADDRESS"`
	PolygonNftMarketContractAddress string `envconfig:"POLYGON_NFT_MARKET_CONTRACT_ADDRESS"`
	NftStorageKey		 			string `envconfig:"NFT_STORAGE_KEY"`

}

var c = new(cfg)

func init() {
	_ = godotenv.Overload(".env", ".env.local")
	_ = envconfig.Process("", c)
}

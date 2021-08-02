package main

import (
	"github.com/rs/zerolog/log"
	"math/big"
)

const nearTestnetExplorer = "https://explorer.testnet.near.org/accounts/aurora/"
const nearBlockchainName = "near"

func mintNearNft(name string, description string, imageCid string) (txHash string, err error) {
	// Make json and upload
	nftJson := NftJson{Name: name, Description: description, Image: imageCid}
	metadataURI, err := uploadJsonToIpfs(nftJson)
	if err != nil {
		log.Error().Err(err).Msg("Error uploading json to Ipfs")
		return
	}

	// Create tx
	instance, txOptions, err := createInstance(c.AuroraEndpointUrl, c.AuroraDeployWalletPk, c.AuroraContractAddress, big.NewInt(-1))
	if err != nil {
		log.Error().Err(err).Msg("Failed to create contract instance")
		return
	}

	txHash, err = mintNft(instance, txOptions, metadataURI)
	if err != nil {
		log.Error().Err(err)
		return
	}

	txHash = nearTestnetExplorer + txHash
	return
}

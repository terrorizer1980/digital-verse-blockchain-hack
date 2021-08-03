package main

import (
	"github.com/rs/zerolog/log"
	"math/big"
)

const hecoTestnetExplorer = "https://testnet.hecoinfo.com/tx/"
const hecoBlockchainName = "heco"


func mintHecoNft(name string, description string, imageCid string) (txHash string, err error) {
	// Make json and upload
	nftJson := NftJson{Name: name, Description: description, Image: imageCid}
	metadataURI, err := uploadJsonToIpfs(nftJson)
	if err != nil {
		log.Error().Err(err).Msg("Error uploading json to Ipfs")
		return
	}

	// Create tx
	instance, txOptions, err := createInstance(c.HecoEndpointUrl, c.HecoDeployWalletPk, c.HecoNftContractAddress, big.NewInt(-1))
	if err != nil {
		log.Error().Err(err).Msg("Failed to create contract instance")
		return
	}

	txHash, err = mintNft(instance, txOptions, metadataURI)
	if err != nil {
		log.Error().Err(err)
		return
	}

	txHash = hecoTestnetExplorer + txHash
	return
}



package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// https://api-dev.rarible.com/protocol/v0.1/ethereum/nft/collections/0xB0EA149212Eb707a1E5FC1D2d3fD318a8d94cf05/generate_token_id?minter=0x2fB3Fdf2C5ffB1cbe8D51761FD4C4398c99ac88f
// https://api-staging.rarible.com/protocol/v0.1/ethereum/nft/items/{itemId}/meta
// https://api-dev.rarible.com/protocol/v0.1/ethereum/nft/items/21576722194993390593864209416790113291260823118183058861532018344159959056398/meta

func mintRaribleNft(name string, description string, imageCid string) (txHash string, err error) {
	// Make json and upload
	// nftJson := NftJson{Name: name, Description: description, Image: imageCid}
	// metadataURI, err := uploadJsonToIpfs(nftJson)
	// if err != nil {
	//	log.Error().Err(err).Msg("Error uploading json to Ipfs")
	//	return
	//}


	return
}

type RaribleTokenIdResponse struct {
	TokenId string
}

func getRaribleNftTokenId() (tokenId string, err error) {
	resp, err := http.Get("https://api-dev.rarible.com/protocol/v0.1/ethereum/nft/collections/" + c.RaribleRopstenNftContractAddress + "/generate_token_id?minter=" + c.RopstenDeployWalletAddress)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		var response RaribleTokenIdResponse
		json.Unmarshal(bodyBytes, &response)
		return response.TokenId, nil
	}
	return
}

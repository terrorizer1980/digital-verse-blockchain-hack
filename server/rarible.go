package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"math/big"
	"net/http"
)

const ropstenTestnetExplorer = "https://ropsten.etherscan.io/tx/"

// https://api-dev.rarible.com/protocol/v0.1/ethereum/nft/collections/0xB0EA149212Eb707a1E5FC1D2d3fD318a8d94cf05/generate_token_id?minter=0x2fB3Fdf2C5ffB1cbe8D51761FD4C4398c99ac88f
// https://api-staging.rarible.com/protocol/v0.1/ethereum/nft/items/{itemId}/meta
// https://api-dev.rarible.com/protocol/v0.1/ethereum/nft/items/21576722194993390593864209416790113291260823118183058861532018344159959056398/meta
// https://app.rarible.com/{tokenId}
// https://app-dev.rarible.com/{tokenId}

func createRaribleNftInstance(endpointUrl string, walletPk string, smartContractAddress string, chainId *big.Int) (instance *RaribleNft, txOptions *bind.TransactOpts, err error) {

	client, err := ethclient.Dial(endpointUrl)
	if err != nil {
		log.Error().Err(err)
		return
	}

	privateKey, err := crypto.HexToECDSA(walletPk)
	if err != nil {
		log.Error().Err(err)
		return
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Error().Err(err).Msg("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Error().Err(err)
		return
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Error().Err(err)
		return
	}

	// NewKeyedTransactorWithChainID
	if chainId.Int64() != -1 {
		txOptions, err = bind.NewKeyedTransactorWithChainID(privateKey, chainId)
		if err != nil {
			log.Error().Err(err)
			return
		}
	} else {
		txOptions = bind.NewKeyedTransactor(privateKey)
	}

	txOptions.Nonce = big.NewInt(int64(nonce))
	txOptions.Value = big.NewInt(0)     // in wei
	txOptions.GasLimit = uint64(300000) // in units
	txOptions.GasPrice = gasPrice

	address := common.HexToAddress(smartContractAddress)
	instance, err = NewRaribleNft(address, client)
	if err != nil {
		log.Error().Err(err)
		return
	}
	return instance, txOptions, nil
}


func mintRaribleNft(name string, description string, imageCid string, minter string) (txHash string, tokenId *big.Int, metadataURI string, err error) {
	tokenId, err = getRaribleNftTokenId()
	if err != nil {
		log.Error().Err(err)
		return
	}
	// Make json and upload
	nftJson := NftJson{Name: name, Description: description, Image: imageCid }
	metadataURI, err = uploadJsonToIpfs(nftJson)
	if err != nil {
		log.Error().Err(err).Msg("Error uploading json to Ipfs")
		return
	}
	// Create tx
	instance, txOptions, err := createRaribleNftInstance(c.RopstenEndpointUrl, c.RopstenDeployWalletPk, c.RaribleRopstenNftContractAddress, big.NewInt(-1))
	if err != nil {
		log.Error().Err(err)
		return
	}

	data := LibERC721LazyMintMint721Data{
		Uri: metadataURI,
		TokenId: tokenId,
		Creators: []LibPartPart{{Account: common.HexToAddress(minter), Value: new(big.Int).SetInt64(10000)}},
		Royalties: []LibPartPart{{Account: common.HexToAddress(minter), Value: new(big.Int).SetInt64(1000)}},
		Signatures: [][]byte{[]byte("0x")},
	}
	fmt.Print(data)
	tx, err := instance.MintAndTransfer(txOptions, data, common.HexToAddress(minter))
	if err != nil {
		log.Error().Err(err)
		return
	}
	txHash = tx.Hash().Hex()
	fmt.Printf("tx sent: %s", txHash)

	txHash = ropstenTestnetExplorer + txHash
	return
}

type RaribleTokenIdResponse struct {
	TokenId string
}

// https://api-dev.rarible.com/protocol/v0.1/ethereum/nft/collections/0xB0EA149212Eb707a1E5FC1D2d3fD318a8d94cf05/generate_token_id?minter=0x2fB3Fdf2C5ffB1cbe8D51761FD4C4398c99ac88f
func getRaribleNftTokenId() (tokenId *big.Int, err error) {
	resp, err := http.Get("https://api-dev.rarible.com/protocol/v0.1/ethereum/nft/collections/" + c.RaribleRopstenNftContractAddress + "/generate_token_id?minter=" + c.RopstenDeployWalletAddress)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return tokenId, err
		}
		var response RaribleTokenIdResponse
		json.Unmarshal(bodyBytes, &response)
		tokenId := new(big.Int)
		tokenId, ok := tokenId.SetString(response.TokenId, 10)
		if !ok {
			return tokenId, errors.New("Error converting rarible tokenId")
		}
		return tokenId, nil
	}
	return
}

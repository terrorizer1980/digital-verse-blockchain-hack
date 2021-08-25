package main

// https://github.com/nftstorage/go-client/blob/main/docs/NFTStorageAPI.md#store

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
)

type NeoStorageResponse struct {
	Ok bool
	Value ValueField
}

type ValueField struct {
	Cid string
	Created string
}

type NftJson struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

func uploadJsonToIpfs(nftJson NftJson) (cid string, err error) {
	b, err := json.Marshal(nftJson)
	if err != nil {
		log.Error().Err(err).Msg("Error converting to json")
		return
	}
	cid, err = postToNftStorage(bytes.NewReader(b))
	if err != nil {
		log.Error().Err(err).Msg("Error uploading json to Ipfs")
		return
	}
	return
}

func uploadFileFromUrlToIpfs(fileUrl string) (cid string, err error) {
	localFilePath := fmt.Sprintf("./%s", path.Base(fileUrl))

	if _, err := os.Stat(localFilePath); os.IsNotExist(err) {
		file, err := os.Create(path.Base(localFilePath))
		if err != nil {
			log.Error().Err(err).Msg("Error while creating file")
			return "", err
		}

		defer file.Close()

		if err := downloadFromS3(fileUrl, file); err != nil {
			return "", err
		}
	}

	// Upload localFilePath
	f, err := os.Open(localFilePath)
	if err != nil {
		return
	}
	defer f.Close()
	cid, err = postToNftStorage(f)
	if err != nil {
		log.Error().Err(err).Msg("Error while posting to NFT storage")
		return "", err
	}
	return
}

func postToNftStorage(body io.Reader) (cid string, err error) {
	req, err := http.NewRequest("POST", "https://api.nft.storage/upload", body)
	if err != nil {
		return
	}
	req.Header.Set("Authorization", "Bearer " + c.NftStorageKey)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", err
		}
		var response NeoStorageResponse
		json.Unmarshal(bodyBytes, &response)
		return response.Value.Cid, nil
	}
	return
}


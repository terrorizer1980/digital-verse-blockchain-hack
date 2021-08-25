package main

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type CreateNFTRequest struct {
	Name        	string `json:"name"`
	Description 	string `json:"description"`
	URL         	string `json:"url"`
	Blockchain      string `json:"blockchain"`
}

func getCreateNftRequest(ctx *gin.Context) (createNftRequest CreateNFTRequest, err error) {
	if err = ctx.BindJSON(&createNftRequest); err != nil {
		log.Error().Err(err).Msg("Failed to encode CreateNFTRequest")
		return
	}

	if len(createNftRequest.Name) == 0 {
		createNftRequest.Name = "Digital Verse"
	}
	if len(createNftRequest.Description) == 0 {
		createNftRequest.Description = "Celebrity video"
	}
	if len(createNftRequest.URL) == 0 {
		log.Error().Msg("File url not specified")
		return createNftRequest, errors.New("file url not specified")
	}
	return createNftRequest, nil
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/create_heco_nft", func(ctx *gin.Context) {
		createReq, err := getCreateNftRequest(ctx)
		if err != nil {
			log.Error().Err(err).Msg("Failed to get params from request")
			return
		}

		ipfsCid, err := uploadFileFromUrlToIpfs(createReq.URL)
		if err != nil {
			log.Error().Err(err).Msg("Failed to upload file to IPFS")
			return
		}

		txHash, err := mintHecoNft(createReq.Name, createReq.Description, ipfsCid)
		if err != nil {
			log.Error().Err(err).Msg("Failed to mint NFT")
			return
		}

		ctx.JSON(200, gin.H{
			"tx_hash":      txHash,
			"url":          txHash,
			"fileUrl": 		"https://ipfs.io/ipfs/" + ipfsCid,
			"error":        err,
		})
	})

	r.POST("/create_eth_nft", func(ctx *gin.Context) {
		createReq, err := getCreateNftRequest(ctx)
		if err != nil {
			log.Error().Err(err).Msg("Failed to get params from request")
			return
		}

		ipfsCid, err := uploadFileFromUrlToIpfs(createReq.URL)
		if err != nil {
			log.Error().Err(err).Msg("Failed to upload file to IPFS")
			return
		}

		txHash, err := mintRinkebyNft(createReq.Name, createReq.Description, ipfsCid)
		if err != nil {
			log.Error().Err(err).Msg("Failed to mint NFT")
			return
		}

		ctx.JSON(200, gin.H{
			"tx_hash":      txHash,
			"url":          txHash,
			"fileUrl": 		"https://ipfs.io/ipfs/" + ipfsCid,
			"error":        err,
			"unmarshal":    "https://stg-api.unmarshal.io/v1/eth/address/" + c.RinkebyDeployWalletAddress + "/nft-assets?auth_key=VGVtcEtleQ==",
		})
	})

	r.POST("/create_near_nft", func(ctx *gin.Context) {
		createReq, err := getCreateNftRequest(ctx)
		if err != nil {
			log.Error().Err(err).Msg("Failed to get params from request")
			return
		}

		ipfsCid, err := uploadFileFromUrlToIpfs(createReq.URL)
		if err != nil {
			log.Error().Err(err).Msg("Failed to upload file to IPFS")
			return
		}

		txHash, err := mintNearNft(createReq.Name, createReq.Description, ipfsCid)
		if err != nil {
			log.Error().Err(err).Msg("Failed to mint NFT")
			return
		}

		ctx.JSON(200, gin.H{
			"tx_hash":      txHash,
			"url":          nearTestnetExplorer,
			"fileUrl": 		"https://ipfs.io/ipfs/" + ipfsCid,
			"error":        err,
		})
	})

	r.POST("/create_polygon_nft", func(ctx *gin.Context) {
		createReq, err := getCreateNftRequest(ctx)
		if err != nil {
			log.Error().Err(err).Msg("Failed to get params from request")
			return
		}

		ipfsCid, err := uploadFileFromUrlToIpfs(createReq.URL)
		if err != nil {
			log.Error().Err(err).Msg("Failed to upload file to IPFS")
			return
		}

		txHash, err := mintPolygonNft(createReq.Name, createReq.Description, ipfsCid)

		if err != nil {
			log.Error().Err(err).Msg("Failed to mint NFT")
			return
		}

		ctx.JSON(200, gin.H{
			"tx_hash":      txHash,
			"url":          txHash,
			"fileUrl": 		"https://ipfs.io/ipfs/" + ipfsCid,
			"error":        err,
		})
	})


	r.POST("/create_rarible_nft", func(ctx *gin.Context) {
		createReq, err := getCreateNftRequest(ctx)
		if err != nil {
			log.Error().Err(err).Msg("Failed to get params from request")
			return
		}

		ipfsCid, err := uploadFileFromUrlToIpfs(createReq.URL)
		if err != nil {
			log.Error().Err(err).Msg("Failed to upload file to IPFS")
			return
		}

		txHash, tokenId, metadataURI, err := mintRaribleNft(createReq.Name, createReq.Description, ipfsCid, c.RopstenDeployWalletPk)
		if err != nil {
			log.Error().Err(err).Msg("Failed to mint NFT")
			return
		}

		ctx.JSON(200, gin.H{
			"tx_hash":      txHash,
			"url":          "https://app.rarible.com/" + tokenId.String(),
			"fileUrl": 		"https://ipfs.io/ipfs/" + metadataURI,
			"error":        err,
		})
	})

	r.POST("/upload_file_to_ipfs", func(c *gin.Context) {
		fileUrl := c.PostForm("fileUrl")
		cid, err := uploadFileFromUrlToIpfs(fileUrl)
		c.JSON(200, gin.H{
			"cid":  cid,
			"fileUrl":  "https://ipfs.io/ipfs/" + cid,
			"error": 	err,
		})
	})

	r.Run()
}

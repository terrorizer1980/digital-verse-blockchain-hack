## Setup:
docker-compose build

docker-compose up

server runs on port -> localhost:49160

## Demo access
Url: https://app.digitalverse.ai/

Email: demo@digitalverse.com

Password: qwerty

Support: @panichm (telegram account)

Recommends for good results: upload small videos (5-20 seconds long), keep your face away from the camera for better quality, use videos with standard lighting.

## Deployed smart contracts

- ./contracts/solidity/ - path

- https://mumbai.polygonscan.com/token/0x6Aabd2Bce4bd40f109A5753BFc5c82bDb2Ad86Bd - NFT
https://mumbai.polygonscan.com/token/0xfa98276B5D6e0C0118f1BB0B82E9d4A656454Fe2 - NFT Market

- https://explorer.testnet.near.org/accounts/aurora
https://explorer.testnet.near.org/transactions/3KH8gCwg1CV1ZwxkTMZM85qnd2gft6NP8eJsZAkxcyq7, 0x4495cf49C0AD032AD691df1ebF850eAa3847E381 - NFT
https://explorer.testnet.near.org/transactions/AGz6CHzXf3miTdU9JhRkLMUea4zfjBvHQtpXowaQUPGR , 0x98567e77F3214BEe5F34E8b9596C472C67721336 - NFT Market

- https://rinkeby.etherscan.io/address/0x0ED06150f3Bb1E164d0065fAa4EAbC4843659Ae8 - NFT
https://rinkeby.etherscan.io/address/0x37d6d59448778AFAcc14E880D0A2dfb3A36253d9 - NFT Market

- https://testnet.hecoinfo.com/address/0x3b2e02ABB28d8a124748e78E1fDE31B4F7520a7c - NFT
https://testnet.hecoinfo.com/address/0x3b3ab07FabF9AFCb20952CE936673f1c27f9654C - NFT Market

Our contracts has all needed basic functions to work with NFT tokens:

- **MintTokenToAddress(address owner, string memory metadataURI)** 

    Allows you to mint a NFT token.
- **TransferFrom(address from, address to, uint256 tokenId)**

    Allows you to transfer any NFT token on other address.
- **Implementation of access control**

- **NFT Market**

    Allows you to list and sell created NFT tokens. Functions:

    createMarketItem(address nftContract,uint256 tokenId,uint256 price)
    
    createMarketSale(address nftContract,uint256 itemId)
    
    fetchMarketItems() public view returns (MarketItem[] memory)
    
    fetchMyNFTs() public view returns (MarketItem[] memory)
    
## API

{server} = 165.22.72.149:49160 

- **{server}/create_{eth}_nft** 

    Allows you to mint a NFT token

- **{server}/upload_file_to_ipfs**

    Uploading a media file to IPFS storage
    
## Open Issues

- Code refactor 

## Roadmap

- Integrate copyrights mechanism
- Integrate NFT sales/transfer options in UI 
    
## Help Tools

### Filecoin/IPFS

- https://nft.storage/#docs

### Heco

- https://docs.hecochain.com/#/en-us/testnet

### Ethereum

- https://wizard.openzeppelin.com
- https://docs.openzeppelin.com/contracts/3.x/api/presets
- https://forum.openzeppelin.com/t/create-an-nft-and-deploy-to-a-public-testnet-using-remix/6358
- https://goethereumbook.org/en/smart-contract-compile/
- https://www.freeformatter.com/json-formatter.html
- https://testnets-api.opensea.io/asset/0xBf1550d38FADc10afa665C0F4bDA3cA0dc77BDBb/1/validate/

### Other

- https://github.com/oligot/go-mod-upgrade
- go mod tidy
- https://mholt.github.io/curl-to-go/

## License
Apache License, Version 2.0
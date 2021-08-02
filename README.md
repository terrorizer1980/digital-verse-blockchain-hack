## Setup:
docker-compose build

docker-compose up

server runs on port -> localhost:49160

## Demo access
Url: https://dev.app.digitalverse.ai/

Email: demo@digitalverse.com

Password: qwerty

Support: @panichm (telegram account)

Recommends for good results: upload small videos (5-20 seconds long), keep your face away from the camera for better quality, use videos with standard lighting.

## Deployed smart contracts

- ./contracts/solidity/ - path
- https://testnet.hecoinfo.com/address/0x1d2a0ba44f522de56bf8fa83c312b8bdacfc20eb - info

Our contract has all needed basic functions to work with NFT tokens:

- **MintTokenToAddress(address owner, string memory metadataURI)** 

    Allows you to mint a NFT token.
- **TransferFrom(address from, address to, uint256 tokenId)**

    Allows you to transfer any NFT token on other address.
- **Implementation of access control**

## API

{server} = 165.22.72.149:49160 

- **{server}/create_nft** 

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

## License
Apache License, Version 2.0
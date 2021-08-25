// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// LibERC721LazyMintMint721Data is an auto generated low-level Go binding around an user-defined struct.
type LibERC721LazyMintMint721Data struct {
	TokenId    *big.Int
	Uri        string
	Creators   []LibPartPart
	Royalties  []LibPartPart
	Signatures [][]byte
}

// LibPartPart is an auto generated low-level Go binding around an user-defined struct.
type LibPartPart struct {
	Account common.Address
	Value   *big.Int
}

// RaribleNftABI is the input ABI used to generate the binding from.
const RaribleNftABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"uri\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibPart.Part[]\",\"name\":\"creators\",\"type\":\"tuple[]\"},{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLibPart.Part[]\",\"name\":\"royalties\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes[]\",\"name\":\"signatures\",\"type\":\"bytes[]\"}],\"internalType\":\"structLibERC721LazyMint.Mint721Data\",\"name\":\"data\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mintAndTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// RaribleNft is an auto generated Go binding around an Ethereum contract.
type RaribleNft struct {
	RaribleNftCaller     // Read-only binding to the contract
	RaribleNftTransactor // Write-only binding to the contract
	RaribleNftFilterer   // Log filterer for contract events
}

// RaribleNftCaller is an auto generated read-only Go binding around an Ethereum contract.
type RaribleNftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RaribleNftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RaribleNftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RaribleNftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RaribleNftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RaribleNftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RaribleNftSession struct {
	Contract     *RaribleNft       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RaribleNftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RaribleNftCallerSession struct {
	Contract *RaribleNftCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// RaribleNftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RaribleNftTransactorSession struct {
	Contract     *RaribleNftTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// RaribleNftRaw is an auto generated low-level Go binding around an Ethereum contract.
type RaribleNftRaw struct {
	Contract *RaribleNft // Generic contract binding to access the raw methods on
}

// RaribleNftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RaribleNftCallerRaw struct {
	Contract *RaribleNftCaller // Generic read-only contract binding to access the raw methods on
}

// RaribleNftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RaribleNftTransactorRaw struct {
	Contract *RaribleNftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRaribleNft creates a new instance of RaribleNft, bound to a specific deployed contract.
func NewRaribleNft(address common.Address, backend bind.ContractBackend) (*RaribleNft, error) {
	contract, err := bindRaribleNft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RaribleNft{RaribleNftCaller: RaribleNftCaller{contract: contract}, RaribleNftTransactor: RaribleNftTransactor{contract: contract}, RaribleNftFilterer: RaribleNftFilterer{contract: contract}}, nil
}

// NewRaribleNftCaller creates a new read-only instance of RaribleNft, bound to a specific deployed contract.
func NewRaribleNftCaller(address common.Address, caller bind.ContractCaller) (*RaribleNftCaller, error) {
	contract, err := bindRaribleNft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RaribleNftCaller{contract: contract}, nil
}

// NewRaribleNftTransactor creates a new write-only instance of RaribleNft, bound to a specific deployed contract.
func NewRaribleNftTransactor(address common.Address, transactor bind.ContractTransactor) (*RaribleNftTransactor, error) {
	contract, err := bindRaribleNft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RaribleNftTransactor{contract: contract}, nil
}

// NewRaribleNftFilterer creates a new log filterer instance of RaribleNft, bound to a specific deployed contract.
func NewRaribleNftFilterer(address common.Address, filterer bind.ContractFilterer) (*RaribleNftFilterer, error) {
	contract, err := bindRaribleNft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RaribleNftFilterer{contract: contract}, nil
}

// bindRaribleNft binds a generic wrapper to an already deployed contract.
func bindRaribleNft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RaribleNftABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RaribleNft *RaribleNftRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RaribleNft.Contract.RaribleNftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RaribleNft *RaribleNftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RaribleNft.Contract.RaribleNftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RaribleNft *RaribleNftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RaribleNft.Contract.RaribleNftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RaribleNft *RaribleNftCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RaribleNft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RaribleNft *RaribleNftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RaribleNft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RaribleNft *RaribleNftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RaribleNft.Contract.contract.Transact(opts, method, params...)
}

// MintAndTransfer is a paid mutator transaction binding the contract method 0x450346e0.
//
// Solidity: function mintAndTransfer((uint256,string,(address,uint256)[],(address,uint256)[],bytes[]) data, address to) returns()
func (_RaribleNft *RaribleNftTransactor) MintAndTransfer(opts *bind.TransactOpts, data LibERC721LazyMintMint721Data, to common.Address) (*types.Transaction, error) {
	return _RaribleNft.contract.Transact(opts, "mintAndTransfer", data, to)
}

// MintAndTransfer is a paid mutator transaction binding the contract method 0x450346e0.
//
// Solidity: function mintAndTransfer((uint256,string,(address,uint256)[],(address,uint256)[],bytes[]) data, address to) returns()
func (_RaribleNft *RaribleNftSession) MintAndTransfer(data LibERC721LazyMintMint721Data, to common.Address) (*types.Transaction, error) {
	return _RaribleNft.Contract.MintAndTransfer(&_RaribleNft.TransactOpts, data, to)
}

// MintAndTransfer is a paid mutator transaction binding the contract method 0x450346e0.
//
// Solidity: function mintAndTransfer((uint256,string,(address,uint256)[],(address,uint256)[],bytes[]) data, address to) returns()
func (_RaribleNft *RaribleNftTransactorSession) MintAndTransfer(data LibERC721LazyMintMint721Data, to common.Address) (*types.Transaction, error) {
	return _RaribleNft.Contract.MintAndTransfer(&_RaribleNft.TransactOpts, data, to)
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cyber

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

// CarABI is the input ABI used to generate the binding from.
const CarABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"capacity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"mintPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"phase\",\"outputs\":[{\"internalType\":\"int8\",\"name\":\"\",\"type\":\"int8\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserved\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"whitelistCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int8\",\"name\":\"newPhase\",\"type\":\"int8\"}],\"name\":\"setPhase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"amount\",\"type\":\"uint8\"}],\"name\":\"claim\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reserve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"amount\",\"type\":\"uint8\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\",\"payable\":true},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"amount\",\"type\":\"uint8\"}],\"name\":\"addWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"amount\",\"type\":\"uint8\"}],\"name\":\"addAirdrop\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"addrs\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"amount\",\"type\":\"uint8\"}],\"name\":\"addReserve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"mintQuota\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"minted\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"cap\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"airdropQuota\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"minted\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"cap\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true}]"

// Car is an auto generated Go binding around an Ethereum contract.
type Car struct {
	CarCaller     // Read-only binding to the contract
	CarTransactor // Write-only binding to the contract
	CarFilterer   // Log filterer for contract events
}

// CarCaller is an auto generated read-only Go binding around an Ethereum contract.
type CarCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CarTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CarTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CarFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CarFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CarSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CarSession struct {
	Contract     *Car              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CarCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CarCallerSession struct {
	Contract *CarCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CarTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CarTransactorSession struct {
	Contract     *CarTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CarRaw is an auto generated low-level Go binding around an Ethereum contract.
type CarRaw struct {
	Contract *Car // Generic contract binding to access the raw methods on
}

// CarCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CarCallerRaw struct {
	Contract *CarCaller // Generic read-only contract binding to access the raw methods on
}

// CarTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CarTransactorRaw struct {
	Contract *CarTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCar creates a new instance of Car, bound to a specific deployed contract.
func NewCar(address common.Address, backend bind.ContractBackend) (*Car, error) {
	contract, err := bindCar(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Car{CarCaller: CarCaller{contract: contract}, CarTransactor: CarTransactor{contract: contract}, CarFilterer: CarFilterer{contract: contract}}, nil
}

// NewCarCaller creates a new read-only instance of Car, bound to a specific deployed contract.
func NewCarCaller(address common.Address, caller bind.ContractCaller) (*CarCaller, error) {
	contract, err := bindCar(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CarCaller{contract: contract}, nil
}

// NewCarTransactor creates a new write-only instance of Car, bound to a specific deployed contract.
func NewCarTransactor(address common.Address, transactor bind.ContractTransactor) (*CarTransactor, error) {
	contract, err := bindCar(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CarTransactor{contract: contract}, nil
}

// NewCarFilterer creates a new log filterer instance of Car, bound to a specific deployed contract.
func NewCarFilterer(address common.Address, filterer bind.ContractFilterer) (*CarFilterer, error) {
	contract, err := bindCar(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CarFilterer{contract: contract}, nil
}

// bindCar binds a generic wrapper to an already deployed contract.
func bindCar(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CarABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Car *CarRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Car.Contract.CarCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Car *CarRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Car.Contract.CarTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Car *CarRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Car.Contract.CarTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Car *CarCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Car.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Car *CarTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Car.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Car *CarTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Car.Contract.contract.Transact(opts, method, params...)
}

// AirdropQuota is a free data retrieval call binding the contract method 0x2bce1d35.
//
// Solidity: function airdropQuota(address addr) view returns(uint8 minted, uint8 cap)
func (_Car *CarCaller) AirdropQuota(opts *bind.CallOpts, addr common.Address) (struct {
	Minted uint8
	Cap    uint8
}, error) {
	var out []interface{}
	err := _Car.contract.Call(opts, &out, "airdropQuota", addr)

	outstruct := new(struct {
		Minted uint8
		Cap    uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Minted = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Cap = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

// AirdropQuota is a free data retrieval call binding the contract method 0x2bce1d35.
//
// Solidity: function airdropQuota(address addr) view returns(uint8 minted, uint8 cap)
func (_Car *CarSession) AirdropQuota(addr common.Address) (struct {
	Minted uint8
	Cap    uint8
}, error) {
	return _Car.Contract.AirdropQuota(&_Car.CallOpts, addr)
}

// AirdropQuota is a free data retrieval call binding the contract method 0x2bce1d35.
//
// Solidity: function airdropQuota(address addr) view returns(uint8 minted, uint8 cap)
func (_Car *CarCallerSession) AirdropQuota(addr common.Address) (struct {
	Minted uint8
	Cap    uint8
}, error) {
	return _Car.Contract.AirdropQuota(&_Car.CallOpts, addr)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Car *CarCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Car.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Car *CarSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Car.Contract.BalanceOf(&_Car.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Car *CarCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Car.Contract.BalanceOf(&_Car.CallOpts, owner)
}

// Capacity is a free data retrieval call binding the contract method 0x5cfc1a51.
//
// Solidity: function capacity() view returns(uint256)
func (_Car *CarCaller) Capacity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Car.contract.Call(opts, &out, "capacity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Capacity is a free data retrieval call binding the contract method 0x5cfc1a51.
//
// Solidity: function capacity() view returns(uint256)
func (_Car *CarSession) Capacity() (*big.Int, error) {
	return _Car.Contract.Capacity(&_Car.CallOpts)
}

// Capacity is a free data retrieval call binding the contract method 0x5cfc1a51.
//
// Solidity: function capacity() view returns(uint256)
func (_Car *CarCallerSession) Capacity() (*big.Int, error) {
	return _Car.Contract.Capacity(&_Car.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Car *CarCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Car.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Car *CarSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Car.Contract.GetApproved(&_Car.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Car *CarCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Car.Contract.GetApproved(&_Car.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Car *CarCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Car.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Car *CarSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Car.Contract.IsApprovedForAll(&_Car.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Car *CarCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Car.Contract.IsApprovedForAll(&_Car.CallOpts, owner, operator)
}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_Car *CarCaller) MintPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Car.contract.Call(opts, &out, "mintPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_Car *CarSession) MintPrice() (*big.Int, error) {
	return _Car.Contract.MintPrice(&_Car.CallOpts)
}

// MintPrice is a free data retrieval call binding the contract method 0x6817c76c.
//
// Solidity: function mintPrice() view returns(uint256)
func (_Car *CarCallerSession) MintPrice() (*big.Int, error) {
	return _Car.Contract.MintPrice(&_Car.CallOpts)
}

// MintQuota is a free data retrieval call binding the contract method 0x2c915a22.
//
// Solidity: function mintQuota(address addr) view returns(uint8 minted, uint8 cap)
func (_Car *CarCaller) MintQuota(opts *bind.CallOpts, addr common.Address) (struct {
	Minted uint8
	Cap    uint8
}, error) {
	var out []interface{}
	err := _Car.contract.Call(opts, &out, "mintQuota", addr)

	outstruct := new(struct {
		Minted uint8
		Cap    uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Minted = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Cap = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

// MintQuota is a free data retrieval call binding the contract method 0x2c915a22.
//
// Solidity: function mintQuota(address addr) view returns(uint8 minted, uint8 cap)
func (_Car *CarSession) MintQuota(addr common.Address) (struct {
	Minted uint8
	Cap    uint8
}, error) {
	return _Car.Contract.MintQuota(&_Car.CallOpts, addr)
}

// MintQuota is a free data retrieval call binding the contract method 0x2c915a22.
//
// Solidity: function mintQuota(address addr) view returns(uint8 minted, uint8 cap)
func (_Car *CarCallerSession) MintQuota(addr common.Address) (struct {
	Minted uint8
	Cap    uint8
}, error) {
	return _Car.Contract.MintQuota(&_Car.CallOpts, addr)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Car *CarCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Car.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Car *CarSession) Name() (string, error) {
	return _Car.Contract.Name(&_Car.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Car *CarCallerSession) Name() (string, error) {
	return _Car.Contract.Name(&_Car.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Car *CarCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Car.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Car *CarSession) Owner() (common.Address, error) {
	return _Car.Contract.Owner(&_Car.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Car *CarCallerSession) Owner() (common.Address, error) {
	return _Car.Contract.Owner(&_Car.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Car *CarCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Car.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Car *CarSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Car.Contract.OwnerOf(&_Car.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Car *CarCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Car.Contract.OwnerOf(&_Car.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Car *CarCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Car.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Car *CarSession) Paused() (bool, error) {
	return _Car.Contract.Paused(&_Car.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Car *CarCallerSession) Paused() (bool, error) {
	return _Car.Contract.Paused(&_Car.CallOpts)
}

// Phase is a free data retrieval call binding the contract method 0xb1c9fe6e.
//
// Solidity: function phase() view returns(int8)
func (_Car *CarCaller) Phase(opts *bind.CallOpts) (int8, error) {
	var out []interface{}
	err := _Car.contract.Call(opts, &out, "phase")

	if err != nil {
		return *new(int8), err
	}

	out0 := *abi.ConvertType(out[0], new(int8)).(*int8)

	return out0, err

}

// Phase is a free data retrieval call binding the contract method 0xb1c9fe6e.
//
// Solidity: function phase() view returns(int8)
func (_Car *CarSession) Phase() (int8, error) {
	return _Car.Contract.Phase(&_Car.CallOpts)
}

// Phase is a free data retrieval call binding the contract method 0xb1c9fe6e.
//
// Solidity: function phase() view returns(int8)
func (_Car *CarCallerSession) Phase() (int8, error) {
	return _Car.Contract.Phase(&_Car.CallOpts)
}

// Reserved is a free data retrieval call binding the contract method 0xfe60d12c.
//
// Solidity: function reserved() view returns(uint256)
func (_Car *CarCaller) Reserved(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Car.contract.Call(opts, &out, "reserved")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Reserved is a free data retrieval call binding the contract method 0xfe60d12c.
//
// Solidity: function reserved() view returns(uint256)
func (_Car *CarSession) Reserved() (*big.Int, error) {
	return _Car.Contract.Reserved(&_Car.CallOpts)
}

// Reserved is a free data retrieval call binding the contract method 0xfe60d12c.
//
// Solidity: function reserved() view returns(uint256)
func (_Car *CarCallerSession) Reserved() (*big.Int, error) {
	return _Car.Contract.Reserved(&_Car.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Car *CarCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Car.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Car *CarSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Car.Contract.SupportsInterface(&_Car.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Car *CarCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Car.Contract.SupportsInterface(&_Car.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Car *CarCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Car.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Car *CarSession) Symbol() (string, error) {
	return _Car.Contract.Symbol(&_Car.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Car *CarCallerSession) Symbol() (string, error) {
	return _Car.Contract.Symbol(&_Car.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Car *CarCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Car.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Car *CarSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Car.Contract.TokenByIndex(&_Car.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Car *CarCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Car.Contract.TokenByIndex(&_Car.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Car *CarCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Car.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Car *CarSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Car.Contract.TokenOfOwnerByIndex(&_Car.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Car *CarCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Car.Contract.TokenOfOwnerByIndex(&_Car.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Car *CarCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Car.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Car *CarSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Car.Contract.TokenURI(&_Car.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Car *CarCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Car.Contract.TokenURI(&_Car.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Car *CarCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Car.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Car *CarSession) TotalSupply() (*big.Int, error) {
	return _Car.Contract.TotalSupply(&_Car.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Car *CarCallerSession) TotalSupply() (*big.Int, error) {
	return _Car.Contract.TotalSupply(&_Car.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Car *CarCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Car.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Car *CarSession) Version() (string, error) {
	return _Car.Contract.Version(&_Car.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_Car *CarCallerSession) Version() (string, error) {
	return _Car.Contract.Version(&_Car.CallOpts)
}

// WhitelistCap is a free data retrieval call binding the contract method 0xf634d519.
//
// Solidity: function whitelistCap() view returns(uint256)
func (_Car *CarCaller) WhitelistCap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Car.contract.Call(opts, &out, "whitelistCap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WhitelistCap is a free data retrieval call binding the contract method 0xf634d519.
//
// Solidity: function whitelistCap() view returns(uint256)
func (_Car *CarSession) WhitelistCap() (*big.Int, error) {
	return _Car.Contract.WhitelistCap(&_Car.CallOpts)
}

// WhitelistCap is a free data retrieval call binding the contract method 0xf634d519.
//
// Solidity: function whitelistCap() view returns(uint256)
func (_Car *CarCallerSession) WhitelistCap() (*big.Int, error) {
	return _Car.Contract.WhitelistCap(&_Car.CallOpts)
}

// AddAirdrop is a paid mutator transaction binding the contract method 0x7509447d.
//
// Solidity: function addAirdrop(address[] addrs, uint8 amount) returns()
func (_Car *CarTransactor) AddAirdrop(opts *bind.TransactOpts, addrs []common.Address, amount uint8) (*types.Transaction, error) {
	return _Car.contract.Transact(opts, "addAirdrop", addrs, amount)
}

// AddAirdrop is a paid mutator transaction binding the contract method 0x7509447d.
//
// Solidity: function addAirdrop(address[] addrs, uint8 amount) returns()
func (_Car *CarSession) AddAirdrop(addrs []common.Address, amount uint8) (*types.Transaction, error) {
	return _Car.Contract.AddAirdrop(&_Car.TransactOpts, addrs, amount)
}

// AddAirdrop is a paid mutator transaction binding the contract method 0x7509447d.
//
// Solidity: function addAirdrop(address[] addrs, uint8 amount) returns()
func (_Car *CarTransactorSession) AddAirdrop(addrs []common.Address, amount uint8) (*types.Transaction, error) {
	return _Car.Contract.AddAirdrop(&_Car.TransactOpts, addrs, amount)
}

// AddReserve is a paid mutator transaction binding the contract method 0xb2fa739a.
//
// Solidity: function addReserve(address[] addrs, uint8 amount) returns()
func (_Car *CarTransactor) AddReserve(opts *bind.TransactOpts, addrs []common.Address, amount uint8) (*types.Transaction, error) {
	return _Car.contract.Transact(opts, "addReserve", addrs, amount)
}

// AddReserve is a paid mutator transaction binding the contract method 0xb2fa739a.
//
// Solidity: function addReserve(address[] addrs, uint8 amount) returns()
func (_Car *CarSession) AddReserve(addrs []common.Address, amount uint8) (*types.Transaction, error) {
	return _Car.Contract.AddReserve(&_Car.TransactOpts, addrs, amount)
}

// AddReserve is a paid mutator transaction binding the contract method 0xb2fa739a.
//
// Solidity: function addReserve(address[] addrs, uint8 amount) returns()
func (_Car *CarTransactorSession) AddReserve(addrs []common.Address, amount uint8) (*types.Transaction, error) {
	return _Car.Contract.AddReserve(&_Car.TransactOpts, addrs, amount)
}

// AddWhitelist is a paid mutator transaction binding the contract method 0x6eb42e9d.
//
// Solidity: function addWhitelist(address[] addrs, uint8 amount) returns()
func (_Car *CarTransactor) AddWhitelist(opts *bind.TransactOpts, addrs []common.Address, amount uint8) (*types.Transaction, error) {
	return _Car.contract.Transact(opts, "addWhitelist", addrs, amount)
}

// AddWhitelist is a paid mutator transaction binding the contract method 0x6eb42e9d.
//
// Solidity: function addWhitelist(address[] addrs, uint8 amount) returns()
func (_Car *CarSession) AddWhitelist(addrs []common.Address, amount uint8) (*types.Transaction, error) {
	return _Car.Contract.AddWhitelist(&_Car.TransactOpts, addrs, amount)
}

// AddWhitelist is a paid mutator transaction binding the contract method 0x6eb42e9d.
//
// Solidity: function addWhitelist(address[] addrs, uint8 amount) returns()
func (_Car *CarTransactorSession) AddWhitelist(addrs []common.Address, amount uint8) (*types.Transaction, error) {
	return _Car.Contract.AddWhitelist(&_Car.TransactOpts, addrs, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Car *CarTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Car.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Car *CarSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Car.Contract.Approve(&_Car.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Car *CarTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Car.Contract.Approve(&_Car.TransactOpts, to, tokenId)
}

// Claim is a paid mutator transaction binding the contract method 0x95d4063f.
//
// Solidity: function claim(uint8 amount) returns()
func (_Car *CarTransactor) Claim(opts *bind.TransactOpts, amount uint8) (*types.Transaction, error) {
	return _Car.contract.Transact(opts, "claim", amount)
}

// Claim is a paid mutator transaction binding the contract method 0x95d4063f.
//
// Solidity: function claim(uint8 amount) returns()
func (_Car *CarSession) Claim(amount uint8) (*types.Transaction, error) {
	return _Car.Contract.Claim(&_Car.TransactOpts, amount)
}

// Claim is a paid mutator transaction binding the contract method 0x95d4063f.
//
// Solidity: function claim(uint8 amount) returns()
func (_Car *CarTransactorSession) Claim(amount uint8) (*types.Transaction, error) {
	return _Car.Contract.Claim(&_Car.TransactOpts, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Car *CarTransactor) Initialize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Car.contract.Transact(opts, "initialize")
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Car *CarSession) Initialize() (*types.Transaction, error) {
	return _Car.Contract.Initialize(&_Car.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8129fc1c.
//
// Solidity: function initialize() returns()
func (_Car *CarTransactorSession) Initialize() (*types.Transaction, error) {
	return _Car.Contract.Initialize(&_Car.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x6ecd2306.
//
// Solidity: function mint(uint8 amount) payable returns()
func (_Car *CarTransactor) Mint(opts *bind.TransactOpts, amount uint8) (*types.Transaction, error) {
	return _Car.contract.Transact(opts, "mint", amount)
}

// Mint is a paid mutator transaction binding the contract method 0x6ecd2306.
//
// Solidity: function mint(uint8 amount) payable returns()
func (_Car *CarSession) Mint(amount uint8) (*types.Transaction, error) {
	return _Car.Contract.Mint(&_Car.TransactOpts, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x6ecd2306.
//
// Solidity: function mint(uint8 amount) payable returns()
func (_Car *CarTransactorSession) Mint(amount uint8) (*types.Transaction, error) {
	return _Car.Contract.Mint(&_Car.TransactOpts, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Car *CarTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Car.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Car *CarSession) Pause() (*types.Transaction, error) {
	return _Car.Contract.Pause(&_Car.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Car *CarTransactorSession) Pause() (*types.Transaction, error) {
	return _Car.Contract.Pause(&_Car.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Car *CarTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Car.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Car *CarSession) RenounceOwnership() (*types.Transaction, error) {
	return _Car.Contract.RenounceOwnership(&_Car.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Car *CarTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Car.Contract.RenounceOwnership(&_Car.TransactOpts)
}

// Reserve is a paid mutator transaction binding the contract method 0xcd3293de.
//
// Solidity: function reserve() returns()
func (_Car *CarTransactor) Reserve(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Car.contract.Transact(opts, "reserve")
}

// Reserve is a paid mutator transaction binding the contract method 0xcd3293de.
//
// Solidity: function reserve() returns()
func (_Car *CarSession) Reserve() (*types.Transaction, error) {
	return _Car.Contract.Reserve(&_Car.TransactOpts)
}

// Reserve is a paid mutator transaction binding the contract method 0xcd3293de.
//
// Solidity: function reserve() returns()
func (_Car *CarTransactorSession) Reserve() (*types.Transaction, error) {
	return _Car.Contract.Reserve(&_Car.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Car *CarTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Car.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Car *CarSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Car.Contract.SafeTransferFrom(&_Car.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Car *CarTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Car.Contract.SafeTransferFrom(&_Car.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Car *CarTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Car.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Car *CarSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Car.Contract.SafeTransferFrom0(&_Car.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Car *CarTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Car.Contract.SafeTransferFrom0(&_Car.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Car *CarTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Car.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Car *CarSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Car.Contract.SetApprovalForAll(&_Car.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Car *CarTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Car.Contract.SetApprovalForAll(&_Car.TransactOpts, operator, approved)
}

// SetPhase is a paid mutator transaction binding the contract method 0xea41df3d.
//
// Solidity: function setPhase(int8 newPhase) returns()
func (_Car *CarTransactor) SetPhase(opts *bind.TransactOpts, newPhase int8) (*types.Transaction, error) {
	return _Car.contract.Transact(opts, "setPhase", newPhase)
}

// SetPhase is a paid mutator transaction binding the contract method 0xea41df3d.
//
// Solidity: function setPhase(int8 newPhase) returns()
func (_Car *CarSession) SetPhase(newPhase int8) (*types.Transaction, error) {
	return _Car.Contract.SetPhase(&_Car.TransactOpts, newPhase)
}

// SetPhase is a paid mutator transaction binding the contract method 0xea41df3d.
//
// Solidity: function setPhase(int8 newPhase) returns()
func (_Car *CarTransactorSession) SetPhase(newPhase int8) (*types.Transaction, error) {
	return _Car.Contract.SetPhase(&_Car.TransactOpts, newPhase)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Car *CarTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Car.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Car *CarSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Car.Contract.TransferFrom(&_Car.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Car *CarTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Car.Contract.TransferFrom(&_Car.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Car *CarTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Car.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Car *CarSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Car.Contract.TransferOwnership(&_Car.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Car *CarTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Car.Contract.TransferOwnership(&_Car.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Car *CarTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Car.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Car *CarSession) Unpause() (*types.Transaction, error) {
	return _Car.Contract.Unpause(&_Car.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Car *CarTransactorSession) Unpause() (*types.Transaction, error) {
	return _Car.Contract.Unpause(&_Car.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Car *CarTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Car.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Car *CarSession) Withdraw() (*types.Transaction, error) {
	return _Car.Contract.Withdraw(&_Car.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Car *CarTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Car.Contract.Withdraw(&_Car.TransactOpts)
}

// CarApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Car contract.
type CarApprovalIterator struct {
	Event *CarApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CarApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CarApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CarApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CarApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CarApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CarApproval represents a Approval event raised by the Car contract.
type CarApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Car *CarFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*CarApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Car.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CarApprovalIterator{contract: _Car.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Car *CarFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CarApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Car.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CarApproval)
				if err := _Car.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Car *CarFilterer) ParseApproval(log types.Log) (*CarApproval, error) {
	event := new(CarApproval)
	if err := _Car.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CarApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Car contract.
type CarApprovalForAllIterator struct {
	Event *CarApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CarApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CarApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CarApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CarApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CarApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CarApprovalForAll represents a ApprovalForAll event raised by the Car contract.
type CarApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Car *CarFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*CarApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Car.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &CarApprovalForAllIterator{contract: _Car.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Car *CarFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *CarApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Car.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CarApprovalForAll)
				if err := _Car.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Car *CarFilterer) ParseApprovalForAll(log types.Log) (*CarApprovalForAll, error) {
	event := new(CarApprovalForAll)
	if err := _Car.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CarOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Car contract.
type CarOwnershipTransferredIterator struct {
	Event *CarOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CarOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CarOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CarOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CarOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CarOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CarOwnershipTransferred represents a OwnershipTransferred event raised by the Car contract.
type CarOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Car *CarFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CarOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Car.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CarOwnershipTransferredIterator{contract: _Car.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Car *CarFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CarOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Car.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CarOwnershipTransferred)
				if err := _Car.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Car *CarFilterer) ParseOwnershipTransferred(log types.Log) (*CarOwnershipTransferred, error) {
	event := new(CarOwnershipTransferred)
	if err := _Car.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CarPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Car contract.
type CarPausedIterator struct {
	Event *CarPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CarPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CarPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CarPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CarPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CarPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CarPaused represents a Paused event raised by the Car contract.
type CarPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Car *CarFilterer) FilterPaused(opts *bind.FilterOpts) (*CarPausedIterator, error) {

	logs, sub, err := _Car.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CarPausedIterator{contract: _Car.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Car *CarFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CarPaused) (event.Subscription, error) {

	logs, sub, err := _Car.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CarPaused)
				if err := _Car.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Car *CarFilterer) ParsePaused(log types.Log) (*CarPaused, error) {
	event := new(CarPaused)
	if err := _Car.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CarTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Car contract.
type CarTransferIterator struct {
	Event *CarTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CarTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CarTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CarTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CarTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CarTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CarTransfer represents a Transfer event raised by the Car contract.
type CarTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Car *CarFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*CarTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Car.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &CarTransferIterator{contract: _Car.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Car *CarFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CarTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Car.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CarTransfer)
				if err := _Car.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Car *CarFilterer) ParseTransfer(log types.Log) (*CarTransfer, error) {
	event := new(CarTransfer)
	if err := _Car.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CarUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Car contract.
type CarUnpausedIterator struct {
	Event *CarUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *CarUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CarUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(CarUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *CarUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CarUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CarUnpaused represents a Unpaused event raised by the Car contract.
type CarUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Car *CarFilterer) FilterUnpaused(opts *bind.FilterOpts) (*CarUnpausedIterator, error) {

	logs, sub, err := _Car.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CarUnpausedIterator{contract: _Car.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Car *CarFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CarUnpaused) (event.Subscription, error) {

	logs, sub, err := _Car.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CarUnpaused)
				if err := _Car.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Car *CarFilterer) ParseUnpaused(log types.Log) (*CarUnpaused, error) {
	event := new(CarUnpaused)
	if err := _Car.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

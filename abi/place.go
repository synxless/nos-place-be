// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// PixelPixel is an auto generated low-level Go binding around an user-defined struct.
type PixelPixel struct {
	Setter common.Address
	R      uint8
	G      uint8
	B      uint8
	Value  *big.Int
}

// PlaceMetaData contains all meta data concerning the Place contract.
var PlaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"BeingLocked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughValueToClaimPixel\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OutOfBound\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"_newValue\",\"type\":\"uint16\"}],\"name\":\"ClaimingValuePercentageUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int32\",\"name\":\"_newValue\",\"type\":\"int32\"}],\"name\":\"LeftBoundExpansion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint40\",\"name\":\"_newValue\",\"type\":\"uint40\"}],\"name\":\"LockDurationUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int32\",\"name\":\"_newValue\",\"type\":\"int32\"}],\"name\":\"LowerBoundExpansion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"setter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"int64\",\"name\":\"x\",\"type\":\"int64\"},{\"indexed\":true,\"internalType\":\"int64\",\"name\":\"y\",\"type\":\"int64\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"r\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"g\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"b\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"PixelSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int32\",\"name\":\"_newValue\",\"type\":\"int32\"}],\"name\":\"RightBoundExpansion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int32\",\"name\":\"_newValue\",\"type\":\"int32\"}],\"name\":\"UpperBoundExpansion\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"claimingValuePercentage\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int32\",\"name\":\"_leftBound\",\"type\":\"int32\"}],\"name\":\"expandLeftBound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int32\",\"name\":\"_lowerBound\",\"type\":\"int32\"}],\"name\":\"expandLowerBound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int32\",\"name\":\"_rightBound\",\"type\":\"int32\"}],\"name\":\"expandRightBound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int32\",\"name\":\"_upperBound\",\"type\":\"int32\"}],\"name\":\"expandUpperBound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getHeight\",\"outputs\":[{\"internalType\":\"int64\",\"name\":\"\",\"type\":\"int64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int32\",\"name\":\"x\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"y\",\"type\":\"int32\"}],\"name\":\"getPixel\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"setter\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"r\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"g\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"b\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structPixel.Pixel\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWidth\",\"outputs\":[{\"internalType\":\"int64\",\"name\":\"\",\"type\":\"int64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_treasury\",\"type\":\"address\"},{\"internalType\":\"int32\",\"name\":\"_lowerBound\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"_upperBound\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"_leftBound\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"_rightBound\",\"type\":\"int32\"},{\"internalType\":\"uint40\",\"name\":\"_lockDuration\",\"type\":\"uint40\"},{\"internalType\":\"uint16\",\"name\":\"_claimingValuePercentage\",\"type\":\"uint16\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastSet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"leftBound\",\"outputs\":[{\"internalType\":\"int32\",\"name\":\"\",\"type\":\"int32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockDuration\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lowerBound\",\"outputs\":[{\"internalType\":\"int32\",\"name\":\"\",\"type\":\"int32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int64\",\"name\":\"\",\"type\":\"int64\"},{\"internalType\":\"int64\",\"name\":\"\",\"type\":\"int64\"}],\"name\":\"pixels\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"setter\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"r\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"g\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"b\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rightBound\",\"outputs\":[{\"internalType\":\"int32\",\"name\":\"\",\"type\":\"int32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int32\",\"name\":\"x\",\"type\":\"int32\"},{\"internalType\":\"int32\",\"name\":\"y\",\"type\":\"int32\"},{\"internalType\":\"uint8\",\"name\":\"r\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"g\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"b\",\"type\":\"uint8\"}],\"name\":\"setPixel\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treasury\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"_claimingValuePercentage\",\"type\":\"uint16\"}],\"name\":\"updateClaimingValuePercentage\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"_lockDuration\",\"type\":\"uint40\"}],\"name\":\"updateLockDuration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upperBound\",\"outputs\":[{\"internalType\":\"int32\",\"name\":\"\",\"type\":\"int32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PlaceABI is the input ABI used to generate the binding from.
// Deprecated: Use PlaceMetaData.ABI instead.
var PlaceABI = PlaceMetaData.ABI

// Place is an auto generated Go binding around an Ethereum contract.
type Place struct {
	PlaceCaller     // Read-only binding to the contract
	PlaceTransactor // Write-only binding to the contract
	PlaceFilterer   // Log filterer for contract events
}

// PlaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type PlaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PlaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PlaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PlaceSession struct {
	Contract     *Place            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PlaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PlaceCallerSession struct {
	Contract *PlaceCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PlaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PlaceTransactorSession struct {
	Contract     *PlaceTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PlaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type PlaceRaw struct {
	Contract *Place // Generic contract binding to access the raw methods on
}

// PlaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PlaceCallerRaw struct {
	Contract *PlaceCaller // Generic read-only contract binding to access the raw methods on
}

// PlaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PlaceTransactorRaw struct {
	Contract *PlaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPlace creates a new instance of Place, bound to a specific deployed contract.
func NewPlace(address common.Address, backend bind.ContractBackend) (*Place, error) {
	contract, err := bindPlace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Place{PlaceCaller: PlaceCaller{contract: contract}, PlaceTransactor: PlaceTransactor{contract: contract}, PlaceFilterer: PlaceFilterer{contract: contract}}, nil
}

// NewPlaceCaller creates a new read-only instance of Place, bound to a specific deployed contract.
func NewPlaceCaller(address common.Address, caller bind.ContractCaller) (*PlaceCaller, error) {
	contract, err := bindPlace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PlaceCaller{contract: contract}, nil
}

// NewPlaceTransactor creates a new write-only instance of Place, bound to a specific deployed contract.
func NewPlaceTransactor(address common.Address, transactor bind.ContractTransactor) (*PlaceTransactor, error) {
	contract, err := bindPlace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PlaceTransactor{contract: contract}, nil
}

// NewPlaceFilterer creates a new log filterer instance of Place, bound to a specific deployed contract.
func NewPlaceFilterer(address common.Address, filterer bind.ContractFilterer) (*PlaceFilterer, error) {
	contract, err := bindPlace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PlaceFilterer{contract: contract}, nil
}

// bindPlace binds a generic wrapper to an already deployed contract.
func bindPlace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PlaceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Place *PlaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Place.Contract.PlaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Place *PlaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Place.Contract.PlaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Place *PlaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Place.Contract.PlaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Place *PlaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Place.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Place *PlaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Place.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Place *PlaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Place.Contract.contract.Transact(opts, method, params...)
}

// ClaimingValuePercentage is a free data retrieval call binding the contract method 0xbbef719b.
//
// Solidity: function claimingValuePercentage() view returns(uint16)
func (_Place *PlaceCaller) ClaimingValuePercentage(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Place.contract.Call(opts, &out, "claimingValuePercentage")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ClaimingValuePercentage is a free data retrieval call binding the contract method 0xbbef719b.
//
// Solidity: function claimingValuePercentage() view returns(uint16)
func (_Place *PlaceSession) ClaimingValuePercentage() (uint16, error) {
	return _Place.Contract.ClaimingValuePercentage(&_Place.CallOpts)
}

// ClaimingValuePercentage is a free data retrieval call binding the contract method 0xbbef719b.
//
// Solidity: function claimingValuePercentage() view returns(uint16)
func (_Place *PlaceCallerSession) ClaimingValuePercentage() (uint16, error) {
	return _Place.Contract.ClaimingValuePercentage(&_Place.CallOpts)
}

// GetHeight is a free data retrieval call binding the contract method 0x19efb11d.
//
// Solidity: function getHeight() view returns(int64)
func (_Place *PlaceCaller) GetHeight(opts *bind.CallOpts) (int64, error) {
	var out []interface{}
	err := _Place.contract.Call(opts, &out, "getHeight")

	if err != nil {
		return *new(int64), err
	}

	out0 := *abi.ConvertType(out[0], new(int64)).(*int64)

	return out0, err

}

// GetHeight is a free data retrieval call binding the contract method 0x19efb11d.
//
// Solidity: function getHeight() view returns(int64)
func (_Place *PlaceSession) GetHeight() (int64, error) {
	return _Place.Contract.GetHeight(&_Place.CallOpts)
}

// GetHeight is a free data retrieval call binding the contract method 0x19efb11d.
//
// Solidity: function getHeight() view returns(int64)
func (_Place *PlaceCallerSession) GetHeight() (int64, error) {
	return _Place.Contract.GetHeight(&_Place.CallOpts)
}

// GetPixel is a free data retrieval call binding the contract method 0x3c7e748d.
//
// Solidity: function getPixel(int32 x, int32 y) view returns((address,uint8,uint8,uint8,uint256))
func (_Place *PlaceCaller) GetPixel(opts *bind.CallOpts, x int32, y int32) (PixelPixel, error) {
	var out []interface{}
	err := _Place.contract.Call(opts, &out, "getPixel", x, y)

	if err != nil {
		return *new(PixelPixel), err
	}

	out0 := *abi.ConvertType(out[0], new(PixelPixel)).(*PixelPixel)

	return out0, err

}

// GetPixel is a free data retrieval call binding the contract method 0x3c7e748d.
//
// Solidity: function getPixel(int32 x, int32 y) view returns((address,uint8,uint8,uint8,uint256))
func (_Place *PlaceSession) GetPixel(x int32, y int32) (PixelPixel, error) {
	return _Place.Contract.GetPixel(&_Place.CallOpts, x, y)
}

// GetPixel is a free data retrieval call binding the contract method 0x3c7e748d.
//
// Solidity: function getPixel(int32 x, int32 y) view returns((address,uint8,uint8,uint8,uint256))
func (_Place *PlaceCallerSession) GetPixel(x int32, y int32) (PixelPixel, error) {
	return _Place.Contract.GetPixel(&_Place.CallOpts, x, y)
}

// GetWidth is a free data retrieval call binding the contract method 0xf8e195b3.
//
// Solidity: function getWidth() view returns(int64)
func (_Place *PlaceCaller) GetWidth(opts *bind.CallOpts) (int64, error) {
	var out []interface{}
	err := _Place.contract.Call(opts, &out, "getWidth")

	if err != nil {
		return *new(int64), err
	}

	out0 := *abi.ConvertType(out[0], new(int64)).(*int64)

	return out0, err

}

// GetWidth is a free data retrieval call binding the contract method 0xf8e195b3.
//
// Solidity: function getWidth() view returns(int64)
func (_Place *PlaceSession) GetWidth() (int64, error) {
	return _Place.Contract.GetWidth(&_Place.CallOpts)
}

// GetWidth is a free data retrieval call binding the contract method 0xf8e195b3.
//
// Solidity: function getWidth() view returns(int64)
func (_Place *PlaceCallerSession) GetWidth() (int64, error) {
	return _Place.Contract.GetWidth(&_Place.CallOpts)
}

// LastSet is a free data retrieval call binding the contract method 0xe5a000a6.
//
// Solidity: function lastSet(address ) view returns(uint256)
func (_Place *PlaceCaller) LastSet(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Place.contract.Call(opts, &out, "lastSet", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastSet is a free data retrieval call binding the contract method 0xe5a000a6.
//
// Solidity: function lastSet(address ) view returns(uint256)
func (_Place *PlaceSession) LastSet(arg0 common.Address) (*big.Int, error) {
	return _Place.Contract.LastSet(&_Place.CallOpts, arg0)
}

// LastSet is a free data retrieval call binding the contract method 0xe5a000a6.
//
// Solidity: function lastSet(address ) view returns(uint256)
func (_Place *PlaceCallerSession) LastSet(arg0 common.Address) (*big.Int, error) {
	return _Place.Contract.LastSet(&_Place.CallOpts, arg0)
}

// LeftBound is a free data retrieval call binding the contract method 0xf7049a89.
//
// Solidity: function leftBound() view returns(int32)
func (_Place *PlaceCaller) LeftBound(opts *bind.CallOpts) (int32, error) {
	var out []interface{}
	err := _Place.contract.Call(opts, &out, "leftBound")

	if err != nil {
		return *new(int32), err
	}

	out0 := *abi.ConvertType(out[0], new(int32)).(*int32)

	return out0, err

}

// LeftBound is a free data retrieval call binding the contract method 0xf7049a89.
//
// Solidity: function leftBound() view returns(int32)
func (_Place *PlaceSession) LeftBound() (int32, error) {
	return _Place.Contract.LeftBound(&_Place.CallOpts)
}

// LeftBound is a free data retrieval call binding the contract method 0xf7049a89.
//
// Solidity: function leftBound() view returns(int32)
func (_Place *PlaceCallerSession) LeftBound() (int32, error) {
	return _Place.Contract.LeftBound(&_Place.CallOpts)
}

// LockDuration is a free data retrieval call binding the contract method 0x04554443.
//
// Solidity: function lockDuration() view returns(uint40)
func (_Place *PlaceCaller) LockDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Place.contract.Call(opts, &out, "lockDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockDuration is a free data retrieval call binding the contract method 0x04554443.
//
// Solidity: function lockDuration() view returns(uint40)
func (_Place *PlaceSession) LockDuration() (*big.Int, error) {
	return _Place.Contract.LockDuration(&_Place.CallOpts)
}

// LockDuration is a free data retrieval call binding the contract method 0x04554443.
//
// Solidity: function lockDuration() view returns(uint40)
func (_Place *PlaceCallerSession) LockDuration() (*big.Int, error) {
	return _Place.Contract.LockDuration(&_Place.CallOpts)
}

// LowerBound is a free data retrieval call binding the contract method 0xa384d6ff.
//
// Solidity: function lowerBound() view returns(int32)
func (_Place *PlaceCaller) LowerBound(opts *bind.CallOpts) (int32, error) {
	var out []interface{}
	err := _Place.contract.Call(opts, &out, "lowerBound")

	if err != nil {
		return *new(int32), err
	}

	out0 := *abi.ConvertType(out[0], new(int32)).(*int32)

	return out0, err

}

// LowerBound is a free data retrieval call binding the contract method 0xa384d6ff.
//
// Solidity: function lowerBound() view returns(int32)
func (_Place *PlaceSession) LowerBound() (int32, error) {
	return _Place.Contract.LowerBound(&_Place.CallOpts)
}

// LowerBound is a free data retrieval call binding the contract method 0xa384d6ff.
//
// Solidity: function lowerBound() view returns(int32)
func (_Place *PlaceCallerSession) LowerBound() (int32, error) {
	return _Place.Contract.LowerBound(&_Place.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Place *PlaceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Place.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Place *PlaceSession) Owner() (common.Address, error) {
	return _Place.Contract.Owner(&_Place.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Place *PlaceCallerSession) Owner() (common.Address, error) {
	return _Place.Contract.Owner(&_Place.CallOpts)
}

// Pixels is a free data retrieval call binding the contract method 0x9598ba4e.
//
// Solidity: function pixels(int64 , int64 ) view returns(address setter, uint8 r, uint8 g, uint8 b, uint256 value)
func (_Place *PlaceCaller) Pixels(opts *bind.CallOpts, arg0 int64, arg1 int64) (struct {
	Setter common.Address
	R      uint8
	G      uint8
	B      uint8
	Value  *big.Int
}, error) {
	var out []interface{}
	err := _Place.contract.Call(opts, &out, "pixels", arg0, arg1)

	outstruct := new(struct {
		Setter common.Address
		R      uint8
		G      uint8
		B      uint8
		Value  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Setter = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.R = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.G = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.B = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.Value = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Pixels is a free data retrieval call binding the contract method 0x9598ba4e.
//
// Solidity: function pixels(int64 , int64 ) view returns(address setter, uint8 r, uint8 g, uint8 b, uint256 value)
func (_Place *PlaceSession) Pixels(arg0 int64, arg1 int64) (struct {
	Setter common.Address
	R      uint8
	G      uint8
	B      uint8
	Value  *big.Int
}, error) {
	return _Place.Contract.Pixels(&_Place.CallOpts, arg0, arg1)
}

// Pixels is a free data retrieval call binding the contract method 0x9598ba4e.
//
// Solidity: function pixels(int64 , int64 ) view returns(address setter, uint8 r, uint8 g, uint8 b, uint256 value)
func (_Place *PlaceCallerSession) Pixels(arg0 int64, arg1 int64) (struct {
	Setter common.Address
	R      uint8
	G      uint8
	B      uint8
	Value  *big.Int
}, error) {
	return _Place.Contract.Pixels(&_Place.CallOpts, arg0, arg1)
}

// RightBound is a free data retrieval call binding the contract method 0x1e100aa6.
//
// Solidity: function rightBound() view returns(int32)
func (_Place *PlaceCaller) RightBound(opts *bind.CallOpts) (int32, error) {
	var out []interface{}
	err := _Place.contract.Call(opts, &out, "rightBound")

	if err != nil {
		return *new(int32), err
	}

	out0 := *abi.ConvertType(out[0], new(int32)).(*int32)

	return out0, err

}

// RightBound is a free data retrieval call binding the contract method 0x1e100aa6.
//
// Solidity: function rightBound() view returns(int32)
func (_Place *PlaceSession) RightBound() (int32, error) {
	return _Place.Contract.RightBound(&_Place.CallOpts)
}

// RightBound is a free data retrieval call binding the contract method 0x1e100aa6.
//
// Solidity: function rightBound() view returns(int32)
func (_Place *PlaceCallerSession) RightBound() (int32, error) {
	return _Place.Contract.RightBound(&_Place.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Place *PlaceCaller) Treasury(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Place.contract.Call(opts, &out, "treasury")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Place *PlaceSession) Treasury() (common.Address, error) {
	return _Place.Contract.Treasury(&_Place.CallOpts)
}

// Treasury is a free data retrieval call binding the contract method 0x61d027b3.
//
// Solidity: function treasury() view returns(address)
func (_Place *PlaceCallerSession) Treasury() (common.Address, error) {
	return _Place.Contract.Treasury(&_Place.CallOpts)
}

// UpperBound is a free data retrieval call binding the contract method 0xb09ad8a0.
//
// Solidity: function upperBound() view returns(int32)
func (_Place *PlaceCaller) UpperBound(opts *bind.CallOpts) (int32, error) {
	var out []interface{}
	err := _Place.contract.Call(opts, &out, "upperBound")

	if err != nil {
		return *new(int32), err
	}

	out0 := *abi.ConvertType(out[0], new(int32)).(*int32)

	return out0, err

}

// UpperBound is a free data retrieval call binding the contract method 0xb09ad8a0.
//
// Solidity: function upperBound() view returns(int32)
func (_Place *PlaceSession) UpperBound() (int32, error) {
	return _Place.Contract.UpperBound(&_Place.CallOpts)
}

// UpperBound is a free data retrieval call binding the contract method 0xb09ad8a0.
//
// Solidity: function upperBound() view returns(int32)
func (_Place *PlaceCallerSession) UpperBound() (int32, error) {
	return _Place.Contract.UpperBound(&_Place.CallOpts)
}

// ExpandLeftBound is a paid mutator transaction binding the contract method 0x9a0f6786.
//
// Solidity: function expandLeftBound(int32 _leftBound) returns()
func (_Place *PlaceTransactor) ExpandLeftBound(opts *bind.TransactOpts, _leftBound int32) (*types.Transaction, error) {
	return _Place.contract.Transact(opts, "expandLeftBound", _leftBound)
}

// ExpandLeftBound is a paid mutator transaction binding the contract method 0x9a0f6786.
//
// Solidity: function expandLeftBound(int32 _leftBound) returns()
func (_Place *PlaceSession) ExpandLeftBound(_leftBound int32) (*types.Transaction, error) {
	return _Place.Contract.ExpandLeftBound(&_Place.TransactOpts, _leftBound)
}

// ExpandLeftBound is a paid mutator transaction binding the contract method 0x9a0f6786.
//
// Solidity: function expandLeftBound(int32 _leftBound) returns()
func (_Place *PlaceTransactorSession) ExpandLeftBound(_leftBound int32) (*types.Transaction, error) {
	return _Place.Contract.ExpandLeftBound(&_Place.TransactOpts, _leftBound)
}

// ExpandLowerBound is a paid mutator transaction binding the contract method 0x308f8bbf.
//
// Solidity: function expandLowerBound(int32 _lowerBound) returns()
func (_Place *PlaceTransactor) ExpandLowerBound(opts *bind.TransactOpts, _lowerBound int32) (*types.Transaction, error) {
	return _Place.contract.Transact(opts, "expandLowerBound", _lowerBound)
}

// ExpandLowerBound is a paid mutator transaction binding the contract method 0x308f8bbf.
//
// Solidity: function expandLowerBound(int32 _lowerBound) returns()
func (_Place *PlaceSession) ExpandLowerBound(_lowerBound int32) (*types.Transaction, error) {
	return _Place.Contract.ExpandLowerBound(&_Place.TransactOpts, _lowerBound)
}

// ExpandLowerBound is a paid mutator transaction binding the contract method 0x308f8bbf.
//
// Solidity: function expandLowerBound(int32 _lowerBound) returns()
func (_Place *PlaceTransactorSession) ExpandLowerBound(_lowerBound int32) (*types.Transaction, error) {
	return _Place.Contract.ExpandLowerBound(&_Place.TransactOpts, _lowerBound)
}

// ExpandRightBound is a paid mutator transaction binding the contract method 0x94bfcae7.
//
// Solidity: function expandRightBound(int32 _rightBound) returns()
func (_Place *PlaceTransactor) ExpandRightBound(opts *bind.TransactOpts, _rightBound int32) (*types.Transaction, error) {
	return _Place.contract.Transact(opts, "expandRightBound", _rightBound)
}

// ExpandRightBound is a paid mutator transaction binding the contract method 0x94bfcae7.
//
// Solidity: function expandRightBound(int32 _rightBound) returns()
func (_Place *PlaceSession) ExpandRightBound(_rightBound int32) (*types.Transaction, error) {
	return _Place.Contract.ExpandRightBound(&_Place.TransactOpts, _rightBound)
}

// ExpandRightBound is a paid mutator transaction binding the contract method 0x94bfcae7.
//
// Solidity: function expandRightBound(int32 _rightBound) returns()
func (_Place *PlaceTransactorSession) ExpandRightBound(_rightBound int32) (*types.Transaction, error) {
	return _Place.Contract.ExpandRightBound(&_Place.TransactOpts, _rightBound)
}

// ExpandUpperBound is a paid mutator transaction binding the contract method 0xe941fc9d.
//
// Solidity: function expandUpperBound(int32 _upperBound) returns()
func (_Place *PlaceTransactor) ExpandUpperBound(opts *bind.TransactOpts, _upperBound int32) (*types.Transaction, error) {
	return _Place.contract.Transact(opts, "expandUpperBound", _upperBound)
}

// ExpandUpperBound is a paid mutator transaction binding the contract method 0xe941fc9d.
//
// Solidity: function expandUpperBound(int32 _upperBound) returns()
func (_Place *PlaceSession) ExpandUpperBound(_upperBound int32) (*types.Transaction, error) {
	return _Place.Contract.ExpandUpperBound(&_Place.TransactOpts, _upperBound)
}

// ExpandUpperBound is a paid mutator transaction binding the contract method 0xe941fc9d.
//
// Solidity: function expandUpperBound(int32 _upperBound) returns()
func (_Place *PlaceTransactorSession) ExpandUpperBound(_upperBound int32) (*types.Transaction, error) {
	return _Place.Contract.ExpandUpperBound(&_Place.TransactOpts, _upperBound)
}

// Initialize is a paid mutator transaction binding the contract method 0x4052cb8a.
//
// Solidity: function initialize(address _treasury, int32 _lowerBound, int32 _upperBound, int32 _leftBound, int32 _rightBound, uint40 _lockDuration, uint16 _claimingValuePercentage) returns()
func (_Place *PlaceTransactor) Initialize(opts *bind.TransactOpts, _treasury common.Address, _lowerBound int32, _upperBound int32, _leftBound int32, _rightBound int32, _lockDuration *big.Int, _claimingValuePercentage uint16) (*types.Transaction, error) {
	return _Place.contract.Transact(opts, "initialize", _treasury, _lowerBound, _upperBound, _leftBound, _rightBound, _lockDuration, _claimingValuePercentage)
}

// Initialize is a paid mutator transaction binding the contract method 0x4052cb8a.
//
// Solidity: function initialize(address _treasury, int32 _lowerBound, int32 _upperBound, int32 _leftBound, int32 _rightBound, uint40 _lockDuration, uint16 _claimingValuePercentage) returns()
func (_Place *PlaceSession) Initialize(_treasury common.Address, _lowerBound int32, _upperBound int32, _leftBound int32, _rightBound int32, _lockDuration *big.Int, _claimingValuePercentage uint16) (*types.Transaction, error) {
	return _Place.Contract.Initialize(&_Place.TransactOpts, _treasury, _lowerBound, _upperBound, _leftBound, _rightBound, _lockDuration, _claimingValuePercentage)
}

// Initialize is a paid mutator transaction binding the contract method 0x4052cb8a.
//
// Solidity: function initialize(address _treasury, int32 _lowerBound, int32 _upperBound, int32 _leftBound, int32 _rightBound, uint40 _lockDuration, uint16 _claimingValuePercentage) returns()
func (_Place *PlaceTransactorSession) Initialize(_treasury common.Address, _lowerBound int32, _upperBound int32, _leftBound int32, _rightBound int32, _lockDuration *big.Int, _claimingValuePercentage uint16) (*types.Transaction, error) {
	return _Place.Contract.Initialize(&_Place.TransactOpts, _treasury, _lowerBound, _upperBound, _leftBound, _rightBound, _lockDuration, _claimingValuePercentage)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Place *PlaceTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Place.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Place *PlaceSession) RenounceOwnership() (*types.Transaction, error) {
	return _Place.Contract.RenounceOwnership(&_Place.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Place *PlaceTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Place.Contract.RenounceOwnership(&_Place.TransactOpts)
}

// SetPixel is a paid mutator transaction binding the contract method 0x712b6611.
//
// Solidity: function setPixel(int32 x, int32 y, uint8 r, uint8 g, uint8 b) payable returns()
func (_Place *PlaceTransactor) SetPixel(opts *bind.TransactOpts, x int32, y int32, r uint8, g uint8, b uint8) (*types.Transaction, error) {
	return _Place.contract.Transact(opts, "setPixel", x, y, r, g, b)
}

// SetPixel is a paid mutator transaction binding the contract method 0x712b6611.
//
// Solidity: function setPixel(int32 x, int32 y, uint8 r, uint8 g, uint8 b) payable returns()
func (_Place *PlaceSession) SetPixel(x int32, y int32, r uint8, g uint8, b uint8) (*types.Transaction, error) {
	return _Place.Contract.SetPixel(&_Place.TransactOpts, x, y, r, g, b)
}

// SetPixel is a paid mutator transaction binding the contract method 0x712b6611.
//
// Solidity: function setPixel(int32 x, int32 y, uint8 r, uint8 g, uint8 b) payable returns()
func (_Place *PlaceTransactorSession) SetPixel(x int32, y int32, r uint8, g uint8, b uint8) (*types.Transaction, error) {
	return _Place.Contract.SetPixel(&_Place.TransactOpts, x, y, r, g, b)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Place *PlaceTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Place.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Place *PlaceSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Place.Contract.TransferOwnership(&_Place.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Place *PlaceTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Place.Contract.TransferOwnership(&_Place.TransactOpts, newOwner)
}

// UpdateClaimingValuePercentage is a paid mutator transaction binding the contract method 0xfce0aea5.
//
// Solidity: function updateClaimingValuePercentage(uint16 _claimingValuePercentage) returns()
func (_Place *PlaceTransactor) UpdateClaimingValuePercentage(opts *bind.TransactOpts, _claimingValuePercentage uint16) (*types.Transaction, error) {
	return _Place.contract.Transact(opts, "updateClaimingValuePercentage", _claimingValuePercentage)
}

// UpdateClaimingValuePercentage is a paid mutator transaction binding the contract method 0xfce0aea5.
//
// Solidity: function updateClaimingValuePercentage(uint16 _claimingValuePercentage) returns()
func (_Place *PlaceSession) UpdateClaimingValuePercentage(_claimingValuePercentage uint16) (*types.Transaction, error) {
	return _Place.Contract.UpdateClaimingValuePercentage(&_Place.TransactOpts, _claimingValuePercentage)
}

// UpdateClaimingValuePercentage is a paid mutator transaction binding the contract method 0xfce0aea5.
//
// Solidity: function updateClaimingValuePercentage(uint16 _claimingValuePercentage) returns()
func (_Place *PlaceTransactorSession) UpdateClaimingValuePercentage(_claimingValuePercentage uint16) (*types.Transaction, error) {
	return _Place.Contract.UpdateClaimingValuePercentage(&_Place.TransactOpts, _claimingValuePercentage)
}

// UpdateLockDuration is a paid mutator transaction binding the contract method 0x5c4a26c7.
//
// Solidity: function updateLockDuration(uint40 _lockDuration) returns()
func (_Place *PlaceTransactor) UpdateLockDuration(opts *bind.TransactOpts, _lockDuration *big.Int) (*types.Transaction, error) {
	return _Place.contract.Transact(opts, "updateLockDuration", _lockDuration)
}

// UpdateLockDuration is a paid mutator transaction binding the contract method 0x5c4a26c7.
//
// Solidity: function updateLockDuration(uint40 _lockDuration) returns()
func (_Place *PlaceSession) UpdateLockDuration(_lockDuration *big.Int) (*types.Transaction, error) {
	return _Place.Contract.UpdateLockDuration(&_Place.TransactOpts, _lockDuration)
}

// UpdateLockDuration is a paid mutator transaction binding the contract method 0x5c4a26c7.
//
// Solidity: function updateLockDuration(uint40 _lockDuration) returns()
func (_Place *PlaceTransactorSession) UpdateLockDuration(_lockDuration *big.Int) (*types.Transaction, error) {
	return _Place.Contract.UpdateLockDuration(&_Place.TransactOpts, _lockDuration)
}

// PlaceClaimingValuePercentageUpdateIterator is returned from FilterClaimingValuePercentageUpdate and is used to iterate over the raw logs and unpacked data for ClaimingValuePercentageUpdate events raised by the Place contract.
type PlaceClaimingValuePercentageUpdateIterator struct {
	Event *PlaceClaimingValuePercentageUpdate // Event containing the contract specifics and raw log

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
func (it *PlaceClaimingValuePercentageUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlaceClaimingValuePercentageUpdate)
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
		it.Event = new(PlaceClaimingValuePercentageUpdate)
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
func (it *PlaceClaimingValuePercentageUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlaceClaimingValuePercentageUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlaceClaimingValuePercentageUpdate represents a ClaimingValuePercentageUpdate event raised by the Place contract.
type PlaceClaimingValuePercentageUpdate struct {
	NewValue uint16
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterClaimingValuePercentageUpdate is a free log retrieval operation binding the contract event 0x8ef39cd167324b4bebf933406d4a4d6444694ac145c337744ea03f805a26efe1.
//
// Solidity: event ClaimingValuePercentageUpdate(uint16 _newValue)
func (_Place *PlaceFilterer) FilterClaimingValuePercentageUpdate(opts *bind.FilterOpts) (*PlaceClaimingValuePercentageUpdateIterator, error) {

	logs, sub, err := _Place.contract.FilterLogs(opts, "ClaimingValuePercentageUpdate")
	if err != nil {
		return nil, err
	}
	return &PlaceClaimingValuePercentageUpdateIterator{contract: _Place.contract, event: "ClaimingValuePercentageUpdate", logs: logs, sub: sub}, nil
}

// WatchClaimingValuePercentageUpdate is a free log subscription operation binding the contract event 0x8ef39cd167324b4bebf933406d4a4d6444694ac145c337744ea03f805a26efe1.
//
// Solidity: event ClaimingValuePercentageUpdate(uint16 _newValue)
func (_Place *PlaceFilterer) WatchClaimingValuePercentageUpdate(opts *bind.WatchOpts, sink chan<- *PlaceClaimingValuePercentageUpdate) (event.Subscription, error) {

	logs, sub, err := _Place.contract.WatchLogs(opts, "ClaimingValuePercentageUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlaceClaimingValuePercentageUpdate)
				if err := _Place.contract.UnpackLog(event, "ClaimingValuePercentageUpdate", log); err != nil {
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

// ParseClaimingValuePercentageUpdate is a log parse operation binding the contract event 0x8ef39cd167324b4bebf933406d4a4d6444694ac145c337744ea03f805a26efe1.
//
// Solidity: event ClaimingValuePercentageUpdate(uint16 _newValue)
func (_Place *PlaceFilterer) ParseClaimingValuePercentageUpdate(log types.Log) (*PlaceClaimingValuePercentageUpdate, error) {
	event := new(PlaceClaimingValuePercentageUpdate)
	if err := _Place.contract.UnpackLog(event, "ClaimingValuePercentageUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlaceInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Place contract.
type PlaceInitializedIterator struct {
	Event *PlaceInitialized // Event containing the contract specifics and raw log

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
func (it *PlaceInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlaceInitialized)
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
		it.Event = new(PlaceInitialized)
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
func (it *PlaceInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlaceInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlaceInitialized represents a Initialized event raised by the Place contract.
type PlaceInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Place *PlaceFilterer) FilterInitialized(opts *bind.FilterOpts) (*PlaceInitializedIterator, error) {

	logs, sub, err := _Place.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PlaceInitializedIterator{contract: _Place.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Place *PlaceFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PlaceInitialized) (event.Subscription, error) {

	logs, sub, err := _Place.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlaceInitialized)
				if err := _Place.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Place *PlaceFilterer) ParseInitialized(log types.Log) (*PlaceInitialized, error) {
	event := new(PlaceInitialized)
	if err := _Place.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlaceLeftBoundExpansionIterator is returned from FilterLeftBoundExpansion and is used to iterate over the raw logs and unpacked data for LeftBoundExpansion events raised by the Place contract.
type PlaceLeftBoundExpansionIterator struct {
	Event *PlaceLeftBoundExpansion // Event containing the contract specifics and raw log

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
func (it *PlaceLeftBoundExpansionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlaceLeftBoundExpansion)
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
		it.Event = new(PlaceLeftBoundExpansion)
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
func (it *PlaceLeftBoundExpansionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlaceLeftBoundExpansionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlaceLeftBoundExpansion represents a LeftBoundExpansion event raised by the Place contract.
type PlaceLeftBoundExpansion struct {
	NewValue int32
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLeftBoundExpansion is a free log retrieval operation binding the contract event 0x9a6a7042e4ee080962ea1e847740193d0dcfcbaae29623d4151688cad4c3871a.
//
// Solidity: event LeftBoundExpansion(int32 _newValue)
func (_Place *PlaceFilterer) FilterLeftBoundExpansion(opts *bind.FilterOpts) (*PlaceLeftBoundExpansionIterator, error) {

	logs, sub, err := _Place.contract.FilterLogs(opts, "LeftBoundExpansion")
	if err != nil {
		return nil, err
	}
	return &PlaceLeftBoundExpansionIterator{contract: _Place.contract, event: "LeftBoundExpansion", logs: logs, sub: sub}, nil
}

// WatchLeftBoundExpansion is a free log subscription operation binding the contract event 0x9a6a7042e4ee080962ea1e847740193d0dcfcbaae29623d4151688cad4c3871a.
//
// Solidity: event LeftBoundExpansion(int32 _newValue)
func (_Place *PlaceFilterer) WatchLeftBoundExpansion(opts *bind.WatchOpts, sink chan<- *PlaceLeftBoundExpansion) (event.Subscription, error) {

	logs, sub, err := _Place.contract.WatchLogs(opts, "LeftBoundExpansion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlaceLeftBoundExpansion)
				if err := _Place.contract.UnpackLog(event, "LeftBoundExpansion", log); err != nil {
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

// ParseLeftBoundExpansion is a log parse operation binding the contract event 0x9a6a7042e4ee080962ea1e847740193d0dcfcbaae29623d4151688cad4c3871a.
//
// Solidity: event LeftBoundExpansion(int32 _newValue)
func (_Place *PlaceFilterer) ParseLeftBoundExpansion(log types.Log) (*PlaceLeftBoundExpansion, error) {
	event := new(PlaceLeftBoundExpansion)
	if err := _Place.contract.UnpackLog(event, "LeftBoundExpansion", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlaceLockDurationUpdateIterator is returned from FilterLockDurationUpdate and is used to iterate over the raw logs and unpacked data for LockDurationUpdate events raised by the Place contract.
type PlaceLockDurationUpdateIterator struct {
	Event *PlaceLockDurationUpdate // Event containing the contract specifics and raw log

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
func (it *PlaceLockDurationUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlaceLockDurationUpdate)
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
		it.Event = new(PlaceLockDurationUpdate)
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
func (it *PlaceLockDurationUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlaceLockDurationUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlaceLockDurationUpdate represents a LockDurationUpdate event raised by the Place contract.
type PlaceLockDurationUpdate struct {
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLockDurationUpdate is a free log retrieval operation binding the contract event 0x60f0690657cbf801ae73f50cc900d0910723d1a5fc7f84b572b73c0b4591a78a.
//
// Solidity: event LockDurationUpdate(uint40 _newValue)
func (_Place *PlaceFilterer) FilterLockDurationUpdate(opts *bind.FilterOpts) (*PlaceLockDurationUpdateIterator, error) {

	logs, sub, err := _Place.contract.FilterLogs(opts, "LockDurationUpdate")
	if err != nil {
		return nil, err
	}
	return &PlaceLockDurationUpdateIterator{contract: _Place.contract, event: "LockDurationUpdate", logs: logs, sub: sub}, nil
}

// WatchLockDurationUpdate is a free log subscription operation binding the contract event 0x60f0690657cbf801ae73f50cc900d0910723d1a5fc7f84b572b73c0b4591a78a.
//
// Solidity: event LockDurationUpdate(uint40 _newValue)
func (_Place *PlaceFilterer) WatchLockDurationUpdate(opts *bind.WatchOpts, sink chan<- *PlaceLockDurationUpdate) (event.Subscription, error) {

	logs, sub, err := _Place.contract.WatchLogs(opts, "LockDurationUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlaceLockDurationUpdate)
				if err := _Place.contract.UnpackLog(event, "LockDurationUpdate", log); err != nil {
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

// ParseLockDurationUpdate is a log parse operation binding the contract event 0x60f0690657cbf801ae73f50cc900d0910723d1a5fc7f84b572b73c0b4591a78a.
//
// Solidity: event LockDurationUpdate(uint40 _newValue)
func (_Place *PlaceFilterer) ParseLockDurationUpdate(log types.Log) (*PlaceLockDurationUpdate, error) {
	event := new(PlaceLockDurationUpdate)
	if err := _Place.contract.UnpackLog(event, "LockDurationUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlaceLowerBoundExpansionIterator is returned from FilterLowerBoundExpansion and is used to iterate over the raw logs and unpacked data for LowerBoundExpansion events raised by the Place contract.
type PlaceLowerBoundExpansionIterator struct {
	Event *PlaceLowerBoundExpansion // Event containing the contract specifics and raw log

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
func (it *PlaceLowerBoundExpansionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlaceLowerBoundExpansion)
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
		it.Event = new(PlaceLowerBoundExpansion)
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
func (it *PlaceLowerBoundExpansionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlaceLowerBoundExpansionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlaceLowerBoundExpansion represents a LowerBoundExpansion event raised by the Place contract.
type PlaceLowerBoundExpansion struct {
	NewValue int32
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLowerBoundExpansion is a free log retrieval operation binding the contract event 0x5e3d464e77c826a1fe11b9dbc67fa1c6bc415ed5bdda4dd3863abe545607386f.
//
// Solidity: event LowerBoundExpansion(int32 _newValue)
func (_Place *PlaceFilterer) FilterLowerBoundExpansion(opts *bind.FilterOpts) (*PlaceLowerBoundExpansionIterator, error) {

	logs, sub, err := _Place.contract.FilterLogs(opts, "LowerBoundExpansion")
	if err != nil {
		return nil, err
	}
	return &PlaceLowerBoundExpansionIterator{contract: _Place.contract, event: "LowerBoundExpansion", logs: logs, sub: sub}, nil
}

// WatchLowerBoundExpansion is a free log subscription operation binding the contract event 0x5e3d464e77c826a1fe11b9dbc67fa1c6bc415ed5bdda4dd3863abe545607386f.
//
// Solidity: event LowerBoundExpansion(int32 _newValue)
func (_Place *PlaceFilterer) WatchLowerBoundExpansion(opts *bind.WatchOpts, sink chan<- *PlaceLowerBoundExpansion) (event.Subscription, error) {

	logs, sub, err := _Place.contract.WatchLogs(opts, "LowerBoundExpansion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlaceLowerBoundExpansion)
				if err := _Place.contract.UnpackLog(event, "LowerBoundExpansion", log); err != nil {
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

// ParseLowerBoundExpansion is a log parse operation binding the contract event 0x5e3d464e77c826a1fe11b9dbc67fa1c6bc415ed5bdda4dd3863abe545607386f.
//
// Solidity: event LowerBoundExpansion(int32 _newValue)
func (_Place *PlaceFilterer) ParseLowerBoundExpansion(log types.Log) (*PlaceLowerBoundExpansion, error) {
	event := new(PlaceLowerBoundExpansion)
	if err := _Place.contract.UnpackLog(event, "LowerBoundExpansion", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlaceOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Place contract.
type PlaceOwnershipTransferredIterator struct {
	Event *PlaceOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PlaceOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlaceOwnershipTransferred)
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
		it.Event = new(PlaceOwnershipTransferred)
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
func (it *PlaceOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlaceOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlaceOwnershipTransferred represents a OwnershipTransferred event raised by the Place contract.
type PlaceOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Place *PlaceFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PlaceOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Place.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PlaceOwnershipTransferredIterator{contract: _Place.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Place *PlaceFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PlaceOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Place.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlaceOwnershipTransferred)
				if err := _Place.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Place *PlaceFilterer) ParseOwnershipTransferred(log types.Log) (*PlaceOwnershipTransferred, error) {
	event := new(PlaceOwnershipTransferred)
	if err := _Place.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlacePixelSetIterator is returned from FilterPixelSet and is used to iterate over the raw logs and unpacked data for PixelSet events raised by the Place contract.
type PlacePixelSetIterator struct {
	Event *PlacePixelSet // Event containing the contract specifics and raw log

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
func (it *PlacePixelSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlacePixelSet)
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
		it.Event = new(PlacePixelSet)
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
func (it *PlacePixelSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlacePixelSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlacePixelSet represents a PixelSet event raised by the Place contract.
type PlacePixelSet struct {
	Setter common.Address
	X      int64
	Y      int64
	R      uint8
	G      uint8
	B      uint8
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPixelSet is a free log retrieval operation binding the contract event 0x80542baac7440772826923f7851805d2ba3530780c18eec5b278201321ba14c9.
//
// Solidity: event PixelSet(address indexed setter, int64 indexed x, int64 indexed y, uint8 r, uint8 g, uint8 b, uint256 value)
func (_Place *PlaceFilterer) FilterPixelSet(opts *bind.FilterOpts, setter []common.Address, x []int64, y []int64) (*PlacePixelSetIterator, error) {

	var setterRule []interface{}
	for _, setterItem := range setter {
		setterRule = append(setterRule, setterItem)
	}
	var xRule []interface{}
	for _, xItem := range x {
		xRule = append(xRule, xItem)
	}
	var yRule []interface{}
	for _, yItem := range y {
		yRule = append(yRule, yItem)
	}

	logs, sub, err := _Place.contract.FilterLogs(opts, "PixelSet", setterRule, xRule, yRule)
	if err != nil {
		return nil, err
	}
	return &PlacePixelSetIterator{contract: _Place.contract, event: "PixelSet", logs: logs, sub: sub}, nil
}

// WatchPixelSet is a free log subscription operation binding the contract event 0x80542baac7440772826923f7851805d2ba3530780c18eec5b278201321ba14c9.
//
// Solidity: event PixelSet(address indexed setter, int64 indexed x, int64 indexed y, uint8 r, uint8 g, uint8 b, uint256 value)
func (_Place *PlaceFilterer) WatchPixelSet(opts *bind.WatchOpts, sink chan<- *PlacePixelSet, setter []common.Address, x []int64, y []int64) (event.Subscription, error) {

	var setterRule []interface{}
	for _, setterItem := range setter {
		setterRule = append(setterRule, setterItem)
	}
	var xRule []interface{}
	for _, xItem := range x {
		xRule = append(xRule, xItem)
	}
	var yRule []interface{}
	for _, yItem := range y {
		yRule = append(yRule, yItem)
	}

	logs, sub, err := _Place.contract.WatchLogs(opts, "PixelSet", setterRule, xRule, yRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlacePixelSet)
				if err := _Place.contract.UnpackLog(event, "PixelSet", log); err != nil {
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

// ParsePixelSet is a log parse operation binding the contract event 0x80542baac7440772826923f7851805d2ba3530780c18eec5b278201321ba14c9.
//
// Solidity: event PixelSet(address indexed setter, int64 indexed x, int64 indexed y, uint8 r, uint8 g, uint8 b, uint256 value)
func (_Place *PlaceFilterer) ParsePixelSet(log types.Log) (*PlacePixelSet, error) {
	event := new(PlacePixelSet)
	if err := _Place.contract.UnpackLog(event, "PixelSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlaceRightBoundExpansionIterator is returned from FilterRightBoundExpansion and is used to iterate over the raw logs and unpacked data for RightBoundExpansion events raised by the Place contract.
type PlaceRightBoundExpansionIterator struct {
	Event *PlaceRightBoundExpansion // Event containing the contract specifics and raw log

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
func (it *PlaceRightBoundExpansionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlaceRightBoundExpansion)
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
		it.Event = new(PlaceRightBoundExpansion)
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
func (it *PlaceRightBoundExpansionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlaceRightBoundExpansionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlaceRightBoundExpansion represents a RightBoundExpansion event raised by the Place contract.
type PlaceRightBoundExpansion struct {
	NewValue int32
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRightBoundExpansion is a free log retrieval operation binding the contract event 0x2bf747280d9e11ea908896900e2f615f78f484eaf0471a689f013190d17f8161.
//
// Solidity: event RightBoundExpansion(int32 _newValue)
func (_Place *PlaceFilterer) FilterRightBoundExpansion(opts *bind.FilterOpts) (*PlaceRightBoundExpansionIterator, error) {

	logs, sub, err := _Place.contract.FilterLogs(opts, "RightBoundExpansion")
	if err != nil {
		return nil, err
	}
	return &PlaceRightBoundExpansionIterator{contract: _Place.contract, event: "RightBoundExpansion", logs: logs, sub: sub}, nil
}

// WatchRightBoundExpansion is a free log subscription operation binding the contract event 0x2bf747280d9e11ea908896900e2f615f78f484eaf0471a689f013190d17f8161.
//
// Solidity: event RightBoundExpansion(int32 _newValue)
func (_Place *PlaceFilterer) WatchRightBoundExpansion(opts *bind.WatchOpts, sink chan<- *PlaceRightBoundExpansion) (event.Subscription, error) {

	logs, sub, err := _Place.contract.WatchLogs(opts, "RightBoundExpansion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlaceRightBoundExpansion)
				if err := _Place.contract.UnpackLog(event, "RightBoundExpansion", log); err != nil {
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

// ParseRightBoundExpansion is a log parse operation binding the contract event 0x2bf747280d9e11ea908896900e2f615f78f484eaf0471a689f013190d17f8161.
//
// Solidity: event RightBoundExpansion(int32 _newValue)
func (_Place *PlaceFilterer) ParseRightBoundExpansion(log types.Log) (*PlaceRightBoundExpansion, error) {
	event := new(PlaceRightBoundExpansion)
	if err := _Place.contract.UnpackLog(event, "RightBoundExpansion", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PlaceUpperBoundExpansionIterator is returned from FilterUpperBoundExpansion and is used to iterate over the raw logs and unpacked data for UpperBoundExpansion events raised by the Place contract.
type PlaceUpperBoundExpansionIterator struct {
	Event *PlaceUpperBoundExpansion // Event containing the contract specifics and raw log

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
func (it *PlaceUpperBoundExpansionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlaceUpperBoundExpansion)
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
		it.Event = new(PlaceUpperBoundExpansion)
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
func (it *PlaceUpperBoundExpansionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlaceUpperBoundExpansionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlaceUpperBoundExpansion represents a UpperBoundExpansion event raised by the Place contract.
type PlaceUpperBoundExpansion struct {
	NewValue int32
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterUpperBoundExpansion is a free log retrieval operation binding the contract event 0x73aa33d8c80e2781606757c4b5824419c0e7506b7654c5bff8b2daade69cf4ba.
//
// Solidity: event UpperBoundExpansion(int32 _newValue)
func (_Place *PlaceFilterer) FilterUpperBoundExpansion(opts *bind.FilterOpts) (*PlaceUpperBoundExpansionIterator, error) {

	logs, sub, err := _Place.contract.FilterLogs(opts, "UpperBoundExpansion")
	if err != nil {
		return nil, err
	}
	return &PlaceUpperBoundExpansionIterator{contract: _Place.contract, event: "UpperBoundExpansion", logs: logs, sub: sub}, nil
}

// WatchUpperBoundExpansion is a free log subscription operation binding the contract event 0x73aa33d8c80e2781606757c4b5824419c0e7506b7654c5bff8b2daade69cf4ba.
//
// Solidity: event UpperBoundExpansion(int32 _newValue)
func (_Place *PlaceFilterer) WatchUpperBoundExpansion(opts *bind.WatchOpts, sink chan<- *PlaceUpperBoundExpansion) (event.Subscription, error) {

	logs, sub, err := _Place.contract.WatchLogs(opts, "UpperBoundExpansion")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlaceUpperBoundExpansion)
				if err := _Place.contract.UnpackLog(event, "UpperBoundExpansion", log); err != nil {
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

// ParseUpperBoundExpansion is a log parse operation binding the contract event 0x73aa33d8c80e2781606757c4b5824419c0e7506b7654c5bff8b2daade69cf4ba.
//
// Solidity: event UpperBoundExpansion(int32 _newValue)
func (_Place *PlaceFilterer) ParseUpperBoundExpansion(log types.Log) (*PlaceUpperBoundExpansion, error) {
	event := new(PlaceUpperBoundExpansion)
	if err := _Place.contract.UnpackLog(event, "UpperBoundExpansion", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

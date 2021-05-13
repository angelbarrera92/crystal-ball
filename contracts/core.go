// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// IOrakuruCoreRequest is an auto generated low-level Go binding around an user-defined struct.
type IOrakuruCoreRequest struct {
	Id                 [32]byte
	DataSource         string
	Selector           string
	CallbackAddr       common.Address
	ExecutionTimestamp *big.Int
	IsFulfilled        bool
	AggrType           uint8
	Precision          uint8
}

// IOrakuruCoreResponse is an auto generated low-level Go binding around an user-defined struct.
type IOrakuruCoreResponse struct {
	Id          [32]byte
	RequestId   [32]byte
	Result      []byte
	SubmittedBy common.Address
	SubmittedAt *big.Int
}

// IOrakuruCoreABI is the input ABI used to generate the binding from.
const IOrakuruCoreABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"}],\"name\":\"Canceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"name\":\"Fulfilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dataSource\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"selector\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"callbackAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumIOrakuruCore.Type\",\"name\":\"aggrType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"precision\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executionTimestamp\",\"type\":\"uint256\"}],\"name\":\"Requested\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"submittedResult\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"parsedResult\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"Submitted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"}],\"name\":\"cancelRequest\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"}],\"name\":\"fulfillRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addr\",\"type\":\"address\"}],\"name\":\"getNonceFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPendingRequests\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"}],\"name\":\"getResponses\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"requestId\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"submittedBy\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"submittedAt\",\"type\":\"uint256\"}],\"internalType\":\"structIOrakuruCore.Response[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"}],\"name\":\"getResultsBytes\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"}],\"name\":\"getResultsUint\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_dataSource\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_selector\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_calldataAddr\",\"type\":\"address\"},{\"internalType\":\"enumIOrakuruCore.Type\",\"name\":\"_aggrType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_precision\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_executionTimestamp\",\"type\":\"uint256\"}],\"name\":\"makeRequest\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"requests\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"dataSource\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"selector\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"callbackAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"executionTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isFulfilled\",\"type\":\"bool\"},{\"internalType\":\"enumIOrakuruCore.Type\",\"name\":\"aggrType\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"precision\",\"type\":\"uint8\"}],\"internalType\":\"structIOrakuruCore.Request\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestId\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_result\",\"type\":\"string\"}],\"name\":\"submitResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IOrakuruCoreFuncSigs maps the 4-byte function signature to its string representation.
var IOrakuruCoreFuncSigs = map[string]string{
	"50125546": "cancelRequest(bytes32)",
	"432d0137": "fulfillRequest(bytes32)",
	"bf3125c7": "getNonceFor(address)",
	"80a1f712": "getPendingRequests()",
	"d4b8df22": "getResponses(bytes32)",
	"85bf2c80": "getResultsBytes(bytes32)",
	"cc19fa97": "getResultsUint(bytes32)",
	"588f7a2c": "makeRequest(string,string,address,uint8,uint8,uint256)",
	"9d866985": "requests(bytes32)",
	"cd824ed6": "submitResult(bytes32,string)",
}

// IOrakuruCore is an auto generated Go binding around an Ethereum contract.
type IOrakuruCore struct {
	IOrakuruCoreCaller     // Read-only binding to the contract
	IOrakuruCoreTransactor // Write-only binding to the contract
	IOrakuruCoreFilterer   // Log filterer for contract events
}

// IOrakuruCoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type IOrakuruCoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOrakuruCoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IOrakuruCoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOrakuruCoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IOrakuruCoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IOrakuruCoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IOrakuruCoreSession struct {
	Contract     *IOrakuruCore     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IOrakuruCoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IOrakuruCoreCallerSession struct {
	Contract *IOrakuruCoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IOrakuruCoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IOrakuruCoreTransactorSession struct {
	Contract     *IOrakuruCoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IOrakuruCoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type IOrakuruCoreRaw struct {
	Contract *IOrakuruCore // Generic contract binding to access the raw methods on
}

// IOrakuruCoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IOrakuruCoreCallerRaw struct {
	Contract *IOrakuruCoreCaller // Generic read-only contract binding to access the raw methods on
}

// IOrakuruCoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IOrakuruCoreTransactorRaw struct {
	Contract *IOrakuruCoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIOrakuruCore creates a new instance of IOrakuruCore, bound to a specific deployed contract.
func NewIOrakuruCore(address common.Address, backend bind.ContractBackend) (*IOrakuruCore, error) {
	contract, err := bindIOrakuruCore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IOrakuruCore{IOrakuruCoreCaller: IOrakuruCoreCaller{contract: contract}, IOrakuruCoreTransactor: IOrakuruCoreTransactor{contract: contract}, IOrakuruCoreFilterer: IOrakuruCoreFilterer{contract: contract}}, nil
}

// NewIOrakuruCoreCaller creates a new read-only instance of IOrakuruCore, bound to a specific deployed contract.
func NewIOrakuruCoreCaller(address common.Address, caller bind.ContractCaller) (*IOrakuruCoreCaller, error) {
	contract, err := bindIOrakuruCore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IOrakuruCoreCaller{contract: contract}, nil
}

// NewIOrakuruCoreTransactor creates a new write-only instance of IOrakuruCore, bound to a specific deployed contract.
func NewIOrakuruCoreTransactor(address common.Address, transactor bind.ContractTransactor) (*IOrakuruCoreTransactor, error) {
	contract, err := bindIOrakuruCore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IOrakuruCoreTransactor{contract: contract}, nil
}

// NewIOrakuruCoreFilterer creates a new log filterer instance of IOrakuruCore, bound to a specific deployed contract.
func NewIOrakuruCoreFilterer(address common.Address, filterer bind.ContractFilterer) (*IOrakuruCoreFilterer, error) {
	contract, err := bindIOrakuruCore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IOrakuruCoreFilterer{contract: contract}, nil
}

// bindIOrakuruCore binds a generic wrapper to an already deployed contract.
func bindIOrakuruCore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IOrakuruCoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOrakuruCore *IOrakuruCoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOrakuruCore.Contract.IOrakuruCoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOrakuruCore *IOrakuruCoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOrakuruCore.Contract.IOrakuruCoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOrakuruCore *IOrakuruCoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOrakuruCore.Contract.IOrakuruCoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IOrakuruCore *IOrakuruCoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IOrakuruCore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IOrakuruCore *IOrakuruCoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IOrakuruCore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IOrakuruCore *IOrakuruCoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IOrakuruCore.Contract.contract.Transact(opts, method, params...)
}

// GetNonceFor is a free data retrieval call binding the contract method 0xbf3125c7.
//
// Solidity: function getNonceFor(address _addr) view returns(uint256)
func (_IOrakuruCore *IOrakuruCoreCaller) GetNonceFor(opts *bind.CallOpts, _addr common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IOrakuruCore.contract.Call(opts, &out, "getNonceFor", _addr)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonceFor is a free data retrieval call binding the contract method 0xbf3125c7.
//
// Solidity: function getNonceFor(address _addr) view returns(uint256)
func (_IOrakuruCore *IOrakuruCoreSession) GetNonceFor(_addr common.Address) (*big.Int, error) {
	return _IOrakuruCore.Contract.GetNonceFor(&_IOrakuruCore.CallOpts, _addr)
}

// GetNonceFor is a free data retrieval call binding the contract method 0xbf3125c7.
//
// Solidity: function getNonceFor(address _addr) view returns(uint256)
func (_IOrakuruCore *IOrakuruCoreCallerSession) GetNonceFor(_addr common.Address) (*big.Int, error) {
	return _IOrakuruCore.Contract.GetNonceFor(&_IOrakuruCore.CallOpts, _addr)
}

// GetPendingRequests is a free data retrieval call binding the contract method 0x80a1f712.
//
// Solidity: function getPendingRequests() view returns(bytes32[])
func (_IOrakuruCore *IOrakuruCoreCaller) GetPendingRequests(opts *bind.CallOpts) ([][32]byte, error) {
	var out []interface{}
	err := _IOrakuruCore.contract.Call(opts, &out, "getPendingRequests")

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetPendingRequests is a free data retrieval call binding the contract method 0x80a1f712.
//
// Solidity: function getPendingRequests() view returns(bytes32[])
func (_IOrakuruCore *IOrakuruCoreSession) GetPendingRequests() ([][32]byte, error) {
	return _IOrakuruCore.Contract.GetPendingRequests(&_IOrakuruCore.CallOpts)
}

// GetPendingRequests is a free data retrieval call binding the contract method 0x80a1f712.
//
// Solidity: function getPendingRequests() view returns(bytes32[])
func (_IOrakuruCore *IOrakuruCoreCallerSession) GetPendingRequests() ([][32]byte, error) {
	return _IOrakuruCore.Contract.GetPendingRequests(&_IOrakuruCore.CallOpts)
}

// GetResponses is a free data retrieval call binding the contract method 0xd4b8df22.
//
// Solidity: function getResponses(bytes32 _requestId) view returns((bytes32,bytes32,bytes,address,uint256)[])
func (_IOrakuruCore *IOrakuruCoreCaller) GetResponses(opts *bind.CallOpts, _requestId [32]byte) ([]IOrakuruCoreResponse, error) {
	var out []interface{}
	err := _IOrakuruCore.contract.Call(opts, &out, "getResponses", _requestId)

	if err != nil {
		return *new([]IOrakuruCoreResponse), err
	}

	out0 := *abi.ConvertType(out[0], new([]IOrakuruCoreResponse)).(*[]IOrakuruCoreResponse)

	return out0, err

}

// GetResponses is a free data retrieval call binding the contract method 0xd4b8df22.
//
// Solidity: function getResponses(bytes32 _requestId) view returns((bytes32,bytes32,bytes,address,uint256)[])
func (_IOrakuruCore *IOrakuruCoreSession) GetResponses(_requestId [32]byte) ([]IOrakuruCoreResponse, error) {
	return _IOrakuruCore.Contract.GetResponses(&_IOrakuruCore.CallOpts, _requestId)
}

// GetResponses is a free data retrieval call binding the contract method 0xd4b8df22.
//
// Solidity: function getResponses(bytes32 _requestId) view returns((bytes32,bytes32,bytes,address,uint256)[])
func (_IOrakuruCore *IOrakuruCoreCallerSession) GetResponses(_requestId [32]byte) ([]IOrakuruCoreResponse, error) {
	return _IOrakuruCore.Contract.GetResponses(&_IOrakuruCore.CallOpts, _requestId)
}

// GetResultsBytes is a free data retrieval call binding the contract method 0x85bf2c80.
//
// Solidity: function getResultsBytes(bytes32 _requestId) view returns(bytes[])
func (_IOrakuruCore *IOrakuruCoreCaller) GetResultsBytes(opts *bind.CallOpts, _requestId [32]byte) ([][]byte, error) {
	var out []interface{}
	err := _IOrakuruCore.contract.Call(opts, &out, "getResultsBytes", _requestId)

	if err != nil {
		return *new([][]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][]byte)).(*[][]byte)

	return out0, err

}

// GetResultsBytes is a free data retrieval call binding the contract method 0x85bf2c80.
//
// Solidity: function getResultsBytes(bytes32 _requestId) view returns(bytes[])
func (_IOrakuruCore *IOrakuruCoreSession) GetResultsBytes(_requestId [32]byte) ([][]byte, error) {
	return _IOrakuruCore.Contract.GetResultsBytes(&_IOrakuruCore.CallOpts, _requestId)
}

// GetResultsBytes is a free data retrieval call binding the contract method 0x85bf2c80.
//
// Solidity: function getResultsBytes(bytes32 _requestId) view returns(bytes[])
func (_IOrakuruCore *IOrakuruCoreCallerSession) GetResultsBytes(_requestId [32]byte) ([][]byte, error) {
	return _IOrakuruCore.Contract.GetResultsBytes(&_IOrakuruCore.CallOpts, _requestId)
}

// GetResultsUint is a free data retrieval call binding the contract method 0xcc19fa97.
//
// Solidity: function getResultsUint(bytes32 _requestId) view returns(uint256[])
func (_IOrakuruCore *IOrakuruCoreCaller) GetResultsUint(opts *bind.CallOpts, _requestId [32]byte) ([]*big.Int, error) {
	var out []interface{}
	err := _IOrakuruCore.contract.Call(opts, &out, "getResultsUint", _requestId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetResultsUint is a free data retrieval call binding the contract method 0xcc19fa97.
//
// Solidity: function getResultsUint(bytes32 _requestId) view returns(uint256[])
func (_IOrakuruCore *IOrakuruCoreSession) GetResultsUint(_requestId [32]byte) ([]*big.Int, error) {
	return _IOrakuruCore.Contract.GetResultsUint(&_IOrakuruCore.CallOpts, _requestId)
}

// GetResultsUint is a free data retrieval call binding the contract method 0xcc19fa97.
//
// Solidity: function getResultsUint(bytes32 _requestId) view returns(uint256[])
func (_IOrakuruCore *IOrakuruCoreCallerSession) GetResultsUint(_requestId [32]byte) ([]*big.Int, error) {
	return _IOrakuruCore.Contract.GetResultsUint(&_IOrakuruCore.CallOpts, _requestId)
}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 id) view returns((bytes32,string,string,address,uint256,bool,uint8,uint8))
func (_IOrakuruCore *IOrakuruCoreCaller) Requests(opts *bind.CallOpts, id [32]byte) (IOrakuruCoreRequest, error) {
	var out []interface{}
	err := _IOrakuruCore.contract.Call(opts, &out, "requests", id)

	if err != nil {
		return *new(IOrakuruCoreRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(IOrakuruCoreRequest)).(*IOrakuruCoreRequest)

	return out0, err

}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 id) view returns((bytes32,string,string,address,uint256,bool,uint8,uint8))
func (_IOrakuruCore *IOrakuruCoreSession) Requests(id [32]byte) (IOrakuruCoreRequest, error) {
	return _IOrakuruCore.Contract.Requests(&_IOrakuruCore.CallOpts, id)
}

// Requests is a free data retrieval call binding the contract method 0x9d866985.
//
// Solidity: function requests(bytes32 id) view returns((bytes32,string,string,address,uint256,bool,uint8,uint8))
func (_IOrakuruCore *IOrakuruCoreCallerSession) Requests(id [32]byte) (IOrakuruCoreRequest, error) {
	return _IOrakuruCore.Contract.Requests(&_IOrakuruCore.CallOpts, id)
}

// CancelRequest is a paid mutator transaction binding the contract method 0x50125546.
//
// Solidity: function cancelRequest(bytes32 _requestId) returns(bool)
func (_IOrakuruCore *IOrakuruCoreTransactor) CancelRequest(opts *bind.TransactOpts, _requestId [32]byte) (*types.Transaction, error) {
	return _IOrakuruCore.contract.Transact(opts, "cancelRequest", _requestId)
}

// CancelRequest is a paid mutator transaction binding the contract method 0x50125546.
//
// Solidity: function cancelRequest(bytes32 _requestId) returns(bool)
func (_IOrakuruCore *IOrakuruCoreSession) CancelRequest(_requestId [32]byte) (*types.Transaction, error) {
	return _IOrakuruCore.Contract.CancelRequest(&_IOrakuruCore.TransactOpts, _requestId)
}

// CancelRequest is a paid mutator transaction binding the contract method 0x50125546.
//
// Solidity: function cancelRequest(bytes32 _requestId) returns(bool)
func (_IOrakuruCore *IOrakuruCoreTransactorSession) CancelRequest(_requestId [32]byte) (*types.Transaction, error) {
	return _IOrakuruCore.Contract.CancelRequest(&_IOrakuruCore.TransactOpts, _requestId)
}

// FulfillRequest is a paid mutator transaction binding the contract method 0x432d0137.
//
// Solidity: function fulfillRequest(bytes32 _requestId) returns()
func (_IOrakuruCore *IOrakuruCoreTransactor) FulfillRequest(opts *bind.TransactOpts, _requestId [32]byte) (*types.Transaction, error) {
	return _IOrakuruCore.contract.Transact(opts, "fulfillRequest", _requestId)
}

// FulfillRequest is a paid mutator transaction binding the contract method 0x432d0137.
//
// Solidity: function fulfillRequest(bytes32 _requestId) returns()
func (_IOrakuruCore *IOrakuruCoreSession) FulfillRequest(_requestId [32]byte) (*types.Transaction, error) {
	return _IOrakuruCore.Contract.FulfillRequest(&_IOrakuruCore.TransactOpts, _requestId)
}

// FulfillRequest is a paid mutator transaction binding the contract method 0x432d0137.
//
// Solidity: function fulfillRequest(bytes32 _requestId) returns()
func (_IOrakuruCore *IOrakuruCoreTransactorSession) FulfillRequest(_requestId [32]byte) (*types.Transaction, error) {
	return _IOrakuruCore.Contract.FulfillRequest(&_IOrakuruCore.TransactOpts, _requestId)
}

// MakeRequest is a paid mutator transaction binding the contract method 0x588f7a2c.
//
// Solidity: function makeRequest(string _dataSource, string _selector, address _calldataAddr, uint8 _aggrType, uint8 _precision, uint256 _executionTimestamp) returns(bytes32)
func (_IOrakuruCore *IOrakuruCoreTransactor) MakeRequest(opts *bind.TransactOpts, _dataSource string, _selector string, _calldataAddr common.Address, _aggrType uint8, _precision uint8, _executionTimestamp *big.Int) (*types.Transaction, error) {
	return _IOrakuruCore.contract.Transact(opts, "makeRequest", _dataSource, _selector, _calldataAddr, _aggrType, _precision, _executionTimestamp)
}

// MakeRequest is a paid mutator transaction binding the contract method 0x588f7a2c.
//
// Solidity: function makeRequest(string _dataSource, string _selector, address _calldataAddr, uint8 _aggrType, uint8 _precision, uint256 _executionTimestamp) returns(bytes32)
func (_IOrakuruCore *IOrakuruCoreSession) MakeRequest(_dataSource string, _selector string, _calldataAddr common.Address, _aggrType uint8, _precision uint8, _executionTimestamp *big.Int) (*types.Transaction, error) {
	return _IOrakuruCore.Contract.MakeRequest(&_IOrakuruCore.TransactOpts, _dataSource, _selector, _calldataAddr, _aggrType, _precision, _executionTimestamp)
}

// MakeRequest is a paid mutator transaction binding the contract method 0x588f7a2c.
//
// Solidity: function makeRequest(string _dataSource, string _selector, address _calldataAddr, uint8 _aggrType, uint8 _precision, uint256 _executionTimestamp) returns(bytes32)
func (_IOrakuruCore *IOrakuruCoreTransactorSession) MakeRequest(_dataSource string, _selector string, _calldataAddr common.Address, _aggrType uint8, _precision uint8, _executionTimestamp *big.Int) (*types.Transaction, error) {
	return _IOrakuruCore.Contract.MakeRequest(&_IOrakuruCore.TransactOpts, _dataSource, _selector, _calldataAddr, _aggrType, _precision, _executionTimestamp)
}

// SubmitResult is a paid mutator transaction binding the contract method 0xcd824ed6.
//
// Solidity: function submitResult(bytes32 _requestId, string _result) returns()
func (_IOrakuruCore *IOrakuruCoreTransactor) SubmitResult(opts *bind.TransactOpts, _requestId [32]byte, _result string) (*types.Transaction, error) {
	return _IOrakuruCore.contract.Transact(opts, "submitResult", _requestId, _result)
}

// SubmitResult is a paid mutator transaction binding the contract method 0xcd824ed6.
//
// Solidity: function submitResult(bytes32 _requestId, string _result) returns()
func (_IOrakuruCore *IOrakuruCoreSession) SubmitResult(_requestId [32]byte, _result string) (*types.Transaction, error) {
	return _IOrakuruCore.Contract.SubmitResult(&_IOrakuruCore.TransactOpts, _requestId, _result)
}

// SubmitResult is a paid mutator transaction binding the contract method 0xcd824ed6.
//
// Solidity: function submitResult(bytes32 _requestId, string _result) returns()
func (_IOrakuruCore *IOrakuruCoreTransactorSession) SubmitResult(_requestId [32]byte, _result string) (*types.Transaction, error) {
	return _IOrakuruCore.Contract.SubmitResult(&_IOrakuruCore.TransactOpts, _requestId, _result)
}

// IOrakuruCoreCanceledIterator is returned from FilterCanceled and is used to iterate over the raw logs and unpacked data for Canceled events raised by the IOrakuruCore contract.
type IOrakuruCoreCanceledIterator struct {
	Event *IOrakuruCoreCanceled // Event containing the contract specifics and raw log

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
func (it *IOrakuruCoreCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOrakuruCoreCanceled)
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
		it.Event = new(IOrakuruCoreCanceled)
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
func (it *IOrakuruCoreCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOrakuruCoreCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOrakuruCoreCanceled represents a Canceled event raised by the IOrakuruCore contract.
type IOrakuruCoreCanceled struct {
	RequestId [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCanceled is a free log retrieval operation binding the contract event 0x134fdd648feeaf30251f0157f9624ef8608ff9a042aad6d13e73f35d21d3f88d.
//
// Solidity: event Canceled(bytes32 indexed requestId)
func (_IOrakuruCore *IOrakuruCoreFilterer) FilterCanceled(opts *bind.FilterOpts, requestId [][32]byte) (*IOrakuruCoreCanceledIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _IOrakuruCore.contract.FilterLogs(opts, "Canceled", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &IOrakuruCoreCanceledIterator{contract: _IOrakuruCore.contract, event: "Canceled", logs: logs, sub: sub}, nil
}

// WatchCanceled is a free log subscription operation binding the contract event 0x134fdd648feeaf30251f0157f9624ef8608ff9a042aad6d13e73f35d21d3f88d.
//
// Solidity: event Canceled(bytes32 indexed requestId)
func (_IOrakuruCore *IOrakuruCoreFilterer) WatchCanceled(opts *bind.WatchOpts, sink chan<- *IOrakuruCoreCanceled, requestId [][32]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _IOrakuruCore.contract.WatchLogs(opts, "Canceled", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOrakuruCoreCanceled)
				if err := _IOrakuruCore.contract.UnpackLog(event, "Canceled", log); err != nil {
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

// ParseCanceled is a log parse operation binding the contract event 0x134fdd648feeaf30251f0157f9624ef8608ff9a042aad6d13e73f35d21d3f88d.
//
// Solidity: event Canceled(bytes32 indexed requestId)
func (_IOrakuruCore *IOrakuruCoreFilterer) ParseCanceled(log types.Log) (*IOrakuruCoreCanceled, error) {
	event := new(IOrakuruCoreCanceled)
	if err := _IOrakuruCore.contract.UnpackLog(event, "Canceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOrakuruCoreFulfilledIterator is returned from FilterFulfilled and is used to iterate over the raw logs and unpacked data for Fulfilled events raised by the IOrakuruCore contract.
type IOrakuruCoreFulfilledIterator struct {
	Event *IOrakuruCoreFulfilled // Event containing the contract specifics and raw log

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
func (it *IOrakuruCoreFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOrakuruCoreFulfilled)
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
		it.Event = new(IOrakuruCoreFulfilled)
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
func (it *IOrakuruCoreFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOrakuruCoreFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOrakuruCoreFulfilled represents a Fulfilled event raised by the IOrakuruCore contract.
type IOrakuruCoreFulfilled struct {
	RequestId [32]byte
	Result    []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFulfilled is a free log retrieval operation binding the contract event 0x2f32b2ead759413df4558cb416e9b67b02081989ace899bb231416daae106b45.
//
// Solidity: event Fulfilled(bytes32 indexed requestId, bytes result)
func (_IOrakuruCore *IOrakuruCoreFilterer) FilterFulfilled(opts *bind.FilterOpts, requestId [][32]byte) (*IOrakuruCoreFulfilledIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _IOrakuruCore.contract.FilterLogs(opts, "Fulfilled", requestIdRule)
	if err != nil {
		return nil, err
	}
	return &IOrakuruCoreFulfilledIterator{contract: _IOrakuruCore.contract, event: "Fulfilled", logs: logs, sub: sub}, nil
}

// WatchFulfilled is a free log subscription operation binding the contract event 0x2f32b2ead759413df4558cb416e9b67b02081989ace899bb231416daae106b45.
//
// Solidity: event Fulfilled(bytes32 indexed requestId, bytes result)
func (_IOrakuruCore *IOrakuruCoreFilterer) WatchFulfilled(opts *bind.WatchOpts, sink chan<- *IOrakuruCoreFulfilled, requestId [][32]byte) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	logs, sub, err := _IOrakuruCore.contract.WatchLogs(opts, "Fulfilled", requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOrakuruCoreFulfilled)
				if err := _IOrakuruCore.contract.UnpackLog(event, "Fulfilled", log); err != nil {
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

// ParseFulfilled is a log parse operation binding the contract event 0x2f32b2ead759413df4558cb416e9b67b02081989ace899bb231416daae106b45.
//
// Solidity: event Fulfilled(bytes32 indexed requestId, bytes result)
func (_IOrakuruCore *IOrakuruCoreFilterer) ParseFulfilled(log types.Log) (*IOrakuruCoreFulfilled, error) {
	event := new(IOrakuruCoreFulfilled)
	if err := _IOrakuruCore.contract.UnpackLog(event, "Fulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOrakuruCoreRequestedIterator is returned from FilterRequested and is used to iterate over the raw logs and unpacked data for Requested events raised by the IOrakuruCore contract.
type IOrakuruCoreRequestedIterator struct {
	Event *IOrakuruCoreRequested // Event containing the contract specifics and raw log

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
func (it *IOrakuruCoreRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOrakuruCoreRequested)
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
		it.Event = new(IOrakuruCoreRequested)
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
func (it *IOrakuruCoreRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOrakuruCoreRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOrakuruCoreRequested represents a Requested event raised by the IOrakuruCore contract.
type IOrakuruCoreRequested struct {
	RequestId          [32]byte
	DataSource         string
	Selector           string
	CallbackAddr       common.Address
	AggrType           uint8
	Precision          uint8
	ExecutionTimestamp *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRequested is a free log retrieval operation binding the contract event 0x80fb6ec429779cbda9a8b6a957fbfc7607f5b46a63955bab35c33fc97fbb6687.
//
// Solidity: event Requested(bytes32 indexed requestId, string dataSource, string selector, address indexed callbackAddr, uint8 aggrType, uint8 precision, uint256 executionTimestamp)
func (_IOrakuruCore *IOrakuruCoreFilterer) FilterRequested(opts *bind.FilterOpts, requestId [][32]byte, callbackAddr []common.Address) (*IOrakuruCoreRequestedIterator, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	var callbackAddrRule []interface{}
	for _, callbackAddrItem := range callbackAddr {
		callbackAddrRule = append(callbackAddrRule, callbackAddrItem)
	}

	logs, sub, err := _IOrakuruCore.contract.FilterLogs(opts, "Requested", requestIdRule, callbackAddrRule)
	if err != nil {
		return nil, err
	}
	return &IOrakuruCoreRequestedIterator{contract: _IOrakuruCore.contract, event: "Requested", logs: logs, sub: sub}, nil
}

// WatchRequested is a free log subscription operation binding the contract event 0x80fb6ec429779cbda9a8b6a957fbfc7607f5b46a63955bab35c33fc97fbb6687.
//
// Solidity: event Requested(bytes32 indexed requestId, string dataSource, string selector, address indexed callbackAddr, uint8 aggrType, uint8 precision, uint256 executionTimestamp)
func (_IOrakuruCore *IOrakuruCoreFilterer) WatchRequested(opts *bind.WatchOpts, sink chan<- *IOrakuruCoreRequested, requestId [][32]byte, callbackAddr []common.Address) (event.Subscription, error) {

	var requestIdRule []interface{}
	for _, requestIdItem := range requestId {
		requestIdRule = append(requestIdRule, requestIdItem)
	}

	var callbackAddrRule []interface{}
	for _, callbackAddrItem := range callbackAddr {
		callbackAddrRule = append(callbackAddrRule, callbackAddrItem)
	}

	logs, sub, err := _IOrakuruCore.contract.WatchLogs(opts, "Requested", requestIdRule, callbackAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOrakuruCoreRequested)
				if err := _IOrakuruCore.contract.UnpackLog(event, "Requested", log); err != nil {
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

// ParseRequested is a log parse operation binding the contract event 0x80fb6ec429779cbda9a8b6a957fbfc7607f5b46a63955bab35c33fc97fbb6687.
//
// Solidity: event Requested(bytes32 indexed requestId, string dataSource, string selector, address indexed callbackAddr, uint8 aggrType, uint8 precision, uint256 executionTimestamp)
func (_IOrakuruCore *IOrakuruCoreFilterer) ParseRequested(log types.Log) (*IOrakuruCoreRequested, error) {
	event := new(IOrakuruCoreRequested)
	if err := _IOrakuruCore.contract.UnpackLog(event, "Requested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IOrakuruCoreSubmittedIterator is returned from FilterSubmitted and is used to iterate over the raw logs and unpacked data for Submitted events raised by the IOrakuruCore contract.
type IOrakuruCoreSubmittedIterator struct {
	Event *IOrakuruCoreSubmitted // Event containing the contract specifics and raw log

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
func (it *IOrakuruCoreSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IOrakuruCoreSubmitted)
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
		it.Event = new(IOrakuruCoreSubmitted)
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
func (it *IOrakuruCoreSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IOrakuruCoreSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IOrakuruCoreSubmitted represents a Submitted event raised by the IOrakuruCore contract.
type IOrakuruCoreSubmitted struct {
	RequestId       [32]byte
	SubmittedResult string
	ParsedResult    []byte
	Oracle          common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterSubmitted is a free log retrieval operation binding the contract event 0x87a989dac6bc925fc5a5656f502f963cdb1f2ad71370ac490d9ee0fbff837f4b.
//
// Solidity: event Submitted(bytes32 requestId, string submittedResult, bytes parsedResult, address oracle)
func (_IOrakuruCore *IOrakuruCoreFilterer) FilterSubmitted(opts *bind.FilterOpts) (*IOrakuruCoreSubmittedIterator, error) {

	logs, sub, err := _IOrakuruCore.contract.FilterLogs(opts, "Submitted")
	if err != nil {
		return nil, err
	}
	return &IOrakuruCoreSubmittedIterator{contract: _IOrakuruCore.contract, event: "Submitted", logs: logs, sub: sub}, nil
}

// WatchSubmitted is a free log subscription operation binding the contract event 0x87a989dac6bc925fc5a5656f502f963cdb1f2ad71370ac490d9ee0fbff837f4b.
//
// Solidity: event Submitted(bytes32 requestId, string submittedResult, bytes parsedResult, address oracle)
func (_IOrakuruCore *IOrakuruCoreFilterer) WatchSubmitted(opts *bind.WatchOpts, sink chan<- *IOrakuruCoreSubmitted) (event.Subscription, error) {

	logs, sub, err := _IOrakuruCore.contract.WatchLogs(opts, "Submitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IOrakuruCoreSubmitted)
				if err := _IOrakuruCore.contract.UnpackLog(event, "Submitted", log); err != nil {
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

// ParseSubmitted is a log parse operation binding the contract event 0x87a989dac6bc925fc5a5656f502f963cdb1f2ad71370ac490d9ee0fbff837f4b.
//
// Solidity: event Submitted(bytes32 requestId, string submittedResult, bytes parsedResult, address oracle)
func (_IOrakuruCore *IOrakuruCoreFilterer) ParseSubmitted(log types.Log) (*IOrakuruCoreSubmitted, error) {
	event := new(IOrakuruCoreSubmitted)
	if err := _IOrakuruCore.contract.UnpackLog(event, "Submitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

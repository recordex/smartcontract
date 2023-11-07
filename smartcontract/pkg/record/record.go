// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package record

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

// RecordMetaData contains all meta data concerning the Record contract.
var RecordMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"fileHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"createdBy\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"}],\"name\":\"FileAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"deletedBy\",\"type\":\"address\"}],\"name\":\"FileDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DELETE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"READ_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WRITE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"fileHash\",\"type\":\"bytes32\"}],\"name\":\"addFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"fileName\",\"type\":\"string\"}],\"name\":\"deleteFile\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"files\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"created_by\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"created_at\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// RecordABI is the input ABI used to generate the binding from.
// Deprecated: Use RecordMetaData.ABI instead.
var RecordABI = RecordMetaData.ABI

// Record is an auto generated Go binding around an Ethereum contract.
type Record struct {
	RecordCaller     // Read-only binding to the contract
	RecordTransactor // Write-only binding to the contract
	RecordFilterer   // Log filterer for contract events
}

// RecordCaller is an auto generated read-only Go binding around an Ethereum contract.
type RecordCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecordTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RecordTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecordFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RecordFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RecordSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RecordSession struct {
	Contract     *Record           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RecordCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RecordCallerSession struct {
	Contract *RecordCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RecordTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RecordTransactorSession struct {
	Contract     *RecordTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RecordRaw is an auto generated low-level Go binding around an Ethereum contract.
type RecordRaw struct {
	Contract *Record // Generic contract binding to access the raw methods on
}

// RecordCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RecordCallerRaw struct {
	Contract *RecordCaller // Generic read-only contract binding to access the raw methods on
}

// RecordTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RecordTransactorRaw struct {
	Contract *RecordTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRecord creates a new instance of Record, bound to a specific deployed contract.
func NewRecord(address common.Address, backend bind.ContractBackend) (*Record, error) {
	contract, err := bindRecord(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Record{RecordCaller: RecordCaller{contract: contract}, RecordTransactor: RecordTransactor{contract: contract}, RecordFilterer: RecordFilterer{contract: contract}}, nil
}

// NewRecordCaller creates a new read-only instance of Record, bound to a specific deployed contract.
func NewRecordCaller(address common.Address, caller bind.ContractCaller) (*RecordCaller, error) {
	contract, err := bindRecord(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RecordCaller{contract: contract}, nil
}

// NewRecordTransactor creates a new write-only instance of Record, bound to a specific deployed contract.
func NewRecordTransactor(address common.Address, transactor bind.ContractTransactor) (*RecordTransactor, error) {
	contract, err := bindRecord(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RecordTransactor{contract: contract}, nil
}

// NewRecordFilterer creates a new log filterer instance of Record, bound to a specific deployed contract.
func NewRecordFilterer(address common.Address, filterer bind.ContractFilterer) (*RecordFilterer, error) {
	contract, err := bindRecord(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RecordFilterer{contract: contract}, nil
}

// bindRecord binds a generic wrapper to an already deployed contract.
func bindRecord(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RecordMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Record *RecordRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Record.Contract.RecordCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Record *RecordRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Record.Contract.RecordTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Record *RecordRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Record.Contract.RecordTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Record *RecordCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Record.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Record *RecordTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Record.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Record *RecordTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Record.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Record *RecordCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Record.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Record *RecordSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Record.Contract.DEFAULTADMINROLE(&_Record.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Record *RecordCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Record.Contract.DEFAULTADMINROLE(&_Record.CallOpts)
}

// DELETEROLE is a free data retrieval call binding the contract method 0xa88c8335.
//
// Solidity: function DELETE_ROLE() view returns(bytes32)
func (_Record *RecordCaller) DELETEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Record.contract.Call(opts, &out, "DELETE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DELETEROLE is a free data retrieval call binding the contract method 0xa88c8335.
//
// Solidity: function DELETE_ROLE() view returns(bytes32)
func (_Record *RecordSession) DELETEROLE() ([32]byte, error) {
	return _Record.Contract.DELETEROLE(&_Record.CallOpts)
}

// DELETEROLE is a free data retrieval call binding the contract method 0xa88c8335.
//
// Solidity: function DELETE_ROLE() view returns(bytes32)
func (_Record *RecordCallerSession) DELETEROLE() ([32]byte, error) {
	return _Record.Contract.DELETEROLE(&_Record.CallOpts)
}

// READROLE is a free data retrieval call binding the contract method 0xaf15256a.
//
// Solidity: function READ_ROLE() view returns(bytes32)
func (_Record *RecordCaller) READROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Record.contract.Call(opts, &out, "READ_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// READROLE is a free data retrieval call binding the contract method 0xaf15256a.
//
// Solidity: function READ_ROLE() view returns(bytes32)
func (_Record *RecordSession) READROLE() ([32]byte, error) {
	return _Record.Contract.READROLE(&_Record.CallOpts)
}

// READROLE is a free data retrieval call binding the contract method 0xaf15256a.
//
// Solidity: function READ_ROLE() view returns(bytes32)
func (_Record *RecordCallerSession) READROLE() ([32]byte, error) {
	return _Record.Contract.READROLE(&_Record.CallOpts)
}

// WRITEROLE is a free data retrieval call binding the contract method 0xaf8822a6.
//
// Solidity: function WRITE_ROLE() view returns(bytes32)
func (_Record *RecordCaller) WRITEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Record.contract.Call(opts, &out, "WRITE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// WRITEROLE is a free data retrieval call binding the contract method 0xaf8822a6.
//
// Solidity: function WRITE_ROLE() view returns(bytes32)
func (_Record *RecordSession) WRITEROLE() ([32]byte, error) {
	return _Record.Contract.WRITEROLE(&_Record.CallOpts)
}

// WRITEROLE is a free data retrieval call binding the contract method 0xaf8822a6.
//
// Solidity: function WRITE_ROLE() view returns(bytes32)
func (_Record *RecordCallerSession) WRITEROLE() ([32]byte, error) {
	return _Record.Contract.WRITEROLE(&_Record.CallOpts)
}

// Files is a free data retrieval call binding the contract method 0x8e058482.
//
// Solidity: function files(string , uint256 ) view returns(bytes32 hash, address created_by, uint256 created_at)
func (_Record *RecordCaller) Files(opts *bind.CallOpts, arg0 string, arg1 *big.Int) (struct {
	Hash      [32]byte
	CreatedBy common.Address
	CreatedAt *big.Int
}, error) {
	var out []interface{}
	err := _Record.contract.Call(opts, &out, "files", arg0, arg1)

	outstruct := new(struct {
		Hash      [32]byte
		CreatedBy common.Address
		CreatedAt *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Hash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.CreatedBy = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.CreatedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Files is a free data retrieval call binding the contract method 0x8e058482.
//
// Solidity: function files(string , uint256 ) view returns(bytes32 hash, address created_by, uint256 created_at)
func (_Record *RecordSession) Files(arg0 string, arg1 *big.Int) (struct {
	Hash      [32]byte
	CreatedBy common.Address
	CreatedAt *big.Int
}, error) {
	return _Record.Contract.Files(&_Record.CallOpts, arg0, arg1)
}

// Files is a free data retrieval call binding the contract method 0x8e058482.
//
// Solidity: function files(string , uint256 ) view returns(bytes32 hash, address created_by, uint256 created_at)
func (_Record *RecordCallerSession) Files(arg0 string, arg1 *big.Int) (struct {
	Hash      [32]byte
	CreatedBy common.Address
	CreatedAt *big.Int
}, error) {
	return _Record.Contract.Files(&_Record.CallOpts, arg0, arg1)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Record *RecordCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Record.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Record *RecordSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Record.Contract.GetRoleAdmin(&_Record.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Record *RecordCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Record.Contract.GetRoleAdmin(&_Record.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Record *RecordCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Record.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Record *RecordSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Record.Contract.HasRole(&_Record.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Record *RecordCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Record.Contract.HasRole(&_Record.CallOpts, role, account)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Record *RecordCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Record.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Record *RecordSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Record.Contract.SupportsInterface(&_Record.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Record *RecordCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Record.Contract.SupportsInterface(&_Record.CallOpts, interfaceId)
}

// AddFile is a paid mutator transaction binding the contract method 0x1294495f.
//
// Solidity: function addFile(string fileName, bytes32 fileHash) returns()
func (_Record *RecordTransactor) AddFile(opts *bind.TransactOpts, fileName string, fileHash [32]byte) (*types.Transaction, error) {
	return _Record.contract.Transact(opts, "addFile", fileName, fileHash)
}

// AddFile is a paid mutator transaction binding the contract method 0x1294495f.
//
// Solidity: function addFile(string fileName, bytes32 fileHash) returns()
func (_Record *RecordSession) AddFile(fileName string, fileHash [32]byte) (*types.Transaction, error) {
	return _Record.Contract.AddFile(&_Record.TransactOpts, fileName, fileHash)
}

// AddFile is a paid mutator transaction binding the contract method 0x1294495f.
//
// Solidity: function addFile(string fileName, bytes32 fileHash) returns()
func (_Record *RecordTransactorSession) AddFile(fileName string, fileHash [32]byte) (*types.Transaction, error) {
	return _Record.Contract.AddFile(&_Record.TransactOpts, fileName, fileHash)
}

// DeleteFile is a paid mutator transaction binding the contract method 0xa9910054.
//
// Solidity: function deleteFile(string fileName) returns()
func (_Record *RecordTransactor) DeleteFile(opts *bind.TransactOpts, fileName string) (*types.Transaction, error) {
	return _Record.contract.Transact(opts, "deleteFile", fileName)
}

// DeleteFile is a paid mutator transaction binding the contract method 0xa9910054.
//
// Solidity: function deleteFile(string fileName) returns()
func (_Record *RecordSession) DeleteFile(fileName string) (*types.Transaction, error) {
	return _Record.Contract.DeleteFile(&_Record.TransactOpts, fileName)
}

// DeleteFile is a paid mutator transaction binding the contract method 0xa9910054.
//
// Solidity: function deleteFile(string fileName) returns()
func (_Record *RecordTransactorSession) DeleteFile(fileName string) (*types.Transaction, error) {
	return _Record.Contract.DeleteFile(&_Record.TransactOpts, fileName)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Record *RecordTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Record.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Record *RecordSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Record.Contract.GrantRole(&_Record.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Record *RecordTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Record.Contract.GrantRole(&_Record.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Record *RecordTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Record.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Record *RecordSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Record.Contract.RenounceRole(&_Record.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Record *RecordTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Record.Contract.RenounceRole(&_Record.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Record *RecordTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Record.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Record *RecordSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Record.Contract.RevokeRole(&_Record.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Record *RecordTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Record.Contract.RevokeRole(&_Record.TransactOpts, role, account)
}

// RecordFileAddedIterator is returned from FilterFileAdded and is used to iterate over the raw logs and unpacked data for FileAdded events raised by the Record contract.
type RecordFileAddedIterator struct {
	Event *RecordFileAdded // Event containing the contract specifics and raw log

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
func (it *RecordFileAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecordFileAdded)
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
		it.Event = new(RecordFileAdded)
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
func (it *RecordFileAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RecordFileAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RecordFileAdded represents a FileAdded event raised by the Record contract.
type RecordFileAdded struct {
	FileName  string
	FileHash  [32]byte
	CreatedBy common.Address
	CreatedAt *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFileAdded is a free log retrieval operation binding the contract event 0x56e097685df18d60d3dfcd2e6d325e539d13d1a6c8aecee3d2c5f581f4d48f25.
//
// Solidity: event FileAdded(string fileName, bytes32 fileHash, address createdBy, uint256 createdAt)
func (_Record *RecordFilterer) FilterFileAdded(opts *bind.FilterOpts) (*RecordFileAddedIterator, error) {

	logs, sub, err := _Record.contract.FilterLogs(opts, "FileAdded")
	if err != nil {
		return nil, err
	}
	return &RecordFileAddedIterator{contract: _Record.contract, event: "FileAdded", logs: logs, sub: sub}, nil
}

// WatchFileAdded is a free log subscription operation binding the contract event 0x56e097685df18d60d3dfcd2e6d325e539d13d1a6c8aecee3d2c5f581f4d48f25.
//
// Solidity: event FileAdded(string fileName, bytes32 fileHash, address createdBy, uint256 createdAt)
func (_Record *RecordFilterer) WatchFileAdded(opts *bind.WatchOpts, sink chan<- *RecordFileAdded) (event.Subscription, error) {

	logs, sub, err := _Record.contract.WatchLogs(opts, "FileAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RecordFileAdded)
				if err := _Record.contract.UnpackLog(event, "FileAdded", log); err != nil {
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

// ParseFileAdded is a log parse operation binding the contract event 0x56e097685df18d60d3dfcd2e6d325e539d13d1a6c8aecee3d2c5f581f4d48f25.
//
// Solidity: event FileAdded(string fileName, bytes32 fileHash, address createdBy, uint256 createdAt)
func (_Record *RecordFilterer) ParseFileAdded(log types.Log) (*RecordFileAdded, error) {
	event := new(RecordFileAdded)
	if err := _Record.contract.UnpackLog(event, "FileAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RecordFileDeletedIterator is returned from FilterFileDeleted and is used to iterate over the raw logs and unpacked data for FileDeleted events raised by the Record contract.
type RecordFileDeletedIterator struct {
	Event *RecordFileDeleted // Event containing the contract specifics and raw log

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
func (it *RecordFileDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecordFileDeleted)
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
		it.Event = new(RecordFileDeleted)
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
func (it *RecordFileDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RecordFileDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RecordFileDeleted represents a FileDeleted event raised by the Record contract.
type RecordFileDeleted struct {
	FileName  string
	DeletedBy common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFileDeleted is a free log retrieval operation binding the contract event 0xb7f404dbaa1b2ab0528fd20ea00c12bcd6d8d8e57a5e63af5710d9031e62ef1f.
//
// Solidity: event FileDeleted(string fileName, address deletedBy)
func (_Record *RecordFilterer) FilterFileDeleted(opts *bind.FilterOpts) (*RecordFileDeletedIterator, error) {

	logs, sub, err := _Record.contract.FilterLogs(opts, "FileDeleted")
	if err != nil {
		return nil, err
	}
	return &RecordFileDeletedIterator{contract: _Record.contract, event: "FileDeleted", logs: logs, sub: sub}, nil
}

// WatchFileDeleted is a free log subscription operation binding the contract event 0xb7f404dbaa1b2ab0528fd20ea00c12bcd6d8d8e57a5e63af5710d9031e62ef1f.
//
// Solidity: event FileDeleted(string fileName, address deletedBy)
func (_Record *RecordFilterer) WatchFileDeleted(opts *bind.WatchOpts, sink chan<- *RecordFileDeleted) (event.Subscription, error) {

	logs, sub, err := _Record.contract.WatchLogs(opts, "FileDeleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RecordFileDeleted)
				if err := _Record.contract.UnpackLog(event, "FileDeleted", log); err != nil {
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

// ParseFileDeleted is a log parse operation binding the contract event 0xb7f404dbaa1b2ab0528fd20ea00c12bcd6d8d8e57a5e63af5710d9031e62ef1f.
//
// Solidity: event FileDeleted(string fileName, address deletedBy)
func (_Record *RecordFilterer) ParseFileDeleted(log types.Log) (*RecordFileDeleted, error) {
	event := new(RecordFileDeleted)
	if err := _Record.contract.UnpackLog(event, "FileDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RecordRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Record contract.
type RecordRoleAdminChangedIterator struct {
	Event *RecordRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *RecordRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecordRoleAdminChanged)
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
		it.Event = new(RecordRoleAdminChanged)
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
func (it *RecordRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RecordRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RecordRoleAdminChanged represents a RoleAdminChanged event raised by the Record contract.
type RecordRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Record *RecordFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*RecordRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Record.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &RecordRoleAdminChangedIterator{contract: _Record.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Record *RecordFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *RecordRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Record.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RecordRoleAdminChanged)
				if err := _Record.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Record *RecordFilterer) ParseRoleAdminChanged(log types.Log) (*RecordRoleAdminChanged, error) {
	event := new(RecordRoleAdminChanged)
	if err := _Record.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RecordRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Record contract.
type RecordRoleGrantedIterator struct {
	Event *RecordRoleGranted // Event containing the contract specifics and raw log

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
func (it *RecordRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecordRoleGranted)
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
		it.Event = new(RecordRoleGranted)
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
func (it *RecordRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RecordRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RecordRoleGranted represents a RoleGranted event raised by the Record contract.
type RecordRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Record *RecordFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RecordRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Record.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RecordRoleGrantedIterator{contract: _Record.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Record *RecordFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *RecordRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Record.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RecordRoleGranted)
				if err := _Record.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Record *RecordFilterer) ParseRoleGranted(log types.Log) (*RecordRoleGranted, error) {
	event := new(RecordRoleGranted)
	if err := _Record.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RecordRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Record contract.
type RecordRoleRevokedIterator struct {
	Event *RecordRoleRevoked // Event containing the contract specifics and raw log

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
func (it *RecordRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RecordRoleRevoked)
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
		it.Event = new(RecordRoleRevoked)
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
func (it *RecordRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RecordRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RecordRoleRevoked represents a RoleRevoked event raised by the Record contract.
type RecordRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Record *RecordFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*RecordRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Record.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &RecordRoleRevokedIterator{contract: _Record.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Record *RecordFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *RecordRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Record.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RecordRoleRevoked)
				if err := _Record.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Record *RecordFilterer) ParseRoleRevoked(log types.Log) (*RecordRoleRevoked, error) {
	event := new(RecordRoleRevoked)
	if err := _Record.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

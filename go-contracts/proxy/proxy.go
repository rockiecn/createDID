// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package proxy

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
)

// IAccountDidPublicKey is an auto generated low-level Go binding around an user-defined struct.
type IAccountDidPublicKey struct {
	MethodType  string
	Controller  string
	PubKeyData  []byte
	Deactivated bool
}

// IAccountDidMetaData contains all meta data concerning the IAccountDid contract.
var IAccountDidMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"AddAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"AddAuth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"}],\"name\":\"AddDelegation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"recovery\",\"type\":\"string\"}],\"name\":\"AddRecovery\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"AddVeri\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"CreateDID\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"DeactivateDID\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"DeactivateVeri\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"RemoveAssertion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"RemoveAuth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"RemoveDelegation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"recovery\",\"type\":\"string\"}],\"name\":\"RemoveRecovery\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"UpdateVeri\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"addAssertion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"addAuth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"}],\"name\":\"addDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"addRecovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"methodType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"controller\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"pubKeyData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"deactivated\",\"type\":\"bool\"}],\"internalType\":\"structIAccountDid.PublicKey\",\"name\":\"pubKey\",\"type\":\"tuple\"}],\"name\":\"addVeri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"methodType\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"pubKeyData\",\"type\":\"bytes\"}],\"name\":\"createDID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"deactivate\",\"type\":\"bool\"}],\"name\":\"deactivateDID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"deactivate\",\"type\":\"bool\"}],\"name\":\"deactivateVeri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"getMasterKeyAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"getMasterVerification\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"methodType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"controller\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"pubKeyData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"deactivated\",\"type\":\"bool\"}],\"internalType\":\"structIAccountDid.PublicKey\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getVeri\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"methodType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"controller\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"pubKeyData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"deactivated\",\"type\":\"bool\"}],\"internalType\":\"structIAccountDid.PublicKey\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"getVeriLen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"inAssertion\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"inAuth\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"inDelegation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"inRecovery\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"isDeactivated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"removeAssertion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"removeAuth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"removeDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"removeRecovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"methodType\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"pubKeyData\",\"type\":\"bytes\"}],\"name\":\"updateVeri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"50031b76": "addAssertion(string,string)",
		"f81568bd": "addAuth(string,string)",
		"9045322c": "addDelegation(string,string,uint256)",
		"00d13cf1": "addRecovery(string,string)",
		"ef611e74": "addVeri(string,(string,string,bytes,bool))",
		"e15ddb00": "createDID(string,string,bytes)",
		"c9be3599": "deactivateDID(string,bool)",
		"51ac5022": "deactivateVeri(string,uint256,bool)",
		"0e372676": "getMasterKeyAddr(string)",
		"a174a4cc": "getMasterVerification(string)",
		"f7181a1a": "getVeri(string,uint256)",
		"c7948e4b": "getVeriLen(string)",
		"7c01d259": "inAssertion(string,string)",
		"f888a7de": "inAuth(string,string)",
		"7f0a84cc": "inDelegation(string,string)",
		"b4334b0d": "inRecovery(string,string)",
		"e0064236": "isDeactivated(string)",
		"7904ee08": "removeAssertion(string,string)",
		"e9454a14": "removeAuth(string,string)",
		"f140545b": "removeDelegation(string,string)",
		"c1613ed4": "removeRecovery(string,string)",
		"58ba1928": "updateVeri(string,uint256,string,bytes)",
	},
}

// IAccountDidABI is the input ABI used to generate the binding from.
// Deprecated: Use IAccountDidMetaData.ABI instead.
var IAccountDidABI = IAccountDidMetaData.ABI

// Deprecated: Use IAccountDidMetaData.Sigs instead.
// IAccountDidFuncSigs maps the 4-byte function signature to its string representation.
var IAccountDidFuncSigs = IAccountDidMetaData.Sigs

// IAccountDid is an auto generated Go binding around an Ethereum contract.
type IAccountDid struct {
	IAccountDidCaller     // Read-only binding to the contract
	IAccountDidTransactor // Write-only binding to the contract
	IAccountDidFilterer   // Log filterer for contract events
}

// IAccountDidCaller is an auto generated read-only Go binding around an Ethereum contract.
type IAccountDidCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccountDidTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IAccountDidTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccountDidFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAccountDidFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAccountDidSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAccountDidSession struct {
	Contract     *IAccountDid      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAccountDidCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAccountDidCallerSession struct {
	Contract *IAccountDidCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IAccountDidTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAccountDidTransactorSession struct {
	Contract     *IAccountDidTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IAccountDidRaw is an auto generated low-level Go binding around an Ethereum contract.
type IAccountDidRaw struct {
	Contract *IAccountDid // Generic contract binding to access the raw methods on
}

// IAccountDidCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAccountDidCallerRaw struct {
	Contract *IAccountDidCaller // Generic read-only contract binding to access the raw methods on
}

// IAccountDidTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAccountDidTransactorRaw struct {
	Contract *IAccountDidTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIAccountDid creates a new instance of IAccountDid, bound to a specific deployed contract.
func NewIAccountDid(address common.Address, backend bind.ContractBackend) (*IAccountDid, error) {
	contract, err := bindIAccountDid(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAccountDid{IAccountDidCaller: IAccountDidCaller{contract: contract}, IAccountDidTransactor: IAccountDidTransactor{contract: contract}, IAccountDidFilterer: IAccountDidFilterer{contract: contract}}, nil
}

// NewIAccountDidCaller creates a new read-only instance of IAccountDid, bound to a specific deployed contract.
func NewIAccountDidCaller(address common.Address, caller bind.ContractCaller) (*IAccountDidCaller, error) {
	contract, err := bindIAccountDid(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAccountDidCaller{contract: contract}, nil
}

// NewIAccountDidTransactor creates a new write-only instance of IAccountDid, bound to a specific deployed contract.
func NewIAccountDidTransactor(address common.Address, transactor bind.ContractTransactor) (*IAccountDidTransactor, error) {
	contract, err := bindIAccountDid(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAccountDidTransactor{contract: contract}, nil
}

// NewIAccountDidFilterer creates a new log filterer instance of IAccountDid, bound to a specific deployed contract.
func NewIAccountDidFilterer(address common.Address, filterer bind.ContractFilterer) (*IAccountDidFilterer, error) {
	contract, err := bindIAccountDid(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAccountDidFilterer{contract: contract}, nil
}

// bindIAccountDid binds a generic wrapper to an already deployed contract.
func bindIAccountDid(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IAccountDidABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccountDid *IAccountDidRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccountDid.Contract.IAccountDidCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccountDid *IAccountDidRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccountDid.Contract.IAccountDidTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccountDid *IAccountDidRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccountDid.Contract.IAccountDidTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAccountDid *IAccountDidCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAccountDid.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAccountDid *IAccountDidTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAccountDid.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAccountDid *IAccountDidTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAccountDid.Contract.contract.Transact(opts, method, params...)
}

// GetMasterKeyAddr is a free data retrieval call binding the contract method 0x0e372676.
//
// Solidity: function getMasterKeyAddr(string did) view returns(address)
func (_IAccountDid *IAccountDidCaller) GetMasterKeyAddr(opts *bind.CallOpts, did string) (common.Address, error) {
	var out []interface{}
	err := _IAccountDid.contract.Call(opts, &out, "getMasterKeyAddr", did)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMasterKeyAddr is a free data retrieval call binding the contract method 0x0e372676.
//
// Solidity: function getMasterKeyAddr(string did) view returns(address)
func (_IAccountDid *IAccountDidSession) GetMasterKeyAddr(did string) (common.Address, error) {
	return _IAccountDid.Contract.GetMasterKeyAddr(&_IAccountDid.CallOpts, did)
}

// GetMasterKeyAddr is a free data retrieval call binding the contract method 0x0e372676.
//
// Solidity: function getMasterKeyAddr(string did) view returns(address)
func (_IAccountDid *IAccountDidCallerSession) GetMasterKeyAddr(did string) (common.Address, error) {
	return _IAccountDid.Contract.GetMasterKeyAddr(&_IAccountDid.CallOpts, did)
}

// GetMasterVerification is a free data retrieval call binding the contract method 0xa174a4cc.
//
// Solidity: function getMasterVerification(string did) view returns((string,string,bytes,bool))
func (_IAccountDid *IAccountDidCaller) GetMasterVerification(opts *bind.CallOpts, did string) (IAccountDidPublicKey, error) {
	var out []interface{}
	err := _IAccountDid.contract.Call(opts, &out, "getMasterVerification", did)

	if err != nil {
		return *new(IAccountDidPublicKey), err
	}

	out0 := *abi.ConvertType(out[0], new(IAccountDidPublicKey)).(*IAccountDidPublicKey)

	return out0, err

}

// GetMasterVerification is a free data retrieval call binding the contract method 0xa174a4cc.
//
// Solidity: function getMasterVerification(string did) view returns((string,string,bytes,bool))
func (_IAccountDid *IAccountDidSession) GetMasterVerification(did string) (IAccountDidPublicKey, error) {
	return _IAccountDid.Contract.GetMasterVerification(&_IAccountDid.CallOpts, did)
}

// GetMasterVerification is a free data retrieval call binding the contract method 0xa174a4cc.
//
// Solidity: function getMasterVerification(string did) view returns((string,string,bytes,bool))
func (_IAccountDid *IAccountDidCallerSession) GetMasterVerification(did string) (IAccountDidPublicKey, error) {
	return _IAccountDid.Contract.GetMasterVerification(&_IAccountDid.CallOpts, did)
}

// GetVeri is a free data retrieval call binding the contract method 0xf7181a1a.
//
// Solidity: function getVeri(string did, uint256 index) view returns((string,string,bytes,bool))
func (_IAccountDid *IAccountDidCaller) GetVeri(opts *bind.CallOpts, did string, index *big.Int) (IAccountDidPublicKey, error) {
	var out []interface{}
	err := _IAccountDid.contract.Call(opts, &out, "getVeri", did, index)

	if err != nil {
		return *new(IAccountDidPublicKey), err
	}

	out0 := *abi.ConvertType(out[0], new(IAccountDidPublicKey)).(*IAccountDidPublicKey)

	return out0, err

}

// GetVeri is a free data retrieval call binding the contract method 0xf7181a1a.
//
// Solidity: function getVeri(string did, uint256 index) view returns((string,string,bytes,bool))
func (_IAccountDid *IAccountDidSession) GetVeri(did string, index *big.Int) (IAccountDidPublicKey, error) {
	return _IAccountDid.Contract.GetVeri(&_IAccountDid.CallOpts, did, index)
}

// GetVeri is a free data retrieval call binding the contract method 0xf7181a1a.
//
// Solidity: function getVeri(string did, uint256 index) view returns((string,string,bytes,bool))
func (_IAccountDid *IAccountDidCallerSession) GetVeri(did string, index *big.Int) (IAccountDidPublicKey, error) {
	return _IAccountDid.Contract.GetVeri(&_IAccountDid.CallOpts, did, index)
}

// GetVeriLen is a free data retrieval call binding the contract method 0xc7948e4b.
//
// Solidity: function getVeriLen(string did) view returns(uint256)
func (_IAccountDid *IAccountDidCaller) GetVeriLen(opts *bind.CallOpts, did string) (*big.Int, error) {
	var out []interface{}
	err := _IAccountDid.contract.Call(opts, &out, "getVeriLen", did)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVeriLen is a free data retrieval call binding the contract method 0xc7948e4b.
//
// Solidity: function getVeriLen(string did) view returns(uint256)
func (_IAccountDid *IAccountDidSession) GetVeriLen(did string) (*big.Int, error) {
	return _IAccountDid.Contract.GetVeriLen(&_IAccountDid.CallOpts, did)
}

// GetVeriLen is a free data retrieval call binding the contract method 0xc7948e4b.
//
// Solidity: function getVeriLen(string did) view returns(uint256)
func (_IAccountDid *IAccountDidCallerSession) GetVeriLen(did string) (*big.Int, error) {
	return _IAccountDid.Contract.GetVeriLen(&_IAccountDid.CallOpts, did)
}

// InAssertion is a free data retrieval call binding the contract method 0x7c01d259.
//
// Solidity: function inAssertion(string did, string id) view returns(bool)
func (_IAccountDid *IAccountDidCaller) InAssertion(opts *bind.CallOpts, did string, id string) (bool, error) {
	var out []interface{}
	err := _IAccountDid.contract.Call(opts, &out, "inAssertion", did, id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InAssertion is a free data retrieval call binding the contract method 0x7c01d259.
//
// Solidity: function inAssertion(string did, string id) view returns(bool)
func (_IAccountDid *IAccountDidSession) InAssertion(did string, id string) (bool, error) {
	return _IAccountDid.Contract.InAssertion(&_IAccountDid.CallOpts, did, id)
}

// InAssertion is a free data retrieval call binding the contract method 0x7c01d259.
//
// Solidity: function inAssertion(string did, string id) view returns(bool)
func (_IAccountDid *IAccountDidCallerSession) InAssertion(did string, id string) (bool, error) {
	return _IAccountDid.Contract.InAssertion(&_IAccountDid.CallOpts, did, id)
}

// InAuth is a free data retrieval call binding the contract method 0xf888a7de.
//
// Solidity: function inAuth(string did, string id) view returns(bool)
func (_IAccountDid *IAccountDidCaller) InAuth(opts *bind.CallOpts, did string, id string) (bool, error) {
	var out []interface{}
	err := _IAccountDid.contract.Call(opts, &out, "inAuth", did, id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InAuth is a free data retrieval call binding the contract method 0xf888a7de.
//
// Solidity: function inAuth(string did, string id) view returns(bool)
func (_IAccountDid *IAccountDidSession) InAuth(did string, id string) (bool, error) {
	return _IAccountDid.Contract.InAuth(&_IAccountDid.CallOpts, did, id)
}

// InAuth is a free data retrieval call binding the contract method 0xf888a7de.
//
// Solidity: function inAuth(string did, string id) view returns(bool)
func (_IAccountDid *IAccountDidCallerSession) InAuth(did string, id string) (bool, error) {
	return _IAccountDid.Contract.InAuth(&_IAccountDid.CallOpts, did, id)
}

// InDelegation is a free data retrieval call binding the contract method 0x7f0a84cc.
//
// Solidity: function inDelegation(string did, string id) view returns(uint256)
func (_IAccountDid *IAccountDidCaller) InDelegation(opts *bind.CallOpts, did string, id string) (*big.Int, error) {
	var out []interface{}
	err := _IAccountDid.contract.Call(opts, &out, "inDelegation", did, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InDelegation is a free data retrieval call binding the contract method 0x7f0a84cc.
//
// Solidity: function inDelegation(string did, string id) view returns(uint256)
func (_IAccountDid *IAccountDidSession) InDelegation(did string, id string) (*big.Int, error) {
	return _IAccountDid.Contract.InDelegation(&_IAccountDid.CallOpts, did, id)
}

// InDelegation is a free data retrieval call binding the contract method 0x7f0a84cc.
//
// Solidity: function inDelegation(string did, string id) view returns(uint256)
func (_IAccountDid *IAccountDidCallerSession) InDelegation(did string, id string) (*big.Int, error) {
	return _IAccountDid.Contract.InDelegation(&_IAccountDid.CallOpts, did, id)
}

// InRecovery is a free data retrieval call binding the contract method 0xb4334b0d.
//
// Solidity: function inRecovery(string did, string id) view returns(bool)
func (_IAccountDid *IAccountDidCaller) InRecovery(opts *bind.CallOpts, did string, id string) (bool, error) {
	var out []interface{}
	err := _IAccountDid.contract.Call(opts, &out, "inRecovery", did, id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InRecovery is a free data retrieval call binding the contract method 0xb4334b0d.
//
// Solidity: function inRecovery(string did, string id) view returns(bool)
func (_IAccountDid *IAccountDidSession) InRecovery(did string, id string) (bool, error) {
	return _IAccountDid.Contract.InRecovery(&_IAccountDid.CallOpts, did, id)
}

// InRecovery is a free data retrieval call binding the contract method 0xb4334b0d.
//
// Solidity: function inRecovery(string did, string id) view returns(bool)
func (_IAccountDid *IAccountDidCallerSession) InRecovery(did string, id string) (bool, error) {
	return _IAccountDid.Contract.InRecovery(&_IAccountDid.CallOpts, did, id)
}

// IsDeactivated is a free data retrieval call binding the contract method 0xe0064236.
//
// Solidity: function isDeactivated(string did) view returns(bool)
func (_IAccountDid *IAccountDidCaller) IsDeactivated(opts *bind.CallOpts, did string) (bool, error) {
	var out []interface{}
	err := _IAccountDid.contract.Call(opts, &out, "isDeactivated", did)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDeactivated is a free data retrieval call binding the contract method 0xe0064236.
//
// Solidity: function isDeactivated(string did) view returns(bool)
func (_IAccountDid *IAccountDidSession) IsDeactivated(did string) (bool, error) {
	return _IAccountDid.Contract.IsDeactivated(&_IAccountDid.CallOpts, did)
}

// IsDeactivated is a free data retrieval call binding the contract method 0xe0064236.
//
// Solidity: function isDeactivated(string did) view returns(bool)
func (_IAccountDid *IAccountDidCallerSession) IsDeactivated(did string) (bool, error) {
	return _IAccountDid.Contract.IsDeactivated(&_IAccountDid.CallOpts, did)
}

// AddAssertion is a paid mutator transaction binding the contract method 0x50031b76.
//
// Solidity: function addAssertion(string did, string id) returns()
func (_IAccountDid *IAccountDidTransactor) AddAssertion(opts *bind.TransactOpts, did string, id string) (*types.Transaction, error) {
	return _IAccountDid.contract.Transact(opts, "addAssertion", did, id)
}

// AddAssertion is a paid mutator transaction binding the contract method 0x50031b76.
//
// Solidity: function addAssertion(string did, string id) returns()
func (_IAccountDid *IAccountDidSession) AddAssertion(did string, id string) (*types.Transaction, error) {
	return _IAccountDid.Contract.AddAssertion(&_IAccountDid.TransactOpts, did, id)
}

// AddAssertion is a paid mutator transaction binding the contract method 0x50031b76.
//
// Solidity: function addAssertion(string did, string id) returns()
func (_IAccountDid *IAccountDidTransactorSession) AddAssertion(did string, id string) (*types.Transaction, error) {
	return _IAccountDid.Contract.AddAssertion(&_IAccountDid.TransactOpts, did, id)
}

// AddAuth is a paid mutator transaction binding the contract method 0xf81568bd.
//
// Solidity: function addAuth(string did, string id) returns()
func (_IAccountDid *IAccountDidTransactor) AddAuth(opts *bind.TransactOpts, did string, id string) (*types.Transaction, error) {
	return _IAccountDid.contract.Transact(opts, "addAuth", did, id)
}

// AddAuth is a paid mutator transaction binding the contract method 0xf81568bd.
//
// Solidity: function addAuth(string did, string id) returns()
func (_IAccountDid *IAccountDidSession) AddAuth(did string, id string) (*types.Transaction, error) {
	return _IAccountDid.Contract.AddAuth(&_IAccountDid.TransactOpts, did, id)
}

// AddAuth is a paid mutator transaction binding the contract method 0xf81568bd.
//
// Solidity: function addAuth(string did, string id) returns()
func (_IAccountDid *IAccountDidTransactorSession) AddAuth(did string, id string) (*types.Transaction, error) {
	return _IAccountDid.Contract.AddAuth(&_IAccountDid.TransactOpts, did, id)
}

// AddDelegation is a paid mutator transaction binding the contract method 0x9045322c.
//
// Solidity: function addDelegation(string did, string id, uint256 expiration) returns()
func (_IAccountDid *IAccountDidTransactor) AddDelegation(opts *bind.TransactOpts, did string, id string, expiration *big.Int) (*types.Transaction, error) {
	return _IAccountDid.contract.Transact(opts, "addDelegation", did, id, expiration)
}

// AddDelegation is a paid mutator transaction binding the contract method 0x9045322c.
//
// Solidity: function addDelegation(string did, string id, uint256 expiration) returns()
func (_IAccountDid *IAccountDidSession) AddDelegation(did string, id string, expiration *big.Int) (*types.Transaction, error) {
	return _IAccountDid.Contract.AddDelegation(&_IAccountDid.TransactOpts, did, id, expiration)
}

// AddDelegation is a paid mutator transaction binding the contract method 0x9045322c.
//
// Solidity: function addDelegation(string did, string id, uint256 expiration) returns()
func (_IAccountDid *IAccountDidTransactorSession) AddDelegation(did string, id string, expiration *big.Int) (*types.Transaction, error) {
	return _IAccountDid.Contract.AddDelegation(&_IAccountDid.TransactOpts, did, id, expiration)
}

// AddRecovery is a paid mutator transaction binding the contract method 0x00d13cf1.
//
// Solidity: function addRecovery(string did, string id) returns()
func (_IAccountDid *IAccountDidTransactor) AddRecovery(opts *bind.TransactOpts, did string, id string) (*types.Transaction, error) {
	return _IAccountDid.contract.Transact(opts, "addRecovery", did, id)
}

// AddRecovery is a paid mutator transaction binding the contract method 0x00d13cf1.
//
// Solidity: function addRecovery(string did, string id) returns()
func (_IAccountDid *IAccountDidSession) AddRecovery(did string, id string) (*types.Transaction, error) {
	return _IAccountDid.Contract.AddRecovery(&_IAccountDid.TransactOpts, did, id)
}

// AddRecovery is a paid mutator transaction binding the contract method 0x00d13cf1.
//
// Solidity: function addRecovery(string did, string id) returns()
func (_IAccountDid *IAccountDidTransactorSession) AddRecovery(did string, id string) (*types.Transaction, error) {
	return _IAccountDid.Contract.AddRecovery(&_IAccountDid.TransactOpts, did, id)
}

// AddVeri is a paid mutator transaction binding the contract method 0xef611e74.
//
// Solidity: function addVeri(string did, (string,string,bytes,bool) pubKey) returns()
func (_IAccountDid *IAccountDidTransactor) AddVeri(opts *bind.TransactOpts, did string, pubKey IAccountDidPublicKey) (*types.Transaction, error) {
	return _IAccountDid.contract.Transact(opts, "addVeri", did, pubKey)
}

// AddVeri is a paid mutator transaction binding the contract method 0xef611e74.
//
// Solidity: function addVeri(string did, (string,string,bytes,bool) pubKey) returns()
func (_IAccountDid *IAccountDidSession) AddVeri(did string, pubKey IAccountDidPublicKey) (*types.Transaction, error) {
	return _IAccountDid.Contract.AddVeri(&_IAccountDid.TransactOpts, did, pubKey)
}

// AddVeri is a paid mutator transaction binding the contract method 0xef611e74.
//
// Solidity: function addVeri(string did, (string,string,bytes,bool) pubKey) returns()
func (_IAccountDid *IAccountDidTransactorSession) AddVeri(did string, pubKey IAccountDidPublicKey) (*types.Transaction, error) {
	return _IAccountDid.Contract.AddVeri(&_IAccountDid.TransactOpts, did, pubKey)
}

// CreateDID is a paid mutator transaction binding the contract method 0xe15ddb00.
//
// Solidity: function createDID(string did, string methodType, bytes pubKeyData) returns()
func (_IAccountDid *IAccountDidTransactor) CreateDID(opts *bind.TransactOpts, did string, methodType string, pubKeyData []byte) (*types.Transaction, error) {
	return _IAccountDid.contract.Transact(opts, "createDID", did, methodType, pubKeyData)
}

// CreateDID is a paid mutator transaction binding the contract method 0xe15ddb00.
//
// Solidity: function createDID(string did, string methodType, bytes pubKeyData) returns()
func (_IAccountDid *IAccountDidSession) CreateDID(did string, methodType string, pubKeyData []byte) (*types.Transaction, error) {
	return _IAccountDid.Contract.CreateDID(&_IAccountDid.TransactOpts, did, methodType, pubKeyData)
}

// CreateDID is a paid mutator transaction binding the contract method 0xe15ddb00.
//
// Solidity: function createDID(string did, string methodType, bytes pubKeyData) returns()
func (_IAccountDid *IAccountDidTransactorSession) CreateDID(did string, methodType string, pubKeyData []byte) (*types.Transaction, error) {
	return _IAccountDid.Contract.CreateDID(&_IAccountDid.TransactOpts, did, methodType, pubKeyData)
}

// DeactivateDID is a paid mutator transaction binding the contract method 0xc9be3599.
//
// Solidity: function deactivateDID(string did, bool deactivate) returns()
func (_IAccountDid *IAccountDidTransactor) DeactivateDID(opts *bind.TransactOpts, did string, deactivate bool) (*types.Transaction, error) {
	return _IAccountDid.contract.Transact(opts, "deactivateDID", did, deactivate)
}

// DeactivateDID is a paid mutator transaction binding the contract method 0xc9be3599.
//
// Solidity: function deactivateDID(string did, bool deactivate) returns()
func (_IAccountDid *IAccountDidSession) DeactivateDID(did string, deactivate bool) (*types.Transaction, error) {
	return _IAccountDid.Contract.DeactivateDID(&_IAccountDid.TransactOpts, did, deactivate)
}

// DeactivateDID is a paid mutator transaction binding the contract method 0xc9be3599.
//
// Solidity: function deactivateDID(string did, bool deactivate) returns()
func (_IAccountDid *IAccountDidTransactorSession) DeactivateDID(did string, deactivate bool) (*types.Transaction, error) {
	return _IAccountDid.Contract.DeactivateDID(&_IAccountDid.TransactOpts, did, deactivate)
}

// DeactivateVeri is a paid mutator transaction binding the contract method 0x51ac5022.
//
// Solidity: function deactivateVeri(string did, uint256 index, bool deactivate) returns()
func (_IAccountDid *IAccountDidTransactor) DeactivateVeri(opts *bind.TransactOpts, did string, index *big.Int, deactivate bool) (*types.Transaction, error) {
	return _IAccountDid.contract.Transact(opts, "deactivateVeri", did, index, deactivate)
}

// DeactivateVeri is a paid mutator transaction binding the contract method 0x51ac5022.
//
// Solidity: function deactivateVeri(string did, uint256 index, bool deactivate) returns()
func (_IAccountDid *IAccountDidSession) DeactivateVeri(did string, index *big.Int, deactivate bool) (*types.Transaction, error) {
	return _IAccountDid.Contract.DeactivateVeri(&_IAccountDid.TransactOpts, did, index, deactivate)
}

// DeactivateVeri is a paid mutator transaction binding the contract method 0x51ac5022.
//
// Solidity: function deactivateVeri(string did, uint256 index, bool deactivate) returns()
func (_IAccountDid *IAccountDidTransactorSession) DeactivateVeri(did string, index *big.Int, deactivate bool) (*types.Transaction, error) {
	return _IAccountDid.Contract.DeactivateVeri(&_IAccountDid.TransactOpts, did, index, deactivate)
}

// RemoveAssertion is a paid mutator transaction binding the contract method 0x7904ee08.
//
// Solidity: function removeAssertion(string did, string id) returns()
func (_IAccountDid *IAccountDidTransactor) RemoveAssertion(opts *bind.TransactOpts, did string, id string) (*types.Transaction, error) {
	return _IAccountDid.contract.Transact(opts, "removeAssertion", did, id)
}

// RemoveAssertion is a paid mutator transaction binding the contract method 0x7904ee08.
//
// Solidity: function removeAssertion(string did, string id) returns()
func (_IAccountDid *IAccountDidSession) RemoveAssertion(did string, id string) (*types.Transaction, error) {
	return _IAccountDid.Contract.RemoveAssertion(&_IAccountDid.TransactOpts, did, id)
}

// RemoveAssertion is a paid mutator transaction binding the contract method 0x7904ee08.
//
// Solidity: function removeAssertion(string did, string id) returns()
func (_IAccountDid *IAccountDidTransactorSession) RemoveAssertion(did string, id string) (*types.Transaction, error) {
	return _IAccountDid.Contract.RemoveAssertion(&_IAccountDid.TransactOpts, did, id)
}

// RemoveAuth is a paid mutator transaction binding the contract method 0xe9454a14.
//
// Solidity: function removeAuth(string did, string id) returns()
func (_IAccountDid *IAccountDidTransactor) RemoveAuth(opts *bind.TransactOpts, did string, id string) (*types.Transaction, error) {
	return _IAccountDid.contract.Transact(opts, "removeAuth", did, id)
}

// RemoveAuth is a paid mutator transaction binding the contract method 0xe9454a14.
//
// Solidity: function removeAuth(string did, string id) returns()
func (_IAccountDid *IAccountDidSession) RemoveAuth(did string, id string) (*types.Transaction, error) {
	return _IAccountDid.Contract.RemoveAuth(&_IAccountDid.TransactOpts, did, id)
}

// RemoveAuth is a paid mutator transaction binding the contract method 0xe9454a14.
//
// Solidity: function removeAuth(string did, string id) returns()
func (_IAccountDid *IAccountDidTransactorSession) RemoveAuth(did string, id string) (*types.Transaction, error) {
	return _IAccountDid.Contract.RemoveAuth(&_IAccountDid.TransactOpts, did, id)
}

// RemoveDelegation is a paid mutator transaction binding the contract method 0xf140545b.
//
// Solidity: function removeDelegation(string did, string id) returns()
func (_IAccountDid *IAccountDidTransactor) RemoveDelegation(opts *bind.TransactOpts, did string, id string) (*types.Transaction, error) {
	return _IAccountDid.contract.Transact(opts, "removeDelegation", did, id)
}

// RemoveDelegation is a paid mutator transaction binding the contract method 0xf140545b.
//
// Solidity: function removeDelegation(string did, string id) returns()
func (_IAccountDid *IAccountDidSession) RemoveDelegation(did string, id string) (*types.Transaction, error) {
	return _IAccountDid.Contract.RemoveDelegation(&_IAccountDid.TransactOpts, did, id)
}

// RemoveDelegation is a paid mutator transaction binding the contract method 0xf140545b.
//
// Solidity: function removeDelegation(string did, string id) returns()
func (_IAccountDid *IAccountDidTransactorSession) RemoveDelegation(did string, id string) (*types.Transaction, error) {
	return _IAccountDid.Contract.RemoveDelegation(&_IAccountDid.TransactOpts, did, id)
}

// RemoveRecovery is a paid mutator transaction binding the contract method 0xc1613ed4.
//
// Solidity: function removeRecovery(string did, string id) returns()
func (_IAccountDid *IAccountDidTransactor) RemoveRecovery(opts *bind.TransactOpts, did string, id string) (*types.Transaction, error) {
	return _IAccountDid.contract.Transact(opts, "removeRecovery", did, id)
}

// RemoveRecovery is a paid mutator transaction binding the contract method 0xc1613ed4.
//
// Solidity: function removeRecovery(string did, string id) returns()
func (_IAccountDid *IAccountDidSession) RemoveRecovery(did string, id string) (*types.Transaction, error) {
	return _IAccountDid.Contract.RemoveRecovery(&_IAccountDid.TransactOpts, did, id)
}

// RemoveRecovery is a paid mutator transaction binding the contract method 0xc1613ed4.
//
// Solidity: function removeRecovery(string did, string id) returns()
func (_IAccountDid *IAccountDidTransactorSession) RemoveRecovery(did string, id string) (*types.Transaction, error) {
	return _IAccountDid.Contract.RemoveRecovery(&_IAccountDid.TransactOpts, did, id)
}

// UpdateVeri is a paid mutator transaction binding the contract method 0x58ba1928.
//
// Solidity: function updateVeri(string did, uint256 index, string methodType, bytes pubKeyData) returns()
func (_IAccountDid *IAccountDidTransactor) UpdateVeri(opts *bind.TransactOpts, did string, index *big.Int, methodType string, pubKeyData []byte) (*types.Transaction, error) {
	return _IAccountDid.contract.Transact(opts, "updateVeri", did, index, methodType, pubKeyData)
}

// UpdateVeri is a paid mutator transaction binding the contract method 0x58ba1928.
//
// Solidity: function updateVeri(string did, uint256 index, string methodType, bytes pubKeyData) returns()
func (_IAccountDid *IAccountDidSession) UpdateVeri(did string, index *big.Int, methodType string, pubKeyData []byte) (*types.Transaction, error) {
	return _IAccountDid.Contract.UpdateVeri(&_IAccountDid.TransactOpts, did, index, methodType, pubKeyData)
}

// UpdateVeri is a paid mutator transaction binding the contract method 0x58ba1928.
//
// Solidity: function updateVeri(string did, uint256 index, string methodType, bytes pubKeyData) returns()
func (_IAccountDid *IAccountDidTransactorSession) UpdateVeri(did string, index *big.Int, methodType string, pubKeyData []byte) (*types.Transaction, error) {
	return _IAccountDid.Contract.UpdateVeri(&_IAccountDid.TransactOpts, did, index, methodType, pubKeyData)
}

// IAccountDidAddAssertionIterator is returned from FilterAddAssertion and is used to iterate over the raw logs and unpacked data for AddAssertion events raised by the IAccountDid contract.
type IAccountDidAddAssertionIterator struct {
	Event *IAccountDidAddAssertion // Event containing the contract specifics and raw log

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
func (it *IAccountDidAddAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccountDidAddAssertion)
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
		it.Event = new(IAccountDidAddAssertion)
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
func (it *IAccountDidAddAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccountDidAddAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccountDidAddAssertion represents a AddAssertion event raised by the IAccountDid contract.
type IAccountDidAddAssertion struct {
	Did common.Hash
	Id  string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAddAssertion is a free log retrieval operation binding the contract event 0x3dc8a6dcefb090c20c6f37f01dabde1a17b9c6dd4969ce7e78206888764c2490.
//
// Solidity: event AddAssertion(string indexed did, string id)
func (_IAccountDid *IAccountDidFilterer) FilterAddAssertion(opts *bind.FilterOpts, did []string) (*IAccountDidAddAssertionIterator, error) {

	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _IAccountDid.contract.FilterLogs(opts, "AddAssertion", didRule)
	if err != nil {
		return nil, err
	}
	return &IAccountDidAddAssertionIterator{contract: _IAccountDid.contract, event: "AddAssertion", logs: logs, sub: sub}, nil
}

// WatchAddAssertion is a free log subscription operation binding the contract event 0x3dc8a6dcefb090c20c6f37f01dabde1a17b9c6dd4969ce7e78206888764c2490.
//
// Solidity: event AddAssertion(string indexed did, string id)
func (_IAccountDid *IAccountDidFilterer) WatchAddAssertion(opts *bind.WatchOpts, sink chan<- *IAccountDidAddAssertion, did []string) (event.Subscription, error) {

	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _IAccountDid.contract.WatchLogs(opts, "AddAssertion", didRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccountDidAddAssertion)
				if err := _IAccountDid.contract.UnpackLog(event, "AddAssertion", log); err != nil {
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

// ParseAddAssertion is a log parse operation binding the contract event 0x3dc8a6dcefb090c20c6f37f01dabde1a17b9c6dd4969ce7e78206888764c2490.
//
// Solidity: event AddAssertion(string indexed did, string id)
func (_IAccountDid *IAccountDidFilterer) ParseAddAssertion(log types.Log) (*IAccountDidAddAssertion, error) {
	event := new(IAccountDidAddAssertion)
	if err := _IAccountDid.contract.UnpackLog(event, "AddAssertion", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAccountDidAddAuthIterator is returned from FilterAddAuth and is used to iterate over the raw logs and unpacked data for AddAuth events raised by the IAccountDid contract.
type IAccountDidAddAuthIterator struct {
	Event *IAccountDidAddAuth // Event containing the contract specifics and raw log

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
func (it *IAccountDidAddAuthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccountDidAddAuth)
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
		it.Event = new(IAccountDidAddAuth)
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
func (it *IAccountDidAddAuthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccountDidAddAuthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccountDidAddAuth represents a AddAuth event raised by the IAccountDid contract.
type IAccountDidAddAuth struct {
	Did common.Hash
	Id  string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAddAuth is a free log retrieval operation binding the contract event 0x7b081ab60e2b5916055b661b545c62638d8472b459efdf55571e87a60a83ec0e.
//
// Solidity: event AddAuth(string indexed did, string id)
func (_IAccountDid *IAccountDidFilterer) FilterAddAuth(opts *bind.FilterOpts, did []string) (*IAccountDidAddAuthIterator, error) {

	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _IAccountDid.contract.FilterLogs(opts, "AddAuth", didRule)
	if err != nil {
		return nil, err
	}
	return &IAccountDidAddAuthIterator{contract: _IAccountDid.contract, event: "AddAuth", logs: logs, sub: sub}, nil
}

// WatchAddAuth is a free log subscription operation binding the contract event 0x7b081ab60e2b5916055b661b545c62638d8472b459efdf55571e87a60a83ec0e.
//
// Solidity: event AddAuth(string indexed did, string id)
func (_IAccountDid *IAccountDidFilterer) WatchAddAuth(opts *bind.WatchOpts, sink chan<- *IAccountDidAddAuth, did []string) (event.Subscription, error) {

	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _IAccountDid.contract.WatchLogs(opts, "AddAuth", didRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccountDidAddAuth)
				if err := _IAccountDid.contract.UnpackLog(event, "AddAuth", log); err != nil {
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

// ParseAddAuth is a log parse operation binding the contract event 0x7b081ab60e2b5916055b661b545c62638d8472b459efdf55571e87a60a83ec0e.
//
// Solidity: event AddAuth(string indexed did, string id)
func (_IAccountDid *IAccountDidFilterer) ParseAddAuth(log types.Log) (*IAccountDidAddAuth, error) {
	event := new(IAccountDidAddAuth)
	if err := _IAccountDid.contract.UnpackLog(event, "AddAuth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAccountDidAddDelegationIterator is returned from FilterAddDelegation and is used to iterate over the raw logs and unpacked data for AddDelegation events raised by the IAccountDid contract.
type IAccountDidAddDelegationIterator struct {
	Event *IAccountDidAddDelegation // Event containing the contract specifics and raw log

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
func (it *IAccountDidAddDelegationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccountDidAddDelegation)
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
		it.Event = new(IAccountDidAddDelegation)
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
func (it *IAccountDidAddDelegationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccountDidAddDelegationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccountDidAddDelegation represents a AddDelegation event raised by the IAccountDid contract.
type IAccountDidAddDelegation struct {
	Did        common.Hash
	Id         string
	Expiration *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddDelegation is a free log retrieval operation binding the contract event 0xd8b844a934c50876df4c445f08f47cc0ec2dc234e128be3af72b2f86e89da060.
//
// Solidity: event AddDelegation(string indexed did, string id, uint256 expiration)
func (_IAccountDid *IAccountDidFilterer) FilterAddDelegation(opts *bind.FilterOpts, did []string) (*IAccountDidAddDelegationIterator, error) {

	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _IAccountDid.contract.FilterLogs(opts, "AddDelegation", didRule)
	if err != nil {
		return nil, err
	}
	return &IAccountDidAddDelegationIterator{contract: _IAccountDid.contract, event: "AddDelegation", logs: logs, sub: sub}, nil
}

// WatchAddDelegation is a free log subscription operation binding the contract event 0xd8b844a934c50876df4c445f08f47cc0ec2dc234e128be3af72b2f86e89da060.
//
// Solidity: event AddDelegation(string indexed did, string id, uint256 expiration)
func (_IAccountDid *IAccountDidFilterer) WatchAddDelegation(opts *bind.WatchOpts, sink chan<- *IAccountDidAddDelegation, did []string) (event.Subscription, error) {

	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _IAccountDid.contract.WatchLogs(opts, "AddDelegation", didRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccountDidAddDelegation)
				if err := _IAccountDid.contract.UnpackLog(event, "AddDelegation", log); err != nil {
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

// ParseAddDelegation is a log parse operation binding the contract event 0xd8b844a934c50876df4c445f08f47cc0ec2dc234e128be3af72b2f86e89da060.
//
// Solidity: event AddDelegation(string indexed did, string id, uint256 expiration)
func (_IAccountDid *IAccountDidFilterer) ParseAddDelegation(log types.Log) (*IAccountDidAddDelegation, error) {
	event := new(IAccountDidAddDelegation)
	if err := _IAccountDid.contract.UnpackLog(event, "AddDelegation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAccountDidAddRecoveryIterator is returned from FilterAddRecovery and is used to iterate over the raw logs and unpacked data for AddRecovery events raised by the IAccountDid contract.
type IAccountDidAddRecoveryIterator struct {
	Event *IAccountDidAddRecovery // Event containing the contract specifics and raw log

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
func (it *IAccountDidAddRecoveryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccountDidAddRecovery)
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
		it.Event = new(IAccountDidAddRecovery)
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
func (it *IAccountDidAddRecoveryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccountDidAddRecoveryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccountDidAddRecovery represents a AddRecovery event raised by the IAccountDid contract.
type IAccountDidAddRecovery struct {
	Did      common.Hash
	Recovery string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAddRecovery is a free log retrieval operation binding the contract event 0x3cd3cc4e2cfc49d39689a2f57840dd97e5b351162eacf217a43e35297a7e91f9.
//
// Solidity: event AddRecovery(string indexed did, string recovery)
func (_IAccountDid *IAccountDidFilterer) FilterAddRecovery(opts *bind.FilterOpts, did []string) (*IAccountDidAddRecoveryIterator, error) {

	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _IAccountDid.contract.FilterLogs(opts, "AddRecovery", didRule)
	if err != nil {
		return nil, err
	}
	return &IAccountDidAddRecoveryIterator{contract: _IAccountDid.contract, event: "AddRecovery", logs: logs, sub: sub}, nil
}

// WatchAddRecovery is a free log subscription operation binding the contract event 0x3cd3cc4e2cfc49d39689a2f57840dd97e5b351162eacf217a43e35297a7e91f9.
//
// Solidity: event AddRecovery(string indexed did, string recovery)
func (_IAccountDid *IAccountDidFilterer) WatchAddRecovery(opts *bind.WatchOpts, sink chan<- *IAccountDidAddRecovery, did []string) (event.Subscription, error) {

	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _IAccountDid.contract.WatchLogs(opts, "AddRecovery", didRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccountDidAddRecovery)
				if err := _IAccountDid.contract.UnpackLog(event, "AddRecovery", log); err != nil {
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

// ParseAddRecovery is a log parse operation binding the contract event 0x3cd3cc4e2cfc49d39689a2f57840dd97e5b351162eacf217a43e35297a7e91f9.
//
// Solidity: event AddRecovery(string indexed did, string recovery)
func (_IAccountDid *IAccountDidFilterer) ParseAddRecovery(log types.Log) (*IAccountDidAddRecovery, error) {
	event := new(IAccountDidAddRecovery)
	if err := _IAccountDid.contract.UnpackLog(event, "AddRecovery", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAccountDidAddVeriIterator is returned from FilterAddVeri and is used to iterate over the raw logs and unpacked data for AddVeri events raised by the IAccountDid contract.
type IAccountDidAddVeriIterator struct {
	Event *IAccountDidAddVeri // Event containing the contract specifics and raw log

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
func (it *IAccountDidAddVeriIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccountDidAddVeri)
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
		it.Event = new(IAccountDidAddVeri)
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
func (it *IAccountDidAddVeriIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccountDidAddVeriIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccountDidAddVeri represents a AddVeri event raised by the IAccountDid contract.
type IAccountDidAddVeri struct {
	Did common.Hash
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAddVeri is a free log retrieval operation binding the contract event 0x7d14d7731faefa682df55a71db25727d5a1120c74cd2cd0fb4e2fe58e3adc8dd.
//
// Solidity: event AddVeri(string indexed did, uint256 id)
func (_IAccountDid *IAccountDidFilterer) FilterAddVeri(opts *bind.FilterOpts, did []string) (*IAccountDidAddVeriIterator, error) {

	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _IAccountDid.contract.FilterLogs(opts, "AddVeri", didRule)
	if err != nil {
		return nil, err
	}
	return &IAccountDidAddVeriIterator{contract: _IAccountDid.contract, event: "AddVeri", logs: logs, sub: sub}, nil
}

// WatchAddVeri is a free log subscription operation binding the contract event 0x7d14d7731faefa682df55a71db25727d5a1120c74cd2cd0fb4e2fe58e3adc8dd.
//
// Solidity: event AddVeri(string indexed did, uint256 id)
func (_IAccountDid *IAccountDidFilterer) WatchAddVeri(opts *bind.WatchOpts, sink chan<- *IAccountDidAddVeri, did []string) (event.Subscription, error) {

	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _IAccountDid.contract.WatchLogs(opts, "AddVeri", didRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccountDidAddVeri)
				if err := _IAccountDid.contract.UnpackLog(event, "AddVeri", log); err != nil {
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

// ParseAddVeri is a log parse operation binding the contract event 0x7d14d7731faefa682df55a71db25727d5a1120c74cd2cd0fb4e2fe58e3adc8dd.
//
// Solidity: event AddVeri(string indexed did, uint256 id)
func (_IAccountDid *IAccountDidFilterer) ParseAddVeri(log types.Log) (*IAccountDidAddVeri, error) {
	event := new(IAccountDidAddVeri)
	if err := _IAccountDid.contract.UnpackLog(event, "AddVeri", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAccountDidCreateDIDIterator is returned from FilterCreateDID and is used to iterate over the raw logs and unpacked data for CreateDID events raised by the IAccountDid contract.
type IAccountDidCreateDIDIterator struct {
	Event *IAccountDidCreateDID // Event containing the contract specifics and raw log

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
func (it *IAccountDidCreateDIDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccountDidCreateDID)
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
		it.Event = new(IAccountDidCreateDID)
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
func (it *IAccountDidCreateDIDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccountDidCreateDIDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccountDidCreateDID represents a CreateDID event raised by the IAccountDid contract.
type IAccountDidCreateDID struct {
	Did string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCreateDID is a free log retrieval operation binding the contract event 0x79c88c074754eca102316c9e35da362da35dde8df9c3a975a24c60207646d879.
//
// Solidity: event CreateDID(string did)
func (_IAccountDid *IAccountDidFilterer) FilterCreateDID(opts *bind.FilterOpts) (*IAccountDidCreateDIDIterator, error) {

	logs, sub, err := _IAccountDid.contract.FilterLogs(opts, "CreateDID")
	if err != nil {
		return nil, err
	}
	return &IAccountDidCreateDIDIterator{contract: _IAccountDid.contract, event: "CreateDID", logs: logs, sub: sub}, nil
}

// WatchCreateDID is a free log subscription operation binding the contract event 0x79c88c074754eca102316c9e35da362da35dde8df9c3a975a24c60207646d879.
//
// Solidity: event CreateDID(string did)
func (_IAccountDid *IAccountDidFilterer) WatchCreateDID(opts *bind.WatchOpts, sink chan<- *IAccountDidCreateDID) (event.Subscription, error) {

	logs, sub, err := _IAccountDid.contract.WatchLogs(opts, "CreateDID")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccountDidCreateDID)
				if err := _IAccountDid.contract.UnpackLog(event, "CreateDID", log); err != nil {
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

// ParseCreateDID is a log parse operation binding the contract event 0x79c88c074754eca102316c9e35da362da35dde8df9c3a975a24c60207646d879.
//
// Solidity: event CreateDID(string did)
func (_IAccountDid *IAccountDidFilterer) ParseCreateDID(log types.Log) (*IAccountDidCreateDID, error) {
	event := new(IAccountDidCreateDID)
	if err := _IAccountDid.contract.UnpackLog(event, "CreateDID", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAccountDidDeactivateDIDIterator is returned from FilterDeactivateDID and is used to iterate over the raw logs and unpacked data for DeactivateDID events raised by the IAccountDid contract.
type IAccountDidDeactivateDIDIterator struct {
	Event *IAccountDidDeactivateDID // Event containing the contract specifics and raw log

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
func (it *IAccountDidDeactivateDIDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccountDidDeactivateDID)
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
		it.Event = new(IAccountDidDeactivateDID)
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
func (it *IAccountDidDeactivateDIDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccountDidDeactivateDIDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccountDidDeactivateDID represents a DeactivateDID event raised by the IAccountDid contract.
type IAccountDidDeactivateDID struct {
	Did string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeactivateDID is a free log retrieval operation binding the contract event 0x2057c1bebefc71cf1116bc35105e72c4ec2baa1014a67f924edda51eee4565e8.
//
// Solidity: event DeactivateDID(string did)
func (_IAccountDid *IAccountDidFilterer) FilterDeactivateDID(opts *bind.FilterOpts) (*IAccountDidDeactivateDIDIterator, error) {

	logs, sub, err := _IAccountDid.contract.FilterLogs(opts, "DeactivateDID")
	if err != nil {
		return nil, err
	}
	return &IAccountDidDeactivateDIDIterator{contract: _IAccountDid.contract, event: "DeactivateDID", logs: logs, sub: sub}, nil
}

// WatchDeactivateDID is a free log subscription operation binding the contract event 0x2057c1bebefc71cf1116bc35105e72c4ec2baa1014a67f924edda51eee4565e8.
//
// Solidity: event DeactivateDID(string did)
func (_IAccountDid *IAccountDidFilterer) WatchDeactivateDID(opts *bind.WatchOpts, sink chan<- *IAccountDidDeactivateDID) (event.Subscription, error) {

	logs, sub, err := _IAccountDid.contract.WatchLogs(opts, "DeactivateDID")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccountDidDeactivateDID)
				if err := _IAccountDid.contract.UnpackLog(event, "DeactivateDID", log); err != nil {
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

// ParseDeactivateDID is a log parse operation binding the contract event 0x2057c1bebefc71cf1116bc35105e72c4ec2baa1014a67f924edda51eee4565e8.
//
// Solidity: event DeactivateDID(string did)
func (_IAccountDid *IAccountDidFilterer) ParseDeactivateDID(log types.Log) (*IAccountDidDeactivateDID, error) {
	event := new(IAccountDidDeactivateDID)
	if err := _IAccountDid.contract.UnpackLog(event, "DeactivateDID", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAccountDidDeactivateVeriIterator is returned from FilterDeactivateVeri and is used to iterate over the raw logs and unpacked data for DeactivateVeri events raised by the IAccountDid contract.
type IAccountDidDeactivateVeriIterator struct {
	Event *IAccountDidDeactivateVeri // Event containing the contract specifics and raw log

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
func (it *IAccountDidDeactivateVeriIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccountDidDeactivateVeri)
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
		it.Event = new(IAccountDidDeactivateVeri)
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
func (it *IAccountDidDeactivateVeriIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccountDidDeactivateVeriIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccountDidDeactivateVeri represents a DeactivateVeri event raised by the IAccountDid contract.
type IAccountDidDeactivateVeri struct {
	Did common.Hash
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDeactivateVeri is a free log retrieval operation binding the contract event 0x671fd31386743fa44ce13617f8a8c357387671e6b9886278e67bbd2136d50ec6.
//
// Solidity: event DeactivateVeri(string indexed did, uint256 id)
func (_IAccountDid *IAccountDidFilterer) FilterDeactivateVeri(opts *bind.FilterOpts, did []string) (*IAccountDidDeactivateVeriIterator, error) {

	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _IAccountDid.contract.FilterLogs(opts, "DeactivateVeri", didRule)
	if err != nil {
		return nil, err
	}
	return &IAccountDidDeactivateVeriIterator{contract: _IAccountDid.contract, event: "DeactivateVeri", logs: logs, sub: sub}, nil
}

// WatchDeactivateVeri is a free log subscription operation binding the contract event 0x671fd31386743fa44ce13617f8a8c357387671e6b9886278e67bbd2136d50ec6.
//
// Solidity: event DeactivateVeri(string indexed did, uint256 id)
func (_IAccountDid *IAccountDidFilterer) WatchDeactivateVeri(opts *bind.WatchOpts, sink chan<- *IAccountDidDeactivateVeri, did []string) (event.Subscription, error) {

	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _IAccountDid.contract.WatchLogs(opts, "DeactivateVeri", didRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccountDidDeactivateVeri)
				if err := _IAccountDid.contract.UnpackLog(event, "DeactivateVeri", log); err != nil {
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

// ParseDeactivateVeri is a log parse operation binding the contract event 0x671fd31386743fa44ce13617f8a8c357387671e6b9886278e67bbd2136d50ec6.
//
// Solidity: event DeactivateVeri(string indexed did, uint256 id)
func (_IAccountDid *IAccountDidFilterer) ParseDeactivateVeri(log types.Log) (*IAccountDidDeactivateVeri, error) {
	event := new(IAccountDidDeactivateVeri)
	if err := _IAccountDid.contract.UnpackLog(event, "DeactivateVeri", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAccountDidRemoveAssertionIterator is returned from FilterRemoveAssertion and is used to iterate over the raw logs and unpacked data for RemoveAssertion events raised by the IAccountDid contract.
type IAccountDidRemoveAssertionIterator struct {
	Event *IAccountDidRemoveAssertion // Event containing the contract specifics and raw log

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
func (it *IAccountDidRemoveAssertionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccountDidRemoveAssertion)
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
		it.Event = new(IAccountDidRemoveAssertion)
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
func (it *IAccountDidRemoveAssertionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccountDidRemoveAssertionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccountDidRemoveAssertion represents a RemoveAssertion event raised by the IAccountDid contract.
type IAccountDidRemoveAssertion struct {
	Did common.Hash
	Id  string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRemoveAssertion is a free log retrieval operation binding the contract event 0x3e0a1fa337a30b1d302db25bc619c00f50e32e1ac195055458870ecf6a7662f4.
//
// Solidity: event RemoveAssertion(string indexed did, string id)
func (_IAccountDid *IAccountDidFilterer) FilterRemoveAssertion(opts *bind.FilterOpts, did []string) (*IAccountDidRemoveAssertionIterator, error) {

	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _IAccountDid.contract.FilterLogs(opts, "RemoveAssertion", didRule)
	if err != nil {
		return nil, err
	}
	return &IAccountDidRemoveAssertionIterator{contract: _IAccountDid.contract, event: "RemoveAssertion", logs: logs, sub: sub}, nil
}

// WatchRemoveAssertion is a free log subscription operation binding the contract event 0x3e0a1fa337a30b1d302db25bc619c00f50e32e1ac195055458870ecf6a7662f4.
//
// Solidity: event RemoveAssertion(string indexed did, string id)
func (_IAccountDid *IAccountDidFilterer) WatchRemoveAssertion(opts *bind.WatchOpts, sink chan<- *IAccountDidRemoveAssertion, did []string) (event.Subscription, error) {

	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _IAccountDid.contract.WatchLogs(opts, "RemoveAssertion", didRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccountDidRemoveAssertion)
				if err := _IAccountDid.contract.UnpackLog(event, "RemoveAssertion", log); err != nil {
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

// ParseRemoveAssertion is a log parse operation binding the contract event 0x3e0a1fa337a30b1d302db25bc619c00f50e32e1ac195055458870ecf6a7662f4.
//
// Solidity: event RemoveAssertion(string indexed did, string id)
func (_IAccountDid *IAccountDidFilterer) ParseRemoveAssertion(log types.Log) (*IAccountDidRemoveAssertion, error) {
	event := new(IAccountDidRemoveAssertion)
	if err := _IAccountDid.contract.UnpackLog(event, "RemoveAssertion", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAccountDidRemoveAuthIterator is returned from FilterRemoveAuth and is used to iterate over the raw logs and unpacked data for RemoveAuth events raised by the IAccountDid contract.
type IAccountDidRemoveAuthIterator struct {
	Event *IAccountDidRemoveAuth // Event containing the contract specifics and raw log

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
func (it *IAccountDidRemoveAuthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccountDidRemoveAuth)
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
		it.Event = new(IAccountDidRemoveAuth)
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
func (it *IAccountDidRemoveAuthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccountDidRemoveAuthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccountDidRemoveAuth represents a RemoveAuth event raised by the IAccountDid contract.
type IAccountDidRemoveAuth struct {
	Did common.Hash
	Id  string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRemoveAuth is a free log retrieval operation binding the contract event 0xf2d32726779f6b937767cdd53ca44340d0af922c6c637bdab215b304d902157b.
//
// Solidity: event RemoveAuth(string indexed did, string id)
func (_IAccountDid *IAccountDidFilterer) FilterRemoveAuth(opts *bind.FilterOpts, did []string) (*IAccountDidRemoveAuthIterator, error) {

	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _IAccountDid.contract.FilterLogs(opts, "RemoveAuth", didRule)
	if err != nil {
		return nil, err
	}
	return &IAccountDidRemoveAuthIterator{contract: _IAccountDid.contract, event: "RemoveAuth", logs: logs, sub: sub}, nil
}

// WatchRemoveAuth is a free log subscription operation binding the contract event 0xf2d32726779f6b937767cdd53ca44340d0af922c6c637bdab215b304d902157b.
//
// Solidity: event RemoveAuth(string indexed did, string id)
func (_IAccountDid *IAccountDidFilterer) WatchRemoveAuth(opts *bind.WatchOpts, sink chan<- *IAccountDidRemoveAuth, did []string) (event.Subscription, error) {

	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _IAccountDid.contract.WatchLogs(opts, "RemoveAuth", didRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccountDidRemoveAuth)
				if err := _IAccountDid.contract.UnpackLog(event, "RemoveAuth", log); err != nil {
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

// ParseRemoveAuth is a log parse operation binding the contract event 0xf2d32726779f6b937767cdd53ca44340d0af922c6c637bdab215b304d902157b.
//
// Solidity: event RemoveAuth(string indexed did, string id)
func (_IAccountDid *IAccountDidFilterer) ParseRemoveAuth(log types.Log) (*IAccountDidRemoveAuth, error) {
	event := new(IAccountDidRemoveAuth)
	if err := _IAccountDid.contract.UnpackLog(event, "RemoveAuth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAccountDidRemoveDelegationIterator is returned from FilterRemoveDelegation and is used to iterate over the raw logs and unpacked data for RemoveDelegation events raised by the IAccountDid contract.
type IAccountDidRemoveDelegationIterator struct {
	Event *IAccountDidRemoveDelegation // Event containing the contract specifics and raw log

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
func (it *IAccountDidRemoveDelegationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccountDidRemoveDelegation)
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
		it.Event = new(IAccountDidRemoveDelegation)
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
func (it *IAccountDidRemoveDelegationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccountDidRemoveDelegationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccountDidRemoveDelegation represents a RemoveDelegation event raised by the IAccountDid contract.
type IAccountDidRemoveDelegation struct {
	Did common.Hash
	Id  string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterRemoveDelegation is a free log retrieval operation binding the contract event 0xc8c5be74dbecbc654092affff41a0d8109cef3213923e597b547241dc355cf2e.
//
// Solidity: event RemoveDelegation(string indexed did, string id)
func (_IAccountDid *IAccountDidFilterer) FilterRemoveDelegation(opts *bind.FilterOpts, did []string) (*IAccountDidRemoveDelegationIterator, error) {

	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _IAccountDid.contract.FilterLogs(opts, "RemoveDelegation", didRule)
	if err != nil {
		return nil, err
	}
	return &IAccountDidRemoveDelegationIterator{contract: _IAccountDid.contract, event: "RemoveDelegation", logs: logs, sub: sub}, nil
}

// WatchRemoveDelegation is a free log subscription operation binding the contract event 0xc8c5be74dbecbc654092affff41a0d8109cef3213923e597b547241dc355cf2e.
//
// Solidity: event RemoveDelegation(string indexed did, string id)
func (_IAccountDid *IAccountDidFilterer) WatchRemoveDelegation(opts *bind.WatchOpts, sink chan<- *IAccountDidRemoveDelegation, did []string) (event.Subscription, error) {

	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _IAccountDid.contract.WatchLogs(opts, "RemoveDelegation", didRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccountDidRemoveDelegation)
				if err := _IAccountDid.contract.UnpackLog(event, "RemoveDelegation", log); err != nil {
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

// ParseRemoveDelegation is a log parse operation binding the contract event 0xc8c5be74dbecbc654092affff41a0d8109cef3213923e597b547241dc355cf2e.
//
// Solidity: event RemoveDelegation(string indexed did, string id)
func (_IAccountDid *IAccountDidFilterer) ParseRemoveDelegation(log types.Log) (*IAccountDidRemoveDelegation, error) {
	event := new(IAccountDidRemoveDelegation)
	if err := _IAccountDid.contract.UnpackLog(event, "RemoveDelegation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAccountDidRemoveRecoveryIterator is returned from FilterRemoveRecovery and is used to iterate over the raw logs and unpacked data for RemoveRecovery events raised by the IAccountDid contract.
type IAccountDidRemoveRecoveryIterator struct {
	Event *IAccountDidRemoveRecovery // Event containing the contract specifics and raw log

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
func (it *IAccountDidRemoveRecoveryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccountDidRemoveRecovery)
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
		it.Event = new(IAccountDidRemoveRecovery)
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
func (it *IAccountDidRemoveRecoveryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccountDidRemoveRecoveryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccountDidRemoveRecovery represents a RemoveRecovery event raised by the IAccountDid contract.
type IAccountDidRemoveRecovery struct {
	Did      common.Hash
	Recovery string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRemoveRecovery is a free log retrieval operation binding the contract event 0x40de7c6accd11e8883b27f5c2047d04862324bb82cc21180097aee971255fa40.
//
// Solidity: event RemoveRecovery(string indexed did, string recovery)
func (_IAccountDid *IAccountDidFilterer) FilterRemoveRecovery(opts *bind.FilterOpts, did []string) (*IAccountDidRemoveRecoveryIterator, error) {

	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _IAccountDid.contract.FilterLogs(opts, "RemoveRecovery", didRule)
	if err != nil {
		return nil, err
	}
	return &IAccountDidRemoveRecoveryIterator{contract: _IAccountDid.contract, event: "RemoveRecovery", logs: logs, sub: sub}, nil
}

// WatchRemoveRecovery is a free log subscription operation binding the contract event 0x40de7c6accd11e8883b27f5c2047d04862324bb82cc21180097aee971255fa40.
//
// Solidity: event RemoveRecovery(string indexed did, string recovery)
func (_IAccountDid *IAccountDidFilterer) WatchRemoveRecovery(opts *bind.WatchOpts, sink chan<- *IAccountDidRemoveRecovery, did []string) (event.Subscription, error) {

	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _IAccountDid.contract.WatchLogs(opts, "RemoveRecovery", didRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccountDidRemoveRecovery)
				if err := _IAccountDid.contract.UnpackLog(event, "RemoveRecovery", log); err != nil {
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

// ParseRemoveRecovery is a log parse operation binding the contract event 0x40de7c6accd11e8883b27f5c2047d04862324bb82cc21180097aee971255fa40.
//
// Solidity: event RemoveRecovery(string indexed did, string recovery)
func (_IAccountDid *IAccountDidFilterer) ParseRemoveRecovery(log types.Log) (*IAccountDidRemoveRecovery, error) {
	event := new(IAccountDidRemoveRecovery)
	if err := _IAccountDid.contract.UnpackLog(event, "RemoveRecovery", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IAccountDidUpdateVeriIterator is returned from FilterUpdateVeri and is used to iterate over the raw logs and unpacked data for UpdateVeri events raised by the IAccountDid contract.
type IAccountDidUpdateVeriIterator struct {
	Event *IAccountDidUpdateVeri // Event containing the contract specifics and raw log

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
func (it *IAccountDidUpdateVeriIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IAccountDidUpdateVeri)
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
		it.Event = new(IAccountDidUpdateVeri)
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
func (it *IAccountDidUpdateVeriIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IAccountDidUpdateVeriIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IAccountDidUpdateVeri represents a UpdateVeri event raised by the IAccountDid contract.
type IAccountDidUpdateVeri struct {
	Did common.Hash
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUpdateVeri is a free log retrieval operation binding the contract event 0x0b5c217a46ac23c52cf9217e0c43790834fe751d2e98488c23e5b79fe9c82efa.
//
// Solidity: event UpdateVeri(string indexed did, uint256 id)
func (_IAccountDid *IAccountDidFilterer) FilterUpdateVeri(opts *bind.FilterOpts, did []string) (*IAccountDidUpdateVeriIterator, error) {

	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _IAccountDid.contract.FilterLogs(opts, "UpdateVeri", didRule)
	if err != nil {
		return nil, err
	}
	return &IAccountDidUpdateVeriIterator{contract: _IAccountDid.contract, event: "UpdateVeri", logs: logs, sub: sub}, nil
}

// WatchUpdateVeri is a free log subscription operation binding the contract event 0x0b5c217a46ac23c52cf9217e0c43790834fe751d2e98488c23e5b79fe9c82efa.
//
// Solidity: event UpdateVeri(string indexed did, uint256 id)
func (_IAccountDid *IAccountDidFilterer) WatchUpdateVeri(opts *bind.WatchOpts, sink chan<- *IAccountDidUpdateVeri, did []string) (event.Subscription, error) {

	var didRule []interface{}
	for _, didItem := range did {
		didRule = append(didRule, didItem)
	}

	logs, sub, err := _IAccountDid.contract.WatchLogs(opts, "UpdateVeri", didRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IAccountDidUpdateVeri)
				if err := _IAccountDid.contract.UnpackLog(event, "UpdateVeri", log); err != nil {
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

// ParseUpdateVeri is a log parse operation binding the contract event 0x0b5c217a46ac23c52cf9217e0c43790834fe751d2e98488c23e5b79fe9c82efa.
//
// Solidity: event UpdateVeri(string indexed did, uint256 id)
func (_IAccountDid *IAccountDidFilterer) ParseUpdateVeri(log types.Log) (*IAccountDidUpdateVeri, error) {
	event := new(IAccountDidUpdateVeri)
	if err := _IAccountDid.contract.UnpackLog(event, "UpdateVeri", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IControlMetaData contains all meta data concerning the IControl contract.
var IControlMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"addAssertion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"addAuth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"addDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"addRecovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"methodType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"controller\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"pubKeyData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"deactivated\",\"type\":\"bool\"}],\"internalType\":\"structIAccountDid.PublicKey\",\"name\":\"pubKey\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"addVeri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"methodType\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"pubKeyData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"createDID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"methodType\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"pubKeyData\",\"type\":\"bytes\"}],\"name\":\"createDIDByAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"deactivate\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"deactivateDID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"deactivate\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"deactivateVeri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"removeAssertion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"removeAuth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"removeDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"removeRecovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"methodType\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"pubKeyData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"updateVeri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"62ea1b22": "addAssertion(string,string,bytes)",
		"a5b4eaed": "addAuth(string,string,bytes)",
		"d101ec86": "addDelegation(string,string,uint256,bytes)",
		"4b7897db": "addRecovery(string,string,bytes)",
		"81f2fd87": "addVeri(string,(string,string,bytes,bool),bytes)",
		"5e82d9df": "createDID(string,string,bytes,bytes)",
		"ac4f1d7a": "createDIDByAdmin(string,string,bytes)",
		"917a4b90": "deactivateDID(string,bool,bytes)",
		"21f33be8": "deactivateVeri(string,uint256,bool,bytes)",
		"55b7c6fe": "getNonce(string)",
		"9aae7467": "removeAssertion(string,string,bytes)",
		"eb785ee2": "removeAuth(string,string,bytes)",
		"d3be140e": "removeDelegation(string,string,bytes)",
		"dc649bf4": "removeRecovery(string,string,bytes)",
		"6c95dd03": "updateVeri(string,uint256,string,bytes,bytes)",
	},
}

// IControlABI is the input ABI used to generate the binding from.
// Deprecated: Use IControlMetaData.ABI instead.
var IControlABI = IControlMetaData.ABI

// Deprecated: Use IControlMetaData.Sigs instead.
// IControlFuncSigs maps the 4-byte function signature to its string representation.
var IControlFuncSigs = IControlMetaData.Sigs

// IControl is an auto generated Go binding around an Ethereum contract.
type IControl struct {
	IControlCaller     // Read-only binding to the contract
	IControlTransactor // Write-only binding to the contract
	IControlFilterer   // Log filterer for contract events
}

// IControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type IControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IControlSession struct {
	Contract     *IControl         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IControlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IControlCallerSession struct {
	Contract *IControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IControlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IControlTransactorSession struct {
	Contract     *IControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IControlRaw is an auto generated low-level Go binding around an Ethereum contract.
type IControlRaw struct {
	Contract *IControl // Generic contract binding to access the raw methods on
}

// IControlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IControlCallerRaw struct {
	Contract *IControlCaller // Generic read-only contract binding to access the raw methods on
}

// IControlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IControlTransactorRaw struct {
	Contract *IControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIControl creates a new instance of IControl, bound to a specific deployed contract.
func NewIControl(address common.Address, backend bind.ContractBackend) (*IControl, error) {
	contract, err := bindIControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IControl{IControlCaller: IControlCaller{contract: contract}, IControlTransactor: IControlTransactor{contract: contract}, IControlFilterer: IControlFilterer{contract: contract}}, nil
}

// NewIControlCaller creates a new read-only instance of IControl, bound to a specific deployed contract.
func NewIControlCaller(address common.Address, caller bind.ContractCaller) (*IControlCaller, error) {
	contract, err := bindIControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IControlCaller{contract: contract}, nil
}

// NewIControlTransactor creates a new write-only instance of IControl, bound to a specific deployed contract.
func NewIControlTransactor(address common.Address, transactor bind.ContractTransactor) (*IControlTransactor, error) {
	contract, err := bindIControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IControlTransactor{contract: contract}, nil
}

// NewIControlFilterer creates a new log filterer instance of IControl, bound to a specific deployed contract.
func NewIControlFilterer(address common.Address, filterer bind.ContractFilterer) (*IControlFilterer, error) {
	contract, err := bindIControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IControlFilterer{contract: contract}, nil
}

// bindIControl binds a generic wrapper to an already deployed contract.
func bindIControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IControlABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IControl *IControlRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IControl.Contract.IControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IControl *IControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IControl.Contract.IControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IControl *IControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IControl.Contract.IControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IControl *IControlCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IControl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IControl *IControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IControl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IControl *IControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IControl.Contract.contract.Transact(opts, method, params...)
}

// GetNonce is a free data retrieval call binding the contract method 0x55b7c6fe.
//
// Solidity: function getNonce(string did) view returns(uint64)
func (_IControl *IControlCaller) GetNonce(opts *bind.CallOpts, did string) (uint64, error) {
	var out []interface{}
	err := _IControl.contract.Call(opts, &out, "getNonce", did)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x55b7c6fe.
//
// Solidity: function getNonce(string did) view returns(uint64)
func (_IControl *IControlSession) GetNonce(did string) (uint64, error) {
	return _IControl.Contract.GetNonce(&_IControl.CallOpts, did)
}

// GetNonce is a free data retrieval call binding the contract method 0x55b7c6fe.
//
// Solidity: function getNonce(string did) view returns(uint64)
func (_IControl *IControlCallerSession) GetNonce(did string) (uint64, error) {
	return _IControl.Contract.GetNonce(&_IControl.CallOpts, did)
}

// AddAssertion is a paid mutator transaction binding the contract method 0x62ea1b22.
//
// Solidity: function addAssertion(string did, string id, bytes signature) returns()
func (_IControl *IControlTransactor) AddAssertion(opts *bind.TransactOpts, did string, id string, signature []byte) (*types.Transaction, error) {
	return _IControl.contract.Transact(opts, "addAssertion", did, id, signature)
}

// AddAssertion is a paid mutator transaction binding the contract method 0x62ea1b22.
//
// Solidity: function addAssertion(string did, string id, bytes signature) returns()
func (_IControl *IControlSession) AddAssertion(did string, id string, signature []byte) (*types.Transaction, error) {
	return _IControl.Contract.AddAssertion(&_IControl.TransactOpts, did, id, signature)
}

// AddAssertion is a paid mutator transaction binding the contract method 0x62ea1b22.
//
// Solidity: function addAssertion(string did, string id, bytes signature) returns()
func (_IControl *IControlTransactorSession) AddAssertion(did string, id string, signature []byte) (*types.Transaction, error) {
	return _IControl.Contract.AddAssertion(&_IControl.TransactOpts, did, id, signature)
}

// AddAuth is a paid mutator transaction binding the contract method 0xa5b4eaed.
//
// Solidity: function addAuth(string did, string id, bytes signature) returns()
func (_IControl *IControlTransactor) AddAuth(opts *bind.TransactOpts, did string, id string, signature []byte) (*types.Transaction, error) {
	return _IControl.contract.Transact(opts, "addAuth", did, id, signature)
}

// AddAuth is a paid mutator transaction binding the contract method 0xa5b4eaed.
//
// Solidity: function addAuth(string did, string id, bytes signature) returns()
func (_IControl *IControlSession) AddAuth(did string, id string, signature []byte) (*types.Transaction, error) {
	return _IControl.Contract.AddAuth(&_IControl.TransactOpts, did, id, signature)
}

// AddAuth is a paid mutator transaction binding the contract method 0xa5b4eaed.
//
// Solidity: function addAuth(string did, string id, bytes signature) returns()
func (_IControl *IControlTransactorSession) AddAuth(did string, id string, signature []byte) (*types.Transaction, error) {
	return _IControl.Contract.AddAuth(&_IControl.TransactOpts, did, id, signature)
}

// AddDelegation is a paid mutator transaction binding the contract method 0xd101ec86.
//
// Solidity: function addDelegation(string did, string id, uint256 expiration, bytes signature) returns()
func (_IControl *IControlTransactor) AddDelegation(opts *bind.TransactOpts, did string, id string, expiration *big.Int, signature []byte) (*types.Transaction, error) {
	return _IControl.contract.Transact(opts, "addDelegation", did, id, expiration, signature)
}

// AddDelegation is a paid mutator transaction binding the contract method 0xd101ec86.
//
// Solidity: function addDelegation(string did, string id, uint256 expiration, bytes signature) returns()
func (_IControl *IControlSession) AddDelegation(did string, id string, expiration *big.Int, signature []byte) (*types.Transaction, error) {
	return _IControl.Contract.AddDelegation(&_IControl.TransactOpts, did, id, expiration, signature)
}

// AddDelegation is a paid mutator transaction binding the contract method 0xd101ec86.
//
// Solidity: function addDelegation(string did, string id, uint256 expiration, bytes signature) returns()
func (_IControl *IControlTransactorSession) AddDelegation(did string, id string, expiration *big.Int, signature []byte) (*types.Transaction, error) {
	return _IControl.Contract.AddDelegation(&_IControl.TransactOpts, did, id, expiration, signature)
}

// AddRecovery is a paid mutator transaction binding the contract method 0x4b7897db.
//
// Solidity: function addRecovery(string did, string id, bytes signature) returns()
func (_IControl *IControlTransactor) AddRecovery(opts *bind.TransactOpts, did string, id string, signature []byte) (*types.Transaction, error) {
	return _IControl.contract.Transact(opts, "addRecovery", did, id, signature)
}

// AddRecovery is a paid mutator transaction binding the contract method 0x4b7897db.
//
// Solidity: function addRecovery(string did, string id, bytes signature) returns()
func (_IControl *IControlSession) AddRecovery(did string, id string, signature []byte) (*types.Transaction, error) {
	return _IControl.Contract.AddRecovery(&_IControl.TransactOpts, did, id, signature)
}

// AddRecovery is a paid mutator transaction binding the contract method 0x4b7897db.
//
// Solidity: function addRecovery(string did, string id, bytes signature) returns()
func (_IControl *IControlTransactorSession) AddRecovery(did string, id string, signature []byte) (*types.Transaction, error) {
	return _IControl.Contract.AddRecovery(&_IControl.TransactOpts, did, id, signature)
}

// AddVeri is a paid mutator transaction binding the contract method 0x81f2fd87.
//
// Solidity: function addVeri(string did, (string,string,bytes,bool) pubKey, bytes signature) returns()
func (_IControl *IControlTransactor) AddVeri(opts *bind.TransactOpts, did string, pubKey IAccountDidPublicKey, signature []byte) (*types.Transaction, error) {
	return _IControl.contract.Transact(opts, "addVeri", did, pubKey, signature)
}

// AddVeri is a paid mutator transaction binding the contract method 0x81f2fd87.
//
// Solidity: function addVeri(string did, (string,string,bytes,bool) pubKey, bytes signature) returns()
func (_IControl *IControlSession) AddVeri(did string, pubKey IAccountDidPublicKey, signature []byte) (*types.Transaction, error) {
	return _IControl.Contract.AddVeri(&_IControl.TransactOpts, did, pubKey, signature)
}

// AddVeri is a paid mutator transaction binding the contract method 0x81f2fd87.
//
// Solidity: function addVeri(string did, (string,string,bytes,bool) pubKey, bytes signature) returns()
func (_IControl *IControlTransactorSession) AddVeri(did string, pubKey IAccountDidPublicKey, signature []byte) (*types.Transaction, error) {
	return _IControl.Contract.AddVeri(&_IControl.TransactOpts, did, pubKey, signature)
}

// CreateDID is a paid mutator transaction binding the contract method 0x5e82d9df.
//
// Solidity: function createDID(string did, string methodType, bytes pubKeyData, bytes signature) returns()
func (_IControl *IControlTransactor) CreateDID(opts *bind.TransactOpts, did string, methodType string, pubKeyData []byte, signature []byte) (*types.Transaction, error) {
	return _IControl.contract.Transact(opts, "createDID", did, methodType, pubKeyData, signature)
}

// CreateDID is a paid mutator transaction binding the contract method 0x5e82d9df.
//
// Solidity: function createDID(string did, string methodType, bytes pubKeyData, bytes signature) returns()
func (_IControl *IControlSession) CreateDID(did string, methodType string, pubKeyData []byte, signature []byte) (*types.Transaction, error) {
	return _IControl.Contract.CreateDID(&_IControl.TransactOpts, did, methodType, pubKeyData, signature)
}

// CreateDID is a paid mutator transaction binding the contract method 0x5e82d9df.
//
// Solidity: function createDID(string did, string methodType, bytes pubKeyData, bytes signature) returns()
func (_IControl *IControlTransactorSession) CreateDID(did string, methodType string, pubKeyData []byte, signature []byte) (*types.Transaction, error) {
	return _IControl.Contract.CreateDID(&_IControl.TransactOpts, did, methodType, pubKeyData, signature)
}

// CreateDIDByAdmin is a paid mutator transaction binding the contract method 0xac4f1d7a.
//
// Solidity: function createDIDByAdmin(string did, string methodType, bytes pubKeyData) returns()
func (_IControl *IControlTransactor) CreateDIDByAdmin(opts *bind.TransactOpts, did string, methodType string, pubKeyData []byte) (*types.Transaction, error) {
	return _IControl.contract.Transact(opts, "createDIDByAdmin", did, methodType, pubKeyData)
}

// CreateDIDByAdmin is a paid mutator transaction binding the contract method 0xac4f1d7a.
//
// Solidity: function createDIDByAdmin(string did, string methodType, bytes pubKeyData) returns()
func (_IControl *IControlSession) CreateDIDByAdmin(did string, methodType string, pubKeyData []byte) (*types.Transaction, error) {
	return _IControl.Contract.CreateDIDByAdmin(&_IControl.TransactOpts, did, methodType, pubKeyData)
}

// CreateDIDByAdmin is a paid mutator transaction binding the contract method 0xac4f1d7a.
//
// Solidity: function createDIDByAdmin(string did, string methodType, bytes pubKeyData) returns()
func (_IControl *IControlTransactorSession) CreateDIDByAdmin(did string, methodType string, pubKeyData []byte) (*types.Transaction, error) {
	return _IControl.Contract.CreateDIDByAdmin(&_IControl.TransactOpts, did, methodType, pubKeyData)
}

// DeactivateDID is a paid mutator transaction binding the contract method 0x917a4b90.
//
// Solidity: function deactivateDID(string did, bool deactivate, bytes signature) returns()
func (_IControl *IControlTransactor) DeactivateDID(opts *bind.TransactOpts, did string, deactivate bool, signature []byte) (*types.Transaction, error) {
	return _IControl.contract.Transact(opts, "deactivateDID", did, deactivate, signature)
}

// DeactivateDID is a paid mutator transaction binding the contract method 0x917a4b90.
//
// Solidity: function deactivateDID(string did, bool deactivate, bytes signature) returns()
func (_IControl *IControlSession) DeactivateDID(did string, deactivate bool, signature []byte) (*types.Transaction, error) {
	return _IControl.Contract.DeactivateDID(&_IControl.TransactOpts, did, deactivate, signature)
}

// DeactivateDID is a paid mutator transaction binding the contract method 0x917a4b90.
//
// Solidity: function deactivateDID(string did, bool deactivate, bytes signature) returns()
func (_IControl *IControlTransactorSession) DeactivateDID(did string, deactivate bool, signature []byte) (*types.Transaction, error) {
	return _IControl.Contract.DeactivateDID(&_IControl.TransactOpts, did, deactivate, signature)
}

// DeactivateVeri is a paid mutator transaction binding the contract method 0x21f33be8.
//
// Solidity: function deactivateVeri(string did, uint256 index, bool deactivate, bytes signature) returns()
func (_IControl *IControlTransactor) DeactivateVeri(opts *bind.TransactOpts, did string, index *big.Int, deactivate bool, signature []byte) (*types.Transaction, error) {
	return _IControl.contract.Transact(opts, "deactivateVeri", did, index, deactivate, signature)
}

// DeactivateVeri is a paid mutator transaction binding the contract method 0x21f33be8.
//
// Solidity: function deactivateVeri(string did, uint256 index, bool deactivate, bytes signature) returns()
func (_IControl *IControlSession) DeactivateVeri(did string, index *big.Int, deactivate bool, signature []byte) (*types.Transaction, error) {
	return _IControl.Contract.DeactivateVeri(&_IControl.TransactOpts, did, index, deactivate, signature)
}

// DeactivateVeri is a paid mutator transaction binding the contract method 0x21f33be8.
//
// Solidity: function deactivateVeri(string did, uint256 index, bool deactivate, bytes signature) returns()
func (_IControl *IControlTransactorSession) DeactivateVeri(did string, index *big.Int, deactivate bool, signature []byte) (*types.Transaction, error) {
	return _IControl.Contract.DeactivateVeri(&_IControl.TransactOpts, did, index, deactivate, signature)
}

// RemoveAssertion is a paid mutator transaction binding the contract method 0x9aae7467.
//
// Solidity: function removeAssertion(string did, string id, bytes signature) returns()
func (_IControl *IControlTransactor) RemoveAssertion(opts *bind.TransactOpts, did string, id string, signature []byte) (*types.Transaction, error) {
	return _IControl.contract.Transact(opts, "removeAssertion", did, id, signature)
}

// RemoveAssertion is a paid mutator transaction binding the contract method 0x9aae7467.
//
// Solidity: function removeAssertion(string did, string id, bytes signature) returns()
func (_IControl *IControlSession) RemoveAssertion(did string, id string, signature []byte) (*types.Transaction, error) {
	return _IControl.Contract.RemoveAssertion(&_IControl.TransactOpts, did, id, signature)
}

// RemoveAssertion is a paid mutator transaction binding the contract method 0x9aae7467.
//
// Solidity: function removeAssertion(string did, string id, bytes signature) returns()
func (_IControl *IControlTransactorSession) RemoveAssertion(did string, id string, signature []byte) (*types.Transaction, error) {
	return _IControl.Contract.RemoveAssertion(&_IControl.TransactOpts, did, id, signature)
}

// RemoveAuth is a paid mutator transaction binding the contract method 0xeb785ee2.
//
// Solidity: function removeAuth(string did, string id, bytes signature) returns()
func (_IControl *IControlTransactor) RemoveAuth(opts *bind.TransactOpts, did string, id string, signature []byte) (*types.Transaction, error) {
	return _IControl.contract.Transact(opts, "removeAuth", did, id, signature)
}

// RemoveAuth is a paid mutator transaction binding the contract method 0xeb785ee2.
//
// Solidity: function removeAuth(string did, string id, bytes signature) returns()
func (_IControl *IControlSession) RemoveAuth(did string, id string, signature []byte) (*types.Transaction, error) {
	return _IControl.Contract.RemoveAuth(&_IControl.TransactOpts, did, id, signature)
}

// RemoveAuth is a paid mutator transaction binding the contract method 0xeb785ee2.
//
// Solidity: function removeAuth(string did, string id, bytes signature) returns()
func (_IControl *IControlTransactorSession) RemoveAuth(did string, id string, signature []byte) (*types.Transaction, error) {
	return _IControl.Contract.RemoveAuth(&_IControl.TransactOpts, did, id, signature)
}

// RemoveDelegation is a paid mutator transaction binding the contract method 0xd3be140e.
//
// Solidity: function removeDelegation(string did, string id, bytes signature) returns()
func (_IControl *IControlTransactor) RemoveDelegation(opts *bind.TransactOpts, did string, id string, signature []byte) (*types.Transaction, error) {
	return _IControl.contract.Transact(opts, "removeDelegation", did, id, signature)
}

// RemoveDelegation is a paid mutator transaction binding the contract method 0xd3be140e.
//
// Solidity: function removeDelegation(string did, string id, bytes signature) returns()
func (_IControl *IControlSession) RemoveDelegation(did string, id string, signature []byte) (*types.Transaction, error) {
	return _IControl.Contract.RemoveDelegation(&_IControl.TransactOpts, did, id, signature)
}

// RemoveDelegation is a paid mutator transaction binding the contract method 0xd3be140e.
//
// Solidity: function removeDelegation(string did, string id, bytes signature) returns()
func (_IControl *IControlTransactorSession) RemoveDelegation(did string, id string, signature []byte) (*types.Transaction, error) {
	return _IControl.Contract.RemoveDelegation(&_IControl.TransactOpts, did, id, signature)
}

// RemoveRecovery is a paid mutator transaction binding the contract method 0xdc649bf4.
//
// Solidity: function removeRecovery(string did, string id, bytes signature) returns()
func (_IControl *IControlTransactor) RemoveRecovery(opts *bind.TransactOpts, did string, id string, signature []byte) (*types.Transaction, error) {
	return _IControl.contract.Transact(opts, "removeRecovery", did, id, signature)
}

// RemoveRecovery is a paid mutator transaction binding the contract method 0xdc649bf4.
//
// Solidity: function removeRecovery(string did, string id, bytes signature) returns()
func (_IControl *IControlSession) RemoveRecovery(did string, id string, signature []byte) (*types.Transaction, error) {
	return _IControl.Contract.RemoveRecovery(&_IControl.TransactOpts, did, id, signature)
}

// RemoveRecovery is a paid mutator transaction binding the contract method 0xdc649bf4.
//
// Solidity: function removeRecovery(string did, string id, bytes signature) returns()
func (_IControl *IControlTransactorSession) RemoveRecovery(did string, id string, signature []byte) (*types.Transaction, error) {
	return _IControl.Contract.RemoveRecovery(&_IControl.TransactOpts, did, id, signature)
}

// UpdateVeri is a paid mutator transaction binding the contract method 0x6c95dd03.
//
// Solidity: function updateVeri(string did, uint256 index, string methodType, bytes pubKeyData, bytes signature) returns()
func (_IControl *IControlTransactor) UpdateVeri(opts *bind.TransactOpts, did string, index *big.Int, methodType string, pubKeyData []byte, signature []byte) (*types.Transaction, error) {
	return _IControl.contract.Transact(opts, "updateVeri", did, index, methodType, pubKeyData, signature)
}

// UpdateVeri is a paid mutator transaction binding the contract method 0x6c95dd03.
//
// Solidity: function updateVeri(string did, uint256 index, string methodType, bytes pubKeyData, bytes signature) returns()
func (_IControl *IControlSession) UpdateVeri(did string, index *big.Int, methodType string, pubKeyData []byte, signature []byte) (*types.Transaction, error) {
	return _IControl.Contract.UpdateVeri(&_IControl.TransactOpts, did, index, methodType, pubKeyData, signature)
}

// UpdateVeri is a paid mutator transaction binding the contract method 0x6c95dd03.
//
// Solidity: function updateVeri(string did, uint256 index, string methodType, bytes pubKeyData, bytes signature) returns()
func (_IControl *IControlTransactorSession) UpdateVeri(did string, index *big.Int, methodType string, pubKeyData []byte, signature []byte) (*types.Transaction, error) {
	return _IControl.Contract.UpdateVeri(&_IControl.TransactOpts, did, index, methodType, pubKeyData, signature)
}

// IControlFileDidMetaData contains all meta data concerning the IControlFileDid contract.
var IControlFileDidMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"memoDid\",\"type\":\"string\"}],\"name\":\"buyRead\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"controller\",\"type\":\"string\"}],\"name\":\"changeController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"internalType\":\"enumIFileDid.FileType\",\"name\":\"ftype\",\"type\":\"uint8\"}],\"name\":\"changeFtype\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"keywords\",\"type\":\"string[]\"}],\"name\":\"changeKeywords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"changePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"deactivate\",\"type\":\"bool\"}],\"name\":\"deactivateMfileDid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"memoDid\",\"type\":\"string\"}],\"name\":\"deactivateRead\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"memoDid\",\"type\":\"string\"}],\"name\":\"grantRead\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"encode\",\"type\":\"string\"},{\"internalType\":\"enumIFileDid.FileType\",\"name\":\"ftype\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"controller\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"keywords\",\"type\":\"string[]\"}],\"name\":\"registerMfileDid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"0069c140": "buyRead(string,string)",
		"9aad0566": "changeController(string,string)",
		"52a0fead": "changeFtype(string,uint8)",
		"64c21f2c": "changeKeywords(string,string[])",
		"1fa19596": "changePrice(string,uint256)",
		"86e9dbc6": "deactivateMfileDid(string,bool)",
		"37c52d83": "deactivateRead(string,string)",
		"a870fcc7": "grantRead(string,string)",
		"ca0258dc": "registerMfileDid(string,string,uint8,string,uint256,string[])",
	},
}

// IControlFileDidABI is the input ABI used to generate the binding from.
// Deprecated: Use IControlFileDidMetaData.ABI instead.
var IControlFileDidABI = IControlFileDidMetaData.ABI

// Deprecated: Use IControlFileDidMetaData.Sigs instead.
// IControlFileDidFuncSigs maps the 4-byte function signature to its string representation.
var IControlFileDidFuncSigs = IControlFileDidMetaData.Sigs

// IControlFileDid is an auto generated Go binding around an Ethereum contract.
type IControlFileDid struct {
	IControlFileDidCaller     // Read-only binding to the contract
	IControlFileDidTransactor // Write-only binding to the contract
	IControlFileDidFilterer   // Log filterer for contract events
}

// IControlFileDidCaller is an auto generated read-only Go binding around an Ethereum contract.
type IControlFileDidCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControlFileDidTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IControlFileDidTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControlFileDidFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IControlFileDidFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IControlFileDidSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IControlFileDidSession struct {
	Contract     *IControlFileDid  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IControlFileDidCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IControlFileDidCallerSession struct {
	Contract *IControlFileDidCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// IControlFileDidTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IControlFileDidTransactorSession struct {
	Contract     *IControlFileDidTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// IControlFileDidRaw is an auto generated low-level Go binding around an Ethereum contract.
type IControlFileDidRaw struct {
	Contract *IControlFileDid // Generic contract binding to access the raw methods on
}

// IControlFileDidCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IControlFileDidCallerRaw struct {
	Contract *IControlFileDidCaller // Generic read-only contract binding to access the raw methods on
}

// IControlFileDidTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IControlFileDidTransactorRaw struct {
	Contract *IControlFileDidTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIControlFileDid creates a new instance of IControlFileDid, bound to a specific deployed contract.
func NewIControlFileDid(address common.Address, backend bind.ContractBackend) (*IControlFileDid, error) {
	contract, err := bindIControlFileDid(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IControlFileDid{IControlFileDidCaller: IControlFileDidCaller{contract: contract}, IControlFileDidTransactor: IControlFileDidTransactor{contract: contract}, IControlFileDidFilterer: IControlFileDidFilterer{contract: contract}}, nil
}

// NewIControlFileDidCaller creates a new read-only instance of IControlFileDid, bound to a specific deployed contract.
func NewIControlFileDidCaller(address common.Address, caller bind.ContractCaller) (*IControlFileDidCaller, error) {
	contract, err := bindIControlFileDid(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IControlFileDidCaller{contract: contract}, nil
}

// NewIControlFileDidTransactor creates a new write-only instance of IControlFileDid, bound to a specific deployed contract.
func NewIControlFileDidTransactor(address common.Address, transactor bind.ContractTransactor) (*IControlFileDidTransactor, error) {
	contract, err := bindIControlFileDid(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IControlFileDidTransactor{contract: contract}, nil
}

// NewIControlFileDidFilterer creates a new log filterer instance of IControlFileDid, bound to a specific deployed contract.
func NewIControlFileDidFilterer(address common.Address, filterer bind.ContractFilterer) (*IControlFileDidFilterer, error) {
	contract, err := bindIControlFileDid(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IControlFileDidFilterer{contract: contract}, nil
}

// bindIControlFileDid binds a generic wrapper to an already deployed contract.
func bindIControlFileDid(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IControlFileDidABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IControlFileDid *IControlFileDidRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IControlFileDid.Contract.IControlFileDidCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IControlFileDid *IControlFileDidRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IControlFileDid.Contract.IControlFileDidTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IControlFileDid *IControlFileDidRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IControlFileDid.Contract.IControlFileDidTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IControlFileDid *IControlFileDidCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IControlFileDid.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IControlFileDid *IControlFileDidTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IControlFileDid.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IControlFileDid *IControlFileDidTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IControlFileDid.Contract.contract.Transact(opts, method, params...)
}

// BuyRead is a paid mutator transaction binding the contract method 0x0069c140.
//
// Solidity: function buyRead(string mfileDid, string memoDid) returns()
func (_IControlFileDid *IControlFileDidTransactor) BuyRead(opts *bind.TransactOpts, mfileDid string, memoDid string) (*types.Transaction, error) {
	return _IControlFileDid.contract.Transact(opts, "buyRead", mfileDid, memoDid)
}

// BuyRead is a paid mutator transaction binding the contract method 0x0069c140.
//
// Solidity: function buyRead(string mfileDid, string memoDid) returns()
func (_IControlFileDid *IControlFileDidSession) BuyRead(mfileDid string, memoDid string) (*types.Transaction, error) {
	return _IControlFileDid.Contract.BuyRead(&_IControlFileDid.TransactOpts, mfileDid, memoDid)
}

// BuyRead is a paid mutator transaction binding the contract method 0x0069c140.
//
// Solidity: function buyRead(string mfileDid, string memoDid) returns()
func (_IControlFileDid *IControlFileDidTransactorSession) BuyRead(mfileDid string, memoDid string) (*types.Transaction, error) {
	return _IControlFileDid.Contract.BuyRead(&_IControlFileDid.TransactOpts, mfileDid, memoDid)
}

// ChangeController is a paid mutator transaction binding the contract method 0x9aad0566.
//
// Solidity: function changeController(string mfileDid, string controller) returns()
func (_IControlFileDid *IControlFileDidTransactor) ChangeController(opts *bind.TransactOpts, mfileDid string, controller string) (*types.Transaction, error) {
	return _IControlFileDid.contract.Transact(opts, "changeController", mfileDid, controller)
}

// ChangeController is a paid mutator transaction binding the contract method 0x9aad0566.
//
// Solidity: function changeController(string mfileDid, string controller) returns()
func (_IControlFileDid *IControlFileDidSession) ChangeController(mfileDid string, controller string) (*types.Transaction, error) {
	return _IControlFileDid.Contract.ChangeController(&_IControlFileDid.TransactOpts, mfileDid, controller)
}

// ChangeController is a paid mutator transaction binding the contract method 0x9aad0566.
//
// Solidity: function changeController(string mfileDid, string controller) returns()
func (_IControlFileDid *IControlFileDidTransactorSession) ChangeController(mfileDid string, controller string) (*types.Transaction, error) {
	return _IControlFileDid.Contract.ChangeController(&_IControlFileDid.TransactOpts, mfileDid, controller)
}

// ChangeFtype is a paid mutator transaction binding the contract method 0x52a0fead.
//
// Solidity: function changeFtype(string mfileDid, uint8 ftype) returns()
func (_IControlFileDid *IControlFileDidTransactor) ChangeFtype(opts *bind.TransactOpts, mfileDid string, ftype uint8) (*types.Transaction, error) {
	return _IControlFileDid.contract.Transact(opts, "changeFtype", mfileDid, ftype)
}

// ChangeFtype is a paid mutator transaction binding the contract method 0x52a0fead.
//
// Solidity: function changeFtype(string mfileDid, uint8 ftype) returns()
func (_IControlFileDid *IControlFileDidSession) ChangeFtype(mfileDid string, ftype uint8) (*types.Transaction, error) {
	return _IControlFileDid.Contract.ChangeFtype(&_IControlFileDid.TransactOpts, mfileDid, ftype)
}

// ChangeFtype is a paid mutator transaction binding the contract method 0x52a0fead.
//
// Solidity: function changeFtype(string mfileDid, uint8 ftype) returns()
func (_IControlFileDid *IControlFileDidTransactorSession) ChangeFtype(mfileDid string, ftype uint8) (*types.Transaction, error) {
	return _IControlFileDid.Contract.ChangeFtype(&_IControlFileDid.TransactOpts, mfileDid, ftype)
}

// ChangeKeywords is a paid mutator transaction binding the contract method 0x64c21f2c.
//
// Solidity: function changeKeywords(string mfileDid, string[] keywords) returns()
func (_IControlFileDid *IControlFileDidTransactor) ChangeKeywords(opts *bind.TransactOpts, mfileDid string, keywords []string) (*types.Transaction, error) {
	return _IControlFileDid.contract.Transact(opts, "changeKeywords", mfileDid, keywords)
}

// ChangeKeywords is a paid mutator transaction binding the contract method 0x64c21f2c.
//
// Solidity: function changeKeywords(string mfileDid, string[] keywords) returns()
func (_IControlFileDid *IControlFileDidSession) ChangeKeywords(mfileDid string, keywords []string) (*types.Transaction, error) {
	return _IControlFileDid.Contract.ChangeKeywords(&_IControlFileDid.TransactOpts, mfileDid, keywords)
}

// ChangeKeywords is a paid mutator transaction binding the contract method 0x64c21f2c.
//
// Solidity: function changeKeywords(string mfileDid, string[] keywords) returns()
func (_IControlFileDid *IControlFileDidTransactorSession) ChangeKeywords(mfileDid string, keywords []string) (*types.Transaction, error) {
	return _IControlFileDid.Contract.ChangeKeywords(&_IControlFileDid.TransactOpts, mfileDid, keywords)
}

// ChangePrice is a paid mutator transaction binding the contract method 0x1fa19596.
//
// Solidity: function changePrice(string mfileDid, uint256 price) returns()
func (_IControlFileDid *IControlFileDidTransactor) ChangePrice(opts *bind.TransactOpts, mfileDid string, price *big.Int) (*types.Transaction, error) {
	return _IControlFileDid.contract.Transact(opts, "changePrice", mfileDid, price)
}

// ChangePrice is a paid mutator transaction binding the contract method 0x1fa19596.
//
// Solidity: function changePrice(string mfileDid, uint256 price) returns()
func (_IControlFileDid *IControlFileDidSession) ChangePrice(mfileDid string, price *big.Int) (*types.Transaction, error) {
	return _IControlFileDid.Contract.ChangePrice(&_IControlFileDid.TransactOpts, mfileDid, price)
}

// ChangePrice is a paid mutator transaction binding the contract method 0x1fa19596.
//
// Solidity: function changePrice(string mfileDid, uint256 price) returns()
func (_IControlFileDid *IControlFileDidTransactorSession) ChangePrice(mfileDid string, price *big.Int) (*types.Transaction, error) {
	return _IControlFileDid.Contract.ChangePrice(&_IControlFileDid.TransactOpts, mfileDid, price)
}

// DeactivateMfileDid is a paid mutator transaction binding the contract method 0x86e9dbc6.
//
// Solidity: function deactivateMfileDid(string mfileDid, bool deactivate) returns()
func (_IControlFileDid *IControlFileDidTransactor) DeactivateMfileDid(opts *bind.TransactOpts, mfileDid string, deactivate bool) (*types.Transaction, error) {
	return _IControlFileDid.contract.Transact(opts, "deactivateMfileDid", mfileDid, deactivate)
}

// DeactivateMfileDid is a paid mutator transaction binding the contract method 0x86e9dbc6.
//
// Solidity: function deactivateMfileDid(string mfileDid, bool deactivate) returns()
func (_IControlFileDid *IControlFileDidSession) DeactivateMfileDid(mfileDid string, deactivate bool) (*types.Transaction, error) {
	return _IControlFileDid.Contract.DeactivateMfileDid(&_IControlFileDid.TransactOpts, mfileDid, deactivate)
}

// DeactivateMfileDid is a paid mutator transaction binding the contract method 0x86e9dbc6.
//
// Solidity: function deactivateMfileDid(string mfileDid, bool deactivate) returns()
func (_IControlFileDid *IControlFileDidTransactorSession) DeactivateMfileDid(mfileDid string, deactivate bool) (*types.Transaction, error) {
	return _IControlFileDid.Contract.DeactivateMfileDid(&_IControlFileDid.TransactOpts, mfileDid, deactivate)
}

// DeactivateRead is a paid mutator transaction binding the contract method 0x37c52d83.
//
// Solidity: function deactivateRead(string mfileDid, string memoDid) returns()
func (_IControlFileDid *IControlFileDidTransactor) DeactivateRead(opts *bind.TransactOpts, mfileDid string, memoDid string) (*types.Transaction, error) {
	return _IControlFileDid.contract.Transact(opts, "deactivateRead", mfileDid, memoDid)
}

// DeactivateRead is a paid mutator transaction binding the contract method 0x37c52d83.
//
// Solidity: function deactivateRead(string mfileDid, string memoDid) returns()
func (_IControlFileDid *IControlFileDidSession) DeactivateRead(mfileDid string, memoDid string) (*types.Transaction, error) {
	return _IControlFileDid.Contract.DeactivateRead(&_IControlFileDid.TransactOpts, mfileDid, memoDid)
}

// DeactivateRead is a paid mutator transaction binding the contract method 0x37c52d83.
//
// Solidity: function deactivateRead(string mfileDid, string memoDid) returns()
func (_IControlFileDid *IControlFileDidTransactorSession) DeactivateRead(mfileDid string, memoDid string) (*types.Transaction, error) {
	return _IControlFileDid.Contract.DeactivateRead(&_IControlFileDid.TransactOpts, mfileDid, memoDid)
}

// GrantRead is a paid mutator transaction binding the contract method 0xa870fcc7.
//
// Solidity: function grantRead(string mfileDid, string memoDid) returns()
func (_IControlFileDid *IControlFileDidTransactor) GrantRead(opts *bind.TransactOpts, mfileDid string, memoDid string) (*types.Transaction, error) {
	return _IControlFileDid.contract.Transact(opts, "grantRead", mfileDid, memoDid)
}

// GrantRead is a paid mutator transaction binding the contract method 0xa870fcc7.
//
// Solidity: function grantRead(string mfileDid, string memoDid) returns()
func (_IControlFileDid *IControlFileDidSession) GrantRead(mfileDid string, memoDid string) (*types.Transaction, error) {
	return _IControlFileDid.Contract.GrantRead(&_IControlFileDid.TransactOpts, mfileDid, memoDid)
}

// GrantRead is a paid mutator transaction binding the contract method 0xa870fcc7.
//
// Solidity: function grantRead(string mfileDid, string memoDid) returns()
func (_IControlFileDid *IControlFileDidTransactorSession) GrantRead(mfileDid string, memoDid string) (*types.Transaction, error) {
	return _IControlFileDid.Contract.GrantRead(&_IControlFileDid.TransactOpts, mfileDid, memoDid)
}

// RegisterMfileDid is a paid mutator transaction binding the contract method 0xca0258dc.
//
// Solidity: function registerMfileDid(string mfileDid, string encode, uint8 ftype, string controller, uint256 price, string[] keywords) returns()
func (_IControlFileDid *IControlFileDidTransactor) RegisterMfileDid(opts *bind.TransactOpts, mfileDid string, encode string, ftype uint8, controller string, price *big.Int, keywords []string) (*types.Transaction, error) {
	return _IControlFileDid.contract.Transact(opts, "registerMfileDid", mfileDid, encode, ftype, controller, price, keywords)
}

// RegisterMfileDid is a paid mutator transaction binding the contract method 0xca0258dc.
//
// Solidity: function registerMfileDid(string mfileDid, string encode, uint8 ftype, string controller, uint256 price, string[] keywords) returns()
func (_IControlFileDid *IControlFileDidSession) RegisterMfileDid(mfileDid string, encode string, ftype uint8, controller string, price *big.Int, keywords []string) (*types.Transaction, error) {
	return _IControlFileDid.Contract.RegisterMfileDid(&_IControlFileDid.TransactOpts, mfileDid, encode, ftype, controller, price, keywords)
}

// RegisterMfileDid is a paid mutator transaction binding the contract method 0xca0258dc.
//
// Solidity: function registerMfileDid(string mfileDid, string encode, uint8 ftype, string controller, uint256 price, string[] keywords) returns()
func (_IControlFileDid *IControlFileDidTransactorSession) RegisterMfileDid(mfileDid string, encode string, ftype uint8, controller string, price *big.Int, keywords []string) (*types.Transaction, error) {
	return _IControlFileDid.Contract.RegisterMfileDid(&_IControlFileDid.TransactOpts, mfileDid, encode, ftype, controller, price, keywords)
}

// IFileDidMetaData contains all meta data concerning the IFileDid contract.
var IFileDidMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memoDid\",\"type\":\"string\"}],\"name\":\"BuyRead\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"controller\",\"type\":\"string\"}],\"name\":\"ChangeController\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"enumIFileDid.FileType\",\"name\":\"ftype\",\"type\":\"uint8\"}],\"name\":\"ChangeFtype\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"keywords\",\"type\":\"string[]\"}],\"name\":\"ChangeKeywords\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"ChangePrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"deactivate\",\"type\":\"bool\"}],\"name\":\"DeactivateMfileDid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memoDid\",\"type\":\"string\"}],\"name\":\"DeactivateRead\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"memoDid\",\"type\":\"string\"}],\"name\":\"GrantRead\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"}],\"name\":\"RegisterMfileDid\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"memoDid\",\"type\":\"string\"}],\"name\":\"buyRead\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"controller\",\"type\":\"string\"}],\"name\":\"changeController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"internalType\":\"enumIFileDid.FileType\",\"name\":\"ftype\",\"type\":\"uint8\"}],\"name\":\"changeFtype\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"keywords\",\"type\":\"string[]\"}],\"name\":\"changeKeywords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"changePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"deactivate\",\"type\":\"bool\"}],\"name\":\"deactivateMfileDid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"memoDid\",\"type\":\"string\"}],\"name\":\"deactivateRead\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"}],\"name\":\"deactivated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"}],\"name\":\"getController\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"}],\"name\":\"getEncode\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"}],\"name\":\"getFtype\",\"outputs\":[{\"internalType\":\"enumIFileDid.FileType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"}],\"name\":\"getKeywords\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"memoDid\",\"type\":\"string\"}],\"name\":\"grantRead\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"memoDid\",\"type\":\"string\"}],\"name\":\"read\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"encode\",\"type\":\"string\"},{\"internalType\":\"enumIFileDid.FileType\",\"name\":\"ftype\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"controller\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"keywords\",\"type\":\"string[]\"}],\"name\":\"registerMfileDid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"0069c140": "buyRead(string,string)",
		"9aad0566": "changeController(string,string)",
		"52a0fead": "changeFtype(string,uint8)",
		"64c21f2c": "changeKeywords(string,string[])",
		"1fa19596": "changePrice(string,uint256)",
		"86e9dbc6": "deactivateMfileDid(string,bool)",
		"37c52d83": "deactivateRead(string,string)",
		"84435ce2": "deactivated(string)",
		"63a27111": "getController(string)",
		"f100cfd3": "getEncode(string)",
		"dff5b013": "getFtype(string)",
		"380cc539": "getKeywords(string)",
		"524f3889": "getPrice(string)",
		"a870fcc7": "grantRead(string,string)",
		"8c97f99e": "read(string,string)",
		"ca0258dc": "registerMfileDid(string,string,uint8,string,uint256,string[])",
	},
}

// IFileDidABI is the input ABI used to generate the binding from.
// Deprecated: Use IFileDidMetaData.ABI instead.
var IFileDidABI = IFileDidMetaData.ABI

// Deprecated: Use IFileDidMetaData.Sigs instead.
// IFileDidFuncSigs maps the 4-byte function signature to its string representation.
var IFileDidFuncSigs = IFileDidMetaData.Sigs

// IFileDid is an auto generated Go binding around an Ethereum contract.
type IFileDid struct {
	IFileDidCaller     // Read-only binding to the contract
	IFileDidTransactor // Write-only binding to the contract
	IFileDidFilterer   // Log filterer for contract events
}

// IFileDidCaller is an auto generated read-only Go binding around an Ethereum contract.
type IFileDidCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFileDidTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IFileDidTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFileDidFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFileDidFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFileDidSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFileDidSession struct {
	Contract     *IFileDid         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IFileDidCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFileDidCallerSession struct {
	Contract *IFileDidCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IFileDidTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFileDidTransactorSession struct {
	Contract     *IFileDidTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IFileDidRaw is an auto generated low-level Go binding around an Ethereum contract.
type IFileDidRaw struct {
	Contract *IFileDid // Generic contract binding to access the raw methods on
}

// IFileDidCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFileDidCallerRaw struct {
	Contract *IFileDidCaller // Generic read-only contract binding to access the raw methods on
}

// IFileDidTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFileDidTransactorRaw struct {
	Contract *IFileDidTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIFileDid creates a new instance of IFileDid, bound to a specific deployed contract.
func NewIFileDid(address common.Address, backend bind.ContractBackend) (*IFileDid, error) {
	contract, err := bindIFileDid(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFileDid{IFileDidCaller: IFileDidCaller{contract: contract}, IFileDidTransactor: IFileDidTransactor{contract: contract}, IFileDidFilterer: IFileDidFilterer{contract: contract}}, nil
}

// NewIFileDidCaller creates a new read-only instance of IFileDid, bound to a specific deployed contract.
func NewIFileDidCaller(address common.Address, caller bind.ContractCaller) (*IFileDidCaller, error) {
	contract, err := bindIFileDid(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFileDidCaller{contract: contract}, nil
}

// NewIFileDidTransactor creates a new write-only instance of IFileDid, bound to a specific deployed contract.
func NewIFileDidTransactor(address common.Address, transactor bind.ContractTransactor) (*IFileDidTransactor, error) {
	contract, err := bindIFileDid(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFileDidTransactor{contract: contract}, nil
}

// NewIFileDidFilterer creates a new log filterer instance of IFileDid, bound to a specific deployed contract.
func NewIFileDidFilterer(address common.Address, filterer bind.ContractFilterer) (*IFileDidFilterer, error) {
	contract, err := bindIFileDid(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFileDidFilterer{contract: contract}, nil
}

// bindIFileDid binds a generic wrapper to an already deployed contract.
func bindIFileDid(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IFileDidABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFileDid *IFileDidRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFileDid.Contract.IFileDidCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFileDid *IFileDidRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFileDid.Contract.IFileDidTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFileDid *IFileDidRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFileDid.Contract.IFileDidTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFileDid *IFileDidCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFileDid.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFileDid *IFileDidTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFileDid.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFileDid *IFileDidTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFileDid.Contract.contract.Transact(opts, method, params...)
}

// Deactivated is a free data retrieval call binding the contract method 0x84435ce2.
//
// Solidity: function deactivated(string mfileDid) view returns(bool)
func (_IFileDid *IFileDidCaller) Deactivated(opts *bind.CallOpts, mfileDid string) (bool, error) {
	var out []interface{}
	err := _IFileDid.contract.Call(opts, &out, "deactivated", mfileDid)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Deactivated is a free data retrieval call binding the contract method 0x84435ce2.
//
// Solidity: function deactivated(string mfileDid) view returns(bool)
func (_IFileDid *IFileDidSession) Deactivated(mfileDid string) (bool, error) {
	return _IFileDid.Contract.Deactivated(&_IFileDid.CallOpts, mfileDid)
}

// Deactivated is a free data retrieval call binding the contract method 0x84435ce2.
//
// Solidity: function deactivated(string mfileDid) view returns(bool)
func (_IFileDid *IFileDidCallerSession) Deactivated(mfileDid string) (bool, error) {
	return _IFileDid.Contract.Deactivated(&_IFileDid.CallOpts, mfileDid)
}

// GetController is a free data retrieval call binding the contract method 0x63a27111.
//
// Solidity: function getController(string mfileDid) view returns(string)
func (_IFileDid *IFileDidCaller) GetController(opts *bind.CallOpts, mfileDid string) (string, error) {
	var out []interface{}
	err := _IFileDid.contract.Call(opts, &out, "getController", mfileDid)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetController is a free data retrieval call binding the contract method 0x63a27111.
//
// Solidity: function getController(string mfileDid) view returns(string)
func (_IFileDid *IFileDidSession) GetController(mfileDid string) (string, error) {
	return _IFileDid.Contract.GetController(&_IFileDid.CallOpts, mfileDid)
}

// GetController is a free data retrieval call binding the contract method 0x63a27111.
//
// Solidity: function getController(string mfileDid) view returns(string)
func (_IFileDid *IFileDidCallerSession) GetController(mfileDid string) (string, error) {
	return _IFileDid.Contract.GetController(&_IFileDid.CallOpts, mfileDid)
}

// GetEncode is a free data retrieval call binding the contract method 0xf100cfd3.
//
// Solidity: function getEncode(string mfileDid) view returns(string)
func (_IFileDid *IFileDidCaller) GetEncode(opts *bind.CallOpts, mfileDid string) (string, error) {
	var out []interface{}
	err := _IFileDid.contract.Call(opts, &out, "getEncode", mfileDid)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetEncode is a free data retrieval call binding the contract method 0xf100cfd3.
//
// Solidity: function getEncode(string mfileDid) view returns(string)
func (_IFileDid *IFileDidSession) GetEncode(mfileDid string) (string, error) {
	return _IFileDid.Contract.GetEncode(&_IFileDid.CallOpts, mfileDid)
}

// GetEncode is a free data retrieval call binding the contract method 0xf100cfd3.
//
// Solidity: function getEncode(string mfileDid) view returns(string)
func (_IFileDid *IFileDidCallerSession) GetEncode(mfileDid string) (string, error) {
	return _IFileDid.Contract.GetEncode(&_IFileDid.CallOpts, mfileDid)
}

// GetFtype is a free data retrieval call binding the contract method 0xdff5b013.
//
// Solidity: function getFtype(string mfileDid) view returns(uint8)
func (_IFileDid *IFileDidCaller) GetFtype(opts *bind.CallOpts, mfileDid string) (uint8, error) {
	var out []interface{}
	err := _IFileDid.contract.Call(opts, &out, "getFtype", mfileDid)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetFtype is a free data retrieval call binding the contract method 0xdff5b013.
//
// Solidity: function getFtype(string mfileDid) view returns(uint8)
func (_IFileDid *IFileDidSession) GetFtype(mfileDid string) (uint8, error) {
	return _IFileDid.Contract.GetFtype(&_IFileDid.CallOpts, mfileDid)
}

// GetFtype is a free data retrieval call binding the contract method 0xdff5b013.
//
// Solidity: function getFtype(string mfileDid) view returns(uint8)
func (_IFileDid *IFileDidCallerSession) GetFtype(mfileDid string) (uint8, error) {
	return _IFileDid.Contract.GetFtype(&_IFileDid.CallOpts, mfileDid)
}

// GetKeywords is a free data retrieval call binding the contract method 0x380cc539.
//
// Solidity: function getKeywords(string mfileDid) view returns(string[])
func (_IFileDid *IFileDidCaller) GetKeywords(opts *bind.CallOpts, mfileDid string) ([]string, error) {
	var out []interface{}
	err := _IFileDid.contract.Call(opts, &out, "getKeywords", mfileDid)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetKeywords is a free data retrieval call binding the contract method 0x380cc539.
//
// Solidity: function getKeywords(string mfileDid) view returns(string[])
func (_IFileDid *IFileDidSession) GetKeywords(mfileDid string) ([]string, error) {
	return _IFileDid.Contract.GetKeywords(&_IFileDid.CallOpts, mfileDid)
}

// GetKeywords is a free data retrieval call binding the contract method 0x380cc539.
//
// Solidity: function getKeywords(string mfileDid) view returns(string[])
func (_IFileDid *IFileDidCallerSession) GetKeywords(mfileDid string) ([]string, error) {
	return _IFileDid.Contract.GetKeywords(&_IFileDid.CallOpts, mfileDid)
}

// GetPrice is a free data retrieval call binding the contract method 0x524f3889.
//
// Solidity: function getPrice(string mfileDid) view returns(uint256)
func (_IFileDid *IFileDidCaller) GetPrice(opts *bind.CallOpts, mfileDid string) (*big.Int, error) {
	var out []interface{}
	err := _IFileDid.contract.Call(opts, &out, "getPrice", mfileDid)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x524f3889.
//
// Solidity: function getPrice(string mfileDid) view returns(uint256)
func (_IFileDid *IFileDidSession) GetPrice(mfileDid string) (*big.Int, error) {
	return _IFileDid.Contract.GetPrice(&_IFileDid.CallOpts, mfileDid)
}

// GetPrice is a free data retrieval call binding the contract method 0x524f3889.
//
// Solidity: function getPrice(string mfileDid) view returns(uint256)
func (_IFileDid *IFileDidCallerSession) GetPrice(mfileDid string) (*big.Int, error) {
	return _IFileDid.Contract.GetPrice(&_IFileDid.CallOpts, mfileDid)
}

// Read is a free data retrieval call binding the contract method 0x8c97f99e.
//
// Solidity: function read(string mfileDid, string memoDid) view returns(uint8)
func (_IFileDid *IFileDidCaller) Read(opts *bind.CallOpts, mfileDid string, memoDid string) (uint8, error) {
	var out []interface{}
	err := _IFileDid.contract.Call(opts, &out, "read", mfileDid, memoDid)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Read is a free data retrieval call binding the contract method 0x8c97f99e.
//
// Solidity: function read(string mfileDid, string memoDid) view returns(uint8)
func (_IFileDid *IFileDidSession) Read(mfileDid string, memoDid string) (uint8, error) {
	return _IFileDid.Contract.Read(&_IFileDid.CallOpts, mfileDid, memoDid)
}

// Read is a free data retrieval call binding the contract method 0x8c97f99e.
//
// Solidity: function read(string mfileDid, string memoDid) view returns(uint8)
func (_IFileDid *IFileDidCallerSession) Read(mfileDid string, memoDid string) (uint8, error) {
	return _IFileDid.Contract.Read(&_IFileDid.CallOpts, mfileDid, memoDid)
}

// BuyRead is a paid mutator transaction binding the contract method 0x0069c140.
//
// Solidity: function buyRead(string mfileDid, string memoDid) returns()
func (_IFileDid *IFileDidTransactor) BuyRead(opts *bind.TransactOpts, mfileDid string, memoDid string) (*types.Transaction, error) {
	return _IFileDid.contract.Transact(opts, "buyRead", mfileDid, memoDid)
}

// BuyRead is a paid mutator transaction binding the contract method 0x0069c140.
//
// Solidity: function buyRead(string mfileDid, string memoDid) returns()
func (_IFileDid *IFileDidSession) BuyRead(mfileDid string, memoDid string) (*types.Transaction, error) {
	return _IFileDid.Contract.BuyRead(&_IFileDid.TransactOpts, mfileDid, memoDid)
}

// BuyRead is a paid mutator transaction binding the contract method 0x0069c140.
//
// Solidity: function buyRead(string mfileDid, string memoDid) returns()
func (_IFileDid *IFileDidTransactorSession) BuyRead(mfileDid string, memoDid string) (*types.Transaction, error) {
	return _IFileDid.Contract.BuyRead(&_IFileDid.TransactOpts, mfileDid, memoDid)
}

// ChangeController is a paid mutator transaction binding the contract method 0x9aad0566.
//
// Solidity: function changeController(string mfileDid, string controller) returns()
func (_IFileDid *IFileDidTransactor) ChangeController(opts *bind.TransactOpts, mfileDid string, controller string) (*types.Transaction, error) {
	return _IFileDid.contract.Transact(opts, "changeController", mfileDid, controller)
}

// ChangeController is a paid mutator transaction binding the contract method 0x9aad0566.
//
// Solidity: function changeController(string mfileDid, string controller) returns()
func (_IFileDid *IFileDidSession) ChangeController(mfileDid string, controller string) (*types.Transaction, error) {
	return _IFileDid.Contract.ChangeController(&_IFileDid.TransactOpts, mfileDid, controller)
}

// ChangeController is a paid mutator transaction binding the contract method 0x9aad0566.
//
// Solidity: function changeController(string mfileDid, string controller) returns()
func (_IFileDid *IFileDidTransactorSession) ChangeController(mfileDid string, controller string) (*types.Transaction, error) {
	return _IFileDid.Contract.ChangeController(&_IFileDid.TransactOpts, mfileDid, controller)
}

// ChangeFtype is a paid mutator transaction binding the contract method 0x52a0fead.
//
// Solidity: function changeFtype(string mfileDid, uint8 ftype) returns()
func (_IFileDid *IFileDidTransactor) ChangeFtype(opts *bind.TransactOpts, mfileDid string, ftype uint8) (*types.Transaction, error) {
	return _IFileDid.contract.Transact(opts, "changeFtype", mfileDid, ftype)
}

// ChangeFtype is a paid mutator transaction binding the contract method 0x52a0fead.
//
// Solidity: function changeFtype(string mfileDid, uint8 ftype) returns()
func (_IFileDid *IFileDidSession) ChangeFtype(mfileDid string, ftype uint8) (*types.Transaction, error) {
	return _IFileDid.Contract.ChangeFtype(&_IFileDid.TransactOpts, mfileDid, ftype)
}

// ChangeFtype is a paid mutator transaction binding the contract method 0x52a0fead.
//
// Solidity: function changeFtype(string mfileDid, uint8 ftype) returns()
func (_IFileDid *IFileDidTransactorSession) ChangeFtype(mfileDid string, ftype uint8) (*types.Transaction, error) {
	return _IFileDid.Contract.ChangeFtype(&_IFileDid.TransactOpts, mfileDid, ftype)
}

// ChangeKeywords is a paid mutator transaction binding the contract method 0x64c21f2c.
//
// Solidity: function changeKeywords(string mfileDid, string[] keywords) returns()
func (_IFileDid *IFileDidTransactor) ChangeKeywords(opts *bind.TransactOpts, mfileDid string, keywords []string) (*types.Transaction, error) {
	return _IFileDid.contract.Transact(opts, "changeKeywords", mfileDid, keywords)
}

// ChangeKeywords is a paid mutator transaction binding the contract method 0x64c21f2c.
//
// Solidity: function changeKeywords(string mfileDid, string[] keywords) returns()
func (_IFileDid *IFileDidSession) ChangeKeywords(mfileDid string, keywords []string) (*types.Transaction, error) {
	return _IFileDid.Contract.ChangeKeywords(&_IFileDid.TransactOpts, mfileDid, keywords)
}

// ChangeKeywords is a paid mutator transaction binding the contract method 0x64c21f2c.
//
// Solidity: function changeKeywords(string mfileDid, string[] keywords) returns()
func (_IFileDid *IFileDidTransactorSession) ChangeKeywords(mfileDid string, keywords []string) (*types.Transaction, error) {
	return _IFileDid.Contract.ChangeKeywords(&_IFileDid.TransactOpts, mfileDid, keywords)
}

// ChangePrice is a paid mutator transaction binding the contract method 0x1fa19596.
//
// Solidity: function changePrice(string mfileDid, uint256 price) returns()
func (_IFileDid *IFileDidTransactor) ChangePrice(opts *bind.TransactOpts, mfileDid string, price *big.Int) (*types.Transaction, error) {
	return _IFileDid.contract.Transact(opts, "changePrice", mfileDid, price)
}

// ChangePrice is a paid mutator transaction binding the contract method 0x1fa19596.
//
// Solidity: function changePrice(string mfileDid, uint256 price) returns()
func (_IFileDid *IFileDidSession) ChangePrice(mfileDid string, price *big.Int) (*types.Transaction, error) {
	return _IFileDid.Contract.ChangePrice(&_IFileDid.TransactOpts, mfileDid, price)
}

// ChangePrice is a paid mutator transaction binding the contract method 0x1fa19596.
//
// Solidity: function changePrice(string mfileDid, uint256 price) returns()
func (_IFileDid *IFileDidTransactorSession) ChangePrice(mfileDid string, price *big.Int) (*types.Transaction, error) {
	return _IFileDid.Contract.ChangePrice(&_IFileDid.TransactOpts, mfileDid, price)
}

// DeactivateMfileDid is a paid mutator transaction binding the contract method 0x86e9dbc6.
//
// Solidity: function deactivateMfileDid(string mfileDid, bool deactivate) returns()
func (_IFileDid *IFileDidTransactor) DeactivateMfileDid(opts *bind.TransactOpts, mfileDid string, deactivate bool) (*types.Transaction, error) {
	return _IFileDid.contract.Transact(opts, "deactivateMfileDid", mfileDid, deactivate)
}

// DeactivateMfileDid is a paid mutator transaction binding the contract method 0x86e9dbc6.
//
// Solidity: function deactivateMfileDid(string mfileDid, bool deactivate) returns()
func (_IFileDid *IFileDidSession) DeactivateMfileDid(mfileDid string, deactivate bool) (*types.Transaction, error) {
	return _IFileDid.Contract.DeactivateMfileDid(&_IFileDid.TransactOpts, mfileDid, deactivate)
}

// DeactivateMfileDid is a paid mutator transaction binding the contract method 0x86e9dbc6.
//
// Solidity: function deactivateMfileDid(string mfileDid, bool deactivate) returns()
func (_IFileDid *IFileDidTransactorSession) DeactivateMfileDid(mfileDid string, deactivate bool) (*types.Transaction, error) {
	return _IFileDid.Contract.DeactivateMfileDid(&_IFileDid.TransactOpts, mfileDid, deactivate)
}

// DeactivateRead is a paid mutator transaction binding the contract method 0x37c52d83.
//
// Solidity: function deactivateRead(string mfileDid, string memoDid) returns()
func (_IFileDid *IFileDidTransactor) DeactivateRead(opts *bind.TransactOpts, mfileDid string, memoDid string) (*types.Transaction, error) {
	return _IFileDid.contract.Transact(opts, "deactivateRead", mfileDid, memoDid)
}

// DeactivateRead is a paid mutator transaction binding the contract method 0x37c52d83.
//
// Solidity: function deactivateRead(string mfileDid, string memoDid) returns()
func (_IFileDid *IFileDidSession) DeactivateRead(mfileDid string, memoDid string) (*types.Transaction, error) {
	return _IFileDid.Contract.DeactivateRead(&_IFileDid.TransactOpts, mfileDid, memoDid)
}

// DeactivateRead is a paid mutator transaction binding the contract method 0x37c52d83.
//
// Solidity: function deactivateRead(string mfileDid, string memoDid) returns()
func (_IFileDid *IFileDidTransactorSession) DeactivateRead(mfileDid string, memoDid string) (*types.Transaction, error) {
	return _IFileDid.Contract.DeactivateRead(&_IFileDid.TransactOpts, mfileDid, memoDid)
}

// GrantRead is a paid mutator transaction binding the contract method 0xa870fcc7.
//
// Solidity: function grantRead(string mfileDid, string memoDid) returns()
func (_IFileDid *IFileDidTransactor) GrantRead(opts *bind.TransactOpts, mfileDid string, memoDid string) (*types.Transaction, error) {
	return _IFileDid.contract.Transact(opts, "grantRead", mfileDid, memoDid)
}

// GrantRead is a paid mutator transaction binding the contract method 0xa870fcc7.
//
// Solidity: function grantRead(string mfileDid, string memoDid) returns()
func (_IFileDid *IFileDidSession) GrantRead(mfileDid string, memoDid string) (*types.Transaction, error) {
	return _IFileDid.Contract.GrantRead(&_IFileDid.TransactOpts, mfileDid, memoDid)
}

// GrantRead is a paid mutator transaction binding the contract method 0xa870fcc7.
//
// Solidity: function grantRead(string mfileDid, string memoDid) returns()
func (_IFileDid *IFileDidTransactorSession) GrantRead(mfileDid string, memoDid string) (*types.Transaction, error) {
	return _IFileDid.Contract.GrantRead(&_IFileDid.TransactOpts, mfileDid, memoDid)
}

// RegisterMfileDid is a paid mutator transaction binding the contract method 0xca0258dc.
//
// Solidity: function registerMfileDid(string mfileDid, string encode, uint8 ftype, string controller, uint256 price, string[] keywords) returns()
func (_IFileDid *IFileDidTransactor) RegisterMfileDid(opts *bind.TransactOpts, mfileDid string, encode string, ftype uint8, controller string, price *big.Int, keywords []string) (*types.Transaction, error) {
	return _IFileDid.contract.Transact(opts, "registerMfileDid", mfileDid, encode, ftype, controller, price, keywords)
}

// RegisterMfileDid is a paid mutator transaction binding the contract method 0xca0258dc.
//
// Solidity: function registerMfileDid(string mfileDid, string encode, uint8 ftype, string controller, uint256 price, string[] keywords) returns()
func (_IFileDid *IFileDidSession) RegisterMfileDid(mfileDid string, encode string, ftype uint8, controller string, price *big.Int, keywords []string) (*types.Transaction, error) {
	return _IFileDid.Contract.RegisterMfileDid(&_IFileDid.TransactOpts, mfileDid, encode, ftype, controller, price, keywords)
}

// RegisterMfileDid is a paid mutator transaction binding the contract method 0xca0258dc.
//
// Solidity: function registerMfileDid(string mfileDid, string encode, uint8 ftype, string controller, uint256 price, string[] keywords) returns()
func (_IFileDid *IFileDidTransactorSession) RegisterMfileDid(mfileDid string, encode string, ftype uint8, controller string, price *big.Int, keywords []string) (*types.Transaction, error) {
	return _IFileDid.Contract.RegisterMfileDid(&_IFileDid.TransactOpts, mfileDid, encode, ftype, controller, price, keywords)
}

// IFileDidBuyReadIterator is returned from FilterBuyRead and is used to iterate over the raw logs and unpacked data for BuyRead events raised by the IFileDid contract.
type IFileDidBuyReadIterator struct {
	Event *IFileDidBuyRead // Event containing the contract specifics and raw log

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
func (it *IFileDidBuyReadIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFileDidBuyRead)
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
		it.Event = new(IFileDidBuyRead)
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
func (it *IFileDidBuyReadIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFileDidBuyReadIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFileDidBuyRead represents a BuyRead event raised by the IFileDid contract.
type IFileDidBuyRead struct {
	MfileDid common.Hash
	MemoDid  string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBuyRead is a free log retrieval operation binding the contract event 0x9a23ee694031b7714b282fb4f89b349db581b6c5fb20fb299897b2dbb5b6510a.
//
// Solidity: event BuyRead(string indexed mfileDid, string memoDid)
func (_IFileDid *IFileDidFilterer) FilterBuyRead(opts *bind.FilterOpts, mfileDid []string) (*IFileDidBuyReadIterator, error) {

	var mfileDidRule []interface{}
	for _, mfileDidItem := range mfileDid {
		mfileDidRule = append(mfileDidRule, mfileDidItem)
	}

	logs, sub, err := _IFileDid.contract.FilterLogs(opts, "BuyRead", mfileDidRule)
	if err != nil {
		return nil, err
	}
	return &IFileDidBuyReadIterator{contract: _IFileDid.contract, event: "BuyRead", logs: logs, sub: sub}, nil
}

// WatchBuyRead is a free log subscription operation binding the contract event 0x9a23ee694031b7714b282fb4f89b349db581b6c5fb20fb299897b2dbb5b6510a.
//
// Solidity: event BuyRead(string indexed mfileDid, string memoDid)
func (_IFileDid *IFileDidFilterer) WatchBuyRead(opts *bind.WatchOpts, sink chan<- *IFileDidBuyRead, mfileDid []string) (event.Subscription, error) {

	var mfileDidRule []interface{}
	for _, mfileDidItem := range mfileDid {
		mfileDidRule = append(mfileDidRule, mfileDidItem)
	}

	logs, sub, err := _IFileDid.contract.WatchLogs(opts, "BuyRead", mfileDidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFileDidBuyRead)
				if err := _IFileDid.contract.UnpackLog(event, "BuyRead", log); err != nil {
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

// ParseBuyRead is a log parse operation binding the contract event 0x9a23ee694031b7714b282fb4f89b349db581b6c5fb20fb299897b2dbb5b6510a.
//
// Solidity: event BuyRead(string indexed mfileDid, string memoDid)
func (_IFileDid *IFileDidFilterer) ParseBuyRead(log types.Log) (*IFileDidBuyRead, error) {
	event := new(IFileDidBuyRead)
	if err := _IFileDid.contract.UnpackLog(event, "BuyRead", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFileDidChangeControllerIterator is returned from FilterChangeController and is used to iterate over the raw logs and unpacked data for ChangeController events raised by the IFileDid contract.
type IFileDidChangeControllerIterator struct {
	Event *IFileDidChangeController // Event containing the contract specifics and raw log

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
func (it *IFileDidChangeControllerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFileDidChangeController)
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
		it.Event = new(IFileDidChangeController)
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
func (it *IFileDidChangeControllerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFileDidChangeControllerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFileDidChangeController represents a ChangeController event raised by the IFileDid contract.
type IFileDidChangeController struct {
	MfileDid   common.Hash
	Controller string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterChangeController is a free log retrieval operation binding the contract event 0xad2c5e83d2fbe2975f5ba70ab5c8c351ef2da897e41ce843a2eb6de9ffbf59ee.
//
// Solidity: event ChangeController(string indexed mfileDid, string controller)
func (_IFileDid *IFileDidFilterer) FilterChangeController(opts *bind.FilterOpts, mfileDid []string) (*IFileDidChangeControllerIterator, error) {

	var mfileDidRule []interface{}
	for _, mfileDidItem := range mfileDid {
		mfileDidRule = append(mfileDidRule, mfileDidItem)
	}

	logs, sub, err := _IFileDid.contract.FilterLogs(opts, "ChangeController", mfileDidRule)
	if err != nil {
		return nil, err
	}
	return &IFileDidChangeControllerIterator{contract: _IFileDid.contract, event: "ChangeController", logs: logs, sub: sub}, nil
}

// WatchChangeController is a free log subscription operation binding the contract event 0xad2c5e83d2fbe2975f5ba70ab5c8c351ef2da897e41ce843a2eb6de9ffbf59ee.
//
// Solidity: event ChangeController(string indexed mfileDid, string controller)
func (_IFileDid *IFileDidFilterer) WatchChangeController(opts *bind.WatchOpts, sink chan<- *IFileDidChangeController, mfileDid []string) (event.Subscription, error) {

	var mfileDidRule []interface{}
	for _, mfileDidItem := range mfileDid {
		mfileDidRule = append(mfileDidRule, mfileDidItem)
	}

	logs, sub, err := _IFileDid.contract.WatchLogs(opts, "ChangeController", mfileDidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFileDidChangeController)
				if err := _IFileDid.contract.UnpackLog(event, "ChangeController", log); err != nil {
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

// ParseChangeController is a log parse operation binding the contract event 0xad2c5e83d2fbe2975f5ba70ab5c8c351ef2da897e41ce843a2eb6de9ffbf59ee.
//
// Solidity: event ChangeController(string indexed mfileDid, string controller)
func (_IFileDid *IFileDidFilterer) ParseChangeController(log types.Log) (*IFileDidChangeController, error) {
	event := new(IFileDidChangeController)
	if err := _IFileDid.contract.UnpackLog(event, "ChangeController", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFileDidChangeFtypeIterator is returned from FilterChangeFtype and is used to iterate over the raw logs and unpacked data for ChangeFtype events raised by the IFileDid contract.
type IFileDidChangeFtypeIterator struct {
	Event *IFileDidChangeFtype // Event containing the contract specifics and raw log

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
func (it *IFileDidChangeFtypeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFileDidChangeFtype)
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
		it.Event = new(IFileDidChangeFtype)
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
func (it *IFileDidChangeFtypeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFileDidChangeFtypeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFileDidChangeFtype represents a ChangeFtype event raised by the IFileDid contract.
type IFileDidChangeFtype struct {
	MfileDid common.Hash
	Ftype    uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChangeFtype is a free log retrieval operation binding the contract event 0x935721402a5ae34f86d9500f87d6acfbef247c62cc5c845f49e7ec94a0dd4896.
//
// Solidity: event ChangeFtype(string indexed mfileDid, uint8 ftype)
func (_IFileDid *IFileDidFilterer) FilterChangeFtype(opts *bind.FilterOpts, mfileDid []string) (*IFileDidChangeFtypeIterator, error) {

	var mfileDidRule []interface{}
	for _, mfileDidItem := range mfileDid {
		mfileDidRule = append(mfileDidRule, mfileDidItem)
	}

	logs, sub, err := _IFileDid.contract.FilterLogs(opts, "ChangeFtype", mfileDidRule)
	if err != nil {
		return nil, err
	}
	return &IFileDidChangeFtypeIterator{contract: _IFileDid.contract, event: "ChangeFtype", logs: logs, sub: sub}, nil
}

// WatchChangeFtype is a free log subscription operation binding the contract event 0x935721402a5ae34f86d9500f87d6acfbef247c62cc5c845f49e7ec94a0dd4896.
//
// Solidity: event ChangeFtype(string indexed mfileDid, uint8 ftype)
func (_IFileDid *IFileDidFilterer) WatchChangeFtype(opts *bind.WatchOpts, sink chan<- *IFileDidChangeFtype, mfileDid []string) (event.Subscription, error) {

	var mfileDidRule []interface{}
	for _, mfileDidItem := range mfileDid {
		mfileDidRule = append(mfileDidRule, mfileDidItem)
	}

	logs, sub, err := _IFileDid.contract.WatchLogs(opts, "ChangeFtype", mfileDidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFileDidChangeFtype)
				if err := _IFileDid.contract.UnpackLog(event, "ChangeFtype", log); err != nil {
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

// ParseChangeFtype is a log parse operation binding the contract event 0x935721402a5ae34f86d9500f87d6acfbef247c62cc5c845f49e7ec94a0dd4896.
//
// Solidity: event ChangeFtype(string indexed mfileDid, uint8 ftype)
func (_IFileDid *IFileDidFilterer) ParseChangeFtype(log types.Log) (*IFileDidChangeFtype, error) {
	event := new(IFileDidChangeFtype)
	if err := _IFileDid.contract.UnpackLog(event, "ChangeFtype", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFileDidChangeKeywordsIterator is returned from FilterChangeKeywords and is used to iterate over the raw logs and unpacked data for ChangeKeywords events raised by the IFileDid contract.
type IFileDidChangeKeywordsIterator struct {
	Event *IFileDidChangeKeywords // Event containing the contract specifics and raw log

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
func (it *IFileDidChangeKeywordsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFileDidChangeKeywords)
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
		it.Event = new(IFileDidChangeKeywords)
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
func (it *IFileDidChangeKeywordsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFileDidChangeKeywordsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFileDidChangeKeywords represents a ChangeKeywords event raised by the IFileDid contract.
type IFileDidChangeKeywords struct {
	MfileDid common.Hash
	Keywords []string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChangeKeywords is a free log retrieval operation binding the contract event 0x0d0ecbde8ff7f0cbfa8a5aba6209011f194a2ab60ceb48e047f704427ebbed2e.
//
// Solidity: event ChangeKeywords(string indexed mfileDid, string[] keywords)
func (_IFileDid *IFileDidFilterer) FilterChangeKeywords(opts *bind.FilterOpts, mfileDid []string) (*IFileDidChangeKeywordsIterator, error) {

	var mfileDidRule []interface{}
	for _, mfileDidItem := range mfileDid {
		mfileDidRule = append(mfileDidRule, mfileDidItem)
	}

	logs, sub, err := _IFileDid.contract.FilterLogs(opts, "ChangeKeywords", mfileDidRule)
	if err != nil {
		return nil, err
	}
	return &IFileDidChangeKeywordsIterator{contract: _IFileDid.contract, event: "ChangeKeywords", logs: logs, sub: sub}, nil
}

// WatchChangeKeywords is a free log subscription operation binding the contract event 0x0d0ecbde8ff7f0cbfa8a5aba6209011f194a2ab60ceb48e047f704427ebbed2e.
//
// Solidity: event ChangeKeywords(string indexed mfileDid, string[] keywords)
func (_IFileDid *IFileDidFilterer) WatchChangeKeywords(opts *bind.WatchOpts, sink chan<- *IFileDidChangeKeywords, mfileDid []string) (event.Subscription, error) {

	var mfileDidRule []interface{}
	for _, mfileDidItem := range mfileDid {
		mfileDidRule = append(mfileDidRule, mfileDidItem)
	}

	logs, sub, err := _IFileDid.contract.WatchLogs(opts, "ChangeKeywords", mfileDidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFileDidChangeKeywords)
				if err := _IFileDid.contract.UnpackLog(event, "ChangeKeywords", log); err != nil {
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

// ParseChangeKeywords is a log parse operation binding the contract event 0x0d0ecbde8ff7f0cbfa8a5aba6209011f194a2ab60ceb48e047f704427ebbed2e.
//
// Solidity: event ChangeKeywords(string indexed mfileDid, string[] keywords)
func (_IFileDid *IFileDidFilterer) ParseChangeKeywords(log types.Log) (*IFileDidChangeKeywords, error) {
	event := new(IFileDidChangeKeywords)
	if err := _IFileDid.contract.UnpackLog(event, "ChangeKeywords", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFileDidChangePriceIterator is returned from FilterChangePrice and is used to iterate over the raw logs and unpacked data for ChangePrice events raised by the IFileDid contract.
type IFileDidChangePriceIterator struct {
	Event *IFileDidChangePrice // Event containing the contract specifics and raw log

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
func (it *IFileDidChangePriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFileDidChangePrice)
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
		it.Event = new(IFileDidChangePrice)
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
func (it *IFileDidChangePriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFileDidChangePriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFileDidChangePrice represents a ChangePrice event raised by the IFileDid contract.
type IFileDidChangePrice struct {
	MfileDid common.Hash
	Price    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterChangePrice is a free log retrieval operation binding the contract event 0x0868ddc93d8bff1d2b8931574f1250319ab20dc6cdf2a7f86d5eee192ea1c986.
//
// Solidity: event ChangePrice(string indexed mfileDid, uint256 price)
func (_IFileDid *IFileDidFilterer) FilterChangePrice(opts *bind.FilterOpts, mfileDid []string) (*IFileDidChangePriceIterator, error) {

	var mfileDidRule []interface{}
	for _, mfileDidItem := range mfileDid {
		mfileDidRule = append(mfileDidRule, mfileDidItem)
	}

	logs, sub, err := _IFileDid.contract.FilterLogs(opts, "ChangePrice", mfileDidRule)
	if err != nil {
		return nil, err
	}
	return &IFileDidChangePriceIterator{contract: _IFileDid.contract, event: "ChangePrice", logs: logs, sub: sub}, nil
}

// WatchChangePrice is a free log subscription operation binding the contract event 0x0868ddc93d8bff1d2b8931574f1250319ab20dc6cdf2a7f86d5eee192ea1c986.
//
// Solidity: event ChangePrice(string indexed mfileDid, uint256 price)
func (_IFileDid *IFileDidFilterer) WatchChangePrice(opts *bind.WatchOpts, sink chan<- *IFileDidChangePrice, mfileDid []string) (event.Subscription, error) {

	var mfileDidRule []interface{}
	for _, mfileDidItem := range mfileDid {
		mfileDidRule = append(mfileDidRule, mfileDidItem)
	}

	logs, sub, err := _IFileDid.contract.WatchLogs(opts, "ChangePrice", mfileDidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFileDidChangePrice)
				if err := _IFileDid.contract.UnpackLog(event, "ChangePrice", log); err != nil {
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

// ParseChangePrice is a log parse operation binding the contract event 0x0868ddc93d8bff1d2b8931574f1250319ab20dc6cdf2a7f86d5eee192ea1c986.
//
// Solidity: event ChangePrice(string indexed mfileDid, uint256 price)
func (_IFileDid *IFileDidFilterer) ParseChangePrice(log types.Log) (*IFileDidChangePrice, error) {
	event := new(IFileDidChangePrice)
	if err := _IFileDid.contract.UnpackLog(event, "ChangePrice", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFileDidDeactivateMfileDidIterator is returned from FilterDeactivateMfileDid and is used to iterate over the raw logs and unpacked data for DeactivateMfileDid events raised by the IFileDid contract.
type IFileDidDeactivateMfileDidIterator struct {
	Event *IFileDidDeactivateMfileDid // Event containing the contract specifics and raw log

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
func (it *IFileDidDeactivateMfileDidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFileDidDeactivateMfileDid)
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
		it.Event = new(IFileDidDeactivateMfileDid)
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
func (it *IFileDidDeactivateMfileDidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFileDidDeactivateMfileDidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFileDidDeactivateMfileDid represents a DeactivateMfileDid event raised by the IFileDid contract.
type IFileDidDeactivateMfileDid struct {
	MfileDid   string
	Deactivate bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDeactivateMfileDid is a free log retrieval operation binding the contract event 0x5e17ae7a2d292f907502879bb1abf12db289923734c91a05b5d0651531aebf2a.
//
// Solidity: event DeactivateMfileDid(string mfileDid, bool deactivate)
func (_IFileDid *IFileDidFilterer) FilterDeactivateMfileDid(opts *bind.FilterOpts) (*IFileDidDeactivateMfileDidIterator, error) {

	logs, sub, err := _IFileDid.contract.FilterLogs(opts, "DeactivateMfileDid")
	if err != nil {
		return nil, err
	}
	return &IFileDidDeactivateMfileDidIterator{contract: _IFileDid.contract, event: "DeactivateMfileDid", logs: logs, sub: sub}, nil
}

// WatchDeactivateMfileDid is a free log subscription operation binding the contract event 0x5e17ae7a2d292f907502879bb1abf12db289923734c91a05b5d0651531aebf2a.
//
// Solidity: event DeactivateMfileDid(string mfileDid, bool deactivate)
func (_IFileDid *IFileDidFilterer) WatchDeactivateMfileDid(opts *bind.WatchOpts, sink chan<- *IFileDidDeactivateMfileDid) (event.Subscription, error) {

	logs, sub, err := _IFileDid.contract.WatchLogs(opts, "DeactivateMfileDid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFileDidDeactivateMfileDid)
				if err := _IFileDid.contract.UnpackLog(event, "DeactivateMfileDid", log); err != nil {
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

// ParseDeactivateMfileDid is a log parse operation binding the contract event 0x5e17ae7a2d292f907502879bb1abf12db289923734c91a05b5d0651531aebf2a.
//
// Solidity: event DeactivateMfileDid(string mfileDid, bool deactivate)
func (_IFileDid *IFileDidFilterer) ParseDeactivateMfileDid(log types.Log) (*IFileDidDeactivateMfileDid, error) {
	event := new(IFileDidDeactivateMfileDid)
	if err := _IFileDid.contract.UnpackLog(event, "DeactivateMfileDid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFileDidDeactivateReadIterator is returned from FilterDeactivateRead and is used to iterate over the raw logs and unpacked data for DeactivateRead events raised by the IFileDid contract.
type IFileDidDeactivateReadIterator struct {
	Event *IFileDidDeactivateRead // Event containing the contract specifics and raw log

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
func (it *IFileDidDeactivateReadIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFileDidDeactivateRead)
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
		it.Event = new(IFileDidDeactivateRead)
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
func (it *IFileDidDeactivateReadIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFileDidDeactivateReadIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFileDidDeactivateRead represents a DeactivateRead event raised by the IFileDid contract.
type IFileDidDeactivateRead struct {
	MfileDid common.Hash
	MemoDid  string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDeactivateRead is a free log retrieval operation binding the contract event 0x9d679cd38c09c06f29e83ac51c57df2602453a9336575359346cfab8a8f32606.
//
// Solidity: event DeactivateRead(string indexed mfileDid, string memoDid)
func (_IFileDid *IFileDidFilterer) FilterDeactivateRead(opts *bind.FilterOpts, mfileDid []string) (*IFileDidDeactivateReadIterator, error) {

	var mfileDidRule []interface{}
	for _, mfileDidItem := range mfileDid {
		mfileDidRule = append(mfileDidRule, mfileDidItem)
	}

	logs, sub, err := _IFileDid.contract.FilterLogs(opts, "DeactivateRead", mfileDidRule)
	if err != nil {
		return nil, err
	}
	return &IFileDidDeactivateReadIterator{contract: _IFileDid.contract, event: "DeactivateRead", logs: logs, sub: sub}, nil
}

// WatchDeactivateRead is a free log subscription operation binding the contract event 0x9d679cd38c09c06f29e83ac51c57df2602453a9336575359346cfab8a8f32606.
//
// Solidity: event DeactivateRead(string indexed mfileDid, string memoDid)
func (_IFileDid *IFileDidFilterer) WatchDeactivateRead(opts *bind.WatchOpts, sink chan<- *IFileDidDeactivateRead, mfileDid []string) (event.Subscription, error) {

	var mfileDidRule []interface{}
	for _, mfileDidItem := range mfileDid {
		mfileDidRule = append(mfileDidRule, mfileDidItem)
	}

	logs, sub, err := _IFileDid.contract.WatchLogs(opts, "DeactivateRead", mfileDidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFileDidDeactivateRead)
				if err := _IFileDid.contract.UnpackLog(event, "DeactivateRead", log); err != nil {
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

// ParseDeactivateRead is a log parse operation binding the contract event 0x9d679cd38c09c06f29e83ac51c57df2602453a9336575359346cfab8a8f32606.
//
// Solidity: event DeactivateRead(string indexed mfileDid, string memoDid)
func (_IFileDid *IFileDidFilterer) ParseDeactivateRead(log types.Log) (*IFileDidDeactivateRead, error) {
	event := new(IFileDidDeactivateRead)
	if err := _IFileDid.contract.UnpackLog(event, "DeactivateRead", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFileDidGrantReadIterator is returned from FilterGrantRead and is used to iterate over the raw logs and unpacked data for GrantRead events raised by the IFileDid contract.
type IFileDidGrantReadIterator struct {
	Event *IFileDidGrantRead // Event containing the contract specifics and raw log

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
func (it *IFileDidGrantReadIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFileDidGrantRead)
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
		it.Event = new(IFileDidGrantRead)
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
func (it *IFileDidGrantReadIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFileDidGrantReadIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFileDidGrantRead represents a GrantRead event raised by the IFileDid contract.
type IFileDidGrantRead struct {
	MfileDid common.Hash
	MemoDid  string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGrantRead is a free log retrieval operation binding the contract event 0x94b2915a0e10afcab72fdc47f046e2d9f8c45380746ef36ed2f10afaff5390be.
//
// Solidity: event GrantRead(string indexed mfileDid, string memoDid)
func (_IFileDid *IFileDidFilterer) FilterGrantRead(opts *bind.FilterOpts, mfileDid []string) (*IFileDidGrantReadIterator, error) {

	var mfileDidRule []interface{}
	for _, mfileDidItem := range mfileDid {
		mfileDidRule = append(mfileDidRule, mfileDidItem)
	}

	logs, sub, err := _IFileDid.contract.FilterLogs(opts, "GrantRead", mfileDidRule)
	if err != nil {
		return nil, err
	}
	return &IFileDidGrantReadIterator{contract: _IFileDid.contract, event: "GrantRead", logs: logs, sub: sub}, nil
}

// WatchGrantRead is a free log subscription operation binding the contract event 0x94b2915a0e10afcab72fdc47f046e2d9f8c45380746ef36ed2f10afaff5390be.
//
// Solidity: event GrantRead(string indexed mfileDid, string memoDid)
func (_IFileDid *IFileDidFilterer) WatchGrantRead(opts *bind.WatchOpts, sink chan<- *IFileDidGrantRead, mfileDid []string) (event.Subscription, error) {

	var mfileDidRule []interface{}
	for _, mfileDidItem := range mfileDid {
		mfileDidRule = append(mfileDidRule, mfileDidItem)
	}

	logs, sub, err := _IFileDid.contract.WatchLogs(opts, "GrantRead", mfileDidRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFileDidGrantRead)
				if err := _IFileDid.contract.UnpackLog(event, "GrantRead", log); err != nil {
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

// ParseGrantRead is a log parse operation binding the contract event 0x94b2915a0e10afcab72fdc47f046e2d9f8c45380746ef36ed2f10afaff5390be.
//
// Solidity: event GrantRead(string indexed mfileDid, string memoDid)
func (_IFileDid *IFileDidFilterer) ParseGrantRead(log types.Log) (*IFileDidGrantRead, error) {
	event := new(IFileDidGrantRead)
	if err := _IFileDid.contract.UnpackLog(event, "GrantRead", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IFileDidRegisterMfileDidIterator is returned from FilterRegisterMfileDid and is used to iterate over the raw logs and unpacked data for RegisterMfileDid events raised by the IFileDid contract.
type IFileDidRegisterMfileDidIterator struct {
	Event *IFileDidRegisterMfileDid // Event containing the contract specifics and raw log

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
func (it *IFileDidRegisterMfileDidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IFileDidRegisterMfileDid)
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
		it.Event = new(IFileDidRegisterMfileDid)
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
func (it *IFileDidRegisterMfileDidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IFileDidRegisterMfileDidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IFileDidRegisterMfileDid represents a RegisterMfileDid event raised by the IFileDid contract.
type IFileDidRegisterMfileDid struct {
	MfileDid string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRegisterMfileDid is a free log retrieval operation binding the contract event 0x09f50a6181226540ba96339e43e56ab9004a370db287d9dd7db64891edc494fe.
//
// Solidity: event RegisterMfileDid(string mfileDid)
func (_IFileDid *IFileDidFilterer) FilterRegisterMfileDid(opts *bind.FilterOpts) (*IFileDidRegisterMfileDidIterator, error) {

	logs, sub, err := _IFileDid.contract.FilterLogs(opts, "RegisterMfileDid")
	if err != nil {
		return nil, err
	}
	return &IFileDidRegisterMfileDidIterator{contract: _IFileDid.contract, event: "RegisterMfileDid", logs: logs, sub: sub}, nil
}

// WatchRegisterMfileDid is a free log subscription operation binding the contract event 0x09f50a6181226540ba96339e43e56ab9004a370db287d9dd7db64891edc494fe.
//
// Solidity: event RegisterMfileDid(string mfileDid)
func (_IFileDid *IFileDidFilterer) WatchRegisterMfileDid(opts *bind.WatchOpts, sink chan<- *IFileDidRegisterMfileDid) (event.Subscription, error) {

	logs, sub, err := _IFileDid.contract.WatchLogs(opts, "RegisterMfileDid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IFileDidRegisterMfileDid)
				if err := _IFileDid.contract.UnpackLog(event, "RegisterMfileDid", log); err != nil {
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

// ParseRegisterMfileDid is a log parse operation binding the contract event 0x09f50a6181226540ba96339e43e56ab9004a370db287d9dd7db64891edc494fe.
//
// Solidity: event RegisterMfileDid(string mfileDid)
func (_IFileDid *IFileDidFilterer) ParseRegisterMfileDid(log types.Log) (*IFileDidRegisterMfileDid, error) {
	event := new(IFileDidRegisterMfileDid)
	if err := _IFileDid.contract.UnpackLog(event, "RegisterMfileDid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IInstanceMetaData contains all meta data concerning the IInstance contract.
var IInstanceMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"Alter\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_type\",\"type\":\"uint8\"}],\"name\":\"instances\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"3ec7d5b9": "instances(uint8)",
	},
}

// IInstanceABI is the input ABI used to generate the binding from.
// Deprecated: Use IInstanceMetaData.ABI instead.
var IInstanceABI = IInstanceMetaData.ABI

// Deprecated: Use IInstanceMetaData.Sigs instead.
// IInstanceFuncSigs maps the 4-byte function signature to its string representation.
var IInstanceFuncSigs = IInstanceMetaData.Sigs

// IInstance is an auto generated Go binding around an Ethereum contract.
type IInstance struct {
	IInstanceCaller     // Read-only binding to the contract
	IInstanceTransactor // Write-only binding to the contract
	IInstanceFilterer   // Log filterer for contract events
}

// IInstanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type IInstanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInstanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IInstanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInstanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IInstanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IInstanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IInstanceSession struct {
	Contract     *IInstance        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IInstanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IInstanceCallerSession struct {
	Contract *IInstanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// IInstanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IInstanceTransactorSession struct {
	Contract     *IInstanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IInstanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type IInstanceRaw struct {
	Contract *IInstance // Generic contract binding to access the raw methods on
}

// IInstanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IInstanceCallerRaw struct {
	Contract *IInstanceCaller // Generic read-only contract binding to access the raw methods on
}

// IInstanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IInstanceTransactorRaw struct {
	Contract *IInstanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIInstance creates a new instance of IInstance, bound to a specific deployed contract.
func NewIInstance(address common.Address, backend bind.ContractBackend) (*IInstance, error) {
	contract, err := bindIInstance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IInstance{IInstanceCaller: IInstanceCaller{contract: contract}, IInstanceTransactor: IInstanceTransactor{contract: contract}, IInstanceFilterer: IInstanceFilterer{contract: contract}}, nil
}

// NewIInstanceCaller creates a new read-only instance of IInstance, bound to a specific deployed contract.
func NewIInstanceCaller(address common.Address, caller bind.ContractCaller) (*IInstanceCaller, error) {
	contract, err := bindIInstance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IInstanceCaller{contract: contract}, nil
}

// NewIInstanceTransactor creates a new write-only instance of IInstance, bound to a specific deployed contract.
func NewIInstanceTransactor(address common.Address, transactor bind.ContractTransactor) (*IInstanceTransactor, error) {
	contract, err := bindIInstance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IInstanceTransactor{contract: contract}, nil
}

// NewIInstanceFilterer creates a new log filterer instance of IInstance, bound to a specific deployed contract.
func NewIInstanceFilterer(address common.Address, filterer bind.ContractFilterer) (*IInstanceFilterer, error) {
	contract, err := bindIInstance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IInstanceFilterer{contract: contract}, nil
}

// bindIInstance binds a generic wrapper to an already deployed contract.
func bindIInstance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IInstanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInstance *IInstanceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInstance.Contract.IInstanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInstance *IInstanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInstance.Contract.IInstanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInstance *IInstanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInstance.Contract.IInstanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IInstance *IInstanceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IInstance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IInstance *IInstanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IInstance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IInstance *IInstanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IInstance.Contract.contract.Transact(opts, method, params...)
}

// Instances is a free data retrieval call binding the contract method 0x3ec7d5b9.
//
// Solidity: function instances(uint8 _type) view returns(address)
func (_IInstance *IInstanceCaller) Instances(opts *bind.CallOpts, _type uint8) (common.Address, error) {
	var out []interface{}
	err := _IInstance.contract.Call(opts, &out, "instances", _type)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Instances is a free data retrieval call binding the contract method 0x3ec7d5b9.
//
// Solidity: function instances(uint8 _type) view returns(address)
func (_IInstance *IInstanceSession) Instances(_type uint8) (common.Address, error) {
	return _IInstance.Contract.Instances(&_IInstance.CallOpts, _type)
}

// Instances is a free data retrieval call binding the contract method 0x3ec7d5b9.
//
// Solidity: function instances(uint8 _type) view returns(address)
func (_IInstance *IInstanceCallerSession) Instances(_type uint8) (common.Address, error) {
	return _IInstance.Contract.Instances(&_IInstance.CallOpts, _type)
}

// IInstanceAlterIterator is returned from FilterAlter and is used to iterate over the raw logs and unpacked data for Alter events raised by the IInstance contract.
type IInstanceAlterIterator struct {
	Event *IInstanceAlter // Event containing the contract specifics and raw log

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
func (it *IInstanceAlterIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IInstanceAlter)
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
		it.Event = new(IInstanceAlter)
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
func (it *IInstanceAlterIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IInstanceAlterIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IInstanceAlter represents a Alter event raised by the IInstance contract.
type IInstanceAlter struct {
	Type uint8
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAlter is a free log retrieval operation binding the contract event 0x69da8aaa18d0d64ebdd4d982e4d32ac4aaab1fe0c1034b19846e51a61fcc0f02.
//
// Solidity: event Alter(uint8 _type, address from, address to)
func (_IInstance *IInstanceFilterer) FilterAlter(opts *bind.FilterOpts) (*IInstanceAlterIterator, error) {

	logs, sub, err := _IInstance.contract.FilterLogs(opts, "Alter")
	if err != nil {
		return nil, err
	}
	return &IInstanceAlterIterator{contract: _IInstance.contract, event: "Alter", logs: logs, sub: sub}, nil
}

// WatchAlter is a free log subscription operation binding the contract event 0x69da8aaa18d0d64ebdd4d982e4d32ac4aaab1fe0c1034b19846e51a61fcc0f02.
//
// Solidity: event Alter(uint8 _type, address from, address to)
func (_IInstance *IInstanceFilterer) WatchAlter(opts *bind.WatchOpts, sink chan<- *IInstanceAlter) (event.Subscription, error) {

	logs, sub, err := _IInstance.contract.WatchLogs(opts, "Alter")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IInstanceAlter)
				if err := _IInstance.contract.UnpackLog(event, "Alter", log); err != nil {
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

// ParseAlter is a log parse operation binding the contract event 0x69da8aaa18d0d64ebdd4d982e4d32ac4aaab1fe0c1034b19846e51a61fcc0f02.
//
// Solidity: event Alter(uint8 _type, address from, address to)
func (_IInstance *IInstanceFilterer) ParseAlter(log types.Log) (*IInstanceAlter, error) {
	event := new(IInstanceAlter)
	if err := _IInstance.contract.UnpackLog(event, "Alter", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ProxyMetaData contains all meta data concerning the Proxy contract.
var ProxyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_instance\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"addAssertion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"addAuth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"addDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"addRecovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"methodType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"controller\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"pubKeyData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"deactivated\",\"type\":\"bool\"}],\"internalType\":\"structIAccountDid.PublicKey\",\"name\":\"pubKey\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"addVeri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"memoDid\",\"type\":\"string\"}],\"name\":\"buyRead\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"controller\",\"type\":\"string\"}],\"name\":\"changeController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"internalType\":\"enumIFileDid.FileType\",\"name\":\"ftype\",\"type\":\"uint8\"}],\"name\":\"changeFtype\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"internalType\":\"string[]\",\"name\":\"keywords\",\"type\":\"string[]\"}],\"name\":\"changeKeywords\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"changePrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"methodType\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"pubKeyData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"createDID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"methodType\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"pubKeyData\",\"type\":\"bytes\"}],\"name\":\"createDIDByAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"deactivate\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"deactivateDID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"deactivate\",\"type\":\"bool\"}],\"name\":\"deactivateMfileDid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"memoDid\",\"type\":\"string\"}],\"name\":\"deactivateRead\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"deactivate\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"deactivateVeri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"}],\"name\":\"deactivated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"}],\"name\":\"getController\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"}],\"name\":\"getEncode\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"}],\"name\":\"getFtype\",\"outputs\":[{\"internalType\":\"enumIFileDid.FileType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"}],\"name\":\"getKeywords\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"getMasterKeyAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getVeri\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"methodType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"controller\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"pubKeyData\",\"type\":\"bytes\"},{\"internalType\":\"bool\",\"name\":\"deactivated\",\"type\":\"bool\"}],\"internalType\":\"structIAccountDid.PublicKey\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"getVeriLen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"memoDid\",\"type\":\"string\"}],\"name\":\"grantRead\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"inAssertion\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"inAuth\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"inDelegation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"inRecovery\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"inst\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"}],\"name\":\"isDeactivated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"memoDid\",\"type\":\"string\"}],\"name\":\"read\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"mfileDid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"encode\",\"type\":\"string\"},{\"internalType\":\"enumIFileDid.FileType\",\"name\":\"ftype\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"controller\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"string[]\",\"name\":\"keywords\",\"type\":\"string[]\"}],\"name\":\"registerMfileDid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"removeAssertion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"removeAuth\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"removeDelegation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"removeRecovery\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"did\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"methodType\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"pubKeyData\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"updateVeri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Sigs: map[string]string{
		"62ea1b22": "addAssertion(string,string,bytes)",
		"a5b4eaed": "addAuth(string,string,bytes)",
		"d101ec86": "addDelegation(string,string,uint256,bytes)",
		"4b7897db": "addRecovery(string,string,bytes)",
		"81f2fd87": "addVeri(string,(string,string,bytes,bool),bytes)",
		"0069c140": "buyRead(string,string)",
		"9aad0566": "changeController(string,string)",
		"52a0fead": "changeFtype(string,uint8)",
		"64c21f2c": "changeKeywords(string,string[])",
		"1fa19596": "changePrice(string,uint256)",
		"5e82d9df": "createDID(string,string,bytes,bytes)",
		"ac4f1d7a": "createDIDByAdmin(string,string,bytes)",
		"917a4b90": "deactivateDID(string,bool,bytes)",
		"86e9dbc6": "deactivateMfileDid(string,bool)",
		"37c52d83": "deactivateRead(string,string)",
		"21f33be8": "deactivateVeri(string,uint256,bool,bytes)",
		"84435ce2": "deactivated(string)",
		"63a27111": "getController(string)",
		"f100cfd3": "getEncode(string)",
		"dff5b013": "getFtype(string)",
		"380cc539": "getKeywords(string)",
		"0e372676": "getMasterKeyAddr(string)",
		"55b7c6fe": "getNonce(string)",
		"524f3889": "getPrice(string)",
		"f7181a1a": "getVeri(string,uint256)",
		"c7948e4b": "getVeriLen(string)",
		"a870fcc7": "grantRead(string,string)",
		"7c01d259": "inAssertion(string,string)",
		"f888a7de": "inAuth(string,string)",
		"7f0a84cc": "inDelegation(string,string)",
		"b4334b0d": "inRecovery(string,string)",
		"bd6b2222": "inst()",
		"e0064236": "isDeactivated(string)",
		"8c97f99e": "read(string,string)",
		"ca0258dc": "registerMfileDid(string,string,uint8,string,uint256,string[])",
		"9aae7467": "removeAssertion(string,string,bytes)",
		"eb785ee2": "removeAuth(string,string,bytes)",
		"d3be140e": "removeDelegation(string,string,bytes)",
		"dc649bf4": "removeRecovery(string,string,bytes)",
		"6c95dd03": "updateVeri(string,uint256,string,bytes,bytes)",
	},
	Bin: "0x60806040523480156200001157600080fd5b50604051620032363803806200323683398101604081905262000034916200005a565b600080546001600160a01b0319166001600160a01b03929092169190911790556200008c565b6000602082840312156200006d57600080fd5b81516001600160a01b03811681146200008557600080fd5b9392505050565b61319a806200009c6000396000f3fe608060405234801561001057600080fd5b50600436106102475760003560e01c80638c97f99e1161013b578063ca0258dc116100b8578063e00642361161007c578063e006423614610573578063eb785ee214610586578063f100cfd314610599578063f7181a1a146105ac578063f888a7de146105cc57600080fd5b8063ca0258dc14610507578063d101ec861461051a578063d3be140e1461052d578063dc649bf414610540578063dff5b0131461055357600080fd5b8063a870fcc7116100ff578063a870fcc7146104a8578063ac4f1d7a146104bb578063b4334b0d146104ce578063bd6b2222146104e1578063c7948e4b146104f457600080fd5b80638c97f99e14610437578063917a4b901461045c5780639aad05661461046f5780639aae746714610482578063a5b4eaed1461049557600080fd5b80635e82d9df116101c95780637c01d2591161018d5780637c01d259146103c85780637f0a84cc146103eb57806381f2fd87146103fe57806384435ce21461041157806386e9dbc61461042457600080fd5b80635e82d9df1461035c57806362ea1b221461036f57806363a271111461038257806364c21f2c146103a25780636c95dd03146103b557600080fd5b8063380cc53911610210578063380cc539146102ca5780634b7897db146102ea578063524f3889146102fd57806352a0fead1461031e57806355b7c6fe1461033157600080fd5b806269c1401461024c5780630e372676146102615780631fa195961461029157806321f33be8146102a457806337c52d83146102b7575b600080fd5b61025f61025a366004612291565b6105df565b005b61027461026f3660046122f4565b6106ae565b6040516001600160a01b0390911681526020015b60405180910390f35b61025f61029f366004612330565b61078e565b61025f6102b2366004612385565b610828565b61025f6102c5366004612291565b6108fe565b6102dd6102d83660046122f4565b610998565b60405161028891906124ae565b61025f6102f83660046124c1565b610a76565b61031061030b3660046122f4565b610b49565b604051908152602001610288565b61025f61032c366004612565565b610c23565b61034461033f3660046122f4565b610cbd565b6040516001600160401b039091168152602001610288565b61025f61036a3660046125b6565b610d97565b61025f61037d3660046124c1565b610e35565b6103956103903660046122f4565b610ed1565b6040516102889190612649565b61025f6103b0366004612709565b610faf565b61025f6103c3366004612762565b611049565b6103db6103d6366004612291565b611122565b6040519015158152602001610288565b6103106103f9366004612291565b611205565b61025f61040c366004612818565b6112e1565b6103db61041f3660046122f4565b61137d565b61025f610432366004612916565b611457565b61044a610445366004612291565b6114f1565b60405160ff9091168152602001610288565b61025f61046a36600461295c565b6115cd565b61025f61047d366004612291565b611669565b61025f6104903660046124c1565b611703565b61025f6104a33660046124c1565b61179f565b61025f6104b6366004612291565b61183b565b61025f6104c93660046124c1565b6118d5565b6103db6104dc366004612291565b611971565b600054610274906001600160a01b031681565b6103106105023660046122f4565b611a0c565b61025f6105153660046129bc565b611aa5565b61025f610528366004612a81565b611b81565b61025f61053b3660046124c1565b611c1f565b61025f61054e3660046124c1565b611cbb565b6105666105613660046122f4565b611d57565b6040516102889190612b1b565b6103db6105813660046122f4565b611e31565b61025f6105943660046124c1565b611eca565b6103956105a73660046122f4565b611f66565b6105bf6105ba366004612330565b611fff565b6040516102889190612b8a565b6103db6105da366004612291565b61210b565b600054604051633ec7d5b960e01b8152602360048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa158015610628573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061064c9190612b9d565b6001600160a01b03166269c14083836040518363ffffffff1660e01b8152600401610678929190612bc6565b600060405180830381600087803b15801561069257600080fd5b505af11580156106a6573d6000803e3d6000fd5b505050505050565b60008054604051633ec7d5b960e01b8152601e60048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa1580156106f8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061071c9190612b9d565b6001600160a01b0316630e372676836040518263ffffffff1660e01b81526004016107479190612649565b602060405180830381865afa158015610764573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107889190612b9d565b92915050565b600054604051633ec7d5b960e01b8152602360048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa1580156107d7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107fb9190612b9d565b6001600160a01b0316631fa1959683836040518363ffffffff1660e01b8152600401610678929190612bf4565b600054604051633ec7d5b960e01b8152602060048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa158015610871573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108959190612b9d565b6001600160a01b03166321f33be8858585856040518563ffffffff1660e01b81526004016108c69493929190612c16565b600060405180830381600087803b1580156108e057600080fd5b505af11580156108f4573d6000803e3d6000fd5b5050505050505050565b600054604051633ec7d5b960e01b8152602360048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa158015610947573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061096b9190612b9d565b6001600160a01b03166337c52d8383836040518363ffffffff1660e01b8152600401610678929190612bc6565b600054604051633ec7d5b960e01b8152602260048201526060916001600160a01b031690633ec7d5b990602401602060405180830381865afa1580156109e2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a069190612b9d565b6001600160a01b031663380cc539836040518263ffffffff1660e01b8152600401610a319190612649565b600060405180830381865afa158015610a4e573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526107889190810190612ca4565b600054604051633ec7d5b960e01b8152602060048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa158015610abf573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ae39190612b9d565b6001600160a01b0316634b7897db8484846040518463ffffffff1660e01b8152600401610b1293929190612d54565b600060405180830381600087803b158015610b2c57600080fd5b505af1158015610b40573d6000803e3d6000fd5b50505050505050565b60008054604051633ec7d5b960e01b8152602260048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa158015610b93573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bb79190612b9d565b6001600160a01b031663524f3889836040518263ffffffff1660e01b8152600401610be29190612649565b602060405180830381865afa158015610bff573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107889190612d97565b600054604051633ec7d5b960e01b8152602360048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa158015610c6c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c909190612b9d565b6001600160a01b03166352a0fead83836040518363ffffffff1660e01b8152600401610678929190612db0565b60008054604051633ec7d5b960e01b8152602060048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa158015610d07573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d2b9190612b9d565b6001600160a01b03166355b7c6fe836040518263ffffffff1660e01b8152600401610d569190612649565b602060405180830381865afa158015610d73573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107889190612dd2565b600054604051633ec7d5b960e01b8152602060048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa158015610de0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e049190612b9d565b6001600160a01b0316635e82d9df858585856040518563ffffffff1660e01b81526004016108c69493929190612dfb565b600054604051633ec7d5b960e01b8152602060048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa158015610e7e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ea29190612b9d565b6001600160a01b03166362ea1b228484846040518463ffffffff1660e01b8152600401610b1293929190612d54565b600054604051633ec7d5b960e01b8152602260048201526060916001600160a01b031690633ec7d5b990602401602060405180830381865afa158015610f1b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f3f9190612b9d565b6001600160a01b03166363a27111836040518263ffffffff1660e01b8152600401610f6a9190612649565b600060405180830381865afa158015610f87573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526107889190810190612e48565b600054604051633ec7d5b960e01b8152602360048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa158015610ff8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061101c9190612b9d565b6001600160a01b03166364c21f2c83836040518363ffffffff1660e01b8152600401610678929190612e7c565b600054604051633ec7d5b960e01b8152602060048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa158015611092573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110b69190612b9d565b6001600160a01b0316636c95dd0386868686866040518663ffffffff1660e01b81526004016110e9959493929190612ea1565b600060405180830381600087803b15801561110357600080fd5b505af1158015611117573d6000803e3d6000fd5b505050505050505050565b60008054604051633ec7d5b960e01b8152601e60048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa15801561116c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111909190612b9d565b6001600160a01b0316637c01d25984846040518363ffffffff1660e01b81526004016111bd929190612bc6565b602060405180830381865afa1580156111da573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111fe9190612f0b565b9392505050565b60008054604051633ec7d5b960e01b8152601e60048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa15801561124f573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112739190612b9d565b6001600160a01b0316637f0a84cc84846040518363ffffffff1660e01b81526004016112a0929190612bc6565b602060405180830381865afa1580156112bd573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111fe9190612d97565b600054604051633ec7d5b960e01b8152602060048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa15801561132a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061134e9190612b9d565b6001600160a01b03166381f2fd878484846040518463ffffffff1660e01b8152600401610b1293929190612f28565b60008054604051633ec7d5b960e01b8152602260048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa1580156113c7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906113eb9190612b9d565b6001600160a01b03166384435ce2836040518263ffffffff1660e01b81526004016114169190612649565b602060405180830381865afa158015611433573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107889190612f0b565b600054604051633ec7d5b960e01b8152602360048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa1580156114a0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114c49190612b9d565b6001600160a01b03166386e9dbc683836040518363ffffffff1660e01b8152600401610678929190612f4d565b60008054604051633ec7d5b960e01b8152602260048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa15801561153b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061155f9190612b9d565b6001600160a01b0316638c97f99e84846040518363ffffffff1660e01b815260040161158c929190612bc6565b602060405180830381865afa1580156115a9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111fe9190612f71565b600054604051633ec7d5b960e01b8152602060048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa158015611616573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061163a9190612b9d565b6001600160a01b031663917a4b908484846040518463ffffffff1660e01b8152600401610b1293929190612f94565b600054604051633ec7d5b960e01b8152602360048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa1580156116b2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116d69190612b9d565b6001600160a01b0316639aad056683836040518363ffffffff1660e01b8152600401610678929190612bc6565b600054604051633ec7d5b960e01b8152602060048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa15801561174c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117709190612b9d565b6001600160a01b0316639aae74678484846040518463ffffffff1660e01b8152600401610b1293929190612d54565b600054604051633ec7d5b960e01b8152602060048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa1580156117e8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061180c9190612b9d565b6001600160a01b031663a5b4eaed8484846040518463ffffffff1660e01b8152600401610b1293929190612d54565b600054604051633ec7d5b960e01b8152602360048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa158015611884573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118a89190612b9d565b6001600160a01b031663a870fcc783836040518363ffffffff1660e01b8152600401610678929190612bc6565b600054604051633ec7d5b960e01b8152602060048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa15801561191e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119429190612b9d565b6001600160a01b031663ac4f1d7a8484846040518463ffffffff1660e01b8152600401610b1293929190612d54565b60008054604051633ec7d5b960e01b8152601e60048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa1580156119bb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119df9190612b9d565b6001600160a01b031663b4334b0d84846040518363ffffffff1660e01b81526004016111bd929190612bc6565b60008054604051633ec7d5b960e01b8152601e60048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa158015611a56573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a7a9190612b9d565b6001600160a01b031663c7948e4b836040518263ffffffff1660e01b8152600401610be29190612649565b600054604051633ec7d5b960e01b8152602360048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa158015611aee573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b129190612b9d565b6001600160a01b031663ca0258dc8787878787876040518763ffffffff1660e01b8152600401611b4796959493929190612fc1565b600060405180830381600087803b158015611b6157600080fd5b505af1158015611b75573d6000803e3d6000fd5b50505050505050505050565b600054604051633ec7d5b960e01b8152602060048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa158015611bca573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611bee9190612b9d565b6001600160a01b031663d101ec86858585856040518563ffffffff1660e01b81526004016108c6949392919061302e565b600054604051633ec7d5b960e01b8152602060048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa158015611c68573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c8c9190612b9d565b6001600160a01b031663d3be140e8484846040518463ffffffff1660e01b8152600401610b1293929190612d54565b600054604051633ec7d5b960e01b8152602060048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa158015611d04573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d289190612b9d565b6001600160a01b031663dc649bf48484846040518463ffffffff1660e01b8152600401610b1293929190612d54565b60008054604051633ec7d5b960e01b8152602260048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa158015611da1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611dc59190612b9d565b6001600160a01b031663dff5b013836040518263ffffffff1660e01b8152600401611df09190612649565b602060405180830381865afa158015611e0d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610788919061306d565b60008054604051633ec7d5b960e01b8152601e60048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa158015611e7b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611e9f9190612b9d565b6001600160a01b031663e0064236836040518263ffffffff1660e01b81526004016114169190612649565b600054604051633ec7d5b960e01b8152602060048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa158015611f13573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f379190612b9d565b6001600160a01b031663eb785ee28484846040518463ffffffff1660e01b8152600401610b1293929190612d54565b600054604051633ec7d5b960e01b8152602260048201526060916001600160a01b031690633ec7d5b990602401602060405180830381865afa158015611fb0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611fd49190612b9d565b6001600160a01b031663f100cfd3836040518263ffffffff1660e01b8152600401610f6a9190612649565b61202c60405180608001604052806060815260200160608152602001606081526020016000151581525090565b600054604051633ec7d5b960e01b8152601e60048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa158015612075573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120999190612b9d565b6001600160a01b031663f7181a1a84846040518363ffffffff1660e01b81526004016120c6929190612bf4565b600060405180830381865afa1580156120e3573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526111fe919081019061308a565b60008054604051633ec7d5b960e01b8152601e60048201526001600160a01b0390911690633ec7d5b990602401602060405180830381865afa158015612155573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121799190612b9d565b6001600160a01b031663f888a7de84846040518363ffffffff1660e01b81526004016111bd929190612bc6565b634e487b7160e01b600052604160045260246000fd5b604051608081016001600160401b03811182821017156121de576121de6121a6565b60405290565b604051601f8201601f191681016001600160401b038111828210171561220c5761220c6121a6565b604052919050565b60006001600160401b0382111561222d5761222d6121a6565b50601f01601f191660200190565b600082601f83011261224c57600080fd5b813561225f61225a82612214565b6121e4565b81815284602083860101111561227457600080fd5b816020850160208301376000918101602001919091529392505050565b600080604083850312156122a457600080fd5b82356001600160401b03808211156122bb57600080fd5b6122c78683870161223b565b935060208501359150808211156122dd57600080fd5b506122ea8582860161223b565b9150509250929050565b60006020828403121561230657600080fd5b81356001600160401b0381111561231c57600080fd5b6123288482850161223b565b949350505050565b6000806040838503121561234357600080fd5b82356001600160401b0381111561235957600080fd5b6123658582860161223b565b95602094909401359450505050565b801515811461238257600080fd5b50565b6000806000806080858703121561239b57600080fd5b84356001600160401b03808211156123b257600080fd5b6123be8883890161223b565b955060208701359450604087013591506123d782612374565b909250606086013590808211156123ed57600080fd5b506123fa8782880161223b565b91505092959194509250565b60005b83811015612421578181015183820152602001612409565b50506000910152565b60008151808452612442816020860160208601612406565b601f01601f19169290920160200192915050565b600082825180855260208086019550808260051b84010181860160005b848110156124a157601f1986840301895261248f83835161242a565b98840198925090830190600101612473565b5090979650505050505050565b6020815260006111fe6020830184612456565b6000806000606084860312156124d657600080fd5b83356001600160401b03808211156124ed57600080fd5b6124f98783880161223b565b9450602086013591508082111561250f57600080fd5b61251b8783880161223b565b9350604086013591508082111561253157600080fd5b5061253e8682870161223b565b9150509250925092565b6002811061238257600080fd5b803561256081612548565b919050565b6000806040838503121561257857600080fd5b82356001600160401b0381111561258e57600080fd5b61259a8582860161223b565b92505060208301356125ab81612548565b809150509250929050565b600080600080608085870312156125cc57600080fd5b84356001600160401b03808211156125e357600080fd5b6125ef8883890161223b565b9550602087013591508082111561260557600080fd5b6126118883890161223b565b9450604087013591508082111561262757600080fd5b6126338883890161223b565b935060608701359150808211156123ed57600080fd5b6020815260006111fe602083018461242a565b60006001600160401b03821115612675576126756121a6565b5060051b60200190565b600082601f83011261269057600080fd5b813560206126a061225a8361265c565b82815260059290921b840181019181810190868411156126bf57600080fd5b8286015b848110156126fe5780356001600160401b038111156126e25760008081fd5b6126f08986838b010161223b565b8452509183019183016126c3565b509695505050505050565b6000806040838503121561271c57600080fd5b82356001600160401b038082111561273357600080fd5b61273f8683870161223b565b9350602085013591508082111561275557600080fd5b506122ea8582860161267f565b600080600080600060a0868803121561277a57600080fd5b85356001600160401b038082111561279157600080fd5b61279d89838a0161223b565b96506020880135955060408801359150808211156127ba57600080fd5b6127c689838a0161223b565b945060608801359150808211156127dc57600080fd5b6127e889838a0161223b565b935060808801359150808211156127fe57600080fd5b5061280b8882890161223b565b9150509295509295909350565b60008060006060848603121561282d57600080fd5b83356001600160401b038082111561284457600080fd5b6128508783880161223b565b9450602086013591508082111561286657600080fd5b908501906080828803121561287a57600080fd5b6128826121bc565b82358281111561289157600080fd5b61289d8982860161223b565b8252506020830135828111156128b257600080fd5b6128be8982860161223b565b6020830152506040830135828111156128d657600080fd5b6128e28982860161223b565b604083015250606083013592506128f883612374565b82606082015280945050604086013591508082111561253157600080fd5b6000806040838503121561292957600080fd5b82356001600160401b0381111561293f57600080fd5b61294b8582860161223b565b92505060208301356125ab81612374565b60008060006060848603121561297157600080fd5b83356001600160401b038082111561298857600080fd5b6129948783880161223b565b9450602086013591506129a682612374565b9092506040850135908082111561253157600080fd5b60008060008060008060c087890312156129d557600080fd5b86356001600160401b03808211156129ec57600080fd5b6129f88a838b0161223b565b97506020890135915080821115612a0e57600080fd5b612a1a8a838b0161223b565b9650612a2860408a01612555565b95506060890135915080821115612a3e57600080fd5b612a4a8a838b0161223b565b94506080890135935060a0890135915080821115612a6757600080fd5b50612a7489828a0161267f565b9150509295509295509295565b60008060008060808587031215612a9757600080fd5b84356001600160401b0380821115612aae57600080fd5b612aba8883890161223b565b95506020870135915080821115612ad057600080fd5b612adc8883890161223b565b94506040870135935060608701359150808211156123ed57600080fd5b60028110612b1757634e487b7160e01b600052602160045260246000fd5b9052565b602081016107888284612af9565b6000815160808452612b3e608085018261242a565b905060208301518482036020860152612b57828261242a565b91505060408301518482036040860152612b71828261242a565b9150506060830151151560608501528091505092915050565b6020815260006111fe6020830184612b29565b600060208284031215612baf57600080fd5b81516001600160a01b03811681146111fe57600080fd5b604081526000612bd9604083018561242a565b8281036020840152612beb818561242a565b95945050505050565b604081526000612c07604083018561242a565b90508260208301529392505050565b608081526000612c29608083018761242a565b85602084015284151560408401528281036060840152612c49818561242a565b979650505050505050565b6000612c6261225a84612214565b9050828152838383011115612c7657600080fd5b6111fe836020830184612406565b600082601f830112612c9557600080fd5b6111fe83835160208501612c54565b60006020808385031215612cb757600080fd5b82516001600160401b0380821115612cce57600080fd5b818501915085601f830112612ce257600080fd5b8151612cf061225a8261265c565b81815260059190911b83018401908481019088831115612d0f57600080fd5b8585015b83811015612d4757805185811115612d2b5760008081fd5b612d398b89838a0101612c84565b845250918601918601612d13565b5098975050505050505050565b606081526000612d67606083018661242a565b8281036020840152612d79818661242a565b90508281036040840152612d8d818561242a565b9695505050505050565b600060208284031215612da957600080fd5b5051919050565b604081526000612dc3604083018561242a565b90506111fe6020830184612af9565b600060208284031215612de457600080fd5b81516001600160401b03811681146111fe57600080fd5b608081526000612e0e608083018761242a565b8281036020840152612e20818761242a565b90508281036040840152612e34818661242a565b90508281036060840152612c49818561242a565b600060208284031215612e5a57600080fd5b81516001600160401b03811115612e7057600080fd5b61232884828501612c84565b604081526000612e8f604083018561242a565b8281036020840152612beb8185612456565b60a081526000612eb460a083018861242a565b8660208401528281036040840152612ecc818761242a565b90508281036060840152612ee0818661242a565b90508281036080840152612ef4818561242a565b98975050505050505050565b805161256081612374565b600060208284031215612f1d57600080fd5b81516111fe81612374565b606081526000612f3b606083018661242a565b8281036020840152612d798186612b29565b604081526000612f60604083018561242a565b905082151560208301529392505050565b600060208284031215612f8357600080fd5b815160ff811681146111fe57600080fd5b606081526000612fa7606083018661242a565b84151560208401528281036040840152612d8d818561242a565b60c081526000612fd460c083018961242a565b8281036020840152612fe6818961242a565b9050612ff56040840188612af9565b8281036060840152613007818761242a565b905084608084015282810360a08401526130218185612456565b9998505050505050505050565b608081526000613041608083018761242a565b8281036020840152613053818761242a565b90508460408401528281036060840152612c49818561242a565b60006020828403121561307f57600080fd5b81516111fe81612548565b60006020828403121561309c57600080fd5b81516001600160401b03808211156130b357600080fd5b90830190608082860312156130c757600080fd5b6130cf6121bc565b8251828111156130de57600080fd5b6130ea87828601612c84565b8252506020830151828111156130ff57600080fd5b61310b87828601612c84565b60208301525060408301518281111561312357600080fd5b83019150601f8201861361313657600080fd5b61314586835160208501612c54565b604082015261315660608401612f00565b60608201529594505050505056fea2646970667358221220421e76fa0446303770ee2dbb009eaa6b77aee319bb0a082860cc26ced2d71f0364736f6c63430008130033",
}

// ProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use ProxyMetaData.ABI instead.
var ProxyABI = ProxyMetaData.ABI

// Deprecated: Use ProxyMetaData.Sigs instead.
// ProxyFuncSigs maps the 4-byte function signature to its string representation.
var ProxyFuncSigs = ProxyMetaData.Sigs

// ProxyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ProxyMetaData.Bin instead.
var ProxyBin = ProxyMetaData.Bin

// DeployProxy deploys a new Ethereum contract, binding an instance of Proxy to it.
func DeployProxy(auth *bind.TransactOpts, backend bind.ContractBackend, _instance common.Address) (common.Address, *types.Transaction, *Proxy, error) {
	parsed, err := ProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ProxyBin), backend, _instance)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Proxy{ProxyCaller: ProxyCaller{contract: contract}, ProxyTransactor: ProxyTransactor{contract: contract}, ProxyFilterer: ProxyFilterer{contract: contract}}, nil
}

// Proxy is an auto generated Go binding around an Ethereum contract.
type Proxy struct {
	ProxyCaller     // Read-only binding to the contract
	ProxyTransactor // Write-only binding to the contract
	ProxyFilterer   // Log filterer for contract events
}

// ProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type ProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ProxySession struct {
	Contract     *Proxy            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ProxyCallerSession struct {
	Contract *ProxyCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ProxyTransactorSession struct {
	Contract     *ProxyTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type ProxyRaw struct {
	Contract *Proxy // Generic contract binding to access the raw methods on
}

// ProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ProxyCallerRaw struct {
	Contract *ProxyCaller // Generic read-only contract binding to access the raw methods on
}

// ProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ProxyTransactorRaw struct {
	Contract *ProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewProxy creates a new instance of Proxy, bound to a specific deployed contract.
func NewProxy(address common.Address, backend bind.ContractBackend) (*Proxy, error) {
	contract, err := bindProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Proxy{ProxyCaller: ProxyCaller{contract: contract}, ProxyTransactor: ProxyTransactor{contract: contract}, ProxyFilterer: ProxyFilterer{contract: contract}}, nil
}

// NewProxyCaller creates a new read-only instance of Proxy, bound to a specific deployed contract.
func NewProxyCaller(address common.Address, caller bind.ContractCaller) (*ProxyCaller, error) {
	contract, err := bindProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyCaller{contract: contract}, nil
}

// NewProxyTransactor creates a new write-only instance of Proxy, bound to a specific deployed contract.
func NewProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*ProxyTransactor, error) {
	contract, err := bindProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ProxyTransactor{contract: contract}, nil
}

// NewProxyFilterer creates a new log filterer instance of Proxy, bound to a specific deployed contract.
func NewProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*ProxyFilterer, error) {
	contract, err := bindProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ProxyFilterer{contract: contract}, nil
}

// bindProxy binds a generic wrapper to an already deployed contract.
func bindProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxy *ProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxy.Contract.ProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxy *ProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxy *ProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxy.Contract.ProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Proxy *ProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Proxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Proxy *ProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Proxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Proxy *ProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Proxy.Contract.contract.Transact(opts, method, params...)
}

// Deactivated is a free data retrieval call binding the contract method 0x84435ce2.
//
// Solidity: function deactivated(string mfileDid) view returns(bool)
func (_Proxy *ProxyCaller) Deactivated(opts *bind.CallOpts, mfileDid string) (bool, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "deactivated", mfileDid)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Deactivated is a free data retrieval call binding the contract method 0x84435ce2.
//
// Solidity: function deactivated(string mfileDid) view returns(bool)
func (_Proxy *ProxySession) Deactivated(mfileDid string) (bool, error) {
	return _Proxy.Contract.Deactivated(&_Proxy.CallOpts, mfileDid)
}

// Deactivated is a free data retrieval call binding the contract method 0x84435ce2.
//
// Solidity: function deactivated(string mfileDid) view returns(bool)
func (_Proxy *ProxyCallerSession) Deactivated(mfileDid string) (bool, error) {
	return _Proxy.Contract.Deactivated(&_Proxy.CallOpts, mfileDid)
}

// GetController is a free data retrieval call binding the contract method 0x63a27111.
//
// Solidity: function getController(string mfileDid) view returns(string)
func (_Proxy *ProxyCaller) GetController(opts *bind.CallOpts, mfileDid string) (string, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "getController", mfileDid)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetController is a free data retrieval call binding the contract method 0x63a27111.
//
// Solidity: function getController(string mfileDid) view returns(string)
func (_Proxy *ProxySession) GetController(mfileDid string) (string, error) {
	return _Proxy.Contract.GetController(&_Proxy.CallOpts, mfileDid)
}

// GetController is a free data retrieval call binding the contract method 0x63a27111.
//
// Solidity: function getController(string mfileDid) view returns(string)
func (_Proxy *ProxyCallerSession) GetController(mfileDid string) (string, error) {
	return _Proxy.Contract.GetController(&_Proxy.CallOpts, mfileDid)
}

// GetEncode is a free data retrieval call binding the contract method 0xf100cfd3.
//
// Solidity: function getEncode(string mfileDid) view returns(string)
func (_Proxy *ProxyCaller) GetEncode(opts *bind.CallOpts, mfileDid string) (string, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "getEncode", mfileDid)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetEncode is a free data retrieval call binding the contract method 0xf100cfd3.
//
// Solidity: function getEncode(string mfileDid) view returns(string)
func (_Proxy *ProxySession) GetEncode(mfileDid string) (string, error) {
	return _Proxy.Contract.GetEncode(&_Proxy.CallOpts, mfileDid)
}

// GetEncode is a free data retrieval call binding the contract method 0xf100cfd3.
//
// Solidity: function getEncode(string mfileDid) view returns(string)
func (_Proxy *ProxyCallerSession) GetEncode(mfileDid string) (string, error) {
	return _Proxy.Contract.GetEncode(&_Proxy.CallOpts, mfileDid)
}

// GetFtype is a free data retrieval call binding the contract method 0xdff5b013.
//
// Solidity: function getFtype(string mfileDid) view returns(uint8)
func (_Proxy *ProxyCaller) GetFtype(opts *bind.CallOpts, mfileDid string) (uint8, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "getFtype", mfileDid)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetFtype is a free data retrieval call binding the contract method 0xdff5b013.
//
// Solidity: function getFtype(string mfileDid) view returns(uint8)
func (_Proxy *ProxySession) GetFtype(mfileDid string) (uint8, error) {
	return _Proxy.Contract.GetFtype(&_Proxy.CallOpts, mfileDid)
}

// GetFtype is a free data retrieval call binding the contract method 0xdff5b013.
//
// Solidity: function getFtype(string mfileDid) view returns(uint8)
func (_Proxy *ProxyCallerSession) GetFtype(mfileDid string) (uint8, error) {
	return _Proxy.Contract.GetFtype(&_Proxy.CallOpts, mfileDid)
}

// GetKeywords is a free data retrieval call binding the contract method 0x380cc539.
//
// Solidity: function getKeywords(string mfileDid) view returns(string[])
func (_Proxy *ProxyCaller) GetKeywords(opts *bind.CallOpts, mfileDid string) ([]string, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "getKeywords", mfileDid)

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetKeywords is a free data retrieval call binding the contract method 0x380cc539.
//
// Solidity: function getKeywords(string mfileDid) view returns(string[])
func (_Proxy *ProxySession) GetKeywords(mfileDid string) ([]string, error) {
	return _Proxy.Contract.GetKeywords(&_Proxy.CallOpts, mfileDid)
}

// GetKeywords is a free data retrieval call binding the contract method 0x380cc539.
//
// Solidity: function getKeywords(string mfileDid) view returns(string[])
func (_Proxy *ProxyCallerSession) GetKeywords(mfileDid string) ([]string, error) {
	return _Proxy.Contract.GetKeywords(&_Proxy.CallOpts, mfileDid)
}

// GetMasterKeyAddr is a free data retrieval call binding the contract method 0x0e372676.
//
// Solidity: function getMasterKeyAddr(string did) view returns(address)
func (_Proxy *ProxyCaller) GetMasterKeyAddr(opts *bind.CallOpts, did string) (common.Address, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "getMasterKeyAddr", did)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetMasterKeyAddr is a free data retrieval call binding the contract method 0x0e372676.
//
// Solidity: function getMasterKeyAddr(string did) view returns(address)
func (_Proxy *ProxySession) GetMasterKeyAddr(did string) (common.Address, error) {
	return _Proxy.Contract.GetMasterKeyAddr(&_Proxy.CallOpts, did)
}

// GetMasterKeyAddr is a free data retrieval call binding the contract method 0x0e372676.
//
// Solidity: function getMasterKeyAddr(string did) view returns(address)
func (_Proxy *ProxyCallerSession) GetMasterKeyAddr(did string) (common.Address, error) {
	return _Proxy.Contract.GetMasterKeyAddr(&_Proxy.CallOpts, did)
}

// GetNonce is a free data retrieval call binding the contract method 0x55b7c6fe.
//
// Solidity: function getNonce(string did) view returns(uint64)
func (_Proxy *ProxyCaller) GetNonce(opts *bind.CallOpts, did string) (uint64, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "getNonce", did)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x55b7c6fe.
//
// Solidity: function getNonce(string did) view returns(uint64)
func (_Proxy *ProxySession) GetNonce(did string) (uint64, error) {
	return _Proxy.Contract.GetNonce(&_Proxy.CallOpts, did)
}

// GetNonce is a free data retrieval call binding the contract method 0x55b7c6fe.
//
// Solidity: function getNonce(string did) view returns(uint64)
func (_Proxy *ProxyCallerSession) GetNonce(did string) (uint64, error) {
	return _Proxy.Contract.GetNonce(&_Proxy.CallOpts, did)
}

// GetPrice is a free data retrieval call binding the contract method 0x524f3889.
//
// Solidity: function getPrice(string mfileDid) view returns(uint256)
func (_Proxy *ProxyCaller) GetPrice(opts *bind.CallOpts, mfileDid string) (*big.Int, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "getPrice", mfileDid)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrice is a free data retrieval call binding the contract method 0x524f3889.
//
// Solidity: function getPrice(string mfileDid) view returns(uint256)
func (_Proxy *ProxySession) GetPrice(mfileDid string) (*big.Int, error) {
	return _Proxy.Contract.GetPrice(&_Proxy.CallOpts, mfileDid)
}

// GetPrice is a free data retrieval call binding the contract method 0x524f3889.
//
// Solidity: function getPrice(string mfileDid) view returns(uint256)
func (_Proxy *ProxyCallerSession) GetPrice(mfileDid string) (*big.Int, error) {
	return _Proxy.Contract.GetPrice(&_Proxy.CallOpts, mfileDid)
}

// GetVeri is a free data retrieval call binding the contract method 0xf7181a1a.
//
// Solidity: function getVeri(string did, uint256 index) view returns((string,string,bytes,bool))
func (_Proxy *ProxyCaller) GetVeri(opts *bind.CallOpts, did string, index *big.Int) (IAccountDidPublicKey, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "getVeri", did, index)

	if err != nil {
		return *new(IAccountDidPublicKey), err
	}

	out0 := *abi.ConvertType(out[0], new(IAccountDidPublicKey)).(*IAccountDidPublicKey)

	return out0, err

}

// GetVeri is a free data retrieval call binding the contract method 0xf7181a1a.
//
// Solidity: function getVeri(string did, uint256 index) view returns((string,string,bytes,bool))
func (_Proxy *ProxySession) GetVeri(did string, index *big.Int) (IAccountDidPublicKey, error) {
	return _Proxy.Contract.GetVeri(&_Proxy.CallOpts, did, index)
}

// GetVeri is a free data retrieval call binding the contract method 0xf7181a1a.
//
// Solidity: function getVeri(string did, uint256 index) view returns((string,string,bytes,bool))
func (_Proxy *ProxyCallerSession) GetVeri(did string, index *big.Int) (IAccountDidPublicKey, error) {
	return _Proxy.Contract.GetVeri(&_Proxy.CallOpts, did, index)
}

// GetVeriLen is a free data retrieval call binding the contract method 0xc7948e4b.
//
// Solidity: function getVeriLen(string did) view returns(uint256)
func (_Proxy *ProxyCaller) GetVeriLen(opts *bind.CallOpts, did string) (*big.Int, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "getVeriLen", did)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVeriLen is a free data retrieval call binding the contract method 0xc7948e4b.
//
// Solidity: function getVeriLen(string did) view returns(uint256)
func (_Proxy *ProxySession) GetVeriLen(did string) (*big.Int, error) {
	return _Proxy.Contract.GetVeriLen(&_Proxy.CallOpts, did)
}

// GetVeriLen is a free data retrieval call binding the contract method 0xc7948e4b.
//
// Solidity: function getVeriLen(string did) view returns(uint256)
func (_Proxy *ProxyCallerSession) GetVeriLen(did string) (*big.Int, error) {
	return _Proxy.Contract.GetVeriLen(&_Proxy.CallOpts, did)
}

// InAssertion is a free data retrieval call binding the contract method 0x7c01d259.
//
// Solidity: function inAssertion(string did, string id) view returns(bool)
func (_Proxy *ProxyCaller) InAssertion(opts *bind.CallOpts, did string, id string) (bool, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "inAssertion", did, id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InAssertion is a free data retrieval call binding the contract method 0x7c01d259.
//
// Solidity: function inAssertion(string did, string id) view returns(bool)
func (_Proxy *ProxySession) InAssertion(did string, id string) (bool, error) {
	return _Proxy.Contract.InAssertion(&_Proxy.CallOpts, did, id)
}

// InAssertion is a free data retrieval call binding the contract method 0x7c01d259.
//
// Solidity: function inAssertion(string did, string id) view returns(bool)
func (_Proxy *ProxyCallerSession) InAssertion(did string, id string) (bool, error) {
	return _Proxy.Contract.InAssertion(&_Proxy.CallOpts, did, id)
}

// InAuth is a free data retrieval call binding the contract method 0xf888a7de.
//
// Solidity: function inAuth(string did, string id) view returns(bool)
func (_Proxy *ProxyCaller) InAuth(opts *bind.CallOpts, did string, id string) (bool, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "inAuth", did, id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InAuth is a free data retrieval call binding the contract method 0xf888a7de.
//
// Solidity: function inAuth(string did, string id) view returns(bool)
func (_Proxy *ProxySession) InAuth(did string, id string) (bool, error) {
	return _Proxy.Contract.InAuth(&_Proxy.CallOpts, did, id)
}

// InAuth is a free data retrieval call binding the contract method 0xf888a7de.
//
// Solidity: function inAuth(string did, string id) view returns(bool)
func (_Proxy *ProxyCallerSession) InAuth(did string, id string) (bool, error) {
	return _Proxy.Contract.InAuth(&_Proxy.CallOpts, did, id)
}

// InDelegation is a free data retrieval call binding the contract method 0x7f0a84cc.
//
// Solidity: function inDelegation(string did, string id) view returns(uint256)
func (_Proxy *ProxyCaller) InDelegation(opts *bind.CallOpts, did string, id string) (*big.Int, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "inDelegation", did, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InDelegation is a free data retrieval call binding the contract method 0x7f0a84cc.
//
// Solidity: function inDelegation(string did, string id) view returns(uint256)
func (_Proxy *ProxySession) InDelegation(did string, id string) (*big.Int, error) {
	return _Proxy.Contract.InDelegation(&_Proxy.CallOpts, did, id)
}

// InDelegation is a free data retrieval call binding the contract method 0x7f0a84cc.
//
// Solidity: function inDelegation(string did, string id) view returns(uint256)
func (_Proxy *ProxyCallerSession) InDelegation(did string, id string) (*big.Int, error) {
	return _Proxy.Contract.InDelegation(&_Proxy.CallOpts, did, id)
}

// InRecovery is a free data retrieval call binding the contract method 0xb4334b0d.
//
// Solidity: function inRecovery(string did, string id) view returns(bool)
func (_Proxy *ProxyCaller) InRecovery(opts *bind.CallOpts, did string, id string) (bool, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "inRecovery", did, id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// InRecovery is a free data retrieval call binding the contract method 0xb4334b0d.
//
// Solidity: function inRecovery(string did, string id) view returns(bool)
func (_Proxy *ProxySession) InRecovery(did string, id string) (bool, error) {
	return _Proxy.Contract.InRecovery(&_Proxy.CallOpts, did, id)
}

// InRecovery is a free data retrieval call binding the contract method 0xb4334b0d.
//
// Solidity: function inRecovery(string did, string id) view returns(bool)
func (_Proxy *ProxyCallerSession) InRecovery(did string, id string) (bool, error) {
	return _Proxy.Contract.InRecovery(&_Proxy.CallOpts, did, id)
}

// Inst is a free data retrieval call binding the contract method 0xbd6b2222.
//
// Solidity: function inst() view returns(address)
func (_Proxy *ProxyCaller) Inst(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "inst")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Inst is a free data retrieval call binding the contract method 0xbd6b2222.
//
// Solidity: function inst() view returns(address)
func (_Proxy *ProxySession) Inst() (common.Address, error) {
	return _Proxy.Contract.Inst(&_Proxy.CallOpts)
}

// Inst is a free data retrieval call binding the contract method 0xbd6b2222.
//
// Solidity: function inst() view returns(address)
func (_Proxy *ProxyCallerSession) Inst() (common.Address, error) {
	return _Proxy.Contract.Inst(&_Proxy.CallOpts)
}

// IsDeactivated is a free data retrieval call binding the contract method 0xe0064236.
//
// Solidity: function isDeactivated(string did) view returns(bool)
func (_Proxy *ProxyCaller) IsDeactivated(opts *bind.CallOpts, did string) (bool, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "isDeactivated", did)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDeactivated is a free data retrieval call binding the contract method 0xe0064236.
//
// Solidity: function isDeactivated(string did) view returns(bool)
func (_Proxy *ProxySession) IsDeactivated(did string) (bool, error) {
	return _Proxy.Contract.IsDeactivated(&_Proxy.CallOpts, did)
}

// IsDeactivated is a free data retrieval call binding the contract method 0xe0064236.
//
// Solidity: function isDeactivated(string did) view returns(bool)
func (_Proxy *ProxyCallerSession) IsDeactivated(did string) (bool, error) {
	return _Proxy.Contract.IsDeactivated(&_Proxy.CallOpts, did)
}

// Read is a free data retrieval call binding the contract method 0x8c97f99e.
//
// Solidity: function read(string mfileDid, string memoDid) view returns(uint8)
func (_Proxy *ProxyCaller) Read(opts *bind.CallOpts, mfileDid string, memoDid string) (uint8, error) {
	var out []interface{}
	err := _Proxy.contract.Call(opts, &out, "read", mfileDid, memoDid)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Read is a free data retrieval call binding the contract method 0x8c97f99e.
//
// Solidity: function read(string mfileDid, string memoDid) view returns(uint8)
func (_Proxy *ProxySession) Read(mfileDid string, memoDid string) (uint8, error) {
	return _Proxy.Contract.Read(&_Proxy.CallOpts, mfileDid, memoDid)
}

// Read is a free data retrieval call binding the contract method 0x8c97f99e.
//
// Solidity: function read(string mfileDid, string memoDid) view returns(uint8)
func (_Proxy *ProxyCallerSession) Read(mfileDid string, memoDid string) (uint8, error) {
	return _Proxy.Contract.Read(&_Proxy.CallOpts, mfileDid, memoDid)
}

// AddAssertion is a paid mutator transaction binding the contract method 0x62ea1b22.
//
// Solidity: function addAssertion(string did, string id, bytes signature) returns()
func (_Proxy *ProxyTransactor) AddAssertion(opts *bind.TransactOpts, did string, id string, signature []byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "addAssertion", did, id, signature)
}

// AddAssertion is a paid mutator transaction binding the contract method 0x62ea1b22.
//
// Solidity: function addAssertion(string did, string id, bytes signature) returns()
func (_Proxy *ProxySession) AddAssertion(did string, id string, signature []byte) (*types.Transaction, error) {
	return _Proxy.Contract.AddAssertion(&_Proxy.TransactOpts, did, id, signature)
}

// AddAssertion is a paid mutator transaction binding the contract method 0x62ea1b22.
//
// Solidity: function addAssertion(string did, string id, bytes signature) returns()
func (_Proxy *ProxyTransactorSession) AddAssertion(did string, id string, signature []byte) (*types.Transaction, error) {
	return _Proxy.Contract.AddAssertion(&_Proxy.TransactOpts, did, id, signature)
}

// AddAuth is a paid mutator transaction binding the contract method 0xa5b4eaed.
//
// Solidity: function addAuth(string did, string id, bytes signature) returns()
func (_Proxy *ProxyTransactor) AddAuth(opts *bind.TransactOpts, did string, id string, signature []byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "addAuth", did, id, signature)
}

// AddAuth is a paid mutator transaction binding the contract method 0xa5b4eaed.
//
// Solidity: function addAuth(string did, string id, bytes signature) returns()
func (_Proxy *ProxySession) AddAuth(did string, id string, signature []byte) (*types.Transaction, error) {
	return _Proxy.Contract.AddAuth(&_Proxy.TransactOpts, did, id, signature)
}

// AddAuth is a paid mutator transaction binding the contract method 0xa5b4eaed.
//
// Solidity: function addAuth(string did, string id, bytes signature) returns()
func (_Proxy *ProxyTransactorSession) AddAuth(did string, id string, signature []byte) (*types.Transaction, error) {
	return _Proxy.Contract.AddAuth(&_Proxy.TransactOpts, did, id, signature)
}

// AddDelegation is a paid mutator transaction binding the contract method 0xd101ec86.
//
// Solidity: function addDelegation(string did, string id, uint256 expiration, bytes signature) returns()
func (_Proxy *ProxyTransactor) AddDelegation(opts *bind.TransactOpts, did string, id string, expiration *big.Int, signature []byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "addDelegation", did, id, expiration, signature)
}

// AddDelegation is a paid mutator transaction binding the contract method 0xd101ec86.
//
// Solidity: function addDelegation(string did, string id, uint256 expiration, bytes signature) returns()
func (_Proxy *ProxySession) AddDelegation(did string, id string, expiration *big.Int, signature []byte) (*types.Transaction, error) {
	return _Proxy.Contract.AddDelegation(&_Proxy.TransactOpts, did, id, expiration, signature)
}

// AddDelegation is a paid mutator transaction binding the contract method 0xd101ec86.
//
// Solidity: function addDelegation(string did, string id, uint256 expiration, bytes signature) returns()
func (_Proxy *ProxyTransactorSession) AddDelegation(did string, id string, expiration *big.Int, signature []byte) (*types.Transaction, error) {
	return _Proxy.Contract.AddDelegation(&_Proxy.TransactOpts, did, id, expiration, signature)
}

// AddRecovery is a paid mutator transaction binding the contract method 0x4b7897db.
//
// Solidity: function addRecovery(string did, string id, bytes signature) returns()
func (_Proxy *ProxyTransactor) AddRecovery(opts *bind.TransactOpts, did string, id string, signature []byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "addRecovery", did, id, signature)
}

// AddRecovery is a paid mutator transaction binding the contract method 0x4b7897db.
//
// Solidity: function addRecovery(string did, string id, bytes signature) returns()
func (_Proxy *ProxySession) AddRecovery(did string, id string, signature []byte) (*types.Transaction, error) {
	return _Proxy.Contract.AddRecovery(&_Proxy.TransactOpts, did, id, signature)
}

// AddRecovery is a paid mutator transaction binding the contract method 0x4b7897db.
//
// Solidity: function addRecovery(string did, string id, bytes signature) returns()
func (_Proxy *ProxyTransactorSession) AddRecovery(did string, id string, signature []byte) (*types.Transaction, error) {
	return _Proxy.Contract.AddRecovery(&_Proxy.TransactOpts, did, id, signature)
}

// AddVeri is a paid mutator transaction binding the contract method 0x81f2fd87.
//
// Solidity: function addVeri(string did, (string,string,bytes,bool) pubKey, bytes signature) returns()
func (_Proxy *ProxyTransactor) AddVeri(opts *bind.TransactOpts, did string, pubKey IAccountDidPublicKey, signature []byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "addVeri", did, pubKey, signature)
}

// AddVeri is a paid mutator transaction binding the contract method 0x81f2fd87.
//
// Solidity: function addVeri(string did, (string,string,bytes,bool) pubKey, bytes signature) returns()
func (_Proxy *ProxySession) AddVeri(did string, pubKey IAccountDidPublicKey, signature []byte) (*types.Transaction, error) {
	return _Proxy.Contract.AddVeri(&_Proxy.TransactOpts, did, pubKey, signature)
}

// AddVeri is a paid mutator transaction binding the contract method 0x81f2fd87.
//
// Solidity: function addVeri(string did, (string,string,bytes,bool) pubKey, bytes signature) returns()
func (_Proxy *ProxyTransactorSession) AddVeri(did string, pubKey IAccountDidPublicKey, signature []byte) (*types.Transaction, error) {
	return _Proxy.Contract.AddVeri(&_Proxy.TransactOpts, did, pubKey, signature)
}

// BuyRead is a paid mutator transaction binding the contract method 0x0069c140.
//
// Solidity: function buyRead(string mfileDid, string memoDid) returns()
func (_Proxy *ProxyTransactor) BuyRead(opts *bind.TransactOpts, mfileDid string, memoDid string) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "buyRead", mfileDid, memoDid)
}

// BuyRead is a paid mutator transaction binding the contract method 0x0069c140.
//
// Solidity: function buyRead(string mfileDid, string memoDid) returns()
func (_Proxy *ProxySession) BuyRead(mfileDid string, memoDid string) (*types.Transaction, error) {
	return _Proxy.Contract.BuyRead(&_Proxy.TransactOpts, mfileDid, memoDid)
}

// BuyRead is a paid mutator transaction binding the contract method 0x0069c140.
//
// Solidity: function buyRead(string mfileDid, string memoDid) returns()
func (_Proxy *ProxyTransactorSession) BuyRead(mfileDid string, memoDid string) (*types.Transaction, error) {
	return _Proxy.Contract.BuyRead(&_Proxy.TransactOpts, mfileDid, memoDid)
}

// ChangeController is a paid mutator transaction binding the contract method 0x9aad0566.
//
// Solidity: function changeController(string mfileDid, string controller) returns()
func (_Proxy *ProxyTransactor) ChangeController(opts *bind.TransactOpts, mfileDid string, controller string) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "changeController", mfileDid, controller)
}

// ChangeController is a paid mutator transaction binding the contract method 0x9aad0566.
//
// Solidity: function changeController(string mfileDid, string controller) returns()
func (_Proxy *ProxySession) ChangeController(mfileDid string, controller string) (*types.Transaction, error) {
	return _Proxy.Contract.ChangeController(&_Proxy.TransactOpts, mfileDid, controller)
}

// ChangeController is a paid mutator transaction binding the contract method 0x9aad0566.
//
// Solidity: function changeController(string mfileDid, string controller) returns()
func (_Proxy *ProxyTransactorSession) ChangeController(mfileDid string, controller string) (*types.Transaction, error) {
	return _Proxy.Contract.ChangeController(&_Proxy.TransactOpts, mfileDid, controller)
}

// ChangeFtype is a paid mutator transaction binding the contract method 0x52a0fead.
//
// Solidity: function changeFtype(string mfileDid, uint8 ftype) returns()
func (_Proxy *ProxyTransactor) ChangeFtype(opts *bind.TransactOpts, mfileDid string, ftype uint8) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "changeFtype", mfileDid, ftype)
}

// ChangeFtype is a paid mutator transaction binding the contract method 0x52a0fead.
//
// Solidity: function changeFtype(string mfileDid, uint8 ftype) returns()
func (_Proxy *ProxySession) ChangeFtype(mfileDid string, ftype uint8) (*types.Transaction, error) {
	return _Proxy.Contract.ChangeFtype(&_Proxy.TransactOpts, mfileDid, ftype)
}

// ChangeFtype is a paid mutator transaction binding the contract method 0x52a0fead.
//
// Solidity: function changeFtype(string mfileDid, uint8 ftype) returns()
func (_Proxy *ProxyTransactorSession) ChangeFtype(mfileDid string, ftype uint8) (*types.Transaction, error) {
	return _Proxy.Contract.ChangeFtype(&_Proxy.TransactOpts, mfileDid, ftype)
}

// ChangeKeywords is a paid mutator transaction binding the contract method 0x64c21f2c.
//
// Solidity: function changeKeywords(string mfileDid, string[] keywords) returns()
func (_Proxy *ProxyTransactor) ChangeKeywords(opts *bind.TransactOpts, mfileDid string, keywords []string) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "changeKeywords", mfileDid, keywords)
}

// ChangeKeywords is a paid mutator transaction binding the contract method 0x64c21f2c.
//
// Solidity: function changeKeywords(string mfileDid, string[] keywords) returns()
func (_Proxy *ProxySession) ChangeKeywords(mfileDid string, keywords []string) (*types.Transaction, error) {
	return _Proxy.Contract.ChangeKeywords(&_Proxy.TransactOpts, mfileDid, keywords)
}

// ChangeKeywords is a paid mutator transaction binding the contract method 0x64c21f2c.
//
// Solidity: function changeKeywords(string mfileDid, string[] keywords) returns()
func (_Proxy *ProxyTransactorSession) ChangeKeywords(mfileDid string, keywords []string) (*types.Transaction, error) {
	return _Proxy.Contract.ChangeKeywords(&_Proxy.TransactOpts, mfileDid, keywords)
}

// ChangePrice is a paid mutator transaction binding the contract method 0x1fa19596.
//
// Solidity: function changePrice(string mfileDid, uint256 price) returns()
func (_Proxy *ProxyTransactor) ChangePrice(opts *bind.TransactOpts, mfileDid string, price *big.Int) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "changePrice", mfileDid, price)
}

// ChangePrice is a paid mutator transaction binding the contract method 0x1fa19596.
//
// Solidity: function changePrice(string mfileDid, uint256 price) returns()
func (_Proxy *ProxySession) ChangePrice(mfileDid string, price *big.Int) (*types.Transaction, error) {
	return _Proxy.Contract.ChangePrice(&_Proxy.TransactOpts, mfileDid, price)
}

// ChangePrice is a paid mutator transaction binding the contract method 0x1fa19596.
//
// Solidity: function changePrice(string mfileDid, uint256 price) returns()
func (_Proxy *ProxyTransactorSession) ChangePrice(mfileDid string, price *big.Int) (*types.Transaction, error) {
	return _Proxy.Contract.ChangePrice(&_Proxy.TransactOpts, mfileDid, price)
}

// CreateDID is a paid mutator transaction binding the contract method 0x5e82d9df.
//
// Solidity: function createDID(string did, string methodType, bytes pubKeyData, bytes signature) returns()
func (_Proxy *ProxyTransactor) CreateDID(opts *bind.TransactOpts, did string, methodType string, pubKeyData []byte, signature []byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "createDID", did, methodType, pubKeyData, signature)
}

// CreateDID is a paid mutator transaction binding the contract method 0x5e82d9df.
//
// Solidity: function createDID(string did, string methodType, bytes pubKeyData, bytes signature) returns()
func (_Proxy *ProxySession) CreateDID(did string, methodType string, pubKeyData []byte, signature []byte) (*types.Transaction, error) {
	return _Proxy.Contract.CreateDID(&_Proxy.TransactOpts, did, methodType, pubKeyData, signature)
}

// CreateDID is a paid mutator transaction binding the contract method 0x5e82d9df.
//
// Solidity: function createDID(string did, string methodType, bytes pubKeyData, bytes signature) returns()
func (_Proxy *ProxyTransactorSession) CreateDID(did string, methodType string, pubKeyData []byte, signature []byte) (*types.Transaction, error) {
	return _Proxy.Contract.CreateDID(&_Proxy.TransactOpts, did, methodType, pubKeyData, signature)
}

// CreateDIDByAdmin is a paid mutator transaction binding the contract method 0xac4f1d7a.
//
// Solidity: function createDIDByAdmin(string did, string methodType, bytes pubKeyData) returns()
func (_Proxy *ProxyTransactor) CreateDIDByAdmin(opts *bind.TransactOpts, did string, methodType string, pubKeyData []byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "createDIDByAdmin", did, methodType, pubKeyData)
}

// CreateDIDByAdmin is a paid mutator transaction binding the contract method 0xac4f1d7a.
//
// Solidity: function createDIDByAdmin(string did, string methodType, bytes pubKeyData) returns()
func (_Proxy *ProxySession) CreateDIDByAdmin(did string, methodType string, pubKeyData []byte) (*types.Transaction, error) {
	return _Proxy.Contract.CreateDIDByAdmin(&_Proxy.TransactOpts, did, methodType, pubKeyData)
}

// CreateDIDByAdmin is a paid mutator transaction binding the contract method 0xac4f1d7a.
//
// Solidity: function createDIDByAdmin(string did, string methodType, bytes pubKeyData) returns()
func (_Proxy *ProxyTransactorSession) CreateDIDByAdmin(did string, methodType string, pubKeyData []byte) (*types.Transaction, error) {
	return _Proxy.Contract.CreateDIDByAdmin(&_Proxy.TransactOpts, did, methodType, pubKeyData)
}

// DeactivateDID is a paid mutator transaction binding the contract method 0x917a4b90.
//
// Solidity: function deactivateDID(string did, bool deactivate, bytes signature) returns()
func (_Proxy *ProxyTransactor) DeactivateDID(opts *bind.TransactOpts, did string, deactivate bool, signature []byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "deactivateDID", did, deactivate, signature)
}

// DeactivateDID is a paid mutator transaction binding the contract method 0x917a4b90.
//
// Solidity: function deactivateDID(string did, bool deactivate, bytes signature) returns()
func (_Proxy *ProxySession) DeactivateDID(did string, deactivate bool, signature []byte) (*types.Transaction, error) {
	return _Proxy.Contract.DeactivateDID(&_Proxy.TransactOpts, did, deactivate, signature)
}

// DeactivateDID is a paid mutator transaction binding the contract method 0x917a4b90.
//
// Solidity: function deactivateDID(string did, bool deactivate, bytes signature) returns()
func (_Proxy *ProxyTransactorSession) DeactivateDID(did string, deactivate bool, signature []byte) (*types.Transaction, error) {
	return _Proxy.Contract.DeactivateDID(&_Proxy.TransactOpts, did, deactivate, signature)
}

// DeactivateMfileDid is a paid mutator transaction binding the contract method 0x86e9dbc6.
//
// Solidity: function deactivateMfileDid(string mfileDid, bool deactivate) returns()
func (_Proxy *ProxyTransactor) DeactivateMfileDid(opts *bind.TransactOpts, mfileDid string, deactivate bool) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "deactivateMfileDid", mfileDid, deactivate)
}

// DeactivateMfileDid is a paid mutator transaction binding the contract method 0x86e9dbc6.
//
// Solidity: function deactivateMfileDid(string mfileDid, bool deactivate) returns()
func (_Proxy *ProxySession) DeactivateMfileDid(mfileDid string, deactivate bool) (*types.Transaction, error) {
	return _Proxy.Contract.DeactivateMfileDid(&_Proxy.TransactOpts, mfileDid, deactivate)
}

// DeactivateMfileDid is a paid mutator transaction binding the contract method 0x86e9dbc6.
//
// Solidity: function deactivateMfileDid(string mfileDid, bool deactivate) returns()
func (_Proxy *ProxyTransactorSession) DeactivateMfileDid(mfileDid string, deactivate bool) (*types.Transaction, error) {
	return _Proxy.Contract.DeactivateMfileDid(&_Proxy.TransactOpts, mfileDid, deactivate)
}

// DeactivateRead is a paid mutator transaction binding the contract method 0x37c52d83.
//
// Solidity: function deactivateRead(string mfileDid, string memoDid) returns()
func (_Proxy *ProxyTransactor) DeactivateRead(opts *bind.TransactOpts, mfileDid string, memoDid string) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "deactivateRead", mfileDid, memoDid)
}

// DeactivateRead is a paid mutator transaction binding the contract method 0x37c52d83.
//
// Solidity: function deactivateRead(string mfileDid, string memoDid) returns()
func (_Proxy *ProxySession) DeactivateRead(mfileDid string, memoDid string) (*types.Transaction, error) {
	return _Proxy.Contract.DeactivateRead(&_Proxy.TransactOpts, mfileDid, memoDid)
}

// DeactivateRead is a paid mutator transaction binding the contract method 0x37c52d83.
//
// Solidity: function deactivateRead(string mfileDid, string memoDid) returns()
func (_Proxy *ProxyTransactorSession) DeactivateRead(mfileDid string, memoDid string) (*types.Transaction, error) {
	return _Proxy.Contract.DeactivateRead(&_Proxy.TransactOpts, mfileDid, memoDid)
}

// DeactivateVeri is a paid mutator transaction binding the contract method 0x21f33be8.
//
// Solidity: function deactivateVeri(string did, uint256 index, bool deactivate, bytes signature) returns()
func (_Proxy *ProxyTransactor) DeactivateVeri(opts *bind.TransactOpts, did string, index *big.Int, deactivate bool, signature []byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "deactivateVeri", did, index, deactivate, signature)
}

// DeactivateVeri is a paid mutator transaction binding the contract method 0x21f33be8.
//
// Solidity: function deactivateVeri(string did, uint256 index, bool deactivate, bytes signature) returns()
func (_Proxy *ProxySession) DeactivateVeri(did string, index *big.Int, deactivate bool, signature []byte) (*types.Transaction, error) {
	return _Proxy.Contract.DeactivateVeri(&_Proxy.TransactOpts, did, index, deactivate, signature)
}

// DeactivateVeri is a paid mutator transaction binding the contract method 0x21f33be8.
//
// Solidity: function deactivateVeri(string did, uint256 index, bool deactivate, bytes signature) returns()
func (_Proxy *ProxyTransactorSession) DeactivateVeri(did string, index *big.Int, deactivate bool, signature []byte) (*types.Transaction, error) {
	return _Proxy.Contract.DeactivateVeri(&_Proxy.TransactOpts, did, index, deactivate, signature)
}

// GrantRead is a paid mutator transaction binding the contract method 0xa870fcc7.
//
// Solidity: function grantRead(string mfileDid, string memoDid) returns()
func (_Proxy *ProxyTransactor) GrantRead(opts *bind.TransactOpts, mfileDid string, memoDid string) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "grantRead", mfileDid, memoDid)
}

// GrantRead is a paid mutator transaction binding the contract method 0xa870fcc7.
//
// Solidity: function grantRead(string mfileDid, string memoDid) returns()
func (_Proxy *ProxySession) GrantRead(mfileDid string, memoDid string) (*types.Transaction, error) {
	return _Proxy.Contract.GrantRead(&_Proxy.TransactOpts, mfileDid, memoDid)
}

// GrantRead is a paid mutator transaction binding the contract method 0xa870fcc7.
//
// Solidity: function grantRead(string mfileDid, string memoDid) returns()
func (_Proxy *ProxyTransactorSession) GrantRead(mfileDid string, memoDid string) (*types.Transaction, error) {
	return _Proxy.Contract.GrantRead(&_Proxy.TransactOpts, mfileDid, memoDid)
}

// RegisterMfileDid is a paid mutator transaction binding the contract method 0xca0258dc.
//
// Solidity: function registerMfileDid(string mfileDid, string encode, uint8 ftype, string controller, uint256 price, string[] keywords) returns()
func (_Proxy *ProxyTransactor) RegisterMfileDid(opts *bind.TransactOpts, mfileDid string, encode string, ftype uint8, controller string, price *big.Int, keywords []string) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "registerMfileDid", mfileDid, encode, ftype, controller, price, keywords)
}

// RegisterMfileDid is a paid mutator transaction binding the contract method 0xca0258dc.
//
// Solidity: function registerMfileDid(string mfileDid, string encode, uint8 ftype, string controller, uint256 price, string[] keywords) returns()
func (_Proxy *ProxySession) RegisterMfileDid(mfileDid string, encode string, ftype uint8, controller string, price *big.Int, keywords []string) (*types.Transaction, error) {
	return _Proxy.Contract.RegisterMfileDid(&_Proxy.TransactOpts, mfileDid, encode, ftype, controller, price, keywords)
}

// RegisterMfileDid is a paid mutator transaction binding the contract method 0xca0258dc.
//
// Solidity: function registerMfileDid(string mfileDid, string encode, uint8 ftype, string controller, uint256 price, string[] keywords) returns()
func (_Proxy *ProxyTransactorSession) RegisterMfileDid(mfileDid string, encode string, ftype uint8, controller string, price *big.Int, keywords []string) (*types.Transaction, error) {
	return _Proxy.Contract.RegisterMfileDid(&_Proxy.TransactOpts, mfileDid, encode, ftype, controller, price, keywords)
}

// RemoveAssertion is a paid mutator transaction binding the contract method 0x9aae7467.
//
// Solidity: function removeAssertion(string did, string id, bytes signature) returns()
func (_Proxy *ProxyTransactor) RemoveAssertion(opts *bind.TransactOpts, did string, id string, signature []byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "removeAssertion", did, id, signature)
}

// RemoveAssertion is a paid mutator transaction binding the contract method 0x9aae7467.
//
// Solidity: function removeAssertion(string did, string id, bytes signature) returns()
func (_Proxy *ProxySession) RemoveAssertion(did string, id string, signature []byte) (*types.Transaction, error) {
	return _Proxy.Contract.RemoveAssertion(&_Proxy.TransactOpts, did, id, signature)
}

// RemoveAssertion is a paid mutator transaction binding the contract method 0x9aae7467.
//
// Solidity: function removeAssertion(string did, string id, bytes signature) returns()
func (_Proxy *ProxyTransactorSession) RemoveAssertion(did string, id string, signature []byte) (*types.Transaction, error) {
	return _Proxy.Contract.RemoveAssertion(&_Proxy.TransactOpts, did, id, signature)
}

// RemoveAuth is a paid mutator transaction binding the contract method 0xeb785ee2.
//
// Solidity: function removeAuth(string did, string id, bytes signature) returns()
func (_Proxy *ProxyTransactor) RemoveAuth(opts *bind.TransactOpts, did string, id string, signature []byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "removeAuth", did, id, signature)
}

// RemoveAuth is a paid mutator transaction binding the contract method 0xeb785ee2.
//
// Solidity: function removeAuth(string did, string id, bytes signature) returns()
func (_Proxy *ProxySession) RemoveAuth(did string, id string, signature []byte) (*types.Transaction, error) {
	return _Proxy.Contract.RemoveAuth(&_Proxy.TransactOpts, did, id, signature)
}

// RemoveAuth is a paid mutator transaction binding the contract method 0xeb785ee2.
//
// Solidity: function removeAuth(string did, string id, bytes signature) returns()
func (_Proxy *ProxyTransactorSession) RemoveAuth(did string, id string, signature []byte) (*types.Transaction, error) {
	return _Proxy.Contract.RemoveAuth(&_Proxy.TransactOpts, did, id, signature)
}

// RemoveDelegation is a paid mutator transaction binding the contract method 0xd3be140e.
//
// Solidity: function removeDelegation(string did, string id, bytes signature) returns()
func (_Proxy *ProxyTransactor) RemoveDelegation(opts *bind.TransactOpts, did string, id string, signature []byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "removeDelegation", did, id, signature)
}

// RemoveDelegation is a paid mutator transaction binding the contract method 0xd3be140e.
//
// Solidity: function removeDelegation(string did, string id, bytes signature) returns()
func (_Proxy *ProxySession) RemoveDelegation(did string, id string, signature []byte) (*types.Transaction, error) {
	return _Proxy.Contract.RemoveDelegation(&_Proxy.TransactOpts, did, id, signature)
}

// RemoveDelegation is a paid mutator transaction binding the contract method 0xd3be140e.
//
// Solidity: function removeDelegation(string did, string id, bytes signature) returns()
func (_Proxy *ProxyTransactorSession) RemoveDelegation(did string, id string, signature []byte) (*types.Transaction, error) {
	return _Proxy.Contract.RemoveDelegation(&_Proxy.TransactOpts, did, id, signature)
}

// RemoveRecovery is a paid mutator transaction binding the contract method 0xdc649bf4.
//
// Solidity: function removeRecovery(string did, string id, bytes signature) returns()
func (_Proxy *ProxyTransactor) RemoveRecovery(opts *bind.TransactOpts, did string, id string, signature []byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "removeRecovery", did, id, signature)
}

// RemoveRecovery is a paid mutator transaction binding the contract method 0xdc649bf4.
//
// Solidity: function removeRecovery(string did, string id, bytes signature) returns()
func (_Proxy *ProxySession) RemoveRecovery(did string, id string, signature []byte) (*types.Transaction, error) {
	return _Proxy.Contract.RemoveRecovery(&_Proxy.TransactOpts, did, id, signature)
}

// RemoveRecovery is a paid mutator transaction binding the contract method 0xdc649bf4.
//
// Solidity: function removeRecovery(string did, string id, bytes signature) returns()
func (_Proxy *ProxyTransactorSession) RemoveRecovery(did string, id string, signature []byte) (*types.Transaction, error) {
	return _Proxy.Contract.RemoveRecovery(&_Proxy.TransactOpts, did, id, signature)
}

// UpdateVeri is a paid mutator transaction binding the contract method 0x6c95dd03.
//
// Solidity: function updateVeri(string did, uint256 index, string methodType, bytes pubKeyData, bytes signature) returns()
func (_Proxy *ProxyTransactor) UpdateVeri(opts *bind.TransactOpts, did string, index *big.Int, methodType string, pubKeyData []byte, signature []byte) (*types.Transaction, error) {
	return _Proxy.contract.Transact(opts, "updateVeri", did, index, methodType, pubKeyData, signature)
}

// UpdateVeri is a paid mutator transaction binding the contract method 0x6c95dd03.
//
// Solidity: function updateVeri(string did, uint256 index, string methodType, bytes pubKeyData, bytes signature) returns()
func (_Proxy *ProxySession) UpdateVeri(did string, index *big.Int, methodType string, pubKeyData []byte, signature []byte) (*types.Transaction, error) {
	return _Proxy.Contract.UpdateVeri(&_Proxy.TransactOpts, did, index, methodType, pubKeyData, signature)
}

// UpdateVeri is a paid mutator transaction binding the contract method 0x6c95dd03.
//
// Solidity: function updateVeri(string did, uint256 index, string methodType, bytes pubKeyData, bytes signature) returns()
func (_Proxy *ProxyTransactorSession) UpdateVeri(did string, index *big.Int, methodType string, pubKeyData []byte, signature []byte) (*types.Transaction, error) {
	return _Proxy.Contract.UpdateVeri(&_Proxy.TransactOpts, did, index, methodType, pubKeyData, signature)
}

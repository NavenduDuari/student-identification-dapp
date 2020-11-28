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

// IdentityAcademicDetailsI is an auto generated low-level Go binding around an user-defined struct.
type IdentityAcademicDetailsI struct {
	DegreeId string
}

// IdentityABI is the input ABI used to generate the binding from.
const IdentityABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"studentId\",\"type\":\"string\"}],\"name\":\"getAcademicDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"degreeId\",\"type\":\"string\"}],\"internalType\":\"structIdentity.AcademicDetailsI\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"studentId\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"string\",\"name\":\"degreeId\",\"type\":\"string\"}],\"internalType\":\"structIdentity.AcademicDetailsI\",\"name\":\"academicDetails\",\"type\":\"tuple\"}],\"name\":\"setAcademicDetails\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"studentIdVsAcademicDetails\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"degreeId\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IdentityFuncSigs maps the 4-byte function signature to its string representation.
var IdentityFuncSigs = map[string]string{
	"8dc75241": "getAcademicDetails(string)",
	"9e287953": "setAcademicDetails(string,(string))",
	"64e5bdb2": "studentIdVsAcademicDetails(string)",
}

// IdentityBin is the compiled bytecode used for deploying new contracts.
var IdentityBin = "0x608060405234801561001057600080fd5b506105e9806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806364e5bdb2146100465780638dc752411461006f5780639e2879531461008f575b600080fd5b61005961005436600461039c565b6100a4565b60405161006691906104dd565b60405180910390f35b61008261007d36600461039c565b610151565b60405161006691906104ee565b6100a261009d3660046103d9565b61020e565b005b805160208183018101805160008252928201938201939093209190925280546040805160026001841615610100026000190190931692909204601f81018590048502830185019091528082529192909183918301828280156101475780601f1061011c57610100808354040283529160200191610147565b820191906000526020600020905b81548152906001019060200180831161012a57829003601f168201915b5050505050905081565b610159610250565b60008260405161016991906104ca565b9081526040805160209281900383018120805460026001821615610100026000190190911604601f810185900485028301840184529382018481529193909284929184918401828280156101fe5780601f106101d3576101008083540402835291602001916101fe565b820191906000526020600020905b8154815290600101906020018083116101e157829003601f168201915b5050505050815250509050919050565b8060008360405161021f91906104ca565b90815260200160405180910390206000820151816000019080519060200190610249929190610263565b5050505050565b6040518060200160405280606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106102a457805160ff19168380011785556102d1565b828001600101855582156102d1579182015b828111156102d15782518255916020019190600101906102b6565b506102dd9291506102e1565b5090565b6102fb91905b808211156102dd57600081556001016102e7565b90565b600082601f83011261030f57600080fd5b813561032261031d82610526565b6104ff565b9150808252602083016020830185838301111561033e57600080fd5b610349838284610560565b50505092915050565b60006020828403121561036457600080fd5b61036e60206104ff565b9050813567ffffffffffffffff81111561038757600080fd5b610393848285016102fe565b82525092915050565b6000602082840312156103ae57600080fd5b813567ffffffffffffffff8111156103c557600080fd5b6103d1848285016102fe565b949350505050565b600080604083850312156103ec57600080fd5b823567ffffffffffffffff81111561040357600080fd5b61040f858286016102fe565b925050602083013567ffffffffffffffff81111561042c57600080fd5b61043885828601610352565b9150509250929050565b600061044d8261054e565b610457818561055b565b935061046781856020860161056c565b9290920192915050565b600061047c8261054e565b6104868185610552565b935061049681856020860161056c565b61049f8161059c565b9093019392505050565b80516020808452600091908401906104c18282610471565b95945050505050565b60006104d68284610442565b9392505050565b602080825281016104d68184610471565b602080825281016104d681846104a9565b60405181810167ffffffffffffffff8111828210171561051e57600080fd5b604052919050565b600067ffffffffffffffff82111561053d57600080fd5b506020601f91909101601f19160190565b5190565b90815260200190565b919050565b82818337506000910152565b60005b8381101561058757818101518382015260200161056f565b83811115610596576000848401525b50505050565b601f01601f19169056fea365627a7a72315820bdd3c1c92c00cf7e915e16e0e79210eebb0c25ad6ef17bdebb67fa8d09d13a256c6578706572696d656e74616cf564736f6c63430005100040"

// DeployIdentity deploys a new Ethereum contract, binding an instance of Identity to it.
func DeployIdentity(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Identity, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IdentityBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Identity{IdentityCaller: IdentityCaller{contract: contract}, IdentityTransactor: IdentityTransactor{contract: contract}, IdentityFilterer: IdentityFilterer{contract: contract}}, nil
}

// Identity is an auto generated Go binding around an Ethereum contract.
type Identity struct {
	IdentityCaller     // Read-only binding to the contract
	IdentityTransactor // Write-only binding to the contract
	IdentityFilterer   // Log filterer for contract events
}

// IdentityCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdentityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdentityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdentityFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentitySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdentitySession struct {
	Contract     *Identity         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdentityCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdentityCallerSession struct {
	Contract *IdentityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IdentityTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdentityTransactorSession struct {
	Contract     *IdentityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IdentityRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdentityRaw struct {
	Contract *Identity // Generic contract binding to access the raw methods on
}

// IdentityCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdentityCallerRaw struct {
	Contract *IdentityCaller // Generic read-only contract binding to access the raw methods on
}

// IdentityTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdentityTransactorRaw struct {
	Contract *IdentityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdentity creates a new instance of Identity, bound to a specific deployed contract.
func NewIdentity(address common.Address, backend bind.ContractBackend) (*Identity, error) {
	contract, err := bindIdentity(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Identity{IdentityCaller: IdentityCaller{contract: contract}, IdentityTransactor: IdentityTransactor{contract: contract}, IdentityFilterer: IdentityFilterer{contract: contract}}, nil
}

// NewIdentityCaller creates a new read-only instance of Identity, bound to a specific deployed contract.
func NewIdentityCaller(address common.Address, caller bind.ContractCaller) (*IdentityCaller, error) {
	contract, err := bindIdentity(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityCaller{contract: contract}, nil
}

// NewIdentityTransactor creates a new write-only instance of Identity, bound to a specific deployed contract.
func NewIdentityTransactor(address common.Address, transactor bind.ContractTransactor) (*IdentityTransactor, error) {
	contract, err := bindIdentity(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityTransactor{contract: contract}, nil
}

// NewIdentityFilterer creates a new log filterer instance of Identity, bound to a specific deployed contract.
func NewIdentityFilterer(address common.Address, filterer bind.ContractFilterer) (*IdentityFilterer, error) {
	contract, err := bindIdentity(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdentityFilterer{contract: contract}, nil
}

// bindIdentity binds a generic wrapper to an already deployed contract.
func bindIdentity(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Identity *IdentityRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Identity.Contract.IdentityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Identity *IdentityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identity.Contract.IdentityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Identity *IdentityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Identity.Contract.IdentityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Identity *IdentityCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Identity.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Identity *IdentityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identity.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Identity *IdentityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Identity.Contract.contract.Transact(opts, method, params...)
}

// GetAcademicDetails is a free data retrieval call binding the contract method 0x8dc75241.
//
// Solidity: function getAcademicDetails(string studentId) view returns((string))
func (_Identity *IdentityCaller) GetAcademicDetails(opts *bind.CallOpts, studentId string) (IdentityAcademicDetailsI, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "getAcademicDetails", studentId)

	if err != nil {
		return *new(IdentityAcademicDetailsI), err
	}

	out0 := *abi.ConvertType(out[0], new(IdentityAcademicDetailsI)).(*IdentityAcademicDetailsI)

	return out0, err

}

// GetAcademicDetails is a free data retrieval call binding the contract method 0x8dc75241.
//
// Solidity: function getAcademicDetails(string studentId) view returns((string))
func (_Identity *IdentitySession) GetAcademicDetails(studentId string) (IdentityAcademicDetailsI, error) {
	return _Identity.Contract.GetAcademicDetails(&_Identity.CallOpts, studentId)
}

// GetAcademicDetails is a free data retrieval call binding the contract method 0x8dc75241.
//
// Solidity: function getAcademicDetails(string studentId) view returns((string))
func (_Identity *IdentityCallerSession) GetAcademicDetails(studentId string) (IdentityAcademicDetailsI, error) {
	return _Identity.Contract.GetAcademicDetails(&_Identity.CallOpts, studentId)
}

// StudentIdVsAcademicDetails is a free data retrieval call binding the contract method 0x64e5bdb2.
//
// Solidity: function studentIdVsAcademicDetails(string ) view returns(string degreeId)
func (_Identity *IdentityCaller) StudentIdVsAcademicDetails(opts *bind.CallOpts, arg0 string) (string, error) {
	var out []interface{}
	err := _Identity.contract.Call(opts, &out, "studentIdVsAcademicDetails", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// StudentIdVsAcademicDetails is a free data retrieval call binding the contract method 0x64e5bdb2.
//
// Solidity: function studentIdVsAcademicDetails(string ) view returns(string degreeId)
func (_Identity *IdentitySession) StudentIdVsAcademicDetails(arg0 string) (string, error) {
	return _Identity.Contract.StudentIdVsAcademicDetails(&_Identity.CallOpts, arg0)
}

// StudentIdVsAcademicDetails is a free data retrieval call binding the contract method 0x64e5bdb2.
//
// Solidity: function studentIdVsAcademicDetails(string ) view returns(string degreeId)
func (_Identity *IdentityCallerSession) StudentIdVsAcademicDetails(arg0 string) (string, error) {
	return _Identity.Contract.StudentIdVsAcademicDetails(&_Identity.CallOpts, arg0)
}

// SetAcademicDetails is a paid mutator transaction binding the contract method 0x9e287953.
//
// Solidity: function setAcademicDetails(string studentId, (string) academicDetails) returns()
func (_Identity *IdentityTransactor) SetAcademicDetails(opts *bind.TransactOpts, studentId string, academicDetails IdentityAcademicDetailsI) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "setAcademicDetails", studentId, academicDetails)
}

// SetAcademicDetails is a paid mutator transaction binding the contract method 0x9e287953.
//
// Solidity: function setAcademicDetails(string studentId, (string) academicDetails) returns()
func (_Identity *IdentitySession) SetAcademicDetails(studentId string, academicDetails IdentityAcademicDetailsI) (*types.Transaction, error) {
	return _Identity.Contract.SetAcademicDetails(&_Identity.TransactOpts, studentId, academicDetails)
}

// SetAcademicDetails is a paid mutator transaction binding the contract method 0x9e287953.
//
// Solidity: function setAcademicDetails(string studentId, (string) academicDetails) returns()
func (_Identity *IdentityTransactorSession) SetAcademicDetails(studentId string, academicDetails IdentityAcademicDetailsI) (*types.Transaction, error) {
	return _Identity.Contract.SetAcademicDetails(&_Identity.TransactOpts, studentId, academicDetails)
}

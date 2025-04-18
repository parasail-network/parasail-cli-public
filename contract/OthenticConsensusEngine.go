// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// TaskInfo is an auto generated low-level Go binding around an user-defined struct.
type TaskInfo struct {
	ProofOfTask      string
	Data             []byte
	TaskPerformer    common.Address
	TaskDefinitionId uint16
}

// OthenticConsensusEngineMetaData contains all meta data concerning the OthenticConsensusEngine contract.
var OthenticConsensusEngineMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"},{\"internalType\":\"contractIServiceMonitor\",\"name\":\"_serviceMonitor\",\"type\":\"address\"},{\"internalType\":\"contractIAttestationCenter\",\"name\":\"_attestationCenter\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"proofOfTask\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"taskPerformer\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"taskDefinitionId\",\"type\":\"uint16\"}],\"internalType\":\"structTaskInfo\",\"name\":\"_taskInfo\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"_isApproved\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"},{\"internalType\":\"uint256[2]\",\"name\":\"\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"name\":\"afterTaskSubmission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"attestationCenter\",\"outputs\":[{\"internalType\":\"contractIAttestationCenter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"proofOfTask\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"taskPerformer\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"taskDefinitionId\",\"type\":\"uint16\"}],\"internalType\":\"structTaskInfo\",\"name\":\"_taskInfo\",\"type\":\"tuple\"},{\"internalType\":\"bool\",\"name\":\"_isApproved\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"_tpSignature\",\"type\":\"bytes\"},{\"internalType\":\"uint256[2]\",\"name\":\"_taSignature\",\"type\":\"uint256[2]\"},{\"internalType\":\"uint256[]\",\"name\":\"_operatorIds\",\"type\":\"uint256[]\"}],\"name\":\"beforeTaskSubmission\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"serviceMonitor\",\"outputs\":[{\"internalType\":\"contractIServiceMonitor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIAttestationCenter\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setAttestationCenter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60803460b957601f61061438819003918201601f19168301916001600160401b0383118484101760bd5780849260609460405283398101031260b95780516001600160a01b038116919082900360b95760208101516001600160a01b038116919082900360b957604001516001600160a01b038116929083900360b95760018060a01b0319600254161760025560018060a01b0319600154161760015560018060a01b03195f5416175f5560405161054290816100d28239f35b5f80fd5b634e487b7160e01b5f52604160045260245ffdfe6080806040526004361015610012575f80fd5b5f905f3560e01c90816311c693111461033457508063502f5bd01461029c5780639f4f635814610274578063d92807a21461024d5763dd1a538714610055575f80fd5b346102355760c0366003190112610235576004356001600160401b038111610235576080600319823603011261023557604051608081018181106001600160401b038211176102395760405281600401356001600160401b03811161023557820136602382011215610235576100d5903690602460048201359101610463565b815260248201356001600160401b038111610235578201366023820112156102355761010b903690602460048201359101610463565b60208201908152916044810135906001600160a01b038216820361023557606491604084015201359061ffff82168203610235576060015261014b6103d6565b6044356001600160401b0381116102355761016a9036906004016103e5565b50503660a4116102355760a4356001600160401b03811161023557610193903690600401610412565b50506101a960018060a01b035f541633146104a8565b6101b1575080f35b6001549051906001600160a01b0316803b156102355760205f9160449383604051958680958194631982a10f60e01b83528160048401528051918291826024860152018484015e8181018301849052601f01601f191681010301925af1801561022a5761021c575080f35b61022891505f90610442565b005b6040513d5f823e3d90fd5b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b34610235575f366003190112610235575f546040516001600160a01b039091168152602090f35b34610235575f366003190112610235576001546040516001600160a01b039091168152602090f35b346102355760c0366003190112610235576004356001600160401b038111610235576080906003199036030112610235576102d56103d6565b506044356001600160401b038111610235576102f59036906004016103e5565b50503660a4116102355760a4356001600160401b0381116102355761031e903690600401610412565b505061022860018060a01b035f541633146104a8565b34610235576020366003190112610235576004356001600160a01b0381169190829003610235576002546001600160a01b0316330361038857506bffffffffffffffffffffffff60a01b5f5416175f555f80f35b62461bcd60e51b8152602060048201526024808201527f4f6e6c79206f70657261746f722063616e2063616c6c20746869732066756e636044820152633a34b7b760e11b6064820152608490fd5b60243590811515820361023557565b9181601f84011215610235578235916001600160401b038311610235576020838186019501011161023557565b9181601f84011215610235578235916001600160401b038311610235576020808501948460051b01011161023557565b90601f801991011681019081106001600160401b0382111761023957604052565b9291926001600160401b038211610239576040519161048c601f8201601f191660200184610442565b829481845281830111610235578281602093845f960137010152565b156104af57565b60405162461bcd60e51b815260206004820152602f60248201527f4f6e6c7920617474657374616174696f6e2063656e7465722063616e2063616c60448201526e36103a3434b990333ab731ba34b7b760891b6064820152608490fdfea26469706673582212209650f5fcd76c32a1e85ca266a2a883720e671d55c2f79326a7eba2fcb8dbe5d864736f6c634300081c0033",
}

// OthenticConsensusEngineABI is the input ABI used to generate the binding from.
// Deprecated: Use OthenticConsensusEngineMetaData.ABI instead.
var OthenticConsensusEngineABI = OthenticConsensusEngineMetaData.ABI

// OthenticConsensusEngineBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OthenticConsensusEngineMetaData.Bin instead.
var OthenticConsensusEngineBin = OthenticConsensusEngineMetaData.Bin

// DeployOthenticConsensusEngine deploys a new Ethereum contract, binding an instance of OthenticConsensusEngine to it.
func DeployOthenticConsensusEngine(auth *bind.TransactOpts, backend bind.ContractBackend, _manager common.Address, _serviceMonitor common.Address, _attestationCenter common.Address) (common.Address, *types.Transaction, *OthenticConsensusEngine, error) {
	parsed, err := OthenticConsensusEngineMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OthenticConsensusEngineBin), backend, _manager, _serviceMonitor, _attestationCenter)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OthenticConsensusEngine{OthenticConsensusEngineCaller: OthenticConsensusEngineCaller{contract: contract}, OthenticConsensusEngineTransactor: OthenticConsensusEngineTransactor{contract: contract}, OthenticConsensusEngineFilterer: OthenticConsensusEngineFilterer{contract: contract}}, nil
}

// OthenticConsensusEngine is an auto generated Go binding around an Ethereum contract.
type OthenticConsensusEngine struct {
	OthenticConsensusEngineCaller     // Read-only binding to the contract
	OthenticConsensusEngineTransactor // Write-only binding to the contract
	OthenticConsensusEngineFilterer   // Log filterer for contract events
}

// OthenticConsensusEngineCaller is an auto generated read-only Go binding around an Ethereum contract.
type OthenticConsensusEngineCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OthenticConsensusEngineTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OthenticConsensusEngineTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OthenticConsensusEngineFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OthenticConsensusEngineFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OthenticConsensusEngineSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OthenticConsensusEngineSession struct {
	Contract     *OthenticConsensusEngine // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// OthenticConsensusEngineCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OthenticConsensusEngineCallerSession struct {
	Contract *OthenticConsensusEngineCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// OthenticConsensusEngineTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OthenticConsensusEngineTransactorSession struct {
	Contract     *OthenticConsensusEngineTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// OthenticConsensusEngineRaw is an auto generated low-level Go binding around an Ethereum contract.
type OthenticConsensusEngineRaw struct {
	Contract *OthenticConsensusEngine // Generic contract binding to access the raw methods on
}

// OthenticConsensusEngineCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OthenticConsensusEngineCallerRaw struct {
	Contract *OthenticConsensusEngineCaller // Generic read-only contract binding to access the raw methods on
}

// OthenticConsensusEngineTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OthenticConsensusEngineTransactorRaw struct {
	Contract *OthenticConsensusEngineTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOthenticConsensusEngine creates a new instance of OthenticConsensusEngine, bound to a specific deployed contract.
func NewOthenticConsensusEngine(address common.Address, backend bind.ContractBackend) (*OthenticConsensusEngine, error) {
	contract, err := bindOthenticConsensusEngine(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OthenticConsensusEngine{OthenticConsensusEngineCaller: OthenticConsensusEngineCaller{contract: contract}, OthenticConsensusEngineTransactor: OthenticConsensusEngineTransactor{contract: contract}, OthenticConsensusEngineFilterer: OthenticConsensusEngineFilterer{contract: contract}}, nil
}

// NewOthenticConsensusEngineCaller creates a new read-only instance of OthenticConsensusEngine, bound to a specific deployed contract.
func NewOthenticConsensusEngineCaller(address common.Address, caller bind.ContractCaller) (*OthenticConsensusEngineCaller, error) {
	contract, err := bindOthenticConsensusEngine(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OthenticConsensusEngineCaller{contract: contract}, nil
}

// NewOthenticConsensusEngineTransactor creates a new write-only instance of OthenticConsensusEngine, bound to a specific deployed contract.
func NewOthenticConsensusEngineTransactor(address common.Address, transactor bind.ContractTransactor) (*OthenticConsensusEngineTransactor, error) {
	contract, err := bindOthenticConsensusEngine(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OthenticConsensusEngineTransactor{contract: contract}, nil
}

// NewOthenticConsensusEngineFilterer creates a new log filterer instance of OthenticConsensusEngine, bound to a specific deployed contract.
func NewOthenticConsensusEngineFilterer(address common.Address, filterer bind.ContractFilterer) (*OthenticConsensusEngineFilterer, error) {
	contract, err := bindOthenticConsensusEngine(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OthenticConsensusEngineFilterer{contract: contract}, nil
}

// bindOthenticConsensusEngine binds a generic wrapper to an already deployed contract.
func bindOthenticConsensusEngine(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OthenticConsensusEngineMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OthenticConsensusEngine *OthenticConsensusEngineRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OthenticConsensusEngine.Contract.OthenticConsensusEngineCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OthenticConsensusEngine *OthenticConsensusEngineRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OthenticConsensusEngine.Contract.OthenticConsensusEngineTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OthenticConsensusEngine *OthenticConsensusEngineRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OthenticConsensusEngine.Contract.OthenticConsensusEngineTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OthenticConsensusEngine *OthenticConsensusEngineCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OthenticConsensusEngine.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OthenticConsensusEngine *OthenticConsensusEngineTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OthenticConsensusEngine.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OthenticConsensusEngine *OthenticConsensusEngineTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OthenticConsensusEngine.Contract.contract.Transact(opts, method, params...)
}

// AttestationCenter is a free data retrieval call binding the contract method 0xd92807a2.
//
// Solidity: function attestationCenter() view returns(address)
func (_OthenticConsensusEngine *OthenticConsensusEngineCaller) AttestationCenter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OthenticConsensusEngine.contract.Call(opts, &out, "attestationCenter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AttestationCenter is a free data retrieval call binding the contract method 0xd92807a2.
//
// Solidity: function attestationCenter() view returns(address)
func (_OthenticConsensusEngine *OthenticConsensusEngineSession) AttestationCenter() (common.Address, error) {
	return _OthenticConsensusEngine.Contract.AttestationCenter(&_OthenticConsensusEngine.CallOpts)
}

// AttestationCenter is a free data retrieval call binding the contract method 0xd92807a2.
//
// Solidity: function attestationCenter() view returns(address)
func (_OthenticConsensusEngine *OthenticConsensusEngineCallerSession) AttestationCenter() (common.Address, error) {
	return _OthenticConsensusEngine.Contract.AttestationCenter(&_OthenticConsensusEngine.CallOpts)
}

// ServiceMonitor is a free data retrieval call binding the contract method 0x9f4f6358.
//
// Solidity: function serviceMonitor() view returns(address)
func (_OthenticConsensusEngine *OthenticConsensusEngineCaller) ServiceMonitor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OthenticConsensusEngine.contract.Call(opts, &out, "serviceMonitor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ServiceMonitor is a free data retrieval call binding the contract method 0x9f4f6358.
//
// Solidity: function serviceMonitor() view returns(address)
func (_OthenticConsensusEngine *OthenticConsensusEngineSession) ServiceMonitor() (common.Address, error) {
	return _OthenticConsensusEngine.Contract.ServiceMonitor(&_OthenticConsensusEngine.CallOpts)
}

// ServiceMonitor is a free data retrieval call binding the contract method 0x9f4f6358.
//
// Solidity: function serviceMonitor() view returns(address)
func (_OthenticConsensusEngine *OthenticConsensusEngineCallerSession) ServiceMonitor() (common.Address, error) {
	return _OthenticConsensusEngine.Contract.ServiceMonitor(&_OthenticConsensusEngine.CallOpts)
}

// AfterTaskSubmission is a paid mutator transaction binding the contract method 0xdd1a5387.
//
// Solidity: function afterTaskSubmission((string,bytes,address,uint16) _taskInfo, bool _isApproved, bytes , uint256[2] , uint256[] ) returns()
func (_OthenticConsensusEngine *OthenticConsensusEngineTransactor) AfterTaskSubmission(opts *bind.TransactOpts, _taskInfo TaskInfo, _isApproved bool, arg2 []byte, arg3 [2]*big.Int, arg4 []*big.Int) (*types.Transaction, error) {
	return _OthenticConsensusEngine.contract.Transact(opts, "afterTaskSubmission", _taskInfo, _isApproved, arg2, arg3, arg4)
}

// AfterTaskSubmission is a paid mutator transaction binding the contract method 0xdd1a5387.
//
// Solidity: function afterTaskSubmission((string,bytes,address,uint16) _taskInfo, bool _isApproved, bytes , uint256[2] , uint256[] ) returns()
func (_OthenticConsensusEngine *OthenticConsensusEngineSession) AfterTaskSubmission(_taskInfo TaskInfo, _isApproved bool, arg2 []byte, arg3 [2]*big.Int, arg4 []*big.Int) (*types.Transaction, error) {
	return _OthenticConsensusEngine.Contract.AfterTaskSubmission(&_OthenticConsensusEngine.TransactOpts, _taskInfo, _isApproved, arg2, arg3, arg4)
}

// AfterTaskSubmission is a paid mutator transaction binding the contract method 0xdd1a5387.
//
// Solidity: function afterTaskSubmission((string,bytes,address,uint16) _taskInfo, bool _isApproved, bytes , uint256[2] , uint256[] ) returns()
func (_OthenticConsensusEngine *OthenticConsensusEngineTransactorSession) AfterTaskSubmission(_taskInfo TaskInfo, _isApproved bool, arg2 []byte, arg3 [2]*big.Int, arg4 []*big.Int) (*types.Transaction, error) {
	return _OthenticConsensusEngine.Contract.AfterTaskSubmission(&_OthenticConsensusEngine.TransactOpts, _taskInfo, _isApproved, arg2, arg3, arg4)
}

// BeforeTaskSubmission is a paid mutator transaction binding the contract method 0x502f5bd0.
//
// Solidity: function beforeTaskSubmission((string,bytes,address,uint16) _taskInfo, bool _isApproved, bytes _tpSignature, uint256[2] _taSignature, uint256[] _operatorIds) returns()
func (_OthenticConsensusEngine *OthenticConsensusEngineTransactor) BeforeTaskSubmission(opts *bind.TransactOpts, _taskInfo TaskInfo, _isApproved bool, _tpSignature []byte, _taSignature [2]*big.Int, _operatorIds []*big.Int) (*types.Transaction, error) {
	return _OthenticConsensusEngine.contract.Transact(opts, "beforeTaskSubmission", _taskInfo, _isApproved, _tpSignature, _taSignature, _operatorIds)
}

// BeforeTaskSubmission is a paid mutator transaction binding the contract method 0x502f5bd0.
//
// Solidity: function beforeTaskSubmission((string,bytes,address,uint16) _taskInfo, bool _isApproved, bytes _tpSignature, uint256[2] _taSignature, uint256[] _operatorIds) returns()
func (_OthenticConsensusEngine *OthenticConsensusEngineSession) BeforeTaskSubmission(_taskInfo TaskInfo, _isApproved bool, _tpSignature []byte, _taSignature [2]*big.Int, _operatorIds []*big.Int) (*types.Transaction, error) {
	return _OthenticConsensusEngine.Contract.BeforeTaskSubmission(&_OthenticConsensusEngine.TransactOpts, _taskInfo, _isApproved, _tpSignature, _taSignature, _operatorIds)
}

// BeforeTaskSubmission is a paid mutator transaction binding the contract method 0x502f5bd0.
//
// Solidity: function beforeTaskSubmission((string,bytes,address,uint16) _taskInfo, bool _isApproved, bytes _tpSignature, uint256[2] _taSignature, uint256[] _operatorIds) returns()
func (_OthenticConsensusEngine *OthenticConsensusEngineTransactorSession) BeforeTaskSubmission(_taskInfo TaskInfo, _isApproved bool, _tpSignature []byte, _taSignature [2]*big.Int, _operatorIds []*big.Int) (*types.Transaction, error) {
	return _OthenticConsensusEngine.Contract.BeforeTaskSubmission(&_OthenticConsensusEngine.TransactOpts, _taskInfo, _isApproved, _tpSignature, _taSignature, _operatorIds)
}

// SetAttestationCenter is a paid mutator transaction binding the contract method 0x11c69311.
//
// Solidity: function setAttestationCenter(address addr) returns()
func (_OthenticConsensusEngine *OthenticConsensusEngineTransactor) SetAttestationCenter(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _OthenticConsensusEngine.contract.Transact(opts, "setAttestationCenter", addr)
}

// SetAttestationCenter is a paid mutator transaction binding the contract method 0x11c69311.
//
// Solidity: function setAttestationCenter(address addr) returns()
func (_OthenticConsensusEngine *OthenticConsensusEngineSession) SetAttestationCenter(addr common.Address) (*types.Transaction, error) {
	return _OthenticConsensusEngine.Contract.SetAttestationCenter(&_OthenticConsensusEngine.TransactOpts, addr)
}

// SetAttestationCenter is a paid mutator transaction binding the contract method 0x11c69311.
//
// Solidity: function setAttestationCenter(address addr) returns()
func (_OthenticConsensusEngine *OthenticConsensusEngineTransactorSession) SetAttestationCenter(addr common.Address) (*types.Transaction, error) {
	return _OthenticConsensusEngine.Contract.SetAttestationCenter(&_OthenticConsensusEngine.TransactOpts, addr)
}

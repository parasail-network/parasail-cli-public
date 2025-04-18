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

// ServiceMonitorContractMetaData contains all meta data concerning the ServiceMonitorContract contract.
var ServiceMonitorContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vaultManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_consensusEngineAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_checkInterval\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"vaultAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"dealID\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"failTime\",\"type\":\"uint256\"}],\"name\":\"SLAFailure\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"checkInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"consensusEngine\",\"outputs\":[{\"internalType\":\"contractIConsensusEngine\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCheckInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRoundNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"}],\"name\":\"getRoundTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"roundTimestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_consensusEngineAddress\",\"type\":\"address\"}],\"name\":\"setConsensusEngine\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"slaResults\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"submitSLACheck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultManager\",\"outputs\":[{\"internalType\":\"contractIVaultManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080346100fd57601f610c2438819003918201601f19168301916001600160401b03831184841017610101578084926060946040528339810103126100fd5761004781610115565b90604061005660208301610115565b9101519033156100ea575f8054336001600160a01b0319821681178355604051959290916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09080a360018055600280546001600160a01b039283166001600160a01b03199182161790915560038054939092169216919091179055600455610afa908161012a8239f35b631e4fbdf760e01b5f525f60045260245ffd5b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b03821682036100fd5756fe6080806040526004361015610012575f80fd5b5f3560e01c9081630cde045614610962575080630d8e6e2c1461091257806312725e78146102345780631982a10f1461031457806321762e8e146102c55780634e2786fb146102a8578063715018a614610251578063838c29b6146102345780638a4adf241461020c5780638da5cb5b146101e5578063904a2b38146101bd578063932329c91461014a5780639ba6a02e14610174578063bb9467291461014a5763f2fde38b146100c1575f80fd5b34610146576020366003190112610146576100da6109f2565b6100e2610a9e565b6001600160a01b03168015610133575f80546001600160a01b03198116831782556001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09080a3005b631e4fbdf760e01b5f525f60045260245ffd5b5f80fd5b34610146576020366003190112610146576004355f526005602052602060405f2054604051908152f35b34610146576040366003190112610146576001600160a01b036101956109f2565b165f52600760205260405f206024355f52602052602060ff60405f2054166040519015158152f35b34610146575f366003190112610146576003546040516001600160a01b039091168152602090f35b34610146575f366003190112610146575f546040516001600160a01b039091168152602090f35b34610146575f366003190112610146576002546040516001600160a01b039091168152602090f35b34610146575f366003190112610146576020600454604051908152f35b34610146575f36600319011261014657610269610a9e565b5f80546001600160a01b0319811682556001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b34610146575f366003190112610146576020600654604051908152f35b34610146576020366003190112610146576102de6109f2565b6102e6610a9e565b6102ee610a7e565b600380546001600160a01b0319166001600160a01b039290921691909117905560018055005b346101465760203660031901126101465760043567ffffffffffffffff81116101465736602382011215610146578060040135610350816109d6565b9061035e60405192836109a0565b8082526020820192366024838301011161014657815f926024602093018637830101526003546001600160a01b031633036108b95761039b610a7e565b805181019160a08260208501940312610146575160408201519091906001600160a01b0381169081900361014657606082015167ffffffffffffffff81116101465782019284603f850112156101465760208401516103f981610a08565b9461040760405196876109a0565b8186526020808088019360051b83010101908782116101465760408101925b8284106108765750505050608083015167ffffffffffffffff81116101465783019285603f8501121561014657602084015161046181610a08565b9461046f60405196876109a0565b8186526020808088019360051b830101019088821161014657604001915b8183106108555750505060a08101519067ffffffffffffffff8211610146570185603f82011215610146576020810151906104c782610a08565b966104d560405198896109a0565b828852602080808a019460051b84010101918183116101465760408101935b838510610812575050505050815f52600760205260405f20815f5260205260ff60405f2054166107b857815f52600760205260405f20815f5260205260405f20600160ff1982541617905560065481116107a1575b5f5b845181101561079b5761055e8186610a56565b51906001600160f81b03196105738287610a56565b5116916105808289610a56565b51600160f91b841461060e575b506002546001600160a01b031690813b15610146575f916105d39183604051809781958294636b7bcb2760e11b84528c600485015260606024850152606484019061097c565b90604483015203925af1918215610603576001926105f3575b500161054b565b5f6105fd916109a0565b876105ec565b6040513d5f823e3d90fd5b60025460405163d0c80f1360e01b815290602090829060049082906001600160a01b03165afa908115610603575f9161075a575b506001600160a01b0316801561072357803b15610146576106b25f92918392604051948580948193633004a04d60e11b83528d60048401524260248401528c604484015260c0606484015261069a60c484018b61097c565b908c60848501526003198483030160a485015261097c565b03925af1801561060357610713575b507fa45a414c42a1733f667335788c7c6dc45a98494fac69ea28b7eba7a3faf1327860405186815260806020820152806106fe608082018561097c565b8760408301524260608301520390a18861058d565b5f61071d916109a0565b886106c1565b60405162461bcd60e51b815260206004820152600f60248201526e14db185cda195c881b9bdd081cd95d608a1b6044820152606490fd5b90506020813d8211610793575b81610774602093836109a0565b8101031261014657516001600160a01b0381168103610146578a610642565b3d9150610767565b60018055005b805f5260056020524260405f205580600655610549565b60405162461bcd60e51b815260206004820152602c60248201527f534c4120666f72207468697320726f756e642068617320616c7265616479206260448201526b19595b881c995c1bdc9d195960a21b6064820152608490fd5b845167ffffffffffffffff81116101465760209083010183603f820112156101465760209161084a8583604086809601519101610a20565b8152019401936104f4565b82516001600160f81b0319811681036101465781526020928301920161048d565b835167ffffffffffffffff81116101465760209083010189603f82011215610146576020916108ae8b83604086809601519101610a20565b815201930192610426565b60405162461bcd60e51b815260206004820152602b60248201527f4f6e6c7920436f6e73656e737573456e67696e652063616e2063616c6c20746860448201526a34b990333ab731ba34b7b760a91b6064820152608490fd5b34610146575f3660031901126101465761095e6040516109336040826109a0565b600a8152693132ba3096989718171960b11b602082015260405191829160208352602083019061097c565b0390f35b34610146575f366003190112610146576020906006548152f35b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b90601f8019910116810190811067ffffffffffffffff8211176109c257604052565b634e487b7160e01b5f52604160045260245ffd5b67ffffffffffffffff81116109c257601f01601f191660200190565b600435906001600160a01b038216820361014657565b67ffffffffffffffff81116109c25760051b60200190565b929192610a2c826109d6565b91610a3a60405193846109a0565b829481845281830111610146578281602093845f96015e010152565b8051821015610a6a5760209160051b010190565b634e487b7160e01b5f52603260045260245ffd5b600260015414610a8f576002600155565b633ee5aeb560e01b5f5260045ffd5b5f546001600160a01b03163303610ab157565b63118cdaa760e01b5f523360045260245ffdfea264697066735822122084c1a2d6e3811c5376270f8f24cb6b73ff650b1d60925c8a20b01ecbd89520e964736f6c634300081c0033",
}

// ServiceMonitorContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ServiceMonitorContractMetaData.ABI instead.
var ServiceMonitorContractABI = ServiceMonitorContractMetaData.ABI

// ServiceMonitorContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ServiceMonitorContractMetaData.Bin instead.
var ServiceMonitorContractBin = ServiceMonitorContractMetaData.Bin

// DeployServiceMonitorContract deploys a new Ethereum contract, binding an instance of ServiceMonitorContract to it.
func DeployServiceMonitorContract(auth *bind.TransactOpts, backend bind.ContractBackend, _vaultManager common.Address, _consensusEngineAddress common.Address, _checkInterval *big.Int) (common.Address, *types.Transaction, *ServiceMonitorContract, error) {
	parsed, err := ServiceMonitorContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ServiceMonitorContractBin), backend, _vaultManager, _consensusEngineAddress, _checkInterval)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ServiceMonitorContract{ServiceMonitorContractCaller: ServiceMonitorContractCaller{contract: contract}, ServiceMonitorContractTransactor: ServiceMonitorContractTransactor{contract: contract}, ServiceMonitorContractFilterer: ServiceMonitorContractFilterer{contract: contract}}, nil
}

// ServiceMonitorContract is an auto generated Go binding around an Ethereum contract.
type ServiceMonitorContract struct {
	ServiceMonitorContractCaller     // Read-only binding to the contract
	ServiceMonitorContractTransactor // Write-only binding to the contract
	ServiceMonitorContractFilterer   // Log filterer for contract events
}

// ServiceMonitorContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ServiceMonitorContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServiceMonitorContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ServiceMonitorContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServiceMonitorContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ServiceMonitorContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ServiceMonitorContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ServiceMonitorContractSession struct {
	Contract     *ServiceMonitorContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts           // Call options to use throughout this session
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ServiceMonitorContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ServiceMonitorContractCallerSession struct {
	Contract *ServiceMonitorContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                 // Call options to use throughout this session
}

// ServiceMonitorContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ServiceMonitorContractTransactorSession struct {
	Contract     *ServiceMonitorContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// ServiceMonitorContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ServiceMonitorContractRaw struct {
	Contract *ServiceMonitorContract // Generic contract binding to access the raw methods on
}

// ServiceMonitorContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ServiceMonitorContractCallerRaw struct {
	Contract *ServiceMonitorContractCaller // Generic read-only contract binding to access the raw methods on
}

// ServiceMonitorContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ServiceMonitorContractTransactorRaw struct {
	Contract *ServiceMonitorContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewServiceMonitorContract creates a new instance of ServiceMonitorContract, bound to a specific deployed contract.
func NewServiceMonitorContract(address common.Address, backend bind.ContractBackend) (*ServiceMonitorContract, error) {
	contract, err := bindServiceMonitorContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ServiceMonitorContract{ServiceMonitorContractCaller: ServiceMonitorContractCaller{contract: contract}, ServiceMonitorContractTransactor: ServiceMonitorContractTransactor{contract: contract}, ServiceMonitorContractFilterer: ServiceMonitorContractFilterer{contract: contract}}, nil
}

// NewServiceMonitorContractCaller creates a new read-only instance of ServiceMonitorContract, bound to a specific deployed contract.
func NewServiceMonitorContractCaller(address common.Address, caller bind.ContractCaller) (*ServiceMonitorContractCaller, error) {
	contract, err := bindServiceMonitorContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ServiceMonitorContractCaller{contract: contract}, nil
}

// NewServiceMonitorContractTransactor creates a new write-only instance of ServiceMonitorContract, bound to a specific deployed contract.
func NewServiceMonitorContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ServiceMonitorContractTransactor, error) {
	contract, err := bindServiceMonitorContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ServiceMonitorContractTransactor{contract: contract}, nil
}

// NewServiceMonitorContractFilterer creates a new log filterer instance of ServiceMonitorContract, bound to a specific deployed contract.
func NewServiceMonitorContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ServiceMonitorContractFilterer, error) {
	contract, err := bindServiceMonitorContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ServiceMonitorContractFilterer{contract: contract}, nil
}

// bindServiceMonitorContract binds a generic wrapper to an already deployed contract.
func bindServiceMonitorContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ServiceMonitorContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ServiceMonitorContract *ServiceMonitorContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ServiceMonitorContract.Contract.ServiceMonitorContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ServiceMonitorContract *ServiceMonitorContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ServiceMonitorContract.Contract.ServiceMonitorContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ServiceMonitorContract *ServiceMonitorContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ServiceMonitorContract.Contract.ServiceMonitorContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ServiceMonitorContract *ServiceMonitorContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ServiceMonitorContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ServiceMonitorContract *ServiceMonitorContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ServiceMonitorContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ServiceMonitorContract *ServiceMonitorContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ServiceMonitorContract.Contract.contract.Transact(opts, method, params...)
}

// CheckInterval is a free data retrieval call binding the contract method 0x838c29b6.
//
// Solidity: function checkInterval() view returns(uint256)
func (_ServiceMonitorContract *ServiceMonitorContractCaller) CheckInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ServiceMonitorContract.contract.Call(opts, &out, "checkInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CheckInterval is a free data retrieval call binding the contract method 0x838c29b6.
//
// Solidity: function checkInterval() view returns(uint256)
func (_ServiceMonitorContract *ServiceMonitorContractSession) CheckInterval() (*big.Int, error) {
	return _ServiceMonitorContract.Contract.CheckInterval(&_ServiceMonitorContract.CallOpts)
}

// CheckInterval is a free data retrieval call binding the contract method 0x838c29b6.
//
// Solidity: function checkInterval() view returns(uint256)
func (_ServiceMonitorContract *ServiceMonitorContractCallerSession) CheckInterval() (*big.Int, error) {
	return _ServiceMonitorContract.Contract.CheckInterval(&_ServiceMonitorContract.CallOpts)
}

// ConsensusEngine is a free data retrieval call binding the contract method 0x904a2b38.
//
// Solidity: function consensusEngine() view returns(address)
func (_ServiceMonitorContract *ServiceMonitorContractCaller) ConsensusEngine(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ServiceMonitorContract.contract.Call(opts, &out, "consensusEngine")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConsensusEngine is a free data retrieval call binding the contract method 0x904a2b38.
//
// Solidity: function consensusEngine() view returns(address)
func (_ServiceMonitorContract *ServiceMonitorContractSession) ConsensusEngine() (common.Address, error) {
	return _ServiceMonitorContract.Contract.ConsensusEngine(&_ServiceMonitorContract.CallOpts)
}

// ConsensusEngine is a free data retrieval call binding the contract method 0x904a2b38.
//
// Solidity: function consensusEngine() view returns(address)
func (_ServiceMonitorContract *ServiceMonitorContractCallerSession) ConsensusEngine() (common.Address, error) {
	return _ServiceMonitorContract.Contract.ConsensusEngine(&_ServiceMonitorContract.CallOpts)
}

// GetCheckInterval is a free data retrieval call binding the contract method 0x12725e78.
//
// Solidity: function getCheckInterval() view returns(uint256)
func (_ServiceMonitorContract *ServiceMonitorContractCaller) GetCheckInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ServiceMonitorContract.contract.Call(opts, &out, "getCheckInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCheckInterval is a free data retrieval call binding the contract method 0x12725e78.
//
// Solidity: function getCheckInterval() view returns(uint256)
func (_ServiceMonitorContract *ServiceMonitorContractSession) GetCheckInterval() (*big.Int, error) {
	return _ServiceMonitorContract.Contract.GetCheckInterval(&_ServiceMonitorContract.CallOpts)
}

// GetCheckInterval is a free data retrieval call binding the contract method 0x12725e78.
//
// Solidity: function getCheckInterval() view returns(uint256)
func (_ServiceMonitorContract *ServiceMonitorContractCallerSession) GetCheckInterval() (*big.Int, error) {
	return _ServiceMonitorContract.Contract.GetCheckInterval(&_ServiceMonitorContract.CallOpts)
}

// GetRoundNumber is a free data retrieval call binding the contract method 0x0cde0456.
//
// Solidity: function getRoundNumber() view returns(uint256)
func (_ServiceMonitorContract *ServiceMonitorContractCaller) GetRoundNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ServiceMonitorContract.contract.Call(opts, &out, "getRoundNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoundNumber is a free data retrieval call binding the contract method 0x0cde0456.
//
// Solidity: function getRoundNumber() view returns(uint256)
func (_ServiceMonitorContract *ServiceMonitorContractSession) GetRoundNumber() (*big.Int, error) {
	return _ServiceMonitorContract.Contract.GetRoundNumber(&_ServiceMonitorContract.CallOpts)
}

// GetRoundNumber is a free data retrieval call binding the contract method 0x0cde0456.
//
// Solidity: function getRoundNumber() view returns(uint256)
func (_ServiceMonitorContract *ServiceMonitorContractCallerSession) GetRoundNumber() (*big.Int, error) {
	return _ServiceMonitorContract.Contract.GetRoundNumber(&_ServiceMonitorContract.CallOpts)
}

// GetRoundTimestamp is a free data retrieval call binding the contract method 0x932329c9.
//
// Solidity: function getRoundTimestamp(uint256 round) view returns(uint256)
func (_ServiceMonitorContract *ServiceMonitorContractCaller) GetRoundTimestamp(opts *bind.CallOpts, round *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ServiceMonitorContract.contract.Call(opts, &out, "getRoundTimestamp", round)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoundTimestamp is a free data retrieval call binding the contract method 0x932329c9.
//
// Solidity: function getRoundTimestamp(uint256 round) view returns(uint256)
func (_ServiceMonitorContract *ServiceMonitorContractSession) GetRoundTimestamp(round *big.Int) (*big.Int, error) {
	return _ServiceMonitorContract.Contract.GetRoundTimestamp(&_ServiceMonitorContract.CallOpts, round)
}

// GetRoundTimestamp is a free data retrieval call binding the contract method 0x932329c9.
//
// Solidity: function getRoundTimestamp(uint256 round) view returns(uint256)
func (_ServiceMonitorContract *ServiceMonitorContractCallerSession) GetRoundTimestamp(round *big.Int) (*big.Int, error) {
	return _ServiceMonitorContract.Contract.GetRoundTimestamp(&_ServiceMonitorContract.CallOpts, round)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(string)
func (_ServiceMonitorContract *ServiceMonitorContractCaller) GetVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ServiceMonitorContract.contract.Call(opts, &out, "getVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(string)
func (_ServiceMonitorContract *ServiceMonitorContractSession) GetVersion() (string, error) {
	return _ServiceMonitorContract.Contract.GetVersion(&_ServiceMonitorContract.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() pure returns(string)
func (_ServiceMonitorContract *ServiceMonitorContractCallerSession) GetVersion() (string, error) {
	return _ServiceMonitorContract.Contract.GetVersion(&_ServiceMonitorContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ServiceMonitorContract *ServiceMonitorContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ServiceMonitorContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ServiceMonitorContract *ServiceMonitorContractSession) Owner() (common.Address, error) {
	return _ServiceMonitorContract.Contract.Owner(&_ServiceMonitorContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ServiceMonitorContract *ServiceMonitorContractCallerSession) Owner() (common.Address, error) {
	return _ServiceMonitorContract.Contract.Owner(&_ServiceMonitorContract.CallOpts)
}

// RoundNumber is a free data retrieval call binding the contract method 0x4e2786fb.
//
// Solidity: function roundNumber() view returns(uint256)
func (_ServiceMonitorContract *ServiceMonitorContractCaller) RoundNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ServiceMonitorContract.contract.Call(opts, &out, "roundNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoundNumber is a free data retrieval call binding the contract method 0x4e2786fb.
//
// Solidity: function roundNumber() view returns(uint256)
func (_ServiceMonitorContract *ServiceMonitorContractSession) RoundNumber() (*big.Int, error) {
	return _ServiceMonitorContract.Contract.RoundNumber(&_ServiceMonitorContract.CallOpts)
}

// RoundNumber is a free data retrieval call binding the contract method 0x4e2786fb.
//
// Solidity: function roundNumber() view returns(uint256)
func (_ServiceMonitorContract *ServiceMonitorContractCallerSession) RoundNumber() (*big.Int, error) {
	return _ServiceMonitorContract.Contract.RoundNumber(&_ServiceMonitorContract.CallOpts)
}

// RoundTimestamps is a free data retrieval call binding the contract method 0xbb946729.
//
// Solidity: function roundTimestamps(uint256 ) view returns(uint256)
func (_ServiceMonitorContract *ServiceMonitorContractCaller) RoundTimestamps(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ServiceMonitorContract.contract.Call(opts, &out, "roundTimestamps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoundTimestamps is a free data retrieval call binding the contract method 0xbb946729.
//
// Solidity: function roundTimestamps(uint256 ) view returns(uint256)
func (_ServiceMonitorContract *ServiceMonitorContractSession) RoundTimestamps(arg0 *big.Int) (*big.Int, error) {
	return _ServiceMonitorContract.Contract.RoundTimestamps(&_ServiceMonitorContract.CallOpts, arg0)
}

// RoundTimestamps is a free data retrieval call binding the contract method 0xbb946729.
//
// Solidity: function roundTimestamps(uint256 ) view returns(uint256)
func (_ServiceMonitorContract *ServiceMonitorContractCallerSession) RoundTimestamps(arg0 *big.Int) (*big.Int, error) {
	return _ServiceMonitorContract.Contract.RoundTimestamps(&_ServiceMonitorContract.CallOpts, arg0)
}

// SlaResults is a free data retrieval call binding the contract method 0x9ba6a02e.
//
// Solidity: function slaResults(address , uint256 ) view returns(bool)
func (_ServiceMonitorContract *ServiceMonitorContractCaller) SlaResults(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _ServiceMonitorContract.contract.Call(opts, &out, "slaResults", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SlaResults is a free data retrieval call binding the contract method 0x9ba6a02e.
//
// Solidity: function slaResults(address , uint256 ) view returns(bool)
func (_ServiceMonitorContract *ServiceMonitorContractSession) SlaResults(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _ServiceMonitorContract.Contract.SlaResults(&_ServiceMonitorContract.CallOpts, arg0, arg1)
}

// SlaResults is a free data retrieval call binding the contract method 0x9ba6a02e.
//
// Solidity: function slaResults(address , uint256 ) view returns(bool)
func (_ServiceMonitorContract *ServiceMonitorContractCallerSession) SlaResults(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _ServiceMonitorContract.Contract.SlaResults(&_ServiceMonitorContract.CallOpts, arg0, arg1)
}

// VaultManager is a free data retrieval call binding the contract method 0x8a4adf24.
//
// Solidity: function vaultManager() view returns(address)
func (_ServiceMonitorContract *ServiceMonitorContractCaller) VaultManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ServiceMonitorContract.contract.Call(opts, &out, "vaultManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultManager is a free data retrieval call binding the contract method 0x8a4adf24.
//
// Solidity: function vaultManager() view returns(address)
func (_ServiceMonitorContract *ServiceMonitorContractSession) VaultManager() (common.Address, error) {
	return _ServiceMonitorContract.Contract.VaultManager(&_ServiceMonitorContract.CallOpts)
}

// VaultManager is a free data retrieval call binding the contract method 0x8a4adf24.
//
// Solidity: function vaultManager() view returns(address)
func (_ServiceMonitorContract *ServiceMonitorContractCallerSession) VaultManager() (common.Address, error) {
	return _ServiceMonitorContract.Contract.VaultManager(&_ServiceMonitorContract.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ServiceMonitorContract *ServiceMonitorContractTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ServiceMonitorContract.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ServiceMonitorContract *ServiceMonitorContractSession) RenounceOwnership() (*types.Transaction, error) {
	return _ServiceMonitorContract.Contract.RenounceOwnership(&_ServiceMonitorContract.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ServiceMonitorContract *ServiceMonitorContractTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ServiceMonitorContract.Contract.RenounceOwnership(&_ServiceMonitorContract.TransactOpts)
}

// SetConsensusEngine is a paid mutator transaction binding the contract method 0x21762e8e.
//
// Solidity: function setConsensusEngine(address _consensusEngineAddress) returns()
func (_ServiceMonitorContract *ServiceMonitorContractTransactor) SetConsensusEngine(opts *bind.TransactOpts, _consensusEngineAddress common.Address) (*types.Transaction, error) {
	return _ServiceMonitorContract.contract.Transact(opts, "setConsensusEngine", _consensusEngineAddress)
}

// SetConsensusEngine is a paid mutator transaction binding the contract method 0x21762e8e.
//
// Solidity: function setConsensusEngine(address _consensusEngineAddress) returns()
func (_ServiceMonitorContract *ServiceMonitorContractSession) SetConsensusEngine(_consensusEngineAddress common.Address) (*types.Transaction, error) {
	return _ServiceMonitorContract.Contract.SetConsensusEngine(&_ServiceMonitorContract.TransactOpts, _consensusEngineAddress)
}

// SetConsensusEngine is a paid mutator transaction binding the contract method 0x21762e8e.
//
// Solidity: function setConsensusEngine(address _consensusEngineAddress) returns()
func (_ServiceMonitorContract *ServiceMonitorContractTransactorSession) SetConsensusEngine(_consensusEngineAddress common.Address) (*types.Transaction, error) {
	return _ServiceMonitorContract.Contract.SetConsensusEngine(&_ServiceMonitorContract.TransactOpts, _consensusEngineAddress)
}

// SubmitSLACheck is a paid mutator transaction binding the contract method 0x1982a10f.
//
// Solidity: function submitSLACheck(bytes data) returns()
func (_ServiceMonitorContract *ServiceMonitorContractTransactor) SubmitSLACheck(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _ServiceMonitorContract.contract.Transact(opts, "submitSLACheck", data)
}

// SubmitSLACheck is a paid mutator transaction binding the contract method 0x1982a10f.
//
// Solidity: function submitSLACheck(bytes data) returns()
func (_ServiceMonitorContract *ServiceMonitorContractSession) SubmitSLACheck(data []byte) (*types.Transaction, error) {
	return _ServiceMonitorContract.Contract.SubmitSLACheck(&_ServiceMonitorContract.TransactOpts, data)
}

// SubmitSLACheck is a paid mutator transaction binding the contract method 0x1982a10f.
//
// Solidity: function submitSLACheck(bytes data) returns()
func (_ServiceMonitorContract *ServiceMonitorContractTransactorSession) SubmitSLACheck(data []byte) (*types.Transaction, error) {
	return _ServiceMonitorContract.Contract.SubmitSLACheck(&_ServiceMonitorContract.TransactOpts, data)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ServiceMonitorContract *ServiceMonitorContractTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ServiceMonitorContract.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ServiceMonitorContract *ServiceMonitorContractSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ServiceMonitorContract.Contract.TransferOwnership(&_ServiceMonitorContract.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ServiceMonitorContract *ServiceMonitorContractTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ServiceMonitorContract.Contract.TransferOwnership(&_ServiceMonitorContract.TransactOpts, newOwner)
}

// ServiceMonitorContractOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ServiceMonitorContract contract.
type ServiceMonitorContractOwnershipTransferredIterator struct {
	Event *ServiceMonitorContractOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ServiceMonitorContractOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServiceMonitorContractOwnershipTransferred)
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
		it.Event = new(ServiceMonitorContractOwnershipTransferred)
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
func (it *ServiceMonitorContractOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServiceMonitorContractOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServiceMonitorContractOwnershipTransferred represents a OwnershipTransferred event raised by the ServiceMonitorContract contract.
type ServiceMonitorContractOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ServiceMonitorContract *ServiceMonitorContractFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ServiceMonitorContractOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ServiceMonitorContract.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ServiceMonitorContractOwnershipTransferredIterator{contract: _ServiceMonitorContract.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ServiceMonitorContract *ServiceMonitorContractFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ServiceMonitorContractOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ServiceMonitorContract.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServiceMonitorContractOwnershipTransferred)
				if err := _ServiceMonitorContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ServiceMonitorContract *ServiceMonitorContractFilterer) ParseOwnershipTransferred(log types.Log) (*ServiceMonitorContractOwnershipTransferred, error) {
	event := new(ServiceMonitorContractOwnershipTransferred)
	if err := _ServiceMonitorContract.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ServiceMonitorContractSLAFailureIterator is returned from FilterSLAFailure and is used to iterate over the raw logs and unpacked data for SLAFailure events raised by the ServiceMonitorContract contract.
type ServiceMonitorContractSLAFailureIterator struct {
	Event *ServiceMonitorContractSLAFailure // Event containing the contract specifics and raw log

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
func (it *ServiceMonitorContractSLAFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ServiceMonitorContractSLAFailure)
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
		it.Event = new(ServiceMonitorContractSLAFailure)
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
func (it *ServiceMonitorContractSLAFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ServiceMonitorContractSLAFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ServiceMonitorContractSLAFailure represents a SLAFailure event raised by the ServiceMonitorContract contract.
type ServiceMonitorContractSLAFailure struct {
	VaultAddress common.Address
	DealID       string
	Round        *big.Int
	FailTime     *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSLAFailure is a free log retrieval operation binding the contract event 0xa45a414c42a1733f667335788c7c6dc45a98494fac69ea28b7eba7a3faf13278.
//
// Solidity: event SLAFailure(address vaultAddress, string dealID, uint256 round, uint256 failTime)
func (_ServiceMonitorContract *ServiceMonitorContractFilterer) FilterSLAFailure(opts *bind.FilterOpts) (*ServiceMonitorContractSLAFailureIterator, error) {

	logs, sub, err := _ServiceMonitorContract.contract.FilterLogs(opts, "SLAFailure")
	if err != nil {
		return nil, err
	}
	return &ServiceMonitorContractSLAFailureIterator{contract: _ServiceMonitorContract.contract, event: "SLAFailure", logs: logs, sub: sub}, nil
}

// WatchSLAFailure is a free log subscription operation binding the contract event 0xa45a414c42a1733f667335788c7c6dc45a98494fac69ea28b7eba7a3faf13278.
//
// Solidity: event SLAFailure(address vaultAddress, string dealID, uint256 round, uint256 failTime)
func (_ServiceMonitorContract *ServiceMonitorContractFilterer) WatchSLAFailure(opts *bind.WatchOpts, sink chan<- *ServiceMonitorContractSLAFailure) (event.Subscription, error) {

	logs, sub, err := _ServiceMonitorContract.contract.WatchLogs(opts, "SLAFailure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ServiceMonitorContractSLAFailure)
				if err := _ServiceMonitorContract.contract.UnpackLog(event, "SLAFailure", log); err != nil {
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

// ParseSLAFailure is a log parse operation binding the contract event 0xa45a414c42a1733f667335788c7c6dc45a98494fac69ea28b7eba7a3faf13278.
//
// Solidity: event SLAFailure(address vaultAddress, string dealID, uint256 round, uint256 failTime)
func (_ServiceMonitorContract *ServiceMonitorContractFilterer) ParseSLAFailure(log types.Log) (*ServiceMonitorContractSLAFailure, error) {
	event := new(ServiceMonitorContractSLAFailure)
	if err := _ServiceMonitorContract.contract.UnpackLog(event, "SLAFailure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

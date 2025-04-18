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

// SlasherMetaData contains all meta data concerning the Slasher contract.
var SlasherMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vaultManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_serviceMonitor\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"getSafeCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"serviceMonitor\",\"outputs\":[{\"internalType\":\"contractIServiceMonitor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"checkTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"round\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"dealID\",\"type\":\"string\"},{\"internalType\":\"bytes1\",\"name\":\"checkResult\",\"type\":\"bytes1\"},{\"internalType\":\"bytes\",\"name\":\"slaData\",\"type\":\"bytes\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultManager\",\"outputs\":[{\"internalType\":\"contractIVaultManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608034608a57601f61073538819003918201601f19168301916001600160401b03831184841017608e578084926040948552833981010312608a57604b602060458360a2565b920160a2565b60015f81905580546001600160a01b039384166001600160a01b0319918216179091556002805492909316911617905560405161067f90816100b68239f35b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b0382168203608a5756fe6080806040526004361015610012575f80fd5b5f905f3560e01c908163483e622a14610547575080636009409a1461009e5780638a4adf241461007557639f4f63581461004a575f80fd5b346100725780600319360112610072576002546040516001600160a01b039091168152602090f35b80fd5b50346100725780600319360112610072576001546040516001600160a01b039091168152602090f35b50346102bb5760c03660031901126102bb576100b861056a565b60643567ffffffffffffffff81116102bb57366023820112156102bb576100e99036906024816004013591016105be565b6084356001600160f81b03198116036102bb5760a43567ffffffffffffffff81116102bb57366023820112156102bb5761012d9036906024816004013591016105be565b506002546001600160a01b0316338190036104ee5760025f54146104df5760025f556001546001600160a01b0390811693169081156104a8576040516301ce0f2360e71b8152602060048201525f818061018a60248201886105f4565b0381865afa80156102b0575f90610386575b6004915060206101c460606101b983850151604086015190610618565b930192835190610618565b936040519384809263024e4bcf60e31b82525afa9182156102b0575f92610352575b5051908082101561034457505060015b8015610309576102059161062b565b6040516335392b7160e01b815290602082600481885afa9081156102b0575f916102d3575b610234925061062b565b926064840293808504606414901517156102bf57803b156102bb575f92836102809360405196879586948593634200744360e11b855260048501526060602485015260648401906105f4565b90604483015203925af180156102b05761029d575b506001815580f35b6102a991505f90610580565b5f5f610295565b6040513d5f823e3d90fd5b5f80fd5b634e487b7160e01b5f52601160045260245ffd5b90506020823d602011610301575b816102ee60209383610580565b810103126102bb5761023491519061022a565b3d91506102e1565b60405162461bcd60e51b8152602060048201526013602482015272496e76616c6964206465616c20726f756e647360681b6044820152606490fd5b61034d9161062b565b6101f6565b9091506020813d60201161037e575b8161036e60209383610580565b810103126102bb5751905f6101e6565b3d9150610361565b503d805f833e6103968183610580565b8101906020818303126102bb5780519067ffffffffffffffff82116102bb57019060e0828203126102bb576040519060e0820182811067ffffffffffffffff82111761049457604052825167ffffffffffffffff81116102bb57830181601f820112156102bb57805190610409826105a2565b926104176040519485610580565b828452602083830101116102bb57815f9260208093018386015e83010152815260208281015190820152604080830151908201526060808301519082015260808201516001600160a01b03811681036102bb57608082015260a082015160058110156102bb5760049260c09160a0840152015160c082015261019c565b634e487b7160e01b5f52604160045260245ffd5b60405162461bcd60e51b815260206004820152600f60248201526e15985d5b1d081b9bdd08199bdd5b99608a1b6044820152606490fd5b633ee5aeb560e01b5f5260045ffd5b60405162461bcd60e51b815260206004820152602b60248201527f4f6e6c792073657276696365206d6f6e69746f722063616e2063616c6c20746860448201526a34b990333ab731ba34b7b760a91b6064820152608490fd5b346102bb5760203660031901126102bb5760209061056361056a565b505f198152f35b600435906001600160a01b03821682036102bb57565b90601f8019910116810190811067ffffffffffffffff82111761049457604052565b67ffffffffffffffff811161049457601f01601f191660200190565b9291926105ca826105a2565b916105d86040519384610580565b8294818452818301116102bb578281602093845f960137010152565b805180835260209291819084018484015e5f828201840152601f01601f1916010190565b818102929181159184041417156102bf57565b8115610635570490565b634e487b7160e01b5f52601260045260245ffdfea2646970667358221220b4b573f0be0497fb3d586ddc480a578c4dd19c192d85dcd014f8a20a6c13404864736f6c634300081c0033",
}

// SlasherABI is the input ABI used to generate the binding from.
// Deprecated: Use SlasherMetaData.ABI instead.
var SlasherABI = SlasherMetaData.ABI

// SlasherBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SlasherMetaData.Bin instead.
var SlasherBin = SlasherMetaData.Bin

// DeploySlasher deploys a new Ethereum contract, binding an instance of Slasher to it.
func DeploySlasher(auth *bind.TransactOpts, backend bind.ContractBackend, _vaultManager common.Address, _serviceMonitor common.Address) (common.Address, *types.Transaction, *Slasher, error) {
	parsed, err := SlasherMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SlasherBin), backend, _vaultManager, _serviceMonitor)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Slasher{SlasherCaller: SlasherCaller{contract: contract}, SlasherTransactor: SlasherTransactor{contract: contract}, SlasherFilterer: SlasherFilterer{contract: contract}}, nil
}

// Slasher is an auto generated Go binding around an Ethereum contract.
type Slasher struct {
	SlasherCaller     // Read-only binding to the contract
	SlasherTransactor // Write-only binding to the contract
	SlasherFilterer   // Log filterer for contract events
}

// SlasherCaller is an auto generated read-only Go binding around an Ethereum contract.
type SlasherCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlasherTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SlasherTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlasherFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SlasherFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SlasherSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SlasherSession struct {
	Contract     *Slasher          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SlasherCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SlasherCallerSession struct {
	Contract *SlasherCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// SlasherTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SlasherTransactorSession struct {
	Contract     *SlasherTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// SlasherRaw is an auto generated low-level Go binding around an Ethereum contract.
type SlasherRaw struct {
	Contract *Slasher // Generic contract binding to access the raw methods on
}

// SlasherCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SlasherCallerRaw struct {
	Contract *SlasherCaller // Generic read-only contract binding to access the raw methods on
}

// SlasherTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SlasherTransactorRaw struct {
	Contract *SlasherTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSlasher creates a new instance of Slasher, bound to a specific deployed contract.
func NewSlasher(address common.Address, backend bind.ContractBackend) (*Slasher, error) {
	contract, err := bindSlasher(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Slasher{SlasherCaller: SlasherCaller{contract: contract}, SlasherTransactor: SlasherTransactor{contract: contract}, SlasherFilterer: SlasherFilterer{contract: contract}}, nil
}

// NewSlasherCaller creates a new read-only instance of Slasher, bound to a specific deployed contract.
func NewSlasherCaller(address common.Address, caller bind.ContractCaller) (*SlasherCaller, error) {
	contract, err := bindSlasher(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SlasherCaller{contract: contract}, nil
}

// NewSlasherTransactor creates a new write-only instance of Slasher, bound to a specific deployed contract.
func NewSlasherTransactor(address common.Address, transactor bind.ContractTransactor) (*SlasherTransactor, error) {
	contract, err := bindSlasher(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SlasherTransactor{contract: contract}, nil
}

// NewSlasherFilterer creates a new log filterer instance of Slasher, bound to a specific deployed contract.
func NewSlasherFilterer(address common.Address, filterer bind.ContractFilterer) (*SlasherFilterer, error) {
	contract, err := bindSlasher(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SlasherFilterer{contract: contract}, nil
}

// bindSlasher binds a generic wrapper to an already deployed contract.
func bindSlasher(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SlasherMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Slasher *SlasherRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Slasher.Contract.SlasherCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Slasher *SlasherRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Slasher.Contract.SlasherTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Slasher *SlasherRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Slasher.Contract.SlasherTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Slasher *SlasherCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Slasher.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Slasher *SlasherTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Slasher.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Slasher *SlasherTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Slasher.Contract.contract.Transact(opts, method, params...)
}

// GetSafeCollateral is a free data retrieval call binding the contract method 0x483e622a.
//
// Solidity: function getSafeCollateral(address vault) view returns(uint256)
func (_Slasher *SlasherCaller) GetSafeCollateral(opts *bind.CallOpts, vault common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "getSafeCollateral", vault)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSafeCollateral is a free data retrieval call binding the contract method 0x483e622a.
//
// Solidity: function getSafeCollateral(address vault) view returns(uint256)
func (_Slasher *SlasherSession) GetSafeCollateral(vault common.Address) (*big.Int, error) {
	return _Slasher.Contract.GetSafeCollateral(&_Slasher.CallOpts, vault)
}

// GetSafeCollateral is a free data retrieval call binding the contract method 0x483e622a.
//
// Solidity: function getSafeCollateral(address vault) view returns(uint256)
func (_Slasher *SlasherCallerSession) GetSafeCollateral(vault common.Address) (*big.Int, error) {
	return _Slasher.Contract.GetSafeCollateral(&_Slasher.CallOpts, vault)
}

// ServiceMonitor is a free data retrieval call binding the contract method 0x9f4f6358.
//
// Solidity: function serviceMonitor() view returns(address)
func (_Slasher *SlasherCaller) ServiceMonitor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "serviceMonitor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ServiceMonitor is a free data retrieval call binding the contract method 0x9f4f6358.
//
// Solidity: function serviceMonitor() view returns(address)
func (_Slasher *SlasherSession) ServiceMonitor() (common.Address, error) {
	return _Slasher.Contract.ServiceMonitor(&_Slasher.CallOpts)
}

// ServiceMonitor is a free data retrieval call binding the contract method 0x9f4f6358.
//
// Solidity: function serviceMonitor() view returns(address)
func (_Slasher *SlasherCallerSession) ServiceMonitor() (common.Address, error) {
	return _Slasher.Contract.ServiceMonitor(&_Slasher.CallOpts)
}

// VaultManager is a free data retrieval call binding the contract method 0x8a4adf24.
//
// Solidity: function vaultManager() view returns(address)
func (_Slasher *SlasherCaller) VaultManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Slasher.contract.Call(opts, &out, "vaultManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultManager is a free data retrieval call binding the contract method 0x8a4adf24.
//
// Solidity: function vaultManager() view returns(address)
func (_Slasher *SlasherSession) VaultManager() (common.Address, error) {
	return _Slasher.Contract.VaultManager(&_Slasher.CallOpts)
}

// VaultManager is a free data retrieval call binding the contract method 0x8a4adf24.
//
// Solidity: function vaultManager() view returns(address)
func (_Slasher *SlasherCallerSession) VaultManager() (common.Address, error) {
	return _Slasher.Contract.VaultManager(&_Slasher.CallOpts)
}

// Slash is a paid mutator transaction binding the contract method 0x6009409a.
//
// Solidity: function slash(address _vault, uint256 checkTime, uint256 round, string dealID, bytes1 checkResult, bytes slaData) returns()
func (_Slasher *SlasherTransactor) Slash(opts *bind.TransactOpts, _vault common.Address, checkTime *big.Int, round *big.Int, dealID string, checkResult [1]byte, slaData []byte) (*types.Transaction, error) {
	return _Slasher.contract.Transact(opts, "slash", _vault, checkTime, round, dealID, checkResult, slaData)
}

// Slash is a paid mutator transaction binding the contract method 0x6009409a.
//
// Solidity: function slash(address _vault, uint256 checkTime, uint256 round, string dealID, bytes1 checkResult, bytes slaData) returns()
func (_Slasher *SlasherSession) Slash(_vault common.Address, checkTime *big.Int, round *big.Int, dealID string, checkResult [1]byte, slaData []byte) (*types.Transaction, error) {
	return _Slasher.Contract.Slash(&_Slasher.TransactOpts, _vault, checkTime, round, dealID, checkResult, slaData)
}

// Slash is a paid mutator transaction binding the contract method 0x6009409a.
//
// Solidity: function slash(address _vault, uint256 checkTime, uint256 round, string dealID, bytes1 checkResult, bytes slaData) returns()
func (_Slasher *SlasherTransactorSession) Slash(_vault common.Address, checkTime *big.Int, round *big.Int, dealID string, checkResult [1]byte, slaData []byte) (*types.Transaction, error) {
	return _Slasher.Contract.Slash(&_Slasher.TransactOpts, _vault, checkTime, round, dealID, checkResult, slaData)
}

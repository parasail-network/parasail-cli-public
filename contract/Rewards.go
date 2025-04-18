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

// RewardsMetaData contains all meta data concerning the Rewards contract.
var RewardsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_vaultManager\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyVaultOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardsClaimed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RewardsSlashed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"}],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"vaults\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"rewardAmounts\",\"type\":\"uint256[]\"}],\"name\":\"grantRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"rewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_slasher\",\"type\":\"address\"}],\"name\":\"setSlasher\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"vault\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"slashRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"slasher\",\"outputs\":[{\"internalType\":\"contractISlasher\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultManager\",\"outputs\":[{\"internalType\":\"contractIVaultManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080346100ee57601f610b2538819003918201601f19168301916001600160401b038311848410176100f25780849260409485528339810103126100ee57610052602061004b83610106565b9201610106565b33156100db575f8054336001600160a01b0319821681178355604051949290916001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09080a3600180546001600160a01b039283166001600160a01b03199182161790915560038054939092169216919091179055610a0a908161011b8239f35b631e4fbdf760e01b5f525f60045260245ffd5b5f80fd5b634e487b7160e01b5f52604160045260245ffd5b51906001600160a01b03821682036100ee5756fe6080806040526004361015610012575f80fd5b5f3560e01c9081630700037d1461075f5750806312ce2e92146105075780634c64d21214610393578063715018a61461033c5780638a4adf24146103145780638da5cb5b146102ed578063aabc2496146102aa578063b134427114610282578063ef5cfb8c14610146578063f2fde38b146100c15763f7c618c114610095575f80fd5b346100bd575f3660031901126100bd576001546040516001600160a01b039091168152602090f35b5f80fd5b346100bd5760203660031901126100bd576100da610793565b6100e261089d565b6001600160a01b03168015610133575f80546001600160a01b03198116831782556001600160a01b0316907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09080a3005b631e4fbdf760e01b5f525f60045260245ffd5b346100bd5760203660031901126100bd576001600160a01b03610167610793565b16604051631127a41d60e31b8152602081600481855afa908115610277575f91610248575b506001600160a01b0316330361023957805f52600460205260405f205480156101fe5760207ffc30cddea38e2bf4d6ea7d3f9ed3b6ad7f176419f4963bd81318067a4aee73fe91835f52600482525f60408120556101f5813360018060a01b0360015416610860565b604051908152a2005b60405162461bcd60e51b81526020600482015260136024820152724e6f207265776172647320746f20636c61696d60681b6044820152606490fd5b636a4cad8d60e01b5f5260045ffd5b61026a915060203d602011610270575b61026281836107da565b810190610810565b8261018c565b503d610258565b6040513d5f823e3d90fd5b346100bd575f3660031901126100bd576002546040516001600160a01b039091168152602090f35b346100bd5760203660031901126100bd576102c3610793565b6102cb61089d565b600280546001600160a01b0319166001600160a01b0392909216919091179055005b346100bd575f3660031901126100bd575f546040516001600160a01b039091168152602090f35b346100bd575f3660031901126100bd576003546040516001600160a01b039091168152602090f35b346100bd575f3660031901126100bd5761035461089d565b5f80546001600160a01b0319811682556001600160a01b03167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08280a3005b346100bd5760403660031901126100bd5760043567ffffffffffffffff81116100bd576103c49036906004016107a9565b9060243567ffffffffffffffff81116100bd576103e59036906004016107a9565b92906103ef61089d565b8382036104b75791905f925f925b82841061044c576001546040516323b872dd60e01b60208201523360248201523060448201526064808201889052815261044a916001600160a01b03166104456084836107da565b6108c3565b005b9091929361045b85878561082f565b359061046886868561082f565b356001600160a01b03811691908290036100bd576001926104ad925f52600460205261049960405f20918254610853565b90556104a687898761082f565b3590610853565b94019291906103fd565b60405162461bcd60e51b815260206004820152602260248201527f5661756c747320616e64207265776172647320616d6f756e74206d69736d61746044820152610c6d60f31b6064820152608490fd5b346100bd5760403660031901126100bd57610520610793565b60025460243591906001600160a01b03163381900361070e57156106d7578115610683576001600160a01b03165f81815260046020526040902054909190811161062057815f52600460205260405f2090815481810390811161060c57909155600154600354604051631d8cf42560e11b8152926001600160a01b0392831692909160209185916004918391165afa91821561027757816020936101f5927f92473960d0bf1856e93dde4e6a776ae2423a482ba52e746c605be8290909b4d3965f926105ed575b50610860565b610605919250863d88116102705761026281836107da565b90886105e7565b634e487b7160e01b5f52601160045260245ffd5b60405162461bcd60e51b815260206004820152603560248201527f416d6f756e7420746f20736c617368206d757374206265206c657373207468616044820152746e206f7220657175616c20746f207265776172647360581b6064820152608490fd5b60405162461bcd60e51b815260206004820152602660248201527f416d6f756e7420746f20736c617368206d75737420626520677265617465722060448201526507468616e20360d41b6064820152608490fd5b60405162461bcd60e51b815260206004820152600f60248201526e14db185cda195c881b9bdd081cd95d608a1b6044820152606490fd5b60405162461bcd60e51b815260206004820152602360248201527f4f6e6c7920736c61736865722063616e2063616c6c20746869732066756e637460448201526234b7b760e91b6064820152608490fd5b346100bd5760203660031901126100bd576020906001600160a01b03610783610793565b165f526004825260405f20548152f35b600435906001600160a01b03821682036100bd57565b9181601f840112156100bd5782359167ffffffffffffffff83116100bd576020808501948460051b0101116100bd57565b90601f8019910116810190811067ffffffffffffffff8211176107fc57604052565b634e487b7160e01b5f52604160045260245ffd5b908160209103126100bd57516001600160a01b03811681036100bd5790565b919081101561083f5760051b0190565b634e487b7160e01b5f52603260045260245ffd5b9190820180921161060c57565b60405163a9059cbb60e01b60208201526001600160a01b03909216602483015260448083019390935291815261089b916104456064836107da565b565b5f546001600160a01b031633036108b057565b63118cdaa760e01b5f523360045260245ffd5b81516001600160a01b03909116915f91829160200182855af13d1561096a573d67ffffffffffffffff81116107fc5761091e916040519161090e6020601f19601f84011601846107da565b82523d5f602084013e5b83610976565b8051908115159182610946575b50506109345750565b635274afe760e01b5f5260045260245ffd5b81925090602091810103126100bd57602001518015908115036100bd575f8061092b565b61091e90606090610918565b9061099a575080511561098b57805190602001fd5b630a12f52160e11b5f5260045ffd5b815115806109cb575b6109ab575090565b639996b31560e01b5f9081526001600160a01b0391909116600452602490fd5b50803b156109a356fea264697066735822122035121746e50620857af7102a675b737a1e0a087ccb4569d29176e50364bdab9564736f6c634300081c0033",
}

// RewardsABI is the input ABI used to generate the binding from.
// Deprecated: Use RewardsMetaData.ABI instead.
var RewardsABI = RewardsMetaData.ABI

// RewardsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RewardsMetaData.Bin instead.
var RewardsBin = RewardsMetaData.Bin

// DeployRewards deploys a new Ethereum contract, binding an instance of Rewards to it.
func DeployRewards(auth *bind.TransactOpts, backend bind.ContractBackend, _rewardToken common.Address, _vaultManager common.Address) (common.Address, *types.Transaction, *Rewards, error) {
	parsed, err := RewardsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RewardsBin), backend, _rewardToken, _vaultManager)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Rewards{RewardsCaller: RewardsCaller{contract: contract}, RewardsTransactor: RewardsTransactor{contract: contract}, RewardsFilterer: RewardsFilterer{contract: contract}}, nil
}

// Rewards is an auto generated Go binding around an Ethereum contract.
type Rewards struct {
	RewardsCaller     // Read-only binding to the contract
	RewardsTransactor // Write-only binding to the contract
	RewardsFilterer   // Log filterer for contract events
}

// RewardsCaller is an auto generated read-only Go binding around an Ethereum contract.
type RewardsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RewardsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RewardsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RewardsSession struct {
	Contract     *Rewards          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RewardsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RewardsCallerSession struct {
	Contract *RewardsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// RewardsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RewardsTransactorSession struct {
	Contract     *RewardsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// RewardsRaw is an auto generated low-level Go binding around an Ethereum contract.
type RewardsRaw struct {
	Contract *Rewards // Generic contract binding to access the raw methods on
}

// RewardsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RewardsCallerRaw struct {
	Contract *RewardsCaller // Generic read-only contract binding to access the raw methods on
}

// RewardsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RewardsTransactorRaw struct {
	Contract *RewardsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRewards creates a new instance of Rewards, bound to a specific deployed contract.
func NewRewards(address common.Address, backend bind.ContractBackend) (*Rewards, error) {
	contract, err := bindRewards(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rewards{RewardsCaller: RewardsCaller{contract: contract}, RewardsTransactor: RewardsTransactor{contract: contract}, RewardsFilterer: RewardsFilterer{contract: contract}}, nil
}

// NewRewardsCaller creates a new read-only instance of Rewards, bound to a specific deployed contract.
func NewRewardsCaller(address common.Address, caller bind.ContractCaller) (*RewardsCaller, error) {
	contract, err := bindRewards(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RewardsCaller{contract: contract}, nil
}

// NewRewardsTransactor creates a new write-only instance of Rewards, bound to a specific deployed contract.
func NewRewardsTransactor(address common.Address, transactor bind.ContractTransactor) (*RewardsTransactor, error) {
	contract, err := bindRewards(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RewardsTransactor{contract: contract}, nil
}

// NewRewardsFilterer creates a new log filterer instance of Rewards, bound to a specific deployed contract.
func NewRewardsFilterer(address common.Address, filterer bind.ContractFilterer) (*RewardsFilterer, error) {
	contract, err := bindRewards(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RewardsFilterer{contract: contract}, nil
}

// bindRewards binds a generic wrapper to an already deployed contract.
func bindRewards(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RewardsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rewards *RewardsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rewards.Contract.RewardsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rewards *RewardsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewards.Contract.RewardsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rewards *RewardsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rewards.Contract.RewardsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rewards *RewardsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rewards.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rewards *RewardsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewards.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rewards *RewardsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rewards.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rewards *RewardsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rewards *RewardsSession) Owner() (common.Address, error) {
	return _Rewards.Contract.Owner(&_Rewards.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Rewards *RewardsCallerSession) Owner() (common.Address, error) {
	return _Rewards.Contract.Owner(&_Rewards.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_Rewards *RewardsCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_Rewards *RewardsSession) RewardToken() (common.Address, error) {
	return _Rewards.Contract.RewardToken(&_Rewards.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_Rewards *RewardsCallerSession) RewardToken() (common.Address, error) {
	return _Rewards.Contract.RewardToken(&_Rewards.CallOpts)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_Rewards *RewardsCaller) Rewards(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "rewards", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_Rewards *RewardsSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _Rewards.Contract.Rewards(&_Rewards.CallOpts, arg0)
}

// Rewards is a free data retrieval call binding the contract method 0x0700037d.
//
// Solidity: function rewards(address ) view returns(uint256)
func (_Rewards *RewardsCallerSession) Rewards(arg0 common.Address) (*big.Int, error) {
	return _Rewards.Contract.Rewards(&_Rewards.CallOpts, arg0)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_Rewards *RewardsCaller) Slasher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "slasher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_Rewards *RewardsSession) Slasher() (common.Address, error) {
	return _Rewards.Contract.Slasher(&_Rewards.CallOpts)
}

// Slasher is a free data retrieval call binding the contract method 0xb1344271.
//
// Solidity: function slasher() view returns(address)
func (_Rewards *RewardsCallerSession) Slasher() (common.Address, error) {
	return _Rewards.Contract.Slasher(&_Rewards.CallOpts)
}

// VaultManager is a free data retrieval call binding the contract method 0x8a4adf24.
//
// Solidity: function vaultManager() view returns(address)
func (_Rewards *RewardsCaller) VaultManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Rewards.contract.Call(opts, &out, "vaultManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultManager is a free data retrieval call binding the contract method 0x8a4adf24.
//
// Solidity: function vaultManager() view returns(address)
func (_Rewards *RewardsSession) VaultManager() (common.Address, error) {
	return _Rewards.Contract.VaultManager(&_Rewards.CallOpts)
}

// VaultManager is a free data retrieval call binding the contract method 0x8a4adf24.
//
// Solidity: function vaultManager() view returns(address)
func (_Rewards *RewardsCallerSession) VaultManager() (common.Address, error) {
	return _Rewards.Contract.VaultManager(&_Rewards.CallOpts)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address vault) returns()
func (_Rewards *RewardsTransactor) ClaimRewards(opts *bind.TransactOpts, vault common.Address) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "claimRewards", vault)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address vault) returns()
func (_Rewards *RewardsSession) ClaimRewards(vault common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.ClaimRewards(&_Rewards.TransactOpts, vault)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0xef5cfb8c.
//
// Solidity: function claimRewards(address vault) returns()
func (_Rewards *RewardsTransactorSession) ClaimRewards(vault common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.ClaimRewards(&_Rewards.TransactOpts, vault)
}

// GrantRewards is a paid mutator transaction binding the contract method 0x4c64d212.
//
// Solidity: function grantRewards(address[] vaults, uint256[] rewardAmounts) returns()
func (_Rewards *RewardsTransactor) GrantRewards(opts *bind.TransactOpts, vaults []common.Address, rewardAmounts []*big.Int) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "grantRewards", vaults, rewardAmounts)
}

// GrantRewards is a paid mutator transaction binding the contract method 0x4c64d212.
//
// Solidity: function grantRewards(address[] vaults, uint256[] rewardAmounts) returns()
func (_Rewards *RewardsSession) GrantRewards(vaults []common.Address, rewardAmounts []*big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.GrantRewards(&_Rewards.TransactOpts, vaults, rewardAmounts)
}

// GrantRewards is a paid mutator transaction binding the contract method 0x4c64d212.
//
// Solidity: function grantRewards(address[] vaults, uint256[] rewardAmounts) returns()
func (_Rewards *RewardsTransactorSession) GrantRewards(vaults []common.Address, rewardAmounts []*big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.GrantRewards(&_Rewards.TransactOpts, vaults, rewardAmounts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Rewards *RewardsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Rewards *RewardsSession) RenounceOwnership() (*types.Transaction, error) {
	return _Rewards.Contract.RenounceOwnership(&_Rewards.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Rewards *RewardsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Rewards.Contract.RenounceOwnership(&_Rewards.TransactOpts)
}

// SetSlasher is a paid mutator transaction binding the contract method 0xaabc2496.
//
// Solidity: function setSlasher(address _slasher) returns()
func (_Rewards *RewardsTransactor) SetSlasher(opts *bind.TransactOpts, _slasher common.Address) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "setSlasher", _slasher)
}

// SetSlasher is a paid mutator transaction binding the contract method 0xaabc2496.
//
// Solidity: function setSlasher(address _slasher) returns()
func (_Rewards *RewardsSession) SetSlasher(_slasher common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.SetSlasher(&_Rewards.TransactOpts, _slasher)
}

// SetSlasher is a paid mutator transaction binding the contract method 0xaabc2496.
//
// Solidity: function setSlasher(address _slasher) returns()
func (_Rewards *RewardsTransactorSession) SetSlasher(_slasher common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.SetSlasher(&_Rewards.TransactOpts, _slasher)
}

// SlashRewards is a paid mutator transaction binding the contract method 0x12ce2e92.
//
// Solidity: function slashRewards(address vault, uint256 amount) returns()
func (_Rewards *RewardsTransactor) SlashRewards(opts *bind.TransactOpts, vault common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "slashRewards", vault, amount)
}

// SlashRewards is a paid mutator transaction binding the contract method 0x12ce2e92.
//
// Solidity: function slashRewards(address vault, uint256 amount) returns()
func (_Rewards *RewardsSession) SlashRewards(vault common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.SlashRewards(&_Rewards.TransactOpts, vault, amount)
}

// SlashRewards is a paid mutator transaction binding the contract method 0x12ce2e92.
//
// Solidity: function slashRewards(address vault, uint256 amount) returns()
func (_Rewards *RewardsTransactorSession) SlashRewards(vault common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Rewards.Contract.SlashRewards(&_Rewards.TransactOpts, vault, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rewards *RewardsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Rewards.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rewards *RewardsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.TransferOwnership(&_Rewards.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Rewards *RewardsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Rewards.Contract.TransferOwnership(&_Rewards.TransactOpts, newOwner)
}

// RewardsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Rewards contract.
type RewardsOwnershipTransferredIterator struct {
	Event *RewardsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RewardsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsOwnershipTransferred)
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
		it.Event = new(RewardsOwnershipTransferred)
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
func (it *RewardsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsOwnershipTransferred represents a OwnershipTransferred event raised by the Rewards contract.
type RewardsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Rewards *RewardsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RewardsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Rewards.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RewardsOwnershipTransferredIterator{contract: _Rewards.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Rewards *RewardsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RewardsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Rewards.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsOwnershipTransferred)
				if err := _Rewards.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Rewards *RewardsFilterer) ParseOwnershipTransferred(log types.Log) (*RewardsOwnershipTransferred, error) {
	event := new(RewardsOwnershipTransferred)
	if err := _Rewards.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardsRewardsClaimedIterator is returned from FilterRewardsClaimed and is used to iterate over the raw logs and unpacked data for RewardsClaimed events raised by the Rewards contract.
type RewardsRewardsClaimedIterator struct {
	Event *RewardsRewardsClaimed // Event containing the contract specifics and raw log

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
func (it *RewardsRewardsClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsRewardsClaimed)
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
		it.Event = new(RewardsRewardsClaimed)
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
func (it *RewardsRewardsClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsRewardsClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsRewardsClaimed represents a RewardsClaimed event raised by the Rewards contract.
type RewardsRewardsClaimed struct {
	Vault  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardsClaimed is a free log retrieval operation binding the contract event 0xfc30cddea38e2bf4d6ea7d3f9ed3b6ad7f176419f4963bd81318067a4aee73fe.
//
// Solidity: event RewardsClaimed(address indexed vault, uint256 amount)
func (_Rewards *RewardsFilterer) FilterRewardsClaimed(opts *bind.FilterOpts, vault []common.Address) (*RewardsRewardsClaimedIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _Rewards.contract.FilterLogs(opts, "RewardsClaimed", vaultRule)
	if err != nil {
		return nil, err
	}
	return &RewardsRewardsClaimedIterator{contract: _Rewards.contract, event: "RewardsClaimed", logs: logs, sub: sub}, nil
}

// WatchRewardsClaimed is a free log subscription operation binding the contract event 0xfc30cddea38e2bf4d6ea7d3f9ed3b6ad7f176419f4963bd81318067a4aee73fe.
//
// Solidity: event RewardsClaimed(address indexed vault, uint256 amount)
func (_Rewards *RewardsFilterer) WatchRewardsClaimed(opts *bind.WatchOpts, sink chan<- *RewardsRewardsClaimed, vault []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _Rewards.contract.WatchLogs(opts, "RewardsClaimed", vaultRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsRewardsClaimed)
				if err := _Rewards.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
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

// ParseRewardsClaimed is a log parse operation binding the contract event 0xfc30cddea38e2bf4d6ea7d3f9ed3b6ad7f176419f4963bd81318067a4aee73fe.
//
// Solidity: event RewardsClaimed(address indexed vault, uint256 amount)
func (_Rewards *RewardsFilterer) ParseRewardsClaimed(log types.Log) (*RewardsRewardsClaimed, error) {
	event := new(RewardsRewardsClaimed)
	if err := _Rewards.contract.UnpackLog(event, "RewardsClaimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardsRewardsSlashedIterator is returned from FilterRewardsSlashed and is used to iterate over the raw logs and unpacked data for RewardsSlashed events raised by the Rewards contract.
type RewardsRewardsSlashedIterator struct {
	Event *RewardsRewardsSlashed // Event containing the contract specifics and raw log

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
func (it *RewardsRewardsSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardsRewardsSlashed)
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
		it.Event = new(RewardsRewardsSlashed)
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
func (it *RewardsRewardsSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardsRewardsSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardsRewardsSlashed represents a RewardsSlashed event raised by the Rewards contract.
type RewardsRewardsSlashed struct {
	Vault  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRewardsSlashed is a free log retrieval operation binding the contract event 0x92473960d0bf1856e93dde4e6a776ae2423a482ba52e746c605be8290909b4d3.
//
// Solidity: event RewardsSlashed(address indexed vault, uint256 amount)
func (_Rewards *RewardsFilterer) FilterRewardsSlashed(opts *bind.FilterOpts, vault []common.Address) (*RewardsRewardsSlashedIterator, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _Rewards.contract.FilterLogs(opts, "RewardsSlashed", vaultRule)
	if err != nil {
		return nil, err
	}
	return &RewardsRewardsSlashedIterator{contract: _Rewards.contract, event: "RewardsSlashed", logs: logs, sub: sub}, nil
}

// WatchRewardsSlashed is a free log subscription operation binding the contract event 0x92473960d0bf1856e93dde4e6a776ae2423a482ba52e746c605be8290909b4d3.
//
// Solidity: event RewardsSlashed(address indexed vault, uint256 amount)
func (_Rewards *RewardsFilterer) WatchRewardsSlashed(opts *bind.WatchOpts, sink chan<- *RewardsRewardsSlashed, vault []common.Address) (event.Subscription, error) {

	var vaultRule []interface{}
	for _, vaultItem := range vault {
		vaultRule = append(vaultRule, vaultItem)
	}

	logs, sub, err := _Rewards.contract.WatchLogs(opts, "RewardsSlashed", vaultRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardsRewardsSlashed)
				if err := _Rewards.contract.UnpackLog(event, "RewardsSlashed", log); err != nil {
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

// ParseRewardsSlashed is a log parse operation binding the contract event 0x92473960d0bf1856e93dde4e6a776ae2423a482ba52e746c605be8290909b4d3.
//
// Solidity: event RewardsSlashed(address indexed vault, uint256 amount)
func (_Rewards *RewardsFilterer) ParseRewardsSlashed(log types.Log) (*RewardsRewardsSlashed, error) {
	event := new(RewardsRewardsSlashed)
	if err := _Rewards.contract.UnpackLog(event, "RewardsSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

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
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// OwnableBin is the compiled bytecode used for deploying new contracts.
const OwnableBin = `0x608060405234801561001057600080fd5b50600080546001600160a01b031916331790556101ce806100326000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80638da5cb5b1461003b578063f2fde38b1461005f575b600080fd5b610043610087565b604080516001600160a01b039092168252519081900360200190f35b6100856004803603602081101561007557600080fd5b50356001600160a01b0316610096565b005b6000546001600160a01b031690565b6000546001600160a01b031633146100e75760408051600160e51b62461bcd0281526020600482015260096024820152600160b91b682337b93134b23232b702604482015290519081900360640190fd5b6001600160a01b03811615156101475760408051600160e51b62461bcd02815260206004820152601a60248201527f4e6f6e2d7a65726f20616464726573732072657175697265642e000000000000604482015290519081900360640190fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b039290921691909117905556fea165627a7a72305820b6455a714664a06ef0ed8a9d8728b821dd430e3ca60ebe2fab26ccbeee59fe660029`

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// MerkleRootHashOracleABI is the input ABI used to generate the binding from.
const MerkleRootHashOracleABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"contents\",\"type\":\"bytes32[]\"}],\"name\":\"registerMerkleRootHash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"getRootHashContents\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"date\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"merkleRootHashReceived\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// MerkleRootHashOracleBin is the compiled bytecode used for deploying new contracts.
const MerkleRootHashOracleBin = `0x6080604052600080546001600160a01b0319163317905561046e806100256000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806362e61ea9146100515780638da5cb5b146100fd578063da75878514610121578063f2fde38b1461018e575b600080fd5b6100fb6004803603604081101561006757600080fd5b8135919081019060408101602082013564010000000081111561008957600080fd5b82018360208201111561009b57600080fd5b803590602001918460208302840111640100000000831117156100bd57600080fd5b9190808060200260200160405190810160405280939291908181526020018383602002808284376000920191909152509295506101b4945050505050565b005b61010561025f565b604080516001600160a01b039092168252519081900360200190f35b61013e6004803603602081101561013757600080fd5b503561026f565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561017a578181015183820152602001610162565b505050509050019250505060405180910390f35b6100fb600480360360208110156101a457600080fd5b50356001600160a01b03166102d1565b6000546001600160a01b031633146102055760408051600160e51b62461bcd0281526020600482015260096024820152600160b91b682337b93134b23232b702604482015290519081900360640190fd5b60008281526001602090815260409091208251610224928401906103dd565b5060408051428152905183917f56ff85e93faef434af4ffd25961bf467cb5d30430d2a325483afd913ef847d5f919081900360200190a25050565b6000546001600160a01b03165b90565b6000818152600160209081526040918290208054835181840281018401909452808452606093928301828280156102c557602002820191906000526020600020905b8154815260200190600101908083116102b1575b50505050509050919050565b6000546001600160a01b031633146103225760408051600160e51b62461bcd0281526020600482015260096024820152600160b91b682337b93134b23232b702604482015290519081900360640190fd5b6001600160a01b03811615156103825760408051600160e51b62461bcd02815260206004820152601a60248201527f4e6f6e2d7a65726f20616464726573732072657175697265642e000000000000604482015290519081900360640190fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b828054828255906000526020600020908101928215610418579160200282015b828111156104185782518255916020019190600101906103fd565b50610424929150610428565b5090565b61026c91905b80821115610424576000815560010161042e56fea165627a7a72305820a905f5904755ed8ba466ce19162522f472b7e8a01b3e0fca6a6640fa147378880029`

// DeployMerkleRootHashOracle deploys a new Ethereum contract, binding an instance of MerkleRootHashOracle to it.
func DeployMerkleRootHashOracle(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MerkleRootHashOracle, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleRootHashOracleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(MerkleRootHashOracleBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MerkleRootHashOracle{MerkleRootHashOracleCaller: MerkleRootHashOracleCaller{contract: contract}, MerkleRootHashOracleTransactor: MerkleRootHashOracleTransactor{contract: contract}, MerkleRootHashOracleFilterer: MerkleRootHashOracleFilterer{contract: contract}}, nil
}

// MerkleRootHashOracle is an auto generated Go binding around an Ethereum contract.
type MerkleRootHashOracle struct {
	MerkleRootHashOracleCaller     // Read-only binding to the contract
	MerkleRootHashOracleTransactor // Write-only binding to the contract
	MerkleRootHashOracleFilterer   // Log filterer for contract events
}

// MerkleRootHashOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type MerkleRootHashOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleRootHashOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MerkleRootHashOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleRootHashOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MerkleRootHashOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MerkleRootHashOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MerkleRootHashOracleSession struct {
	Contract     *MerkleRootHashOracle // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// MerkleRootHashOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MerkleRootHashOracleCallerSession struct {
	Contract *MerkleRootHashOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// MerkleRootHashOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MerkleRootHashOracleTransactorSession struct {
	Contract     *MerkleRootHashOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// MerkleRootHashOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type MerkleRootHashOracleRaw struct {
	Contract *MerkleRootHashOracle // Generic contract binding to access the raw methods on
}

// MerkleRootHashOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MerkleRootHashOracleCallerRaw struct {
	Contract *MerkleRootHashOracleCaller // Generic read-only contract binding to access the raw methods on
}

// MerkleRootHashOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MerkleRootHashOracleTransactorRaw struct {
	Contract *MerkleRootHashOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMerkleRootHashOracle creates a new instance of MerkleRootHashOracle, bound to a specific deployed contract.
func NewMerkleRootHashOracle(address common.Address, backend bind.ContractBackend) (*MerkleRootHashOracle, error) {
	contract, err := bindMerkleRootHashOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MerkleRootHashOracle{MerkleRootHashOracleCaller: MerkleRootHashOracleCaller{contract: contract}, MerkleRootHashOracleTransactor: MerkleRootHashOracleTransactor{contract: contract}, MerkleRootHashOracleFilterer: MerkleRootHashOracleFilterer{contract: contract}}, nil
}

// NewMerkleRootHashOracleCaller creates a new read-only instance of MerkleRootHashOracle, bound to a specific deployed contract.
func NewMerkleRootHashOracleCaller(address common.Address, caller bind.ContractCaller) (*MerkleRootHashOracleCaller, error) {
	contract, err := bindMerkleRootHashOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleRootHashOracleCaller{contract: contract}, nil
}

// NewMerkleRootHashOracleTransactor creates a new write-only instance of MerkleRootHashOracle, bound to a specific deployed contract.
func NewMerkleRootHashOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*MerkleRootHashOracleTransactor, error) {
	contract, err := bindMerkleRootHashOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MerkleRootHashOracleTransactor{contract: contract}, nil
}

// NewMerkleRootHashOracleFilterer creates a new log filterer instance of MerkleRootHashOracle, bound to a specific deployed contract.
func NewMerkleRootHashOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*MerkleRootHashOracleFilterer, error) {
	contract, err := bindMerkleRootHashOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MerkleRootHashOracleFilterer{contract: contract}, nil
}

// bindMerkleRootHashOracle binds a generic wrapper to an already deployed contract.
func bindMerkleRootHashOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MerkleRootHashOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleRootHashOracle *MerkleRootHashOracleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MerkleRootHashOracle.Contract.MerkleRootHashOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleRootHashOracle *MerkleRootHashOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleRootHashOracle.Contract.MerkleRootHashOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleRootHashOracle *MerkleRootHashOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleRootHashOracle.Contract.MerkleRootHashOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MerkleRootHashOracle *MerkleRootHashOracleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _MerkleRootHashOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MerkleRootHashOracle *MerkleRootHashOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MerkleRootHashOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MerkleRootHashOracle *MerkleRootHashOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MerkleRootHashOracle.Contract.contract.Transact(opts, method, params...)
}

// GetRootHashContents is a free data retrieval call binding the contract method 0xda758785.
//
// Solidity: function getRootHashContents(hash bytes32) constant returns(bytes32[])
func (_MerkleRootHashOracle *MerkleRootHashOracleCaller) GetRootHashContents(opts *bind.CallOpts, hash [32]byte) ([][32]byte, error) {
	var (
		ret0 = new([][32]byte)
	)
	out := ret0
	err := _MerkleRootHashOracle.contract.Call(opts, out, "getRootHashContents", hash)
	return *ret0, err
}

// GetRootHashContents is a free data retrieval call binding the contract method 0xda758785.
//
// Solidity: function getRootHashContents(hash bytes32) constant returns(bytes32[])
func (_MerkleRootHashOracle *MerkleRootHashOracleSession) GetRootHashContents(hash [32]byte) ([][32]byte, error) {
	return _MerkleRootHashOracle.Contract.GetRootHashContents(&_MerkleRootHashOracle.CallOpts, hash)
}

// GetRootHashContents is a free data retrieval call binding the contract method 0xda758785.
//
// Solidity: function getRootHashContents(hash bytes32) constant returns(bytes32[])
func (_MerkleRootHashOracle *MerkleRootHashOracleCallerSession) GetRootHashContents(hash [32]byte) ([][32]byte, error) {
	return _MerkleRootHashOracle.Contract.GetRootHashContents(&_MerkleRootHashOracle.CallOpts, hash)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MerkleRootHashOracle *MerkleRootHashOracleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _MerkleRootHashOracle.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MerkleRootHashOracle *MerkleRootHashOracleSession) Owner() (common.Address, error) {
	return _MerkleRootHashOracle.Contract.Owner(&_MerkleRootHashOracle.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_MerkleRootHashOracle *MerkleRootHashOracleCallerSession) Owner() (common.Address, error) {
	return _MerkleRootHashOracle.Contract.Owner(&_MerkleRootHashOracle.CallOpts)
}

// RegisterMerkleRootHash is a paid mutator transaction binding the contract method 0x62e61ea9.
//
// Solidity: function registerMerkleRootHash(hash bytes32, contents bytes32[]) returns()
func (_MerkleRootHashOracle *MerkleRootHashOracleTransactor) RegisterMerkleRootHash(opts *bind.TransactOpts, hash [32]byte, contents [][32]byte) (*types.Transaction, error) {
	return _MerkleRootHashOracle.contract.Transact(opts, "registerMerkleRootHash", hash, contents)
}

// RegisterMerkleRootHash is a paid mutator transaction binding the contract method 0x62e61ea9.
//
// Solidity: function registerMerkleRootHash(hash bytes32, contents bytes32[]) returns()
func (_MerkleRootHashOracle *MerkleRootHashOracleSession) RegisterMerkleRootHash(hash [32]byte, contents [][32]byte) (*types.Transaction, error) {
	return _MerkleRootHashOracle.Contract.RegisterMerkleRootHash(&_MerkleRootHashOracle.TransactOpts, hash, contents)
}

// RegisterMerkleRootHash is a paid mutator transaction binding the contract method 0x62e61ea9.
//
// Solidity: function registerMerkleRootHash(hash bytes32, contents bytes32[]) returns()
func (_MerkleRootHashOracle *MerkleRootHashOracleTransactorSession) RegisterMerkleRootHash(hash [32]byte, contents [][32]byte) (*types.Transaction, error) {
	return _MerkleRootHashOracle.Contract.RegisterMerkleRootHash(&_MerkleRootHashOracle.TransactOpts, hash, contents)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_MerkleRootHashOracle *MerkleRootHashOracleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MerkleRootHashOracle.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_MerkleRootHashOracle *MerkleRootHashOracleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MerkleRootHashOracle.Contract.TransferOwnership(&_MerkleRootHashOracle.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_MerkleRootHashOracle *MerkleRootHashOracleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MerkleRootHashOracle.Contract.TransferOwnership(&_MerkleRootHashOracle.TransactOpts, newOwner)
}

// MerkleRootHashOracleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MerkleRootHashOracle contract.
type MerkleRootHashOracleOwnershipTransferredIterator struct {
	Event *MerkleRootHashOracleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MerkleRootHashOracleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerkleRootHashOracleOwnershipTransferred)
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
		it.Event = new(MerkleRootHashOracleOwnershipTransferred)
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
func (it *MerkleRootHashOracleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerkleRootHashOracleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerkleRootHashOracleOwnershipTransferred represents a OwnershipTransferred event raised by the MerkleRootHashOracle contract.
type MerkleRootHashOracleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_MerkleRootHashOracle *MerkleRootHashOracleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MerkleRootHashOracleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MerkleRootHashOracle.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MerkleRootHashOracleOwnershipTransferredIterator{contract: _MerkleRootHashOracle.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_MerkleRootHashOracle *MerkleRootHashOracleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MerkleRootHashOracleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MerkleRootHashOracle.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerkleRootHashOracleOwnershipTransferred)
				if err := _MerkleRootHashOracle.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// MerkleRootHashOracleMerkleRootHashReceivedIterator is returned from FilterMerkleRootHashReceived and is used to iterate over the raw logs and unpacked data for MerkleRootHashReceived events raised by the MerkleRootHashOracle contract.
type MerkleRootHashOracleMerkleRootHashReceivedIterator struct {
	Event *MerkleRootHashOracleMerkleRootHashReceived // Event containing the contract specifics and raw log

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
func (it *MerkleRootHashOracleMerkleRootHashReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MerkleRootHashOracleMerkleRootHashReceived)
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
		it.Event = new(MerkleRootHashOracleMerkleRootHashReceived)
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
func (it *MerkleRootHashOracleMerkleRootHashReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MerkleRootHashOracleMerkleRootHashReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MerkleRootHashOracleMerkleRootHashReceived represents a MerkleRootHashReceived event raised by the MerkleRootHashOracle contract.
type MerkleRootHashOracleMerkleRootHashReceived struct {
	Date *big.Int
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMerkleRootHashReceived is a free log retrieval operation binding the contract event 0x56ff85e93faef434af4ffd25961bf467cb5d30430d2a325483afd913ef847d5f.
//
// Solidity: e merkleRootHashReceived(date uint256, hash indexed bytes32)
func (_MerkleRootHashOracle *MerkleRootHashOracleFilterer) FilterMerkleRootHashReceived(opts *bind.FilterOpts, hash [][32]byte) (*MerkleRootHashOracleMerkleRootHashReceivedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _MerkleRootHashOracle.contract.FilterLogs(opts, "merkleRootHashReceived", hashRule)
	if err != nil {
		return nil, err
	}
	return &MerkleRootHashOracleMerkleRootHashReceivedIterator{contract: _MerkleRootHashOracle.contract, event: "merkleRootHashReceived", logs: logs, sub: sub}, nil
}

// WatchMerkleRootHashReceived is a free log subscription operation binding the contract event 0x56ff85e93faef434af4ffd25961bf467cb5d30430d2a325483afd913ef847d5f.
//
// Solidity: e merkleRootHashReceived(date uint256, hash indexed bytes32)
func (_MerkleRootHashOracle *MerkleRootHashOracleFilterer) WatchMerkleRootHashReceived(opts *bind.WatchOpts, sink chan<- *MerkleRootHashOracleMerkleRootHashReceived, hash [][32]byte) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}

	logs, sub, err := _MerkleRootHashOracle.contract.WatchLogs(opts, "merkleRootHashReceived", hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MerkleRootHashOracleMerkleRootHashReceived)
				if err := _MerkleRootHashOracle.contract.UnpackLog(event, "merkleRootHashReceived", log); err != nil {
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

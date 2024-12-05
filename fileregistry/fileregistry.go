// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fileregistry

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

// FileregistryMetaData contains all meta data concerning the Fileregistry contract.
var FileregistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"get\",\"inputs\":[{\"name\":\"filePath\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"save\",\"inputs\":[{\"name\":\"filePath\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"cid\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"FileSaved\",\"inputs\":[{\"name\":\"filePath\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"cid\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"EmptyCID\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyFilePath\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FileNotFound\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// FileregistryABI is the input ABI used to generate the binding from.
// Deprecated: Use FileregistryMetaData.ABI instead.
var FileregistryABI = FileregistryMetaData.ABI

// Fileregistry is an auto generated Go binding around an Ethereum contract.
type Fileregistry struct {
	FileregistryCaller     // Read-only binding to the contract
	FileregistryTransactor // Write-only binding to the contract
	FileregistryFilterer   // Log filterer for contract events
}

// FileregistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type FileregistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileregistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FileregistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileregistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FileregistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FileregistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FileregistrySession struct {
	Contract     *Fileregistry     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FileregistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FileregistryCallerSession struct {
	Contract *FileregistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// FileregistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FileregistryTransactorSession struct {
	Contract     *FileregistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// FileregistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type FileregistryRaw struct {
	Contract *Fileregistry // Generic contract binding to access the raw methods on
}

// FileregistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FileregistryCallerRaw struct {
	Contract *FileregistryCaller // Generic read-only contract binding to access the raw methods on
}

// FileregistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FileregistryTransactorRaw struct {
	Contract *FileregistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFileregistry creates a new instance of Fileregistry, bound to a specific deployed contract.
func NewFileregistry(address common.Address, backend bind.ContractBackend) (*Fileregistry, error) {
	contract, err := bindFileregistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Fileregistry{FileregistryCaller: FileregistryCaller{contract: contract}, FileregistryTransactor: FileregistryTransactor{contract: contract}, FileregistryFilterer: FileregistryFilterer{contract: contract}}, nil
}

// NewFileregistryCaller creates a new read-only instance of Fileregistry, bound to a specific deployed contract.
func NewFileregistryCaller(address common.Address, caller bind.ContractCaller) (*FileregistryCaller, error) {
	contract, err := bindFileregistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FileregistryCaller{contract: contract}, nil
}

// NewFileregistryTransactor creates a new write-only instance of Fileregistry, bound to a specific deployed contract.
func NewFileregistryTransactor(address common.Address, transactor bind.ContractTransactor) (*FileregistryTransactor, error) {
	contract, err := bindFileregistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FileregistryTransactor{contract: contract}, nil
}

// NewFileregistryFilterer creates a new log filterer instance of Fileregistry, bound to a specific deployed contract.
func NewFileregistryFilterer(address common.Address, filterer bind.ContractFilterer) (*FileregistryFilterer, error) {
	contract, err := bindFileregistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FileregistryFilterer{contract: contract}, nil
}

// bindFileregistry binds a generic wrapper to an already deployed contract.
func bindFileregistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FileregistryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fileregistry *FileregistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Fileregistry.Contract.FileregistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fileregistry *FileregistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fileregistry.Contract.FileregistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fileregistry *FileregistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fileregistry.Contract.FileregistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fileregistry *FileregistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Fileregistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fileregistry *FileregistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fileregistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fileregistry *FileregistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fileregistry.Contract.contract.Transact(opts, method, params...)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string filePath) view returns(string)
func (_Fileregistry *FileregistryCaller) Get(opts *bind.CallOpts, filePath string) (string, error) {
	var out []interface{}
	err := _Fileregistry.contract.Call(opts, &out, "get", filePath)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string filePath) view returns(string)
func (_Fileregistry *FileregistrySession) Get(filePath string) (string, error) {
	return _Fileregistry.Contract.Get(&_Fileregistry.CallOpts, filePath)
}

// Get is a free data retrieval call binding the contract method 0x693ec85e.
//
// Solidity: function get(string filePath) view returns(string)
func (_Fileregistry *FileregistryCallerSession) Get(filePath string) (string, error) {
	return _Fileregistry.Contract.Get(&_Fileregistry.CallOpts, filePath)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Fileregistry *FileregistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fileregistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Fileregistry *FileregistrySession) Owner() (common.Address, error) {
	return _Fileregistry.Contract.Owner(&_Fileregistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Fileregistry *FileregistryCallerSession) Owner() (common.Address, error) {
	return _Fileregistry.Contract.Owner(&_Fileregistry.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Fileregistry *FileregistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fileregistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Fileregistry *FileregistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _Fileregistry.Contract.RenounceOwnership(&_Fileregistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Fileregistry *FileregistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Fileregistry.Contract.RenounceOwnership(&_Fileregistry.TransactOpts)
}

// Save is a paid mutator transaction binding the contract method 0x962939b8.
//
// Solidity: function save(string filePath, string cid) returns()
func (_Fileregistry *FileregistryTransactor) Save(opts *bind.TransactOpts, filePath string, cid string) (*types.Transaction, error) {
	return _Fileregistry.contract.Transact(opts, "save", filePath, cid)
}

// Save is a paid mutator transaction binding the contract method 0x962939b8.
//
// Solidity: function save(string filePath, string cid) returns()
func (_Fileregistry *FileregistrySession) Save(filePath string, cid string) (*types.Transaction, error) {
	return _Fileregistry.Contract.Save(&_Fileregistry.TransactOpts, filePath, cid)
}

// Save is a paid mutator transaction binding the contract method 0x962939b8.
//
// Solidity: function save(string filePath, string cid) returns()
func (_Fileregistry *FileregistryTransactorSession) Save(filePath string, cid string) (*types.Transaction, error) {
	return _Fileregistry.Contract.Save(&_Fileregistry.TransactOpts, filePath, cid)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Fileregistry *FileregistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Fileregistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Fileregistry *FileregistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Fileregistry.Contract.TransferOwnership(&_Fileregistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Fileregistry *FileregistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Fileregistry.Contract.TransferOwnership(&_Fileregistry.TransactOpts, newOwner)
}

// FileregistryFileSavedIterator is returned from FilterFileSaved and is used to iterate over the raw logs and unpacked data for FileSaved events raised by the Fileregistry contract.
type FileregistryFileSavedIterator struct {
	Event *FileregistryFileSaved // Event containing the contract specifics and raw log

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
func (it *FileregistryFileSavedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileregistryFileSaved)
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
		it.Event = new(FileregistryFileSaved)
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
func (it *FileregistryFileSavedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileregistryFileSavedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileregistryFileSaved represents a FileSaved event raised by the Fileregistry contract.
type FileregistryFileSaved struct {
	FilePath string
	Cid      string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFileSaved is a free log retrieval operation binding the contract event 0xcadc7184c55de53d424a8e73df016947523f46250bb8957192d0d084403dfd25.
//
// Solidity: event FileSaved(string filePath, string cid)
func (_Fileregistry *FileregistryFilterer) FilterFileSaved(opts *bind.FilterOpts) (*FileregistryFileSavedIterator, error) {

	logs, sub, err := _Fileregistry.contract.FilterLogs(opts, "FileSaved")
	if err != nil {
		return nil, err
	}
	return &FileregistryFileSavedIterator{contract: _Fileregistry.contract, event: "FileSaved", logs: logs, sub: sub}, nil
}

// WatchFileSaved is a free log subscription operation binding the contract event 0xcadc7184c55de53d424a8e73df016947523f46250bb8957192d0d084403dfd25.
//
// Solidity: event FileSaved(string filePath, string cid)
func (_Fileregistry *FileregistryFilterer) WatchFileSaved(opts *bind.WatchOpts, sink chan<- *FileregistryFileSaved) (event.Subscription, error) {

	logs, sub, err := _Fileregistry.contract.WatchLogs(opts, "FileSaved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileregistryFileSaved)
				if err := _Fileregistry.contract.UnpackLog(event, "FileSaved", log); err != nil {
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

// ParseFileSaved is a log parse operation binding the contract event 0xcadc7184c55de53d424a8e73df016947523f46250bb8957192d0d084403dfd25.
//
// Solidity: event FileSaved(string filePath, string cid)
func (_Fileregistry *FileregistryFilterer) ParseFileSaved(log types.Log) (*FileregistryFileSaved, error) {
	event := new(FileregistryFileSaved)
	if err := _Fileregistry.contract.UnpackLog(event, "FileSaved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FileregistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Fileregistry contract.
type FileregistryOwnershipTransferredIterator struct {
	Event *FileregistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FileregistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FileregistryOwnershipTransferred)
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
		it.Event = new(FileregistryOwnershipTransferred)
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
func (it *FileregistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FileregistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FileregistryOwnershipTransferred represents a OwnershipTransferred event raised by the Fileregistry contract.
type FileregistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Fileregistry *FileregistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FileregistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Fileregistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FileregistryOwnershipTransferredIterator{contract: _Fileregistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Fileregistry *FileregistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FileregistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Fileregistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FileregistryOwnershipTransferred)
				if err := _Fileregistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Fileregistry *FileregistryFilterer) ParseOwnershipTransferred(log types.Log) (*FileregistryOwnershipTransferred, error) {
	event := new(FileregistryOwnershipTransferred)
	if err := _Fileregistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

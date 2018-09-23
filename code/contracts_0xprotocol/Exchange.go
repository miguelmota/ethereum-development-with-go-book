// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package exchange

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

// ExchangeABI is the input ABI used to generate the binding from.
const ExchangeABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"numerator\",\"type\":\"uint256\"},{\"name\":\"denominator\",\"type\":\"uint256\"},{\"name\":\"target\",\"type\":\"uint256\"}],\"name\":\"isRoundingError\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"filled\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"cancelled\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"orderAddresses\",\"type\":\"address[5][]\"},{\"name\":\"orderValues\",\"type\":\"uint256[6][]\"},{\"name\":\"fillTakerTokenAmount\",\"type\":\"uint256\"},{\"name\":\"shouldThrowOnInsufficientBalanceOrAllowance\",\"type\":\"bool\"},{\"name\":\"v\",\"type\":\"uint8[]\"},{\"name\":\"r\",\"type\":\"bytes32[]\"},{\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"fillOrdersUpTo\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"orderAddresses\",\"type\":\"address[5]\"},{\"name\":\"orderValues\",\"type\":\"uint256[6]\"},{\"name\":\"cancelTakerTokenAmount\",\"type\":\"uint256\"}],\"name\":\"cancelOrder\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ZRX_TOKEN_CONTRACT\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"orderAddresses\",\"type\":\"address[5][]\"},{\"name\":\"orderValues\",\"type\":\"uint256[6][]\"},{\"name\":\"fillTakerTokenAmounts\",\"type\":\"uint256[]\"},{\"name\":\"v\",\"type\":\"uint8[]\"},{\"name\":\"r\",\"type\":\"bytes32[]\"},{\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"batchFillOrKillOrders\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"orderAddresses\",\"type\":\"address[5]\"},{\"name\":\"orderValues\",\"type\":\"uint256[6]\"},{\"name\":\"fillTakerTokenAmount\",\"type\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"fillOrKillOrder\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"getUnavailableTakerTokenAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"signer\",\"type\":\"address\"},{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"isValidSignature\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"numerator\",\"type\":\"uint256\"},{\"name\":\"denominator\",\"type\":\"uint256\"},{\"name\":\"target\",\"type\":\"uint256\"}],\"name\":\"getPartialAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"TOKEN_TRANSFER_PROXY_CONTRACT\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"orderAddresses\",\"type\":\"address[5][]\"},{\"name\":\"orderValues\",\"type\":\"uint256[6][]\"},{\"name\":\"fillTakerTokenAmounts\",\"type\":\"uint256[]\"},{\"name\":\"shouldThrowOnInsufficientBalanceOrAllowance\",\"type\":\"bool\"},{\"name\":\"v\",\"type\":\"uint8[]\"},{\"name\":\"r\",\"type\":\"bytes32[]\"},{\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"batchFillOrders\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"orderAddresses\",\"type\":\"address[5][]\"},{\"name\":\"orderValues\",\"type\":\"uint256[6][]\"},{\"name\":\"cancelTakerTokenAmounts\",\"type\":\"uint256[]\"}],\"name\":\"batchCancelOrders\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"orderAddresses\",\"type\":\"address[5]\"},{\"name\":\"orderValues\",\"type\":\"uint256[6]\"},{\"name\":\"fillTakerTokenAmount\",\"type\":\"uint256\"},{\"name\":\"shouldThrowOnInsufficientBalanceOrAllowance\",\"type\":\"bool\"},{\"name\":\"v\",\"type\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"fillOrder\",\"outputs\":[{\"name\":\"filledTakerTokenAmount\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"orderAddresses\",\"type\":\"address[5]\"},{\"name\":\"orderValues\",\"type\":\"uint256[6]\"}],\"name\":\"getOrderHash\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EXTERNAL_QUERY_GAS_LIMIT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"inputs\":[{\"name\":\"_zrxToken\",\"type\":\"address\"},{\"name\":\"_tokenTransferProxy\",\"type\":\"address\"}],\"payable\":false,\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"feeRecipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"makerToken\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"takerToken\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"filledMakerTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"filledTakerTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"paidMakerFee\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"paidTakerFee\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"tokens\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"LogFill\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"feeRecipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"makerToken\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"takerToken\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"cancelledMakerTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"cancelledTakerTokenAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"tokens\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"LogCancel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"errorId\",\"type\":\"uint8\"},{\"indexed\":true,\"name\":\"orderHash\",\"type\":\"bytes32\"}],\"name\":\"LogError\",\"type\":\"event\"}]"

// Exchange is an auto generated Go binding around an Ethereum contract.
type Exchange struct {
	ExchangeCaller     // Read-only binding to the contract
	ExchangeTransactor // Write-only binding to the contract
	ExchangeFilterer   // Log filterer for contract events
}

// ExchangeCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExchangeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExchangeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExchangeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExchangeSession struct {
	Contract     *Exchange         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExchangeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExchangeCallerSession struct {
	Contract *ExchangeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ExchangeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExchangeTransactorSession struct {
	Contract     *ExchangeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ExchangeRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExchangeRaw struct {
	Contract *Exchange // Generic contract binding to access the raw methods on
}

// ExchangeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExchangeCallerRaw struct {
	Contract *ExchangeCaller // Generic read-only contract binding to access the raw methods on
}

// ExchangeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExchangeTransactorRaw struct {
	Contract *ExchangeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExchange creates a new instance of Exchange, bound to a specific deployed contract.
func NewExchange(address common.Address, backend bind.ContractBackend) (*Exchange, error) {
	contract, err := bindExchange(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Exchange{ExchangeCaller: ExchangeCaller{contract: contract}, ExchangeTransactor: ExchangeTransactor{contract: contract}, ExchangeFilterer: ExchangeFilterer{contract: contract}}, nil
}

// NewExchangeCaller creates a new read-only instance of Exchange, bound to a specific deployed contract.
func NewExchangeCaller(address common.Address, caller bind.ContractCaller) (*ExchangeCaller, error) {
	contract, err := bindExchange(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeCaller{contract: contract}, nil
}

// NewExchangeTransactor creates a new write-only instance of Exchange, bound to a specific deployed contract.
func NewExchangeTransactor(address common.Address, transactor bind.ContractTransactor) (*ExchangeTransactor, error) {
	contract, err := bindExchange(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeTransactor{contract: contract}, nil
}

// NewExchangeFilterer creates a new log filterer instance of Exchange, bound to a specific deployed contract.
func NewExchangeFilterer(address common.Address, filterer bind.ContractFilterer) (*ExchangeFilterer, error) {
	contract, err := bindExchange(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExchangeFilterer{contract: contract}, nil
}

// bindExchange binds a generic wrapper to an already deployed contract.
func bindExchange(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExchangeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchange *ExchangeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Exchange.Contract.ExchangeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchange *ExchangeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.Contract.ExchangeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchange *ExchangeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exchange.Contract.ExchangeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exchange *ExchangeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Exchange.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exchange *ExchangeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exchange.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exchange *ExchangeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exchange.Contract.contract.Transact(opts, method, params...)
}

// EXTERNALQUERYGASLIMIT is a free data retrieval call binding the contract method 0xf06bbf75.
//
// Solidity: function EXTERNAL_QUERY_GAS_LIMIT() constant returns(uint16)
func (_Exchange *ExchangeCaller) EXTERNALQUERYGASLIMIT(opts *bind.CallOpts) (uint16, error) {
	var (
		ret0 = new(uint16)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "EXTERNAL_QUERY_GAS_LIMIT")
	return *ret0, err
}

// EXTERNALQUERYGASLIMIT is a free data retrieval call binding the contract method 0xf06bbf75.
//
// Solidity: function EXTERNAL_QUERY_GAS_LIMIT() constant returns(uint16)
func (_Exchange *ExchangeSession) EXTERNALQUERYGASLIMIT() (uint16, error) {
	return _Exchange.Contract.EXTERNALQUERYGASLIMIT(&_Exchange.CallOpts)
}

// EXTERNALQUERYGASLIMIT is a free data retrieval call binding the contract method 0xf06bbf75.
//
// Solidity: function EXTERNAL_QUERY_GAS_LIMIT() constant returns(uint16)
func (_Exchange *ExchangeCallerSession) EXTERNALQUERYGASLIMIT() (uint16, error) {
	return _Exchange.Contract.EXTERNALQUERYGASLIMIT(&_Exchange.CallOpts)
}

// TOKENTRANSFERPROXYCONTRACT is a free data retrieval call binding the contract method 0xadd1cbc5.
//
// Solidity: function TOKEN_TRANSFER_PROXY_CONTRACT() constant returns(address)
func (_Exchange *ExchangeCaller) TOKENTRANSFERPROXYCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "TOKEN_TRANSFER_PROXY_CONTRACT")
	return *ret0, err
}

// TOKENTRANSFERPROXYCONTRACT is a free data retrieval call binding the contract method 0xadd1cbc5.
//
// Solidity: function TOKEN_TRANSFER_PROXY_CONTRACT() constant returns(address)
func (_Exchange *ExchangeSession) TOKENTRANSFERPROXYCONTRACT() (common.Address, error) {
	return _Exchange.Contract.TOKENTRANSFERPROXYCONTRACT(&_Exchange.CallOpts)
}

// TOKENTRANSFERPROXYCONTRACT is a free data retrieval call binding the contract method 0xadd1cbc5.
//
// Solidity: function TOKEN_TRANSFER_PROXY_CONTRACT() constant returns(address)
func (_Exchange *ExchangeCallerSession) TOKENTRANSFERPROXYCONTRACT() (common.Address, error) {
	return _Exchange.Contract.TOKENTRANSFERPROXYCONTRACT(&_Exchange.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_Exchange *ExchangeCaller) VERSION(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "VERSION")
	return *ret0, err
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_Exchange *ExchangeSession) VERSION() (string, error) {
	return _Exchange.Contract.VERSION(&_Exchange.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() constant returns(string)
func (_Exchange *ExchangeCallerSession) VERSION() (string, error) {
	return _Exchange.Contract.VERSION(&_Exchange.CallOpts)
}

// ZRXTOKENCONTRACT is a free data retrieval call binding the contract method 0x3b30ba59.
//
// Solidity: function ZRX_TOKEN_CONTRACT() constant returns(address)
func (_Exchange *ExchangeCaller) ZRXTOKENCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "ZRX_TOKEN_CONTRACT")
	return *ret0, err
}

// ZRXTOKENCONTRACT is a free data retrieval call binding the contract method 0x3b30ba59.
//
// Solidity: function ZRX_TOKEN_CONTRACT() constant returns(address)
func (_Exchange *ExchangeSession) ZRXTOKENCONTRACT() (common.Address, error) {
	return _Exchange.Contract.ZRXTOKENCONTRACT(&_Exchange.CallOpts)
}

// ZRXTOKENCONTRACT is a free data retrieval call binding the contract method 0x3b30ba59.
//
// Solidity: function ZRX_TOKEN_CONTRACT() constant returns(address)
func (_Exchange *ExchangeCallerSession) ZRXTOKENCONTRACT() (common.Address, error) {
	return _Exchange.Contract.ZRXTOKENCONTRACT(&_Exchange.CallOpts)
}

// Cancelled is a free data retrieval call binding the contract method 0x2ac12622.
//
// Solidity: function cancelled( bytes32) constant returns(uint256)
func (_Exchange *ExchangeCaller) Cancelled(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "cancelled", arg0)
	return *ret0, err
}

// Cancelled is a free data retrieval call binding the contract method 0x2ac12622.
//
// Solidity: function cancelled( bytes32) constant returns(uint256)
func (_Exchange *ExchangeSession) Cancelled(arg0 [32]byte) (*big.Int, error) {
	return _Exchange.Contract.Cancelled(&_Exchange.CallOpts, arg0)
}

// Cancelled is a free data retrieval call binding the contract method 0x2ac12622.
//
// Solidity: function cancelled( bytes32) constant returns(uint256)
func (_Exchange *ExchangeCallerSession) Cancelled(arg0 [32]byte) (*big.Int, error) {
	return _Exchange.Contract.Cancelled(&_Exchange.CallOpts, arg0)
}

// Filled is a free data retrieval call binding the contract method 0x288cdc91.
//
// Solidity: function filled( bytes32) constant returns(uint256)
func (_Exchange *ExchangeCaller) Filled(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "filled", arg0)
	return *ret0, err
}

// Filled is a free data retrieval call binding the contract method 0x288cdc91.
//
// Solidity: function filled( bytes32) constant returns(uint256)
func (_Exchange *ExchangeSession) Filled(arg0 [32]byte) (*big.Int, error) {
	return _Exchange.Contract.Filled(&_Exchange.CallOpts, arg0)
}

// Filled is a free data retrieval call binding the contract method 0x288cdc91.
//
// Solidity: function filled( bytes32) constant returns(uint256)
func (_Exchange *ExchangeCallerSession) Filled(arg0 [32]byte) (*big.Int, error) {
	return _Exchange.Contract.Filled(&_Exchange.CallOpts, arg0)
}

// GetOrderHash is a free data retrieval call binding the contract method 0xcfc4d0ec.
//
// Solidity: function getOrderHash(orderAddresses address[5], orderValues uint256[6]) constant returns(bytes32)
func (_Exchange *ExchangeCaller) GetOrderHash(opts *bind.CallOpts, orderAddresses [5]common.Address, orderValues [6]*big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "getOrderHash", orderAddresses, orderValues)
	return *ret0, err
}

// GetOrderHash is a free data retrieval call binding the contract method 0xcfc4d0ec.
//
// Solidity: function getOrderHash(orderAddresses address[5], orderValues uint256[6]) constant returns(bytes32)
func (_Exchange *ExchangeSession) GetOrderHash(orderAddresses [5]common.Address, orderValues [6]*big.Int) ([32]byte, error) {
	return _Exchange.Contract.GetOrderHash(&_Exchange.CallOpts, orderAddresses, orderValues)
}

// GetOrderHash is a free data retrieval call binding the contract method 0xcfc4d0ec.
//
// Solidity: function getOrderHash(orderAddresses address[5], orderValues uint256[6]) constant returns(bytes32)
func (_Exchange *ExchangeCallerSession) GetOrderHash(orderAddresses [5]common.Address, orderValues [6]*big.Int) ([32]byte, error) {
	return _Exchange.Contract.GetOrderHash(&_Exchange.CallOpts, orderAddresses, orderValues)
}

// GetPartialAmount is a free data retrieval call binding the contract method 0x98024a8b.
//
// Solidity: function getPartialAmount(numerator uint256, denominator uint256, target uint256) constant returns(uint256)
func (_Exchange *ExchangeCaller) GetPartialAmount(opts *bind.CallOpts, numerator *big.Int, denominator *big.Int, target *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "getPartialAmount", numerator, denominator, target)
	return *ret0, err
}

// GetPartialAmount is a free data retrieval call binding the contract method 0x98024a8b.
//
// Solidity: function getPartialAmount(numerator uint256, denominator uint256, target uint256) constant returns(uint256)
func (_Exchange *ExchangeSession) GetPartialAmount(numerator *big.Int, denominator *big.Int, target *big.Int) (*big.Int, error) {
	return _Exchange.Contract.GetPartialAmount(&_Exchange.CallOpts, numerator, denominator, target)
}

// GetPartialAmount is a free data retrieval call binding the contract method 0x98024a8b.
//
// Solidity: function getPartialAmount(numerator uint256, denominator uint256, target uint256) constant returns(uint256)
func (_Exchange *ExchangeCallerSession) GetPartialAmount(numerator *big.Int, denominator *big.Int, target *big.Int) (*big.Int, error) {
	return _Exchange.Contract.GetPartialAmount(&_Exchange.CallOpts, numerator, denominator, target)
}

// GetUnavailableTakerTokenAmount is a free data retrieval call binding the contract method 0x7e9abb50.
//
// Solidity: function getUnavailableTakerTokenAmount(orderHash bytes32) constant returns(uint256)
func (_Exchange *ExchangeCaller) GetUnavailableTakerTokenAmount(opts *bind.CallOpts, orderHash [32]byte) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "getUnavailableTakerTokenAmount", orderHash)
	return *ret0, err
}

// GetUnavailableTakerTokenAmount is a free data retrieval call binding the contract method 0x7e9abb50.
//
// Solidity: function getUnavailableTakerTokenAmount(orderHash bytes32) constant returns(uint256)
func (_Exchange *ExchangeSession) GetUnavailableTakerTokenAmount(orderHash [32]byte) (*big.Int, error) {
	return _Exchange.Contract.GetUnavailableTakerTokenAmount(&_Exchange.CallOpts, orderHash)
}

// GetUnavailableTakerTokenAmount is a free data retrieval call binding the contract method 0x7e9abb50.
//
// Solidity: function getUnavailableTakerTokenAmount(orderHash bytes32) constant returns(uint256)
func (_Exchange *ExchangeCallerSession) GetUnavailableTakerTokenAmount(orderHash [32]byte) (*big.Int, error) {
	return _Exchange.Contract.GetUnavailableTakerTokenAmount(&_Exchange.CallOpts, orderHash)
}

// IsRoundingError is a free data retrieval call binding the contract method 0x14df96ee.
//
// Solidity: function isRoundingError(numerator uint256, denominator uint256, target uint256) constant returns(bool)
func (_Exchange *ExchangeCaller) IsRoundingError(opts *bind.CallOpts, numerator *big.Int, denominator *big.Int, target *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "isRoundingError", numerator, denominator, target)
	return *ret0, err
}

// IsRoundingError is a free data retrieval call binding the contract method 0x14df96ee.
//
// Solidity: function isRoundingError(numerator uint256, denominator uint256, target uint256) constant returns(bool)
func (_Exchange *ExchangeSession) IsRoundingError(numerator *big.Int, denominator *big.Int, target *big.Int) (bool, error) {
	return _Exchange.Contract.IsRoundingError(&_Exchange.CallOpts, numerator, denominator, target)
}

// IsRoundingError is a free data retrieval call binding the contract method 0x14df96ee.
//
// Solidity: function isRoundingError(numerator uint256, denominator uint256, target uint256) constant returns(bool)
func (_Exchange *ExchangeCallerSession) IsRoundingError(numerator *big.Int, denominator *big.Int, target *big.Int) (bool, error) {
	return _Exchange.Contract.IsRoundingError(&_Exchange.CallOpts, numerator, denominator, target)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x8163681e.
//
// Solidity: function isValidSignature(signer address, hash bytes32, v uint8, r bytes32, s bytes32) constant returns(bool)
func (_Exchange *ExchangeCaller) IsValidSignature(opts *bind.CallOpts, signer common.Address, hash [32]byte, v uint8, r [32]byte, s [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Exchange.contract.Call(opts, out, "isValidSignature", signer, hash, v, r, s)
	return *ret0, err
}

// IsValidSignature is a free data retrieval call binding the contract method 0x8163681e.
//
// Solidity: function isValidSignature(signer address, hash bytes32, v uint8, r bytes32, s bytes32) constant returns(bool)
func (_Exchange *ExchangeSession) IsValidSignature(signer common.Address, hash [32]byte, v uint8, r [32]byte, s [32]byte) (bool, error) {
	return _Exchange.Contract.IsValidSignature(&_Exchange.CallOpts, signer, hash, v, r, s)
}

// IsValidSignature is a free data retrieval call binding the contract method 0x8163681e.
//
// Solidity: function isValidSignature(signer address, hash bytes32, v uint8, r bytes32, s bytes32) constant returns(bool)
func (_Exchange *ExchangeCallerSession) IsValidSignature(signer common.Address, hash [32]byte, v uint8, r [32]byte, s [32]byte) (bool, error) {
	return _Exchange.Contract.IsValidSignature(&_Exchange.CallOpts, signer, hash, v, r, s)
}

// BatchCancelOrders is a paid mutator transaction binding the contract method 0xbaa0181d.
//
// Solidity: function batchCancelOrders(orderAddresses address[5][], orderValues uint256[6][], cancelTakerTokenAmounts uint256[]) returns()
func (_Exchange *ExchangeTransactor) BatchCancelOrders(opts *bind.TransactOpts, orderAddresses [][5]common.Address, orderValues [][6]*big.Int, cancelTakerTokenAmounts []*big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "batchCancelOrders", orderAddresses, orderValues, cancelTakerTokenAmounts)
}

// BatchCancelOrders is a paid mutator transaction binding the contract method 0xbaa0181d.
//
// Solidity: function batchCancelOrders(orderAddresses address[5][], orderValues uint256[6][], cancelTakerTokenAmounts uint256[]) returns()
func (_Exchange *ExchangeSession) BatchCancelOrders(orderAddresses [][5]common.Address, orderValues [][6]*big.Int, cancelTakerTokenAmounts []*big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.BatchCancelOrders(&_Exchange.TransactOpts, orderAddresses, orderValues, cancelTakerTokenAmounts)
}

// BatchCancelOrders is a paid mutator transaction binding the contract method 0xbaa0181d.
//
// Solidity: function batchCancelOrders(orderAddresses address[5][], orderValues uint256[6][], cancelTakerTokenAmounts uint256[]) returns()
func (_Exchange *ExchangeTransactorSession) BatchCancelOrders(orderAddresses [][5]common.Address, orderValues [][6]*big.Int, cancelTakerTokenAmounts []*big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.BatchCancelOrders(&_Exchange.TransactOpts, orderAddresses, orderValues, cancelTakerTokenAmounts)
}

// BatchFillOrKillOrders is a paid mutator transaction binding the contract method 0x4f150787.
//
// Solidity: function batchFillOrKillOrders(orderAddresses address[5][], orderValues uint256[6][], fillTakerTokenAmounts uint256[], v uint8[], r bytes32[], s bytes32[]) returns()
func (_Exchange *ExchangeTransactor) BatchFillOrKillOrders(opts *bind.TransactOpts, orderAddresses [][5]common.Address, orderValues [][6]*big.Int, fillTakerTokenAmounts []*big.Int, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "batchFillOrKillOrders", orderAddresses, orderValues, fillTakerTokenAmounts, v, r, s)
}

// BatchFillOrKillOrders is a paid mutator transaction binding the contract method 0x4f150787.
//
// Solidity: function batchFillOrKillOrders(orderAddresses address[5][], orderValues uint256[6][], fillTakerTokenAmounts uint256[], v uint8[], r bytes32[], s bytes32[]) returns()
func (_Exchange *ExchangeSession) BatchFillOrKillOrders(orderAddresses [][5]common.Address, orderValues [][6]*big.Int, fillTakerTokenAmounts []*big.Int, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.BatchFillOrKillOrders(&_Exchange.TransactOpts, orderAddresses, orderValues, fillTakerTokenAmounts, v, r, s)
}

// BatchFillOrKillOrders is a paid mutator transaction binding the contract method 0x4f150787.
//
// Solidity: function batchFillOrKillOrders(orderAddresses address[5][], orderValues uint256[6][], fillTakerTokenAmounts uint256[], v uint8[], r bytes32[], s bytes32[]) returns()
func (_Exchange *ExchangeTransactorSession) BatchFillOrKillOrders(orderAddresses [][5]common.Address, orderValues [][6]*big.Int, fillTakerTokenAmounts []*big.Int, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.BatchFillOrKillOrders(&_Exchange.TransactOpts, orderAddresses, orderValues, fillTakerTokenAmounts, v, r, s)
}

// BatchFillOrders is a paid mutator transaction binding the contract method 0xb7b2c7d6.
//
// Solidity: function batchFillOrders(orderAddresses address[5][], orderValues uint256[6][], fillTakerTokenAmounts uint256[], shouldThrowOnInsufficientBalanceOrAllowance bool, v uint8[], r bytes32[], s bytes32[]) returns()
func (_Exchange *ExchangeTransactor) BatchFillOrders(opts *bind.TransactOpts, orderAddresses [][5]common.Address, orderValues [][6]*big.Int, fillTakerTokenAmounts []*big.Int, shouldThrowOnInsufficientBalanceOrAllowance bool, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "batchFillOrders", orderAddresses, orderValues, fillTakerTokenAmounts, shouldThrowOnInsufficientBalanceOrAllowance, v, r, s)
}

// BatchFillOrders is a paid mutator transaction binding the contract method 0xb7b2c7d6.
//
// Solidity: function batchFillOrders(orderAddresses address[5][], orderValues uint256[6][], fillTakerTokenAmounts uint256[], shouldThrowOnInsufficientBalanceOrAllowance bool, v uint8[], r bytes32[], s bytes32[]) returns()
func (_Exchange *ExchangeSession) BatchFillOrders(orderAddresses [][5]common.Address, orderValues [][6]*big.Int, fillTakerTokenAmounts []*big.Int, shouldThrowOnInsufficientBalanceOrAllowance bool, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.BatchFillOrders(&_Exchange.TransactOpts, orderAddresses, orderValues, fillTakerTokenAmounts, shouldThrowOnInsufficientBalanceOrAllowance, v, r, s)
}

// BatchFillOrders is a paid mutator transaction binding the contract method 0xb7b2c7d6.
//
// Solidity: function batchFillOrders(orderAddresses address[5][], orderValues uint256[6][], fillTakerTokenAmounts uint256[], shouldThrowOnInsufficientBalanceOrAllowance bool, v uint8[], r bytes32[], s bytes32[]) returns()
func (_Exchange *ExchangeTransactorSession) BatchFillOrders(orderAddresses [][5]common.Address, orderValues [][6]*big.Int, fillTakerTokenAmounts []*big.Int, shouldThrowOnInsufficientBalanceOrAllowance bool, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.BatchFillOrders(&_Exchange.TransactOpts, orderAddresses, orderValues, fillTakerTokenAmounts, shouldThrowOnInsufficientBalanceOrAllowance, v, r, s)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x394c21e7.
//
// Solidity: function cancelOrder(orderAddresses address[5], orderValues uint256[6], cancelTakerTokenAmount uint256) returns(uint256)
func (_Exchange *ExchangeTransactor) CancelOrder(opts *bind.TransactOpts, orderAddresses [5]common.Address, orderValues [6]*big.Int, cancelTakerTokenAmount *big.Int) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "cancelOrder", orderAddresses, orderValues, cancelTakerTokenAmount)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x394c21e7.
//
// Solidity: function cancelOrder(orderAddresses address[5], orderValues uint256[6], cancelTakerTokenAmount uint256) returns(uint256)
func (_Exchange *ExchangeSession) CancelOrder(orderAddresses [5]common.Address, orderValues [6]*big.Int, cancelTakerTokenAmount *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.CancelOrder(&_Exchange.TransactOpts, orderAddresses, orderValues, cancelTakerTokenAmount)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x394c21e7.
//
// Solidity: function cancelOrder(orderAddresses address[5], orderValues uint256[6], cancelTakerTokenAmount uint256) returns(uint256)
func (_Exchange *ExchangeTransactorSession) CancelOrder(orderAddresses [5]common.Address, orderValues [6]*big.Int, cancelTakerTokenAmount *big.Int) (*types.Transaction, error) {
	return _Exchange.Contract.CancelOrder(&_Exchange.TransactOpts, orderAddresses, orderValues, cancelTakerTokenAmount)
}

// FillOrKillOrder is a paid mutator transaction binding the contract method 0x741bcc93.
//
// Solidity: function fillOrKillOrder(orderAddresses address[5], orderValues uint256[6], fillTakerTokenAmount uint256, v uint8, r bytes32, s bytes32) returns()
func (_Exchange *ExchangeTransactor) FillOrKillOrder(opts *bind.TransactOpts, orderAddresses [5]common.Address, orderValues [6]*big.Int, fillTakerTokenAmount *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "fillOrKillOrder", orderAddresses, orderValues, fillTakerTokenAmount, v, r, s)
}

// FillOrKillOrder is a paid mutator transaction binding the contract method 0x741bcc93.
//
// Solidity: function fillOrKillOrder(orderAddresses address[5], orderValues uint256[6], fillTakerTokenAmount uint256, v uint8, r bytes32, s bytes32) returns()
func (_Exchange *ExchangeSession) FillOrKillOrder(orderAddresses [5]common.Address, orderValues [6]*big.Int, fillTakerTokenAmount *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.FillOrKillOrder(&_Exchange.TransactOpts, orderAddresses, orderValues, fillTakerTokenAmount, v, r, s)
}

// FillOrKillOrder is a paid mutator transaction binding the contract method 0x741bcc93.
//
// Solidity: function fillOrKillOrder(orderAddresses address[5], orderValues uint256[6], fillTakerTokenAmount uint256, v uint8, r bytes32, s bytes32) returns()
func (_Exchange *ExchangeTransactorSession) FillOrKillOrder(orderAddresses [5]common.Address, orderValues [6]*big.Int, fillTakerTokenAmount *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.FillOrKillOrder(&_Exchange.TransactOpts, orderAddresses, orderValues, fillTakerTokenAmount, v, r, s)
}

// FillOrder is a paid mutator transaction binding the contract method 0xbc61394a.
//
// Solidity: function fillOrder(orderAddresses address[5], orderValues uint256[6], fillTakerTokenAmount uint256, shouldThrowOnInsufficientBalanceOrAllowance bool, v uint8, r bytes32, s bytes32) returns(filledTakerTokenAmount uint256)
func (_Exchange *ExchangeTransactor) FillOrder(opts *bind.TransactOpts, orderAddresses [5]common.Address, orderValues [6]*big.Int, fillTakerTokenAmount *big.Int, shouldThrowOnInsufficientBalanceOrAllowance bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "fillOrder", orderAddresses, orderValues, fillTakerTokenAmount, shouldThrowOnInsufficientBalanceOrAllowance, v, r, s)
}

// FillOrder is a paid mutator transaction binding the contract method 0xbc61394a.
//
// Solidity: function fillOrder(orderAddresses address[5], orderValues uint256[6], fillTakerTokenAmount uint256, shouldThrowOnInsufficientBalanceOrAllowance bool, v uint8, r bytes32, s bytes32) returns(filledTakerTokenAmount uint256)
func (_Exchange *ExchangeSession) FillOrder(orderAddresses [5]common.Address, orderValues [6]*big.Int, fillTakerTokenAmount *big.Int, shouldThrowOnInsufficientBalanceOrAllowance bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.FillOrder(&_Exchange.TransactOpts, orderAddresses, orderValues, fillTakerTokenAmount, shouldThrowOnInsufficientBalanceOrAllowance, v, r, s)
}

// FillOrder is a paid mutator transaction binding the contract method 0xbc61394a.
//
// Solidity: function fillOrder(orderAddresses address[5], orderValues uint256[6], fillTakerTokenAmount uint256, shouldThrowOnInsufficientBalanceOrAllowance bool, v uint8, r bytes32, s bytes32) returns(filledTakerTokenAmount uint256)
func (_Exchange *ExchangeTransactorSession) FillOrder(orderAddresses [5]common.Address, orderValues [6]*big.Int, fillTakerTokenAmount *big.Int, shouldThrowOnInsufficientBalanceOrAllowance bool, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.FillOrder(&_Exchange.TransactOpts, orderAddresses, orderValues, fillTakerTokenAmount, shouldThrowOnInsufficientBalanceOrAllowance, v, r, s)
}

// FillOrdersUpTo is a paid mutator transaction binding the contract method 0x363349be.
//
// Solidity: function fillOrdersUpTo(orderAddresses address[5][], orderValues uint256[6][], fillTakerTokenAmount uint256, shouldThrowOnInsufficientBalanceOrAllowance bool, v uint8[], r bytes32[], s bytes32[]) returns(uint256)
func (_Exchange *ExchangeTransactor) FillOrdersUpTo(opts *bind.TransactOpts, orderAddresses [][5]common.Address, orderValues [][6]*big.Int, fillTakerTokenAmount *big.Int, shouldThrowOnInsufficientBalanceOrAllowance bool, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _Exchange.contract.Transact(opts, "fillOrdersUpTo", orderAddresses, orderValues, fillTakerTokenAmount, shouldThrowOnInsufficientBalanceOrAllowance, v, r, s)
}

// FillOrdersUpTo is a paid mutator transaction binding the contract method 0x363349be.
//
// Solidity: function fillOrdersUpTo(orderAddresses address[5][], orderValues uint256[6][], fillTakerTokenAmount uint256, shouldThrowOnInsufficientBalanceOrAllowance bool, v uint8[], r bytes32[], s bytes32[]) returns(uint256)
func (_Exchange *ExchangeSession) FillOrdersUpTo(orderAddresses [][5]common.Address, orderValues [][6]*big.Int, fillTakerTokenAmount *big.Int, shouldThrowOnInsufficientBalanceOrAllowance bool, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.FillOrdersUpTo(&_Exchange.TransactOpts, orderAddresses, orderValues, fillTakerTokenAmount, shouldThrowOnInsufficientBalanceOrAllowance, v, r, s)
}

// FillOrdersUpTo is a paid mutator transaction binding the contract method 0x363349be.
//
// Solidity: function fillOrdersUpTo(orderAddresses address[5][], orderValues uint256[6][], fillTakerTokenAmount uint256, shouldThrowOnInsufficientBalanceOrAllowance bool, v uint8[], r bytes32[], s bytes32[]) returns(uint256)
func (_Exchange *ExchangeTransactorSession) FillOrdersUpTo(orderAddresses [][5]common.Address, orderValues [][6]*big.Int, fillTakerTokenAmount *big.Int, shouldThrowOnInsufficientBalanceOrAllowance bool, v []uint8, r [][32]byte, s [][32]byte) (*types.Transaction, error) {
	return _Exchange.Contract.FillOrdersUpTo(&_Exchange.TransactOpts, orderAddresses, orderValues, fillTakerTokenAmount, shouldThrowOnInsufficientBalanceOrAllowance, v, r, s)
}

// ExchangeLogCancelIterator is returned from FilterLogCancel and is used to iterate over the raw logs and unpacked data for LogCancel events raised by the Exchange contract.
type ExchangeLogCancelIterator struct {
	Event *ExchangeLogCancel // Event containing the contract specifics and raw log

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
func (it *ExchangeLogCancelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogCancel)
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
		it.Event = new(ExchangeLogCancel)
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
func (it *ExchangeLogCancelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogCancelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogCancel represents a LogCancel event raised by the Exchange contract.
type ExchangeLogCancel struct {
	Maker                     common.Address
	FeeRecipient              common.Address
	MakerToken                common.Address
	TakerToken                common.Address
	CancelledMakerTokenAmount *big.Int
	CancelledTakerTokenAmount *big.Int
	Tokens                    [32]byte
	OrderHash                 [32]byte
	Raw                       types.Log // Blockchain specific contextual infos
}

// FilterLogCancel is a free log retrieval operation binding the contract event 0x67d66f160bc93d925d05dae1794c90d2d6d6688b29b84ff069398a9b04587131.
//
// Solidity: e LogCancel(maker indexed address, feeRecipient indexed address, makerToken address, takerToken address, cancelledMakerTokenAmount uint256, cancelledTakerTokenAmount uint256, tokens indexed bytes32, orderHash bytes32)
func (_Exchange *ExchangeFilterer) FilterLogCancel(opts *bind.FilterOpts, maker []common.Address, feeRecipient []common.Address, tokens [][32]byte) (*ExchangeLogCancelIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var feeRecipientRule []interface{}
	for _, feeRecipientItem := range feeRecipient {
		feeRecipientRule = append(feeRecipientRule, feeRecipientItem)
	}

	var tokensRule []interface{}
	for _, tokensItem := range tokens {
		tokensRule = append(tokensRule, tokensItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogCancel", makerRule, feeRecipientRule, tokensRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeLogCancelIterator{contract: _Exchange.contract, event: "LogCancel", logs: logs, sub: sub}, nil
}

// WatchLogCancel is a free log subscription operation binding the contract event 0x67d66f160bc93d925d05dae1794c90d2d6d6688b29b84ff069398a9b04587131.
//
// Solidity: e LogCancel(maker indexed address, feeRecipient indexed address, makerToken address, takerToken address, cancelledMakerTokenAmount uint256, cancelledTakerTokenAmount uint256, tokens indexed bytes32, orderHash bytes32)
func (_Exchange *ExchangeFilterer) WatchLogCancel(opts *bind.WatchOpts, sink chan<- *ExchangeLogCancel, maker []common.Address, feeRecipient []common.Address, tokens [][32]byte) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var feeRecipientRule []interface{}
	for _, feeRecipientItem := range feeRecipient {
		feeRecipientRule = append(feeRecipientRule, feeRecipientItem)
	}

	var tokensRule []interface{}
	for _, tokensItem := range tokens {
		tokensRule = append(tokensRule, tokensItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogCancel", makerRule, feeRecipientRule, tokensRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogCancel)
				if err := _Exchange.contract.UnpackLog(event, "LogCancel", log); err != nil {
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

// ExchangeLogErrorIterator is returned from FilterLogError and is used to iterate over the raw logs and unpacked data for LogError events raised by the Exchange contract.
type ExchangeLogErrorIterator struct {
	Event *ExchangeLogError // Event containing the contract specifics and raw log

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
func (it *ExchangeLogErrorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogError)
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
		it.Event = new(ExchangeLogError)
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
func (it *ExchangeLogErrorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogErrorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogError represents a LogError event raised by the Exchange contract.
type ExchangeLogError struct {
	ErrorId   uint8
	OrderHash [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogError is a free log retrieval operation binding the contract event 0x36d86c59e00bd73dc19ba3adfe068e4b64ac7e92be35546adeddf1b956a87e90.
//
// Solidity: e LogError(errorId indexed uint8, orderHash indexed bytes32)
func (_Exchange *ExchangeFilterer) FilterLogError(opts *bind.FilterOpts, errorId []uint8, orderHash [][32]byte) (*ExchangeLogErrorIterator, error) {

	var errorIdRule []interface{}
	for _, errorIdItem := range errorId {
		errorIdRule = append(errorIdRule, errorIdItem)
	}
	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogError", errorIdRule, orderHashRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeLogErrorIterator{contract: _Exchange.contract, event: "LogError", logs: logs, sub: sub}, nil
}

// WatchLogError is a free log subscription operation binding the contract event 0x36d86c59e00bd73dc19ba3adfe068e4b64ac7e92be35546adeddf1b956a87e90.
//
// Solidity: e LogError(errorId indexed uint8, orderHash indexed bytes32)
func (_Exchange *ExchangeFilterer) WatchLogError(opts *bind.WatchOpts, sink chan<- *ExchangeLogError, errorId []uint8, orderHash [][32]byte) (event.Subscription, error) {

	var errorIdRule []interface{}
	for _, errorIdItem := range errorId {
		errorIdRule = append(errorIdRule, errorIdItem)
	}
	var orderHashRule []interface{}
	for _, orderHashItem := range orderHash {
		orderHashRule = append(orderHashRule, orderHashItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogError", errorIdRule, orderHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogError)
				if err := _Exchange.contract.UnpackLog(event, "LogError", log); err != nil {
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

// ExchangeLogFillIterator is returned from FilterLogFill and is used to iterate over the raw logs and unpacked data for LogFill events raised by the Exchange contract.
type ExchangeLogFillIterator struct {
	Event *ExchangeLogFill // Event containing the contract specifics and raw log

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
func (it *ExchangeLogFillIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeLogFill)
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
		it.Event = new(ExchangeLogFill)
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
func (it *ExchangeLogFillIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeLogFillIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeLogFill represents a LogFill event raised by the Exchange contract.
type ExchangeLogFill struct {
	Maker                  common.Address
	Taker                  common.Address
	FeeRecipient           common.Address
	MakerToken             common.Address
	TakerToken             common.Address
	FilledMakerTokenAmount *big.Int
	FilledTakerTokenAmount *big.Int
	PaidMakerFee           *big.Int
	PaidTakerFee           *big.Int
	Tokens                 [32]byte
	OrderHash              [32]byte
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterLogFill is a free log retrieval operation binding the contract event 0x0d0b9391970d9a25552f37d436d2aae2925e2bfe1b2a923754bada030c498cb3.
//
// Solidity: e LogFill(maker indexed address, taker address, feeRecipient indexed address, makerToken address, takerToken address, filledMakerTokenAmount uint256, filledTakerTokenAmount uint256, paidMakerFee uint256, paidTakerFee uint256, tokens indexed bytes32, orderHash bytes32)
func (_Exchange *ExchangeFilterer) FilterLogFill(opts *bind.FilterOpts, maker []common.Address, feeRecipient []common.Address, tokens [][32]byte) (*ExchangeLogFillIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	var feeRecipientRule []interface{}
	for _, feeRecipientItem := range feeRecipient {
		feeRecipientRule = append(feeRecipientRule, feeRecipientItem)
	}

	var tokensRule []interface{}
	for _, tokensItem := range tokens {
		tokensRule = append(tokensRule, tokensItem)
	}

	logs, sub, err := _Exchange.contract.FilterLogs(opts, "LogFill", makerRule, feeRecipientRule, tokensRule)
	if err != nil {
		return nil, err
	}
	return &ExchangeLogFillIterator{contract: _Exchange.contract, event: "LogFill", logs: logs, sub: sub}, nil
}

// WatchLogFill is a free log subscription operation binding the contract event 0x0d0b9391970d9a25552f37d436d2aae2925e2bfe1b2a923754bada030c498cb3.
//
// Solidity: e LogFill(maker indexed address, taker address, feeRecipient indexed address, makerToken address, takerToken address, filledMakerTokenAmount uint256, filledTakerTokenAmount uint256, paidMakerFee uint256, paidTakerFee uint256, tokens indexed bytes32, orderHash bytes32)
func (_Exchange *ExchangeFilterer) WatchLogFill(opts *bind.WatchOpts, sink chan<- *ExchangeLogFill, maker []common.Address, feeRecipient []common.Address, tokens [][32]byte) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}

	var feeRecipientRule []interface{}
	for _, feeRecipientItem := range feeRecipient {
		feeRecipientRule = append(feeRecipientRule, feeRecipientItem)
	}

	var tokensRule []interface{}
	for _, tokensItem := range tokens {
		tokensRule = append(tokensRule, tokensItem)
	}

	logs, sub, err := _Exchange.contract.WatchLogs(opts, "LogFill", makerRule, feeRecipientRule, tokensRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeLogFill)
				if err := _Exchange.contract.UnpackLog(event, "LogFill", log); err != nil {
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

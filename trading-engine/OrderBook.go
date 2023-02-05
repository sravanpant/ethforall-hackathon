// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package OrderBook

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

// OrderBookMetaData contains all meta data concerning the OrderBook contract.
var OrderBookMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"askAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"askSize\",\"type\":\"uint256\"}],\"name\":\"NewAskEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidSize\",\"type\":\"uint256\"}],\"name\":\"NewBidEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"orderCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_askAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_askSize\",\"type\":\"uint256\"}],\"name\":\"addAsk\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bidAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bidSize\",\"type\":\"uint256\"}],\"name\":\"addBid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_orderId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_bidder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_asker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountFixed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_sizeFixed\",\"type\":\"uint256\"}],\"name\":\"addOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"orders\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"orderId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"asker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountFixed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sizeFixed\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610638806100206000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c806321b23f04146100515780635ff70bbe146100665780638b9b55e614610079578063a85c38ef1461008c575b600080fd5b61006461005f3660046102e9565b6100b9565b005b61006461007436600461033d565b6100ff565b6100646100873660046102e9565b6101ce565b61009f61009a366004610421565b61020c565b6040516100b095949392919061043a565b60405180910390f35b60408051338152602081018490529081018290527f779e06bbce3c0e611f2470dae49dcfb33e158093312fb038097279e599d553ae906060015b60405180910390a15050565b6040805160a0810182528681526001600160a01b038087166020830152851691810191909152606081018390526080810182905260008054600181018255908052815160059091027f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563019081906101769082610542565b5060208201516001820180546001600160a01b039283166001600160a01b0319918216179091556040840151600284018054919093169116179055606082015160038201556080909101516004909101555050505050565b60408051338152602081018490529081018290527f068585832e0e7b6f9edc74eec9a2ccae12d08ee49c155cbaac7825ac8fe72e01906060016100f3565b6000818154811061021c57600080fd5b906000526020600020906005020160009150905080600001805461023f906104b9565b80601f016020809104026020016040519081016040528092919081815260200182805461026b906104b9565b80156102b85780601f1061028d576101008083540402835291602001916102b8565b820191906000526020600020905b81548152906001019060200180831161029b57829003601f168201915b5050505060018301546002840154600385015460049095015493946001600160a01b03928316949290911692509085565b600080604083850312156102fc57600080fd5b50508035926020909101359150565b634e487b7160e01b600052604160045260246000fd5b80356001600160a01b038116811461033857600080fd5b919050565b600080600080600060a0868803121561035557600080fd5b853567ffffffffffffffff8082111561036d57600080fd5b818801915088601f83011261038157600080fd5b8135818111156103935761039361030b565b604051601f8201601f19908116603f011681019083821181831017156103bb576103bb61030b565b816040528281528b60208487010111156103d457600080fd5b8260208601602083013760006020848301015280995050505050506103fb60208701610321565b935061040960408701610321565b94979396509394606081013594506080013592915050565b60006020828403121561043357600080fd5b5035919050565b60a08152600086518060a084015260005b81811015610468576020818a0181015160c086840101520161044b565b50600060c0828501015260c0601f19601f83011684010191505061049760208301876001600160a01b03169052565b6001600160a01b03851660408301526060820193909352608001529392505050565b600181811c908216806104cd57607f821691505b6020821081036104ed57634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561053d57600081815260208120601f850160051c8101602086101561051a5750805b601f850160051c820191505b8181101561053957828155600101610526565b5050505b505050565b815167ffffffffffffffff81111561055c5761055c61030b565b6105708161056a84546104b9565b846104f3565b602080601f8311600181146105a5576000841561058d5750858301515b600019600386901b1c1916600185901b178555610539565b600085815260208120601f198616915b828110156105d4578886015182559484019460019091019084016105b5565b50858210156105f25787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fea2646970667358221220999b0320afe1c7971e67e3c764aaf79f20e69d5324822247aa3cb8183c40a15964736f6c63430008110033",
}

// OrderBookABI is the input ABI used to generate the binding from.
// Deprecated: Use OrderBookMetaData.ABI instead.
var OrderBookABI = OrderBookMetaData.ABI

// OrderBookBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OrderBookMetaData.Bin instead.
var OrderBookBin = OrderBookMetaData.Bin

// DeployOrderBook deploys a new Ethereum contract, binding an instance of OrderBook to it.
func DeployOrderBook(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OrderBook, error) {
	parsed, err := OrderBookMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OrderBookBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OrderBook{OrderBookCaller: OrderBookCaller{contract: contract}, OrderBookTransactor: OrderBookTransactor{contract: contract}, OrderBookFilterer: OrderBookFilterer{contract: contract}}, nil
}

// OrderBook is an auto generated Go binding around an Ethereum contract.
type OrderBook struct {
	OrderBookCaller     // Read-only binding to the contract
	OrderBookTransactor // Write-only binding to the contract
	OrderBookFilterer   // Log filterer for contract events
}

// OrderBookCaller is an auto generated read-only Go binding around an Ethereum contract.
type OrderBookCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderBookTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OrderBookTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderBookFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OrderBookFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OrderBookSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OrderBookSession struct {
	Contract     *OrderBook        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OrderBookCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OrderBookCallerSession struct {
	Contract *OrderBookCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// OrderBookTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OrderBookTransactorSession struct {
	Contract     *OrderBookTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// OrderBookRaw is an auto generated low-level Go binding around an Ethereum contract.
type OrderBookRaw struct {
	Contract *OrderBook // Generic contract binding to access the raw methods on
}

// OrderBookCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OrderBookCallerRaw struct {
	Contract *OrderBookCaller // Generic read-only contract binding to access the raw methods on
}

// OrderBookTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OrderBookTransactorRaw struct {
	Contract *OrderBookTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOrderBook creates a new instance of OrderBook, bound to a specific deployed contract.
func NewOrderBook(address common.Address, backend bind.ContractBackend) (*OrderBook, error) {
	contract, err := bindOrderBook(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OrderBook{OrderBookCaller: OrderBookCaller{contract: contract}, OrderBookTransactor: OrderBookTransactor{contract: contract}, OrderBookFilterer: OrderBookFilterer{contract: contract}}, nil
}

// NewOrderBookCaller creates a new read-only instance of OrderBook, bound to a specific deployed contract.
func NewOrderBookCaller(address common.Address, caller bind.ContractCaller) (*OrderBookCaller, error) {
	contract, err := bindOrderBook(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OrderBookCaller{contract: contract}, nil
}

// NewOrderBookTransactor creates a new write-only instance of OrderBook, bound to a specific deployed contract.
func NewOrderBookTransactor(address common.Address, transactor bind.ContractTransactor) (*OrderBookTransactor, error) {
	contract, err := bindOrderBook(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OrderBookTransactor{contract: contract}, nil
}

// NewOrderBookFilterer creates a new log filterer instance of OrderBook, bound to a specific deployed contract.
func NewOrderBookFilterer(address common.Address, filterer bind.ContractFilterer) (*OrderBookFilterer, error) {
	contract, err := bindOrderBook(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OrderBookFilterer{contract: contract}, nil
}

// bindOrderBook binds a generic wrapper to an already deployed contract.
func bindOrderBook(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OrderBookABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrderBook *OrderBookRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OrderBook.Contract.OrderBookCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrderBook *OrderBookRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderBook.Contract.OrderBookTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrderBook *OrderBookRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrderBook.Contract.OrderBookTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OrderBook *OrderBookCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OrderBook.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OrderBook *OrderBookTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OrderBook.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OrderBook *OrderBookTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OrderBook.Contract.contract.Transact(opts, method, params...)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(string orderId, address bidder, address asker, uint256 amountFixed, uint256 sizeFixed)
func (_OrderBook *OrderBookCaller) Orders(opts *bind.CallOpts, arg0 *big.Int) (struct {
	OrderId     string
	Bidder      common.Address
	Asker       common.Address
	AmountFixed *big.Int
	SizeFixed   *big.Int
}, error) {
	var out []interface{}
	err := _OrderBook.contract.Call(opts, &out, "orders", arg0)

	outstruct := new(struct {
		OrderId     string
		Bidder      common.Address
		Asker       common.Address
		AmountFixed *big.Int
		SizeFixed   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OrderId = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Bidder = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Asker = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.AmountFixed = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.SizeFixed = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(string orderId, address bidder, address asker, uint256 amountFixed, uint256 sizeFixed)
func (_OrderBook *OrderBookSession) Orders(arg0 *big.Int) (struct {
	OrderId     string
	Bidder      common.Address
	Asker       common.Address
	AmountFixed *big.Int
	SizeFixed   *big.Int
}, error) {
	return _OrderBook.Contract.Orders(&_OrderBook.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(string orderId, address bidder, address asker, uint256 amountFixed, uint256 sizeFixed)
func (_OrderBook *OrderBookCallerSession) Orders(arg0 *big.Int) (struct {
	OrderId     string
	Bidder      common.Address
	Asker       common.Address
	AmountFixed *big.Int
	SizeFixed   *big.Int
}, error) {
	return _OrderBook.Contract.Orders(&_OrderBook.CallOpts, arg0)
}

// AddAsk is a paid mutator transaction binding the contract method 0x8b9b55e6.
//
// Solidity: function addAsk(uint256 _askAmount, uint256 _askSize) returns()
func (_OrderBook *OrderBookTransactor) AddAsk(opts *bind.TransactOpts, _askAmount *big.Int, _askSize *big.Int) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "addAsk", _askAmount, _askSize)
}

// AddAsk is a paid mutator transaction binding the contract method 0x8b9b55e6.
//
// Solidity: function addAsk(uint256 _askAmount, uint256 _askSize) returns()
func (_OrderBook *OrderBookSession) AddAsk(_askAmount *big.Int, _askSize *big.Int) (*types.Transaction, error) {
	return _OrderBook.Contract.AddAsk(&_OrderBook.TransactOpts, _askAmount, _askSize)
}

// AddAsk is a paid mutator transaction binding the contract method 0x8b9b55e6.
//
// Solidity: function addAsk(uint256 _askAmount, uint256 _askSize) returns()
func (_OrderBook *OrderBookTransactorSession) AddAsk(_askAmount *big.Int, _askSize *big.Int) (*types.Transaction, error) {
	return _OrderBook.Contract.AddAsk(&_OrderBook.TransactOpts, _askAmount, _askSize)
}

// AddBid is a paid mutator transaction binding the contract method 0x21b23f04.
//
// Solidity: function addBid(uint256 _bidAmount, uint256 _bidSize) returns()
func (_OrderBook *OrderBookTransactor) AddBid(opts *bind.TransactOpts, _bidAmount *big.Int, _bidSize *big.Int) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "addBid", _bidAmount, _bidSize)
}

// AddBid is a paid mutator transaction binding the contract method 0x21b23f04.
//
// Solidity: function addBid(uint256 _bidAmount, uint256 _bidSize) returns()
func (_OrderBook *OrderBookSession) AddBid(_bidAmount *big.Int, _bidSize *big.Int) (*types.Transaction, error) {
	return _OrderBook.Contract.AddBid(&_OrderBook.TransactOpts, _bidAmount, _bidSize)
}

// AddBid is a paid mutator transaction binding the contract method 0x21b23f04.
//
// Solidity: function addBid(uint256 _bidAmount, uint256 _bidSize) returns()
func (_OrderBook *OrderBookTransactorSession) AddBid(_bidAmount *big.Int, _bidSize *big.Int) (*types.Transaction, error) {
	return _OrderBook.Contract.AddBid(&_OrderBook.TransactOpts, _bidAmount, _bidSize)
}

// AddOrder is a paid mutator transaction binding the contract method 0x5ff70bbe.
//
// Solidity: function addOrder(string _orderId, address _bidder, address _asker, uint256 _amountFixed, uint256 _sizeFixed) returns()
func (_OrderBook *OrderBookTransactor) AddOrder(opts *bind.TransactOpts, _orderId string, _bidder common.Address, _asker common.Address, _amountFixed *big.Int, _sizeFixed *big.Int) (*types.Transaction, error) {
	return _OrderBook.contract.Transact(opts, "addOrder", _orderId, _bidder, _asker, _amountFixed, _sizeFixed)
}

// AddOrder is a paid mutator transaction binding the contract method 0x5ff70bbe.
//
// Solidity: function addOrder(string _orderId, address _bidder, address _asker, uint256 _amountFixed, uint256 _sizeFixed) returns()
func (_OrderBook *OrderBookSession) AddOrder(_orderId string, _bidder common.Address, _asker common.Address, _amountFixed *big.Int, _sizeFixed *big.Int) (*types.Transaction, error) {
	return _OrderBook.Contract.AddOrder(&_OrderBook.TransactOpts, _orderId, _bidder, _asker, _amountFixed, _sizeFixed)
}

// AddOrder is a paid mutator transaction binding the contract method 0x5ff70bbe.
//
// Solidity: function addOrder(string _orderId, address _bidder, address _asker, uint256 _amountFixed, uint256 _sizeFixed) returns()
func (_OrderBook *OrderBookTransactorSession) AddOrder(_orderId string, _bidder common.Address, _asker common.Address, _amountFixed *big.Int, _sizeFixed *big.Int) (*types.Transaction, error) {
	return _OrderBook.Contract.AddOrder(&_OrderBook.TransactOpts, _orderId, _bidder, _asker, _amountFixed, _sizeFixed)
}

// OrderBookNewAskEventIterator is returned from FilterNewAskEvent and is used to iterate over the raw logs and unpacked data for NewAskEvent events raised by the OrderBook contract.
type OrderBookNewAskEventIterator struct {
	Event *OrderBookNewAskEvent // Event containing the contract specifics and raw log

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
func (it *OrderBookNewAskEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderBookNewAskEvent)
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
		it.Event = new(OrderBookNewAskEvent)
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
func (it *OrderBookNewAskEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderBookNewAskEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderBookNewAskEvent represents a NewAskEvent event raised by the OrderBook contract.
type OrderBookNewAskEvent struct {
	Asker     common.Address
	AskAmount *big.Int
	AskSize   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewAskEvent is a free log retrieval operation binding the contract event 0x068585832e0e7b6f9edc74eec9a2ccae12d08ee49c155cbaac7825ac8fe72e01.
//
// Solidity: event NewAskEvent(address asker, uint256 askAmount, uint256 askSize)
func (_OrderBook *OrderBookFilterer) FilterNewAskEvent(opts *bind.FilterOpts) (*OrderBookNewAskEventIterator, error) {

	logs, sub, err := _OrderBook.contract.FilterLogs(opts, "NewAskEvent")
	if err != nil {
		return nil, err
	}
	return &OrderBookNewAskEventIterator{contract: _OrderBook.contract, event: "NewAskEvent", logs: logs, sub: sub}, nil
}

// WatchNewAskEvent is a free log subscription operation binding the contract event 0x068585832e0e7b6f9edc74eec9a2ccae12d08ee49c155cbaac7825ac8fe72e01.
//
// Solidity: event NewAskEvent(address asker, uint256 askAmount, uint256 askSize)
func (_OrderBook *OrderBookFilterer) WatchNewAskEvent(opts *bind.WatchOpts, sink chan<- *OrderBookNewAskEvent) (event.Subscription, error) {

	logs, sub, err := _OrderBook.contract.WatchLogs(opts, "NewAskEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderBookNewAskEvent)
				if err := _OrderBook.contract.UnpackLog(event, "NewAskEvent", log); err != nil {
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

// ParseNewAskEvent is a log parse operation binding the contract event 0x068585832e0e7b6f9edc74eec9a2ccae12d08ee49c155cbaac7825ac8fe72e01.
//
// Solidity: event NewAskEvent(address asker, uint256 askAmount, uint256 askSize)
func (_OrderBook *OrderBookFilterer) ParseNewAskEvent(log types.Log) (*OrderBookNewAskEvent, error) {
	event := new(OrderBookNewAskEvent)
	if err := _OrderBook.contract.UnpackLog(event, "NewAskEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderBookNewBidEventIterator is returned from FilterNewBidEvent and is used to iterate over the raw logs and unpacked data for NewBidEvent events raised by the OrderBook contract.
type OrderBookNewBidEventIterator struct {
	Event *OrderBookNewBidEvent // Event containing the contract specifics and raw log

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
func (it *OrderBookNewBidEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderBookNewBidEvent)
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
		it.Event = new(OrderBookNewBidEvent)
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
func (it *OrderBookNewBidEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderBookNewBidEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderBookNewBidEvent represents a NewBidEvent event raised by the OrderBook contract.
type OrderBookNewBidEvent struct {
	Bidder    common.Address
	BidAmount *big.Int
	BidSize   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewBidEvent is a free log retrieval operation binding the contract event 0x779e06bbce3c0e611f2470dae49dcfb33e158093312fb038097279e599d553ae.
//
// Solidity: event NewBidEvent(address bidder, uint256 bidAmount, uint256 bidSize)
func (_OrderBook *OrderBookFilterer) FilterNewBidEvent(opts *bind.FilterOpts) (*OrderBookNewBidEventIterator, error) {

	logs, sub, err := _OrderBook.contract.FilterLogs(opts, "NewBidEvent")
	if err != nil {
		return nil, err
	}
	return &OrderBookNewBidEventIterator{contract: _OrderBook.contract, event: "NewBidEvent", logs: logs, sub: sub}, nil
}

// WatchNewBidEvent is a free log subscription operation binding the contract event 0x779e06bbce3c0e611f2470dae49dcfb33e158093312fb038097279e599d553ae.
//
// Solidity: event NewBidEvent(address bidder, uint256 bidAmount, uint256 bidSize)
func (_OrderBook *OrderBookFilterer) WatchNewBidEvent(opts *bind.WatchOpts, sink chan<- *OrderBookNewBidEvent) (event.Subscription, error) {

	logs, sub, err := _OrderBook.contract.WatchLogs(opts, "NewBidEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderBookNewBidEvent)
				if err := _OrderBook.contract.UnpackLog(event, "NewBidEvent", log); err != nil {
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

// ParseNewBidEvent is a log parse operation binding the contract event 0x779e06bbce3c0e611f2470dae49dcfb33e158093312fb038097279e599d553ae.
//
// Solidity: event NewBidEvent(address bidder, uint256 bidAmount, uint256 bidSize)
func (_OrderBook *OrderBookFilterer) ParseNewBidEvent(log types.Log) (*OrderBookNewBidEvent, error) {
	event := new(OrderBookNewBidEvent)
	if err := _OrderBook.contract.UnpackLog(event, "NewBidEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OrderBookOrderCreatedIterator is returned from FilterOrderCreated and is used to iterate over the raw logs and unpacked data for OrderCreated events raised by the OrderBook contract.
type OrderBookOrderCreatedIterator struct {
	Event *OrderBookOrderCreated // Event containing the contract specifics and raw log

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
func (it *OrderBookOrderCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OrderBookOrderCreated)
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
		it.Event = new(OrderBookOrderCreated)
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
func (it *OrderBookOrderCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OrderBookOrderCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OrderBookOrderCreated represents a OrderCreated event raised by the OrderBook contract.
type OrderBookOrderCreated struct {
	Bidder common.Address
	Asker  common.Address
	Amount *big.Int
	Size   *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOrderCreated is a free log retrieval operation binding the contract event 0xcf172c530bdf90ec43ad3bc18b233626d11364ebfdd97ac30946e2250c1b2dd1.
//
// Solidity: event orderCreated(address bidder, address asker, uint256 amount, uint256 size)
func (_OrderBook *OrderBookFilterer) FilterOrderCreated(opts *bind.FilterOpts) (*OrderBookOrderCreatedIterator, error) {

	logs, sub, err := _OrderBook.contract.FilterLogs(opts, "orderCreated")
	if err != nil {
		return nil, err
	}
	return &OrderBookOrderCreatedIterator{contract: _OrderBook.contract, event: "orderCreated", logs: logs, sub: sub}, nil
}

// WatchOrderCreated is a free log subscription operation binding the contract event 0xcf172c530bdf90ec43ad3bc18b233626d11364ebfdd97ac30946e2250c1b2dd1.
//
// Solidity: event orderCreated(address bidder, address asker, uint256 amount, uint256 size)
func (_OrderBook *OrderBookFilterer) WatchOrderCreated(opts *bind.WatchOpts, sink chan<- *OrderBookOrderCreated) (event.Subscription, error) {

	logs, sub, err := _OrderBook.contract.WatchLogs(opts, "orderCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OrderBookOrderCreated)
				if err := _OrderBook.contract.UnpackLog(event, "orderCreated", log); err != nil {
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

// ParseOrderCreated is a log parse operation binding the contract event 0xcf172c530bdf90ec43ad3bc18b233626d11364ebfdd97ac30946e2250c1b2dd1.
//
// Solidity: event orderCreated(address bidder, address asker, uint256 amount, uint256 size)
func (_OrderBook *OrderBookFilterer) ParseOrderCreated(log types.Log) (*OrderBookOrderCreated, error) {
	event := new(OrderBookOrderCreated)
	if err := _OrderBook.contract.UnpackLog(event, "orderCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

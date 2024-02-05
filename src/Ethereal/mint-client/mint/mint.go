// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mint

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

// PairingG1Point is an auto generated low-level Go binding around an user-defined struct.
type PairingG1Point struct {
	X *big.Int
	Y *big.Int
}

// PairingG2Point is an auto generated low-level Go binding around an user-defined struct.
type PairingG2Point struct {
	X [2]*big.Int
	Y [2]*big.Int
}

// MintMetaData contains all meta data concerning the Mint contract.
var MintMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"SRS_G1_X\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SRS_G1_Y\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SRS_G2_X_0\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SRS_G2_X_1\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SRS_G2_Y_0\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"SRS_G2_Y_1\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"commit\",\"inputs\":[{\"name\":\"coefficients\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPairing.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deposits\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"registered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"commitment\",\"type\":\"tuple\",\"internalType\":\"structPairing.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"publicKey\",\"type\":\"tuple\",\"internalType\":\"structPairing.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"escaped0\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"escaped1\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"evalPolyAt\",\"inputs\":[{\"name\":\"_coefficients\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_index\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"evaluateBarycentricPolynomial\",\"inputs\":[{\"name\":\"x\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"yValues\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"weights\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBarycentricWeights\",\"inputs\":[{\"name\":\"xValues\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCollectedY\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getNonce\",\"inputs\":[{\"name\":\"_user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"_proof\",\"type\":\"tuple\",\"internalType\":\"structPairing.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"_value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"register\",\"inputs\":[{\"name\":\"_commitment\",\"type\":\"tuple\",\"internalType\":\"structPairing.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"_proof\",\"type\":\"tuple\",\"internalType\":\"structPairing.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"_publicKey\",\"type\":\"tuple\",\"internalType\":\"structPairing.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"replay\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"verify\",\"inputs\":[{\"name\":\"_commitment\",\"type\":\"tuple\",\"internalType\":\"structPairing.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"_proof\",\"type\":\"tuple\",\"internalType\":\"structPairing.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"_index\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyMulti\",\"inputs\":[{\"name\":\"_commitment\",\"type\":\"tuple\",\"internalType\":\"structPairing.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"_proof\",\"type\":\"tuple\",\"internalType\":\"structPairing.G2Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"},{\"name\":\"Y\",\"type\":\"uint256[2]\",\"internalType\":\"uint256[2]\"}]},{\"name\":\"_indices\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_values\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_iCoeffs\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"_zCoeffs\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyPublicKey\",\"inputs\":[{\"name\":\"_commitment\",\"type\":\"tuple\",\"internalType\":\"structPairing.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"_proof\",\"type\":\"tuple\",\"internalType\":\"structPairing.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"_publicKey\",\"type\":\"tuple\",\"internalType\":\"structPairing.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Deposit\",\"inputs\":[{\"name\":\"publicKey\",\"type\":\"tuple\",\"indexed\":true,\"internalType\":\"structPairing.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"commitment\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structPairing.G1Point\",\"components\":[{\"name\":\"X\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"Y\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Minted\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"nonce\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Pwned\",\"inputs\":[{\"name\":\"privateKey\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// MintABI is the input ABI used to generate the binding from.
// Deprecated: Use MintMetaData.ABI instead.
var MintABI = MintMetaData.ABI

// Mint is an auto generated Go binding around an Ethereum contract.
type Mint struct {
	MintCaller     // Read-only binding to the contract
	MintTransactor // Write-only binding to the contract
	MintFilterer   // Log filterer for contract events
}

// MintCaller is an auto generated read-only Go binding around an Ethereum contract.
type MintCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MintTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MintFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MintSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MintSession struct {
	Contract     *Mint             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MintCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MintCallerSession struct {
	Contract *MintCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MintTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MintTransactorSession struct {
	Contract     *MintTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MintRaw is an auto generated low-level Go binding around an Ethereum contract.
type MintRaw struct {
	Contract *Mint // Generic contract binding to access the raw methods on
}

// MintCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MintCallerRaw struct {
	Contract *MintCaller // Generic read-only contract binding to access the raw methods on
}

// MintTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MintTransactorRaw struct {
	Contract *MintTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMint creates a new instance of Mint, bound to a specific deployed contract.
func NewMint(address common.Address, backend bind.ContractBackend) (*Mint, error) {
	contract, err := bindMint(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mint{MintCaller: MintCaller{contract: contract}, MintTransactor: MintTransactor{contract: contract}, MintFilterer: MintFilterer{contract: contract}}, nil
}

// NewMintCaller creates a new read-only instance of Mint, bound to a specific deployed contract.
func NewMintCaller(address common.Address, caller bind.ContractCaller) (*MintCaller, error) {
	contract, err := bindMint(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MintCaller{contract: contract}, nil
}

// NewMintTransactor creates a new write-only instance of Mint, bound to a specific deployed contract.
func NewMintTransactor(address common.Address, transactor bind.ContractTransactor) (*MintTransactor, error) {
	contract, err := bindMint(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MintTransactor{contract: contract}, nil
}

// NewMintFilterer creates a new log filterer instance of Mint, bound to a specific deployed contract.
func NewMintFilterer(address common.Address, filterer bind.ContractFilterer) (*MintFilterer, error) {
	contract, err := bindMint(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MintFilterer{contract: contract}, nil
}

// bindMint binds a generic wrapper to an already deployed contract.
func bindMint(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MintMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mint *MintRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mint.Contract.MintCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mint *MintRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mint.Contract.MintTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mint *MintRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mint.Contract.MintTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mint *MintCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mint.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mint *MintTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mint.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mint *MintTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mint.Contract.contract.Transact(opts, method, params...)
}

// SRSG1X is a free data retrieval call binding the contract method 0xdd735c0c.
//
// Solidity: function SRS_G1_X(uint256 ) view returns(uint256)
func (_Mint *MintCaller) SRSG1X(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "SRS_G1_X", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SRSG1X is a free data retrieval call binding the contract method 0xdd735c0c.
//
// Solidity: function SRS_G1_X(uint256 ) view returns(uint256)
func (_Mint *MintSession) SRSG1X(arg0 *big.Int) (*big.Int, error) {
	return _Mint.Contract.SRSG1X(&_Mint.CallOpts, arg0)
}

// SRSG1X is a free data retrieval call binding the contract method 0xdd735c0c.
//
// Solidity: function SRS_G1_X(uint256 ) view returns(uint256)
func (_Mint *MintCallerSession) SRSG1X(arg0 *big.Int) (*big.Int, error) {
	return _Mint.Contract.SRSG1X(&_Mint.CallOpts, arg0)
}

// SRSG1Y is a free data retrieval call binding the contract method 0x082a4770.
//
// Solidity: function SRS_G1_Y(uint256 ) view returns(uint256)
func (_Mint *MintCaller) SRSG1Y(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "SRS_G1_Y", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SRSG1Y is a free data retrieval call binding the contract method 0x082a4770.
//
// Solidity: function SRS_G1_Y(uint256 ) view returns(uint256)
func (_Mint *MintSession) SRSG1Y(arg0 *big.Int) (*big.Int, error) {
	return _Mint.Contract.SRSG1Y(&_Mint.CallOpts, arg0)
}

// SRSG1Y is a free data retrieval call binding the contract method 0x082a4770.
//
// Solidity: function SRS_G1_Y(uint256 ) view returns(uint256)
func (_Mint *MintCallerSession) SRSG1Y(arg0 *big.Int) (*big.Int, error) {
	return _Mint.Contract.SRSG1Y(&_Mint.CallOpts, arg0)
}

// SRSG2X0 is a free data retrieval call binding the contract method 0x71c1441e.
//
// Solidity: function SRS_G2_X_0(uint256 ) view returns(uint256)
func (_Mint *MintCaller) SRSG2X0(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "SRS_G2_X_0", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SRSG2X0 is a free data retrieval call binding the contract method 0x71c1441e.
//
// Solidity: function SRS_G2_X_0(uint256 ) view returns(uint256)
func (_Mint *MintSession) SRSG2X0(arg0 *big.Int) (*big.Int, error) {
	return _Mint.Contract.SRSG2X0(&_Mint.CallOpts, arg0)
}

// SRSG2X0 is a free data retrieval call binding the contract method 0x71c1441e.
//
// Solidity: function SRS_G2_X_0(uint256 ) view returns(uint256)
func (_Mint *MintCallerSession) SRSG2X0(arg0 *big.Int) (*big.Int, error) {
	return _Mint.Contract.SRSG2X0(&_Mint.CallOpts, arg0)
}

// SRSG2X1 is a free data retrieval call binding the contract method 0x82a124f6.
//
// Solidity: function SRS_G2_X_1(uint256 ) view returns(uint256)
func (_Mint *MintCaller) SRSG2X1(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "SRS_G2_X_1", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SRSG2X1 is a free data retrieval call binding the contract method 0x82a124f6.
//
// Solidity: function SRS_G2_X_1(uint256 ) view returns(uint256)
func (_Mint *MintSession) SRSG2X1(arg0 *big.Int) (*big.Int, error) {
	return _Mint.Contract.SRSG2X1(&_Mint.CallOpts, arg0)
}

// SRSG2X1 is a free data retrieval call binding the contract method 0x82a124f6.
//
// Solidity: function SRS_G2_X_1(uint256 ) view returns(uint256)
func (_Mint *MintCallerSession) SRSG2X1(arg0 *big.Int) (*big.Int, error) {
	return _Mint.Contract.SRSG2X1(&_Mint.CallOpts, arg0)
}

// SRSG2Y0 is a free data retrieval call binding the contract method 0xa80a0aca.
//
// Solidity: function SRS_G2_Y_0(uint256 ) view returns(uint256)
func (_Mint *MintCaller) SRSG2Y0(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "SRS_G2_Y_0", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SRSG2Y0 is a free data retrieval call binding the contract method 0xa80a0aca.
//
// Solidity: function SRS_G2_Y_0(uint256 ) view returns(uint256)
func (_Mint *MintSession) SRSG2Y0(arg0 *big.Int) (*big.Int, error) {
	return _Mint.Contract.SRSG2Y0(&_Mint.CallOpts, arg0)
}

// SRSG2Y0 is a free data retrieval call binding the contract method 0xa80a0aca.
//
// Solidity: function SRS_G2_Y_0(uint256 ) view returns(uint256)
func (_Mint *MintCallerSession) SRSG2Y0(arg0 *big.Int) (*big.Int, error) {
	return _Mint.Contract.SRSG2Y0(&_Mint.CallOpts, arg0)
}

// SRSG2Y1 is a free data retrieval call binding the contract method 0x7cf50768.
//
// Solidity: function SRS_G2_Y_1(uint256 ) view returns(uint256)
func (_Mint *MintCaller) SRSG2Y1(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "SRS_G2_Y_1", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SRSG2Y1 is a free data retrieval call binding the contract method 0x7cf50768.
//
// Solidity: function SRS_G2_Y_1(uint256 ) view returns(uint256)
func (_Mint *MintSession) SRSG2Y1(arg0 *big.Int) (*big.Int, error) {
	return _Mint.Contract.SRSG2Y1(&_Mint.CallOpts, arg0)
}

// SRSG2Y1 is a free data retrieval call binding the contract method 0x7cf50768.
//
// Solidity: function SRS_G2_Y_1(uint256 ) view returns(uint256)
func (_Mint *MintCallerSession) SRSG2Y1(arg0 *big.Int) (*big.Int, error) {
	return _Mint.Contract.SRSG2Y1(&_Mint.CallOpts, arg0)
}

// Commit is a free data retrieval call binding the contract method 0x59434d4a.
//
// Solidity: function commit(uint256[] coefficients) view returns((uint256,uint256))
func (_Mint *MintCaller) Commit(opts *bind.CallOpts, coefficients []*big.Int) (PairingG1Point, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "commit", coefficients)

	if err != nil {
		return *new(PairingG1Point), err
	}

	out0 := *abi.ConvertType(out[0], new(PairingG1Point)).(*PairingG1Point)

	return out0, err

}

// Commit is a free data retrieval call binding the contract method 0x59434d4a.
//
// Solidity: function commit(uint256[] coefficients) view returns((uint256,uint256))
func (_Mint *MintSession) Commit(coefficients []*big.Int) (PairingG1Point, error) {
	return _Mint.Contract.Commit(&_Mint.CallOpts, coefficients)
}

// Commit is a free data retrieval call binding the contract method 0x59434d4a.
//
// Solidity: function commit(uint256[] coefficients) view returns((uint256,uint256))
func (_Mint *MintCallerSession) Commit(coefficients []*big.Int) (PairingG1Point, error) {
	return _Mint.Contract.Commit(&_Mint.CallOpts, coefficients)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(bool registered, uint256 nonce, (uint256,uint256) commitment, (uint256,uint256) publicKey, bool escaped0, bool escaped1)
func (_Mint *MintCaller) Deposits(opts *bind.CallOpts, arg0 common.Address) (struct {
	Registered bool
	Nonce      *big.Int
	Commitment PairingG1Point
	PublicKey  PairingG1Point
	Escaped0   bool
	Escaped1   bool
}, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "deposits", arg0)

	outstruct := new(struct {
		Registered bool
		Nonce      *big.Int
		Commitment PairingG1Point
		PublicKey  PairingG1Point
		Escaped0   bool
		Escaped1   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Registered = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Nonce = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Commitment = *abi.ConvertType(out[2], new(PairingG1Point)).(*PairingG1Point)
	outstruct.PublicKey = *abi.ConvertType(out[3], new(PairingG1Point)).(*PairingG1Point)
	outstruct.Escaped0 = *abi.ConvertType(out[4], new(bool)).(*bool)
	outstruct.Escaped1 = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(bool registered, uint256 nonce, (uint256,uint256) commitment, (uint256,uint256) publicKey, bool escaped0, bool escaped1)
func (_Mint *MintSession) Deposits(arg0 common.Address) (struct {
	Registered bool
	Nonce      *big.Int
	Commitment PairingG1Point
	PublicKey  PairingG1Point
	Escaped0   bool
	Escaped1   bool
}, error) {
	return _Mint.Contract.Deposits(&_Mint.CallOpts, arg0)
}

// Deposits is a free data retrieval call binding the contract method 0xfc7e286d.
//
// Solidity: function deposits(address ) view returns(bool registered, uint256 nonce, (uint256,uint256) commitment, (uint256,uint256) publicKey, bool escaped0, bool escaped1)
func (_Mint *MintCallerSession) Deposits(arg0 common.Address) (struct {
	Registered bool
	Nonce      *big.Int
	Commitment PairingG1Point
	PublicKey  PairingG1Point
	Escaped0   bool
	Escaped1   bool
}, error) {
	return _Mint.Contract.Deposits(&_Mint.CallOpts, arg0)
}

// EvalPolyAt is a free data retrieval call binding the contract method 0xc1eb466c.
//
// Solidity: function evalPolyAt(uint256[] _coefficients, uint256 _index) pure returns(uint256)
func (_Mint *MintCaller) EvalPolyAt(opts *bind.CallOpts, _coefficients []*big.Int, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "evalPolyAt", _coefficients, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EvalPolyAt is a free data retrieval call binding the contract method 0xc1eb466c.
//
// Solidity: function evalPolyAt(uint256[] _coefficients, uint256 _index) pure returns(uint256)
func (_Mint *MintSession) EvalPolyAt(_coefficients []*big.Int, _index *big.Int) (*big.Int, error) {
	return _Mint.Contract.EvalPolyAt(&_Mint.CallOpts, _coefficients, _index)
}

// EvalPolyAt is a free data retrieval call binding the contract method 0xc1eb466c.
//
// Solidity: function evalPolyAt(uint256[] _coefficients, uint256 _index) pure returns(uint256)
func (_Mint *MintCallerSession) EvalPolyAt(_coefficients []*big.Int, _index *big.Int) (*big.Int, error) {
	return _Mint.Contract.EvalPolyAt(&_Mint.CallOpts, _coefficients, _index)
}

// EvaluateBarycentricPolynomial is a free data retrieval call binding the contract method 0x0e3d9b93.
//
// Solidity: function evaluateBarycentricPolynomial(uint256 x, uint256[] yValues, uint256[] weights) view returns(uint256)
func (_Mint *MintCaller) EvaluateBarycentricPolynomial(opts *bind.CallOpts, x *big.Int, yValues []*big.Int, weights []*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "evaluateBarycentricPolynomial", x, yValues, weights)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EvaluateBarycentricPolynomial is a free data retrieval call binding the contract method 0x0e3d9b93.
//
// Solidity: function evaluateBarycentricPolynomial(uint256 x, uint256[] yValues, uint256[] weights) view returns(uint256)
func (_Mint *MintSession) EvaluateBarycentricPolynomial(x *big.Int, yValues []*big.Int, weights []*big.Int) (*big.Int, error) {
	return _Mint.Contract.EvaluateBarycentricPolynomial(&_Mint.CallOpts, x, yValues, weights)
}

// EvaluateBarycentricPolynomial is a free data retrieval call binding the contract method 0x0e3d9b93.
//
// Solidity: function evaluateBarycentricPolynomial(uint256 x, uint256[] yValues, uint256[] weights) view returns(uint256)
func (_Mint *MintCallerSession) EvaluateBarycentricPolynomial(x *big.Int, yValues []*big.Int, weights []*big.Int) (*big.Int, error) {
	return _Mint.Contract.EvaluateBarycentricPolynomial(&_Mint.CallOpts, x, yValues, weights)
}

// GetBarycentricWeights is a free data retrieval call binding the contract method 0x07a4d15d.
//
// Solidity: function getBarycentricWeights(uint256[] xValues) view returns(uint256[])
func (_Mint *MintCaller) GetBarycentricWeights(opts *bind.CallOpts, xValues []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "getBarycentricWeights", xValues)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetBarycentricWeights is a free data retrieval call binding the contract method 0x07a4d15d.
//
// Solidity: function getBarycentricWeights(uint256[] xValues) view returns(uint256[])
func (_Mint *MintSession) GetBarycentricWeights(xValues []*big.Int) ([]*big.Int, error) {
	return _Mint.Contract.GetBarycentricWeights(&_Mint.CallOpts, xValues)
}

// GetBarycentricWeights is a free data retrieval call binding the contract method 0x07a4d15d.
//
// Solidity: function getBarycentricWeights(uint256[] xValues) view returns(uint256[])
func (_Mint *MintCallerSession) GetBarycentricWeights(xValues []*big.Int) ([]*big.Int, error) {
	return _Mint.Contract.GetBarycentricWeights(&_Mint.CallOpts, xValues)
}

// GetCollectedY is a free data retrieval call binding the contract method 0x8068d27d.
//
// Solidity: function getCollectedY(address user) view returns(uint256[])
func (_Mint *MintCaller) GetCollectedY(opts *bind.CallOpts, user common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "getCollectedY", user)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetCollectedY is a free data retrieval call binding the contract method 0x8068d27d.
//
// Solidity: function getCollectedY(address user) view returns(uint256[])
func (_Mint *MintSession) GetCollectedY(user common.Address) ([]*big.Int, error) {
	return _Mint.Contract.GetCollectedY(&_Mint.CallOpts, user)
}

// GetCollectedY is a free data retrieval call binding the contract method 0x8068d27d.
//
// Solidity: function getCollectedY(address user) view returns(uint256[])
func (_Mint *MintCallerSession) GetCollectedY(user common.Address) ([]*big.Int, error) {
	return _Mint.Contract.GetCollectedY(&_Mint.CallOpts, user)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address _user) view returns(uint256)
func (_Mint *MintCaller) GetNonce(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "getNonce", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address _user) view returns(uint256)
func (_Mint *MintSession) GetNonce(_user common.Address) (*big.Int, error) {
	return _Mint.Contract.GetNonce(&_Mint.CallOpts, _user)
}

// GetNonce is a free data retrieval call binding the contract method 0x2d0335ab.
//
// Solidity: function getNonce(address _user) view returns(uint256)
func (_Mint *MintCallerSession) GetNonce(_user common.Address) (*big.Int, error) {
	return _Mint.Contract.GetNonce(&_Mint.CallOpts, _user)
}

// Verify is a free data retrieval call binding the contract method 0x903ae890.
//
// Solidity: function verify((uint256,uint256) _commitment, (uint256,uint256) _proof, uint256 _index, uint256 _value) view returns(bool)
func (_Mint *MintCaller) Verify(opts *bind.CallOpts, _commitment PairingG1Point, _proof PairingG1Point, _index *big.Int, _value *big.Int) (bool, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "verify", _commitment, _proof, _index, _value)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verify is a free data retrieval call binding the contract method 0x903ae890.
//
// Solidity: function verify((uint256,uint256) _commitment, (uint256,uint256) _proof, uint256 _index, uint256 _value) view returns(bool)
func (_Mint *MintSession) Verify(_commitment PairingG1Point, _proof PairingG1Point, _index *big.Int, _value *big.Int) (bool, error) {
	return _Mint.Contract.Verify(&_Mint.CallOpts, _commitment, _proof, _index, _value)
}

// Verify is a free data retrieval call binding the contract method 0x903ae890.
//
// Solidity: function verify((uint256,uint256) _commitment, (uint256,uint256) _proof, uint256 _index, uint256 _value) view returns(bool)
func (_Mint *MintCallerSession) Verify(_commitment PairingG1Point, _proof PairingG1Point, _index *big.Int, _value *big.Int) (bool, error) {
	return _Mint.Contract.Verify(&_Mint.CallOpts, _commitment, _proof, _index, _value)
}

// VerifyMulti is a free data retrieval call binding the contract method 0x87a3b8d9.
//
// Solidity: function verifyMulti((uint256,uint256) _commitment, (uint256[2],uint256[2]) _proof, uint256[] _indices, uint256[] _values, uint256[] _iCoeffs, uint256[] _zCoeffs) view returns(bool)
func (_Mint *MintCaller) VerifyMulti(opts *bind.CallOpts, _commitment PairingG1Point, _proof PairingG2Point, _indices []*big.Int, _values []*big.Int, _iCoeffs []*big.Int, _zCoeffs []*big.Int) (bool, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "verifyMulti", _commitment, _proof, _indices, _values, _iCoeffs, _zCoeffs)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyMulti is a free data retrieval call binding the contract method 0x87a3b8d9.
//
// Solidity: function verifyMulti((uint256,uint256) _commitment, (uint256[2],uint256[2]) _proof, uint256[] _indices, uint256[] _values, uint256[] _iCoeffs, uint256[] _zCoeffs) view returns(bool)
func (_Mint *MintSession) VerifyMulti(_commitment PairingG1Point, _proof PairingG2Point, _indices []*big.Int, _values []*big.Int, _iCoeffs []*big.Int, _zCoeffs []*big.Int) (bool, error) {
	return _Mint.Contract.VerifyMulti(&_Mint.CallOpts, _commitment, _proof, _indices, _values, _iCoeffs, _zCoeffs)
}

// VerifyMulti is a free data retrieval call binding the contract method 0x87a3b8d9.
//
// Solidity: function verifyMulti((uint256,uint256) _commitment, (uint256[2],uint256[2]) _proof, uint256[] _indices, uint256[] _values, uint256[] _iCoeffs, uint256[] _zCoeffs) view returns(bool)
func (_Mint *MintCallerSession) VerifyMulti(_commitment PairingG1Point, _proof PairingG2Point, _indices []*big.Int, _values []*big.Int, _iCoeffs []*big.Int, _zCoeffs []*big.Int) (bool, error) {
	return _Mint.Contract.VerifyMulti(&_Mint.CallOpts, _commitment, _proof, _indices, _values, _iCoeffs, _zCoeffs)
}

// VerifyPublicKey is a free data retrieval call binding the contract method 0x1c354fe5.
//
// Solidity: function verifyPublicKey((uint256,uint256) _commitment, (uint256,uint256) _proof, (uint256,uint256) _publicKey) view returns(bool)
func (_Mint *MintCaller) VerifyPublicKey(opts *bind.CallOpts, _commitment PairingG1Point, _proof PairingG1Point, _publicKey PairingG1Point) (bool, error) {
	var out []interface{}
	err := _Mint.contract.Call(opts, &out, "verifyPublicKey", _commitment, _proof, _publicKey)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyPublicKey is a free data retrieval call binding the contract method 0x1c354fe5.
//
// Solidity: function verifyPublicKey((uint256,uint256) _commitment, (uint256,uint256) _proof, (uint256,uint256) _publicKey) view returns(bool)
func (_Mint *MintSession) VerifyPublicKey(_commitment PairingG1Point, _proof PairingG1Point, _publicKey PairingG1Point) (bool, error) {
	return _Mint.Contract.VerifyPublicKey(&_Mint.CallOpts, _commitment, _proof, _publicKey)
}

// VerifyPublicKey is a free data retrieval call binding the contract method 0x1c354fe5.
//
// Solidity: function verifyPublicKey((uint256,uint256) _commitment, (uint256,uint256) _proof, (uint256,uint256) _publicKey) view returns(bool)
func (_Mint *MintCallerSession) VerifyPublicKey(_commitment PairingG1Point, _proof PairingG1Point, _publicKey PairingG1Point) (bool, error) {
	return _Mint.Contract.VerifyPublicKey(&_Mint.CallOpts, _commitment, _proof, _publicKey)
}

// Mint is a paid mutator transaction binding the contract method 0x46409539.
//
// Solidity: function mint((uint256,uint256) _proof, uint256 _value) returns()
func (_Mint *MintTransactor) Mint(opts *bind.TransactOpts, _proof PairingG1Point, _value *big.Int) (*types.Transaction, error) {
	return _Mint.contract.Transact(opts, "mint", _proof, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x46409539.
//
// Solidity: function mint((uint256,uint256) _proof, uint256 _value) returns()
func (_Mint *MintSession) Mint(_proof PairingG1Point, _value *big.Int) (*types.Transaction, error) {
	return _Mint.Contract.Mint(&_Mint.TransactOpts, _proof, _value)
}

// Mint is a paid mutator transaction binding the contract method 0x46409539.
//
// Solidity: function mint((uint256,uint256) _proof, uint256 _value) returns()
func (_Mint *MintTransactorSession) Mint(_proof PairingG1Point, _value *big.Int) (*types.Transaction, error) {
	return _Mint.Contract.Mint(&_Mint.TransactOpts, _proof, _value)
}

// Register is a paid mutator transaction binding the contract method 0x42a67927.
//
// Solidity: function register((uint256,uint256) _commitment, (uint256,uint256) _proof, (uint256,uint256) _publicKey) payable returns()
func (_Mint *MintTransactor) Register(opts *bind.TransactOpts, _commitment PairingG1Point, _proof PairingG1Point, _publicKey PairingG1Point) (*types.Transaction, error) {
	return _Mint.contract.Transact(opts, "register", _commitment, _proof, _publicKey)
}

// Register is a paid mutator transaction binding the contract method 0x42a67927.
//
// Solidity: function register((uint256,uint256) _commitment, (uint256,uint256) _proof, (uint256,uint256) _publicKey) payable returns()
func (_Mint *MintSession) Register(_commitment PairingG1Point, _proof PairingG1Point, _publicKey PairingG1Point) (*types.Transaction, error) {
	return _Mint.Contract.Register(&_Mint.TransactOpts, _commitment, _proof, _publicKey)
}

// Register is a paid mutator transaction binding the contract method 0x42a67927.
//
// Solidity: function register((uint256,uint256) _commitment, (uint256,uint256) _proof, (uint256,uint256) _publicKey) payable returns()
func (_Mint *MintTransactorSession) Register(_commitment PairingG1Point, _proof PairingG1Point, _publicKey PairingG1Point) (*types.Transaction, error) {
	return _Mint.Contract.Register(&_Mint.TransactOpts, _commitment, _proof, _publicKey)
}

// Replay is a paid mutator transaction binding the contract method 0xb5b9cd7f.
//
// Solidity: function replay() returns()
func (_Mint *MintTransactor) Replay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mint.contract.Transact(opts, "replay")
}

// Replay is a paid mutator transaction binding the contract method 0xb5b9cd7f.
//
// Solidity: function replay() returns()
func (_Mint *MintSession) Replay() (*types.Transaction, error) {
	return _Mint.Contract.Replay(&_Mint.TransactOpts)
}

// Replay is a paid mutator transaction binding the contract method 0xb5b9cd7f.
//
// Solidity: function replay() returns()
func (_Mint *MintTransactorSession) Replay() (*types.Transaction, error) {
	return _Mint.Contract.Replay(&_Mint.TransactOpts)
}

// MintDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Mint contract.
type MintDepositIterator struct {
	Event *MintDeposit // Event containing the contract specifics and raw log

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
func (it *MintDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintDeposit)
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
		it.Event = new(MintDeposit)
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
func (it *MintDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintDeposit represents a Deposit event raised by the Mint contract.
type MintDeposit struct {
	PublicKey  PairingG1Point
	Commitment PairingG1Point
	Timestamp  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0x5e02e546dec63caecdebd8692b65ccdbaa30e25ac8395aaef548f4ff09b78aa4.
//
// Solidity: event Deposit((uint256,uint256) indexed publicKey, (uint256,uint256) commitment, uint256 timestamp)
func (_Mint *MintFilterer) FilterDeposit(opts *bind.FilterOpts, publicKey []PairingG1Point) (*MintDepositIterator, error) {

	var publicKeyRule []interface{}
	for _, publicKeyItem := range publicKey {
		publicKeyRule = append(publicKeyRule, publicKeyItem)
	}

	logs, sub, err := _Mint.contract.FilterLogs(opts, "Deposit", publicKeyRule)
	if err != nil {
		return nil, err
	}
	return &MintDepositIterator{contract: _Mint.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0x5e02e546dec63caecdebd8692b65ccdbaa30e25ac8395aaef548f4ff09b78aa4.
//
// Solidity: event Deposit((uint256,uint256) indexed publicKey, (uint256,uint256) commitment, uint256 timestamp)
func (_Mint *MintFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *MintDeposit, publicKey []PairingG1Point) (event.Subscription, error) {

	var publicKeyRule []interface{}
	for _, publicKeyItem := range publicKey {
		publicKeyRule = append(publicKeyRule, publicKeyItem)
	}

	logs, sub, err := _Mint.contract.WatchLogs(opts, "Deposit", publicKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintDeposit)
				if err := _Mint.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0x5e02e546dec63caecdebd8692b65ccdbaa30e25ac8395aaef548f4ff09b78aa4.
//
// Solidity: event Deposit((uint256,uint256) indexed publicKey, (uint256,uint256) commitment, uint256 timestamp)
func (_Mint *MintFilterer) ParseDeposit(log types.Log) (*MintDeposit, error) {
	event := new(MintDeposit)
	if err := _Mint.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintMintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the Mint contract.
type MintMintedIterator struct {
	Event *MintMinted // Event containing the contract specifics and raw log

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
func (it *MintMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintMinted)
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
		it.Event = new(MintMinted)
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
func (it *MintMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintMinted represents a Minted event raised by the Mint contract.
type MintMinted struct {
	User  common.Address
	Nonce *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0x30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe.
//
// Solidity: event Minted(address indexed user, uint256 nonce)
func (_Mint *MintFilterer) FilterMinted(opts *bind.FilterOpts, user []common.Address) (*MintMintedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Mint.contract.FilterLogs(opts, "Minted", userRule)
	if err != nil {
		return nil, err
	}
	return &MintMintedIterator{contract: _Mint.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0x30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe.
//
// Solidity: event Minted(address indexed user, uint256 nonce)
func (_Mint *MintFilterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *MintMinted, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Mint.contract.WatchLogs(opts, "Minted", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintMinted)
				if err := _Mint.contract.UnpackLog(event, "Minted", log); err != nil {
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

// ParseMinted is a log parse operation binding the contract event 0x30385c845b448a36257a6a1716e6ad2e1bc2cbe333cde1e69fe849ad6511adfe.
//
// Solidity: event Minted(address indexed user, uint256 nonce)
func (_Mint *MintFilterer) ParseMinted(log types.Log) (*MintMinted, error) {
	event := new(MintMinted)
	if err := _Mint.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MintPwnedIterator is returned from FilterPwned and is used to iterate over the raw logs and unpacked data for Pwned events raised by the Mint contract.
type MintPwnedIterator struct {
	Event *MintPwned // Event containing the contract specifics and raw log

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
func (it *MintPwnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MintPwned)
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
		it.Event = new(MintPwned)
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
func (it *MintPwnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MintPwnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MintPwned represents a Pwned event raised by the Mint contract.
type MintPwned struct {
	PrivateKey *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPwned is a free log retrieval operation binding the contract event 0x2fa98b20b7848fedb78850292ddfee7e5fbe5584f443fad59212d7aadeda16e6.
//
// Solidity: event Pwned(uint256 indexed privateKey)
func (_Mint *MintFilterer) FilterPwned(opts *bind.FilterOpts, privateKey []*big.Int) (*MintPwnedIterator, error) {

	var privateKeyRule []interface{}
	for _, privateKeyItem := range privateKey {
		privateKeyRule = append(privateKeyRule, privateKeyItem)
	}

	logs, sub, err := _Mint.contract.FilterLogs(opts, "Pwned", privateKeyRule)
	if err != nil {
		return nil, err
	}
	return &MintPwnedIterator{contract: _Mint.contract, event: "Pwned", logs: logs, sub: sub}, nil
}

// WatchPwned is a free log subscription operation binding the contract event 0x2fa98b20b7848fedb78850292ddfee7e5fbe5584f443fad59212d7aadeda16e6.
//
// Solidity: event Pwned(uint256 indexed privateKey)
func (_Mint *MintFilterer) WatchPwned(opts *bind.WatchOpts, sink chan<- *MintPwned, privateKey []*big.Int) (event.Subscription, error) {

	var privateKeyRule []interface{}
	for _, privateKeyItem := range privateKey {
		privateKeyRule = append(privateKeyRule, privateKeyItem)
	}

	logs, sub, err := _Mint.contract.WatchLogs(opts, "Pwned", privateKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MintPwned)
				if err := _Mint.contract.UnpackLog(event, "Pwned", log); err != nil {
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

// ParsePwned is a log parse operation binding the contract event 0x2fa98b20b7848fedb78850292ddfee7e5fbe5584f443fad59212d7aadeda16e6.
//
// Solidity: event Pwned(uint256 indexed privateKey)
func (_Mint *MintFilterer) ParsePwned(log types.Log) (*MintPwned, error) {
	event := new(MintPwned)
	if err := _Mint.contract.UnpackLog(event, "Pwned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

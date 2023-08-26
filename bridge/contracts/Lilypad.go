// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bridge

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

// LilypadMetaData contains all meta data concerning the Lilypad contract.
var LilypadMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeContractAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"}],\"name\":\"JobResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"wallet\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"}],\"name\":\"ResultForUser\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_jobId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"_errorMsg\",\"type\":\"string\"}],\"name\":\"lilypadCancelled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lilypadFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_jobId\",\"type\":\"uint256\"},{\"internalType\":\"enumLilypadResultType\",\"name\":\"_resultType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"_result\",\"type\":\"string\"}],\"name\":\"lilypadFulfilled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"prompts\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_prompt\",\"type\":\"string\"}],\"name\":\"request\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"results\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b50604051620018b0380380620018b08339818101604052810190620000379190620001c9565b806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634dfb43286040518163ffffffff1660e01b8152600401602060405180830381865afa15801562000128573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906200014e919062000236565b905080600281905550505062000268565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620001918262000164565b9050919050565b620001a38162000184565b8114620001af57600080fd5b50565b600081519050620001c38162000198565b92915050565b600060208284031215620001e257620001e16200015f565b5b6000620001f284828501620001b2565b91505092915050565b6000819050919050565b6200021081620001fb565b81146200021c57600080fd5b50565b600081519050620002308162000205565b92915050565b6000602082840312156200024f576200024e6200015f565b5b60006200025f848285016200021f565b91505092915050565b61163880620002786000396000f3fe6080604052600436106100705760003560e01c8063a3c573eb1161004e578063a3c573eb146100f7578063d4cf9af714610122578063d5040d651461015f578063e1dc7a201461018857610070565b80631b0c27da146100755780632c199889146100b25780633487823e146100ce575b600080fd5b34801561008157600080fd5b5061009c6004803603810190610097919061065c565b6101b3565b6040516100a99190610719565b60405180910390f35b6100cc60048036038101906100c791906107a0565b610253565b005b3480156100da57600080fd5b506100f560048036038101906100f0919061084b565b61041f565b005b34801561010357600080fd5b5061010c610481565b60405161011991906108ce565b60405180910390f35b34801561012e57600080fd5b506101496004803603810190610144919061065c565b6104a5565b6040516101569190610719565b60405180910390f35b34801561016b57600080fd5b506101866004803603810190610181919061090e565b610545565b005b34801561019457600080fd5b5061019d610616565b6040516101aa91906109a5565b60405180910390f35b600460205280600052604060002060009150905080546101d2906109ef565b80601f01602080910402602001604051908101604052809291908181526020018280546101fe906109ef565b801561024b5780601f106102205761010080835404028352916020019161024b565b820191906000526020600020905b81548152906001019060200180831161022e57829003601f168201915b505050505081565b600254341015610298576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161028f90610a6c565b60405180910390fd5b600082826040516020016102ad929190611049565b60405160208183030381529060405290506000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663321bbe4960025430856001600381111561031757610316611107565b5b6040518563ffffffff1660e01b815260040161033593929190611152565b60206040518083038185885af1158015610353573d6000803e3d6000fd5b50505050506040513d601f19601f8201168201806040525081019061037891906111a5565b9050600081116103bd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103b49061121e565b60405180910390fd5b83836003600084815260200190815260200160002091826103df929190611424565b507f74b7f39a58489c543add203a10f5e3d610deea59750170794575329e4ab5b7eb8484604051610411929190611521565b60405180910390a150505050565b8181600460008681526020019081526020016000209182610441929190611424565b507f74b7f39a58489c543add203a10f5e3d610deea59750170794575329e4ab5b7eb8282604051610473929190611521565b60405180910390a150505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600360205280600052604060002060009150905080546104c4906109ef565b80601f01602080910402602001604051908101604052809291908181526020018280546104f0906109ef565b801561053d5780601f106105125761010080835404028352916020019161053d565b820191906000526020600020905b81548152906001019060200180831161052057829003601f168201915b505050505081565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff161461059f57600080fd5b81816004600087815260200190815260200160002091826105c1929190611424565b507fb3ab19bd795fbabd419dbf34dda545d0dca86e54c45156d3bc156b184733cd88600360008681526020019081526020016000208383604051610607939291906115c9565b60405180910390a15050505050565b60025481565b600080fd5b600080fd5b6000819050919050565b61063981610626565b811461064457600080fd5b50565b60008135905061065681610630565b92915050565b6000602082840312156106725761067161061c565b5b600061068084828501610647565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156106c35780820151818401526020810190506106a8565b60008484015250505050565b6000601f19601f8301169050919050565b60006106eb82610689565b6106f58185610694565b93506107058185602086016106a5565b61070e816106cf565b840191505092915050565b6000602082019050818103600083015261073381846106e0565b905092915050565b600080fd5b600080fd5b600080fd5b60008083601f8401126107605761075f61073b565b5b8235905067ffffffffffffffff81111561077d5761077c610740565b5b60208301915083600182028301111561079957610798610745565b5b9250929050565b600080602083850312156107b7576107b661061c565b5b600083013567ffffffffffffffff8111156107d5576107d4610621565b5b6107e18582860161074a565b92509250509250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610818826107ed565b9050919050565b6108288161080d565b811461083357600080fd5b50565b6000813590506108458161081f565b92915050565b600080600080606085870312156108655761086461061c565b5b600061087387828801610836565b945050602061088487828801610647565b935050604085013567ffffffffffffffff8111156108a5576108a4610621565b5b6108b18782880161074a565b925092505092959194509250565b6108c88161080d565b82525050565b60006020820190506108e360008301846108bf565b92915050565b600481106108f657600080fd5b50565b600081359050610908816108e9565b92915050565b60008060008060006080868803121561092a5761092961061c565b5b600061093888828901610836565b955050602061094988828901610647565b945050604061095a888289016108f9565b935050606086013567ffffffffffffffff81111561097b5761097a610621565b5b6109878882890161074a565b92509250509295509295909350565b61099f81610626565b82525050565b60006020820190506109ba6000830184610996565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610a0757607f821691505b602082108103610a1a57610a196109c0565b5b50919050565b7f4e6f7420656e6f75676820746f2072756e204c696c79706164206a6f62000000600082015250565b6000610a56601d83610694565b9150610a6182610a20565b602082019050919050565b60006020820190508181036000830152610a8581610a49565b9050919050565b600081905092915050565b7f7b00000000000000000000000000000000000000000000000000000000000000600082015250565b6000610acd600183610a8c565b9150610ad882610a97565b600182019050919050565b7f22456e67696e65223a2022446f636b6572222c00000000000000000000000000600082015250565b6000610b19601383610a8c565b9150610b2482610ae3565b601382019050919050565b7f225665726966696572223a20224e6f6f70222c00000000000000000000000000600082015250565b6000610b65601383610a8c565b9150610b7082610b2f565b601382019050919050565b7f225075626c6973686572223a202245737475617279222c000000000000000000600082015250565b6000610bb1601783610a8c565b9150610bbc82610b7b565b601782019050919050565b7f225075626c697368657253706563223a207b2254797065223a2022457374756160008201527f7279227d2c000000000000000000000000000000000000000000000000000000602082015250565b6000610c23602583610a8c565b9150610c2e82610bc7565b602582019050919050565b7f22446f636b6572223a207b22496d616765223a2022646f67756b616e67756e2f60008201527f646f6e6174696f6e666f72756e69766572736974793a676f5f6c61746573743660208201527f222c2022456e7669726f6e6d656e745661726961626c6573223a205b2257414c60408201527f4c45545f414444524553533d0000000000000000000000000000000000000000606082015250565b6000610ce1606c83610a8c565b9150610cec82610c39565b606c82019050919050565b82818337600083830152505050565b6000610d128385610a8c565b9350610d1f838584610cf7565b82840190509392505050565b7f225d7d2c00000000000000000000000000000000000000000000000000000000600082015250565b6000610d61600483610a8c565b9150610d6c82610d2b565b600482019050919050565b7f224c616e6775616765223a207b224a6f62436f6e74657874223a207b7d7d2c00600082015250565b6000610dad601f83610a8c565b9150610db882610d77565b601f82019050919050565b7f225761736d223a207b22456e7472794d6f64756c65223a207b7d7d2c00000000600082015250565b6000610df9601c83610a8c565b9150610e0482610dc3565b601c82019050919050565b7f225265736f7572636573223a207b22475055223a2022227d2c00000000000000600082015250565b6000610e45601983610a8c565b9150610e5082610e0f565b601982019050919050565b7f224e6574776f726b223a207b2254797065223a202248545450222c2022446f6d60008201527f61696e73223a205b226574682e7075626c69632d7270632e636f6d225d7d2c00602082015250565b6000610eb7603f83610a8c565b9150610ec282610e5b565b603f82019050919050565b7f2254696d656f7574223a20313830302c00000000000000000000000000000000600082015250565b6000610f03601083610a8c565b9150610f0e82610ecd565b601082019050919050565b7f224f757470757473223a205b7b2253746f72616765536f75726365223a20224960008201527f504653222c20224e616d65223a20226f757470757473222c202250617468223a60208201527f20222f6f757470757473227d5d2c000000000000000000000000000000000000604082015250565b6000610f9b604e83610a8c565b9150610fa682610f19565b604e82019050919050565b7f224465616c223a207b22436f6e63757272656e6379223a20317d000000000000600082015250565b6000610fe7601a83610a8c565b9150610ff282610fb1565b601a82019050919050565b7f7d00000000000000000000000000000000000000000000000000000000000000600082015250565b6000611033600183610a8c565b915061103e82610ffd565b600182019050919050565b600061105482610ac0565b915061105f82610b0c565b915061106a82610b58565b915061107582610ba4565b915061108082610c16565b915061108b82610cd4565b9150611098828486610d06565b91506110a382610d54565b91506110ae82610da0565b91506110b982610dec565b91506110c482610e38565b91506110cf82610eaa565b91506110da82610ef6565b91506110e582610f8e565b91506110f082610fda565b91506110fb82611026565b91508190509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600060ff82169050919050565b61114c81611136565b82525050565b600060608201905061116760008301866108bf565b818103602083015261117981856106e0565b90506111886040830184611143565b949350505050565b60008151905061119f81610630565b92915050565b6000602082840312156111bb576111ba61061c565b5b60006111c984828501611190565b91505092915050565b7f6a6f62206469646e27742072657475726e20612076616c756500000000000000600082015250565b6000611208601983610694565b9150611213826111d2565b602082019050919050565b60006020820190508181036000830152611237816111fb565b9050919050565b600082905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026112da7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8261129d565b6112e4868361129d565b95508019841693508086168417925050509392505050565b6000819050919050565b600061132161131c61131784610626565b6112fc565b610626565b9050919050565b6000819050919050565b61133b83611306565b61134f61134782611328565b8484546112aa565b825550505050565b600090565b611364611357565b61136f818484611332565b505050565b5b818110156113935761138860008261135c565b600181019050611375565b5050565b601f8211156113d8576113a981611278565b6113b28461128d565b810160208510156113c1578190505b6113d56113cd8561128d565b830182611374565b50505b505050565b600082821c905092915050565b60006113fb600019846008026113dd565b1980831691505092915050565b600061141483836113ea565b9150826002028217905092915050565b61142e838361123e565b67ffffffffffffffff81111561144757611446611249565b5b61145182546109ef565b61145c828285611397565b6000601f83116001811461148b5760008415611479578287013590505b6114838582611408565b8655506114eb565b601f19841661149986611278565b60005b828110156114c15784890135825560018201915060208501945060208101905061149c565b868310156114de57848901356114da601f8916826113ea565b8355505b6001600288020188555050505b50505050505050565b60006115008385610694565b935061150d838584610cf7565b611516836106cf565b840190509392505050565b6000602082019050818103600083015261153c8184866114f4565b90509392505050565b60008154611552816109ef565b61155c8186610694565b94506001821660008114611577576001811461158d576115c0565b60ff1983168652811515602002860193506115c0565b61159685611278565b60005b838110156115b857815481890152600182019150602081019050611599565b808801955050505b50505092915050565b600060408201905081810360008301526115e38186611545565b905081810360208301526115f88184866114f4565b905094935050505056fea264697066735822122061f4b59246b87981c33da8a46b714e4bb93bebcc1d86fd89ac1a701705a1f98864736f6c63430008130033",
}

// LilypadABI is the input ABI used to generate the binding from.
// Deprecated: Use LilypadMetaData.ABI instead.
var LilypadABI = LilypadMetaData.ABI

// LilypadBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use LilypadMetaData.Bin instead.
var LilypadBin = LilypadMetaData.Bin

// DeployLilypad deploys a new Ethereum contract, binding an instance of Lilypad to it.
func DeployLilypad(auth *bind.TransactOpts, backend bind.ContractBackend, _bridgeContractAddress common.Address) (common.Address, *types.Transaction, *Lilypad, error) {
	parsed, err := LilypadMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(LilypadBin), backend, _bridgeContractAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Lilypad{LilypadCaller: LilypadCaller{contract: contract}, LilypadTransactor: LilypadTransactor{contract: contract}, LilypadFilterer: LilypadFilterer{contract: contract}}, nil
}

// Lilypad is an auto generated Go binding around an Ethereum contract.
type Lilypad struct {
	LilypadCaller     // Read-only binding to the contract
	LilypadTransactor // Write-only binding to the contract
	LilypadFilterer   // Log filterer for contract events
}

// LilypadCaller is an auto generated read-only Go binding around an Ethereum contract.
type LilypadCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LilypadTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LilypadFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LilypadSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LilypadSession struct {
	Contract     *Lilypad          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LilypadCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LilypadCallerSession struct {
	Contract *LilypadCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// LilypadTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LilypadTransactorSession struct {
	Contract     *LilypadTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// LilypadRaw is an auto generated low-level Go binding around an Ethereum contract.
type LilypadRaw struct {
	Contract *Lilypad // Generic contract binding to access the raw methods on
}

// LilypadCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LilypadCallerRaw struct {
	Contract *LilypadCaller // Generic read-only contract binding to access the raw methods on
}

// LilypadTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LilypadTransactorRaw struct {
	Contract *LilypadTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLilypad creates a new instance of Lilypad, bound to a specific deployed contract.
func NewLilypad(address common.Address, backend bind.ContractBackend) (*Lilypad, error) {
	contract, err := bindLilypad(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Lilypad{LilypadCaller: LilypadCaller{contract: contract}, LilypadTransactor: LilypadTransactor{contract: contract}, LilypadFilterer: LilypadFilterer{contract: contract}}, nil
}

// NewLilypadCaller creates a new read-only instance of Lilypad, bound to a specific deployed contract.
func NewLilypadCaller(address common.Address, caller bind.ContractCaller) (*LilypadCaller, error) {
	contract, err := bindLilypad(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LilypadCaller{contract: contract}, nil
}

// NewLilypadTransactor creates a new write-only instance of Lilypad, bound to a specific deployed contract.
func NewLilypadTransactor(address common.Address, transactor bind.ContractTransactor) (*LilypadTransactor, error) {
	contract, err := bindLilypad(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LilypadTransactor{contract: contract}, nil
}

// NewLilypadFilterer creates a new log filterer instance of Lilypad, bound to a specific deployed contract.
func NewLilypadFilterer(address common.Address, filterer bind.ContractFilterer) (*LilypadFilterer, error) {
	contract, err := bindLilypad(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LilypadFilterer{contract: contract}, nil
}

// bindLilypad binds a generic wrapper to an already deployed contract.
func bindLilypad(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := LilypadMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lilypad *LilypadRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lilypad.Contract.LilypadCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lilypad *LilypadRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lilypad.Contract.LilypadTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lilypad *LilypadRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lilypad.Contract.LilypadTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Lilypad *LilypadCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Lilypad.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Lilypad *LilypadTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Lilypad.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Lilypad *LilypadTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Lilypad.Contract.contract.Transact(opts, method, params...)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Lilypad *LilypadCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Lilypad.contract.Call(opts, &out, "bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Lilypad *LilypadSession) BridgeAddress() (common.Address, error) {
	return _Lilypad.Contract.BridgeAddress(&_Lilypad.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0xa3c573eb.
//
// Solidity: function bridgeAddress() view returns(address)
func (_Lilypad *LilypadCallerSession) BridgeAddress() (common.Address, error) {
	return _Lilypad.Contract.BridgeAddress(&_Lilypad.CallOpts)
}

// LilypadFee is a free data retrieval call binding the contract method 0xe1dc7a20.
//
// Solidity: function lilypadFee() view returns(uint256)
func (_Lilypad *LilypadCaller) LilypadFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Lilypad.contract.Call(opts, &out, "lilypadFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LilypadFee is a free data retrieval call binding the contract method 0xe1dc7a20.
//
// Solidity: function lilypadFee() view returns(uint256)
func (_Lilypad *LilypadSession) LilypadFee() (*big.Int, error) {
	return _Lilypad.Contract.LilypadFee(&_Lilypad.CallOpts)
}

// LilypadFee is a free data retrieval call binding the contract method 0xe1dc7a20.
//
// Solidity: function lilypadFee() view returns(uint256)
func (_Lilypad *LilypadCallerSession) LilypadFee() (*big.Int, error) {
	return _Lilypad.Contract.LilypadFee(&_Lilypad.CallOpts)
}

// Prompts is a free data retrieval call binding the contract method 0xd4cf9af7.
//
// Solidity: function prompts(uint256 ) view returns(string)
func (_Lilypad *LilypadCaller) Prompts(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Lilypad.contract.Call(opts, &out, "prompts", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Prompts is a free data retrieval call binding the contract method 0xd4cf9af7.
//
// Solidity: function prompts(uint256 ) view returns(string)
func (_Lilypad *LilypadSession) Prompts(arg0 *big.Int) (string, error) {
	return _Lilypad.Contract.Prompts(&_Lilypad.CallOpts, arg0)
}

// Prompts is a free data retrieval call binding the contract method 0xd4cf9af7.
//
// Solidity: function prompts(uint256 ) view returns(string)
func (_Lilypad *LilypadCallerSession) Prompts(arg0 *big.Int) (string, error) {
	return _Lilypad.Contract.Prompts(&_Lilypad.CallOpts, arg0)
}

// Results is a free data retrieval call binding the contract method 0x1b0c27da.
//
// Solidity: function results(uint256 ) view returns(string)
func (_Lilypad *LilypadCaller) Results(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Lilypad.contract.Call(opts, &out, "results", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Results is a free data retrieval call binding the contract method 0x1b0c27da.
//
// Solidity: function results(uint256 ) view returns(string)
func (_Lilypad *LilypadSession) Results(arg0 *big.Int) (string, error) {
	return _Lilypad.Contract.Results(&_Lilypad.CallOpts, arg0)
}

// Results is a free data retrieval call binding the contract method 0x1b0c27da.
//
// Solidity: function results(uint256 ) view returns(string)
func (_Lilypad *LilypadCallerSession) Results(arg0 *big.Int) (string, error) {
	return _Lilypad.Contract.Results(&_Lilypad.CallOpts, arg0)
}

// LilypadCancelled is a paid mutator transaction binding the contract method 0x3487823e.
//
// Solidity: function lilypadCancelled(address _from, uint256 _jobId, string _errorMsg) returns()
func (_Lilypad *LilypadTransactor) LilypadCancelled(opts *bind.TransactOpts, _from common.Address, _jobId *big.Int, _errorMsg string) (*types.Transaction, error) {
	return _Lilypad.contract.Transact(opts, "lilypadCancelled", _from, _jobId, _errorMsg)
}

// LilypadCancelled is a paid mutator transaction binding the contract method 0x3487823e.
//
// Solidity: function lilypadCancelled(address _from, uint256 _jobId, string _errorMsg) returns()
func (_Lilypad *LilypadSession) LilypadCancelled(_from common.Address, _jobId *big.Int, _errorMsg string) (*types.Transaction, error) {
	return _Lilypad.Contract.LilypadCancelled(&_Lilypad.TransactOpts, _from, _jobId, _errorMsg)
}

// LilypadCancelled is a paid mutator transaction binding the contract method 0x3487823e.
//
// Solidity: function lilypadCancelled(address _from, uint256 _jobId, string _errorMsg) returns()
func (_Lilypad *LilypadTransactorSession) LilypadCancelled(_from common.Address, _jobId *big.Int, _errorMsg string) (*types.Transaction, error) {
	return _Lilypad.Contract.LilypadCancelled(&_Lilypad.TransactOpts, _from, _jobId, _errorMsg)
}

// LilypadFulfilled is a paid mutator transaction binding the contract method 0xd5040d65.
//
// Solidity: function lilypadFulfilled(address _from, uint256 _jobId, uint8 _resultType, string _result) returns()
func (_Lilypad *LilypadTransactor) LilypadFulfilled(opts *bind.TransactOpts, _from common.Address, _jobId *big.Int, _resultType uint8, _result string) (*types.Transaction, error) {
	return _Lilypad.contract.Transact(opts, "lilypadFulfilled", _from, _jobId, _resultType, _result)
}

// LilypadFulfilled is a paid mutator transaction binding the contract method 0xd5040d65.
//
// Solidity: function lilypadFulfilled(address _from, uint256 _jobId, uint8 _resultType, string _result) returns()
func (_Lilypad *LilypadSession) LilypadFulfilled(_from common.Address, _jobId *big.Int, _resultType uint8, _result string) (*types.Transaction, error) {
	return _Lilypad.Contract.LilypadFulfilled(&_Lilypad.TransactOpts, _from, _jobId, _resultType, _result)
}

// LilypadFulfilled is a paid mutator transaction binding the contract method 0xd5040d65.
//
// Solidity: function lilypadFulfilled(address _from, uint256 _jobId, uint8 _resultType, string _result) returns()
func (_Lilypad *LilypadTransactorSession) LilypadFulfilled(_from common.Address, _jobId *big.Int, _resultType uint8, _result string) (*types.Transaction, error) {
	return _Lilypad.Contract.LilypadFulfilled(&_Lilypad.TransactOpts, _from, _jobId, _resultType, _result)
}

// Request is a paid mutator transaction binding the contract method 0x2c199889.
//
// Solidity: function request(string _prompt) payable returns()
func (_Lilypad *LilypadTransactor) Request(opts *bind.TransactOpts, _prompt string) (*types.Transaction, error) {
	return _Lilypad.contract.Transact(opts, "request", _prompt)
}

// Request is a paid mutator transaction binding the contract method 0x2c199889.
//
// Solidity: function request(string _prompt) payable returns()
func (_Lilypad *LilypadSession) Request(_prompt string) (*types.Transaction, error) {
	return _Lilypad.Contract.Request(&_Lilypad.TransactOpts, _prompt)
}

// Request is a paid mutator transaction binding the contract method 0x2c199889.
//
// Solidity: function request(string _prompt) payable returns()
func (_Lilypad *LilypadTransactorSession) Request(_prompt string) (*types.Transaction, error) {
	return _Lilypad.Contract.Request(&_Lilypad.TransactOpts, _prompt)
}

// LilypadJobResultIterator is returned from FilterJobResult and is used to iterate over the raw logs and unpacked data for JobResult events raised by the Lilypad contract.
type LilypadJobResultIterator struct {
	Event *LilypadJobResult // Event containing the contract specifics and raw log

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
func (it *LilypadJobResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadJobResult)
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
		it.Event = new(LilypadJobResult)
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
func (it *LilypadJobResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadJobResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadJobResult represents a JobResult event raised by the Lilypad contract.
type LilypadJobResult struct {
	Result string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterJobResult is a free log retrieval operation binding the contract event 0x74b7f39a58489c543add203a10f5e3d610deea59750170794575329e4ab5b7eb.
//
// Solidity: event JobResult(string result)
func (_Lilypad *LilypadFilterer) FilterJobResult(opts *bind.FilterOpts) (*LilypadJobResultIterator, error) {

	logs, sub, err := _Lilypad.contract.FilterLogs(opts, "JobResult")
	if err != nil {
		return nil, err
	}
	return &LilypadJobResultIterator{contract: _Lilypad.contract, event: "JobResult", logs: logs, sub: sub}, nil
}

// WatchJobResult is a free log subscription operation binding the contract event 0x74b7f39a58489c543add203a10f5e3d610deea59750170794575329e4ab5b7eb.
//
// Solidity: event JobResult(string result)
func (_Lilypad *LilypadFilterer) WatchJobResult(opts *bind.WatchOpts, sink chan<- *LilypadJobResult) (event.Subscription, error) {

	logs, sub, err := _Lilypad.contract.WatchLogs(opts, "JobResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadJobResult)
				if err := _Lilypad.contract.UnpackLog(event, "JobResult", log); err != nil {
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

// ParseJobResult is a log parse operation binding the contract event 0x74b7f39a58489c543add203a10f5e3d610deea59750170794575329e4ab5b7eb.
//
// Solidity: event JobResult(string result)
func (_Lilypad *LilypadFilterer) ParseJobResult(log types.Log) (*LilypadJobResult, error) {
	event := new(LilypadJobResult)
	if err := _Lilypad.contract.UnpackLog(event, "JobResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// LilypadResultForUserIterator is returned from FilterResultForUser and is used to iterate over the raw logs and unpacked data for ResultForUser events raised by the Lilypad contract.
type LilypadResultForUserIterator struct {
	Event *LilypadResultForUser // Event containing the contract specifics and raw log

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
func (it *LilypadResultForUserIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LilypadResultForUser)
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
		it.Event = new(LilypadResultForUser)
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
func (it *LilypadResultForUserIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LilypadResultForUserIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LilypadResultForUser represents a ResultForUser event raised by the Lilypad contract.
type LilypadResultForUser struct {
	Wallet string
	Result string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterResultForUser is a free log retrieval operation binding the contract event 0xb3ab19bd795fbabd419dbf34dda545d0dca86e54c45156d3bc156b184733cd88.
//
// Solidity: event ResultForUser(string wallet, string result)
func (_Lilypad *LilypadFilterer) FilterResultForUser(opts *bind.FilterOpts) (*LilypadResultForUserIterator, error) {

	logs, sub, err := _Lilypad.contract.FilterLogs(opts, "ResultForUser")
	if err != nil {
		return nil, err
	}
	return &LilypadResultForUserIterator{contract: _Lilypad.contract, event: "ResultForUser", logs: logs, sub: sub}, nil
}

// WatchResultForUser is a free log subscription operation binding the contract event 0xb3ab19bd795fbabd419dbf34dda545d0dca86e54c45156d3bc156b184733cd88.
//
// Solidity: event ResultForUser(string wallet, string result)
func (_Lilypad *LilypadFilterer) WatchResultForUser(opts *bind.WatchOpts, sink chan<- *LilypadResultForUser) (event.Subscription, error) {

	logs, sub, err := _Lilypad.contract.WatchLogs(opts, "ResultForUser")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LilypadResultForUser)
				if err := _Lilypad.contract.UnpackLog(event, "ResultForUser", log); err != nil {
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

// ParseResultForUser is a log parse operation binding the contract event 0xb3ab19bd795fbabd419dbf34dda545d0dca86e54c45156d3bc156b184733cd88.
//
// Solidity: event ResultForUser(string wallet, string result)
func (_Lilypad *LilypadFilterer) ParseResultForUser(log types.Log) (*LilypadResultForUser, error) {
	event := new(LilypadResultForUser)
	if err := _Lilypad.contract.UnpackLog(event, "ResultForUser", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

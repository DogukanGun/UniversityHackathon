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

// DonationManagerDonationInfo is an auto generated low-level Go binding around an user-defined struct.
type DonationManagerDonationInfo struct {
	ID              *big.Int
	Title           string
	Description     string
	Owner           common.Address
	ContractAddress common.Address
	CoverImageCID   string
	IsApproved      bool
}

// BridgeMetaData contains all meta data concerning the Bridge contract.
var BridgeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"userWalletAddress\",\"type\":\"address\"}],\"name\":\"ApplyForDonation\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"applyForDonation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approvedAddress\",\"type\":\"address\"}],\"name\":\"approveDonation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"donationFactoryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllDonations\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"coverImageCID\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"}],\"internalType\":\"structDonationManager.DonationInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"donationID\",\"type\":\"uint256\"}],\"name\":\"getDonation\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"coverImageCID\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"}],\"internalType\":\"structDonationManager.DonationInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"destionationWallet\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"coverImageCID\",\"type\":\"string\"}],\"name\":\"startDonation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200234938038062002349833981810160405281019062000037919062000163565b6040516200004590620000eb565b604051809103906000f08015801562000062573d6000803e3d6000fd5b50600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505062000195565b6109fb806200194e83390190565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006200012b82620000fe565b9050919050565b6200013d816200011e565b81146200014957600080fd5b50565b6000815190506200015d8162000132565b92915050565b6000602082840312156200017c576200017b620000f9565b5b60006200018c848285016200014c565b91505092915050565b6117a980620001a56000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c806332f3d5f91461006757806338a59a07146100835780635501e348146100a15780638114b497146100bf5780639eb7477c146100c9578063ef07a81f146100e5575b600080fd5b610081600480360381019061007c9190610d72565b610115565b005b61008b610554565b604051610098919061108d565b60405180910390f35b6100a961083a565b6040516100b691906110be565b60405180910390f35b6100c7610860565b005b6100e360048036038101906100de91906110d9565b610899565b005b6100ff60048036038101906100fa9190611132565b610977565b60405161010c9190611209565b60405180910390f35b600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146101a5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161019c906112ae565b60405180910390fd5b600088846040516024016101ba9291906112ce565b6040516020818303038152906040527fa6f09d53000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050509050600080600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1683604051610283919061133e565b6000604051808303816000865af19150503d80600081146102c0576040519150601f19603f3d011682016040523d82523d6000602084013e6102c5565b606091505b5091509150816102d457600080fd5b6000818060200190518101906102ea9190611393565b905060006040518060e0016040528060008054905081526020018d8d8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505081526020018b8b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505081526020018973ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff16815260200188888080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050815260200160001515815250908060018154018082558091505060019003906000526020600020906007020160009091909190915060008201518160000155602082015181600101908161046991906115fb565b50604082015181600201908161047f91906115fb565b5060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a082015181600501908161052391906115fb565b5060c08201518160060160006101000a81548160ff0219169083151502179055505050505050505050505050505050565b60606000805480602002602001604051908101604052809291908181526020016000905b8282101561083157838290600052602060002090600702016040518060e0016040529081600082015481526020016001820180546105b59061141e565b80601f01602080910402602001604051908101604052809291908181526020018280546105e19061141e565b801561062e5780601f106106035761010080835404028352916020019161062e565b820191906000526020600020905b81548152906001019060200180831161061157829003601f168201915b505050505081526020016002820180546106479061141e565b80601f01602080910402602001604051908101604052809291908181526020018280546106739061141e565b80156106c05780601f10610695576101008083540402835291602001916106c0565b820191906000526020600020905b8154815290600101906020018083116106a357829003601f168201915b505050505081526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016005820180546107859061141e565b80601f01602080910402602001604051908101604052809291908181526020018280546107b19061141e565b80156107fe5780601f106107d3576101008083540402835291602001916107fe565b820191906000526020600020905b8154815290600101906020018083116107e157829003601f168201915b505050505081526020016006820160009054906101000a900460ff16151515158152505081526020019060010190610578565b50505050905090565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b7fff25dd54611daf0db1901cc5bad1d6e38fdc0f5d6cb1e18cf0342e9e2763d78d3360405161088f91906110be565b60405180910390a1565b60005b600080549050811015610973578173ffffffffffffffffffffffffffffffffffffffff16600082815481106108d4576108d36116cd565b5b906000526020600020906007020160030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff160361096057600160008281548110610937576109366116cd565b5b906000526020600020906007020160060160006101000a81548160ff0219169083151502179055505b808061096b9061172b565b91505061089c565b5050565b61097f610c3a565b60008281548110610993576109926116cd565b5b90600052602060002090600702016040518060e0016040529081600082015481526020016001820180546109c69061141e565b80601f01602080910402602001604051908101604052809291908181526020018280546109f29061141e565b8015610a3f5780601f10610a1457610100808354040283529160200191610a3f565b820191906000526020600020905b815481529060010190602001808311610a2257829003601f168201915b50505050508152602001600282018054610a589061141e565b80601f0160208091040260200160405190810160405280929190818152602001828054610a849061141e565b8015610ad15780601f10610aa657610100808354040283529160200191610ad1565b820191906000526020600020905b815481529060010190602001808311610ab457829003601f168201915b505050505081526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600582018054610b969061141e565b80601f0160208091040260200160405190810160405280929190818152602001828054610bc29061141e565b8015610c0f5780601f10610be457610100808354040283529160200191610c0f565b820191906000526020600020905b815481529060010190602001808311610bf257829003601f168201915b505050505081526020016006820160009054906101000a900460ff1615151515815250509050919050565b6040518060e00160405280600081526020016060815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001606081526020016000151581525090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610cda82610caf565b9050919050565b610cea81610ccf565b8114610cf557600080fd5b50565b600081359050610d0781610ce1565b92915050565b600080fd5b600080fd5b600080fd5b60008083601f840112610d3257610d31610d0d565b5b8235905067ffffffffffffffff811115610d4f57610d4e610d12565b5b602083019150836001820283011115610d6b57610d6a610d17565b5b9250929050565b60008060008060008060008060a0898b031215610d9257610d91610ca5565b5b6000610da08b828c01610cf8565b985050602089013567ffffffffffffffff811115610dc157610dc0610caa565b5b610dcd8b828c01610d1c565b9750975050604089013567ffffffffffffffff811115610df057610def610caa565b5b610dfc8b828c01610d1c565b95509550506060610e0f8b828c01610cf8565b935050608089013567ffffffffffffffff811115610e3057610e2f610caa565b5b610e3c8b828c01610d1c565b92509250509295985092959890939650565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6000819050919050565b610e8d81610e7a565b82525050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610ecd578082015181840152602081019050610eb2565b60008484015250505050565b6000601f19601f8301169050919050565b6000610ef582610e93565b610eff8185610e9e565b9350610f0f818560208601610eaf565b610f1881610ed9565b840191505092915050565b610f2c81610ccf565b82525050565b60008115159050919050565b610f4781610f32565b82525050565b600060e083016000830151610f656000860182610e84565b5060208301518482036020860152610f7d8282610eea565b91505060408301518482036040860152610f978282610eea565b9150506060830151610fac6060860182610f23565b506080830151610fbf6080860182610f23565b5060a083015184820360a0860152610fd78282610eea565b91505060c0830151610fec60c0860182610f3e565b508091505092915050565b60006110038383610f4d565b905092915050565b6000602082019050919050565b600061102382610e4e565b61102d8185610e59565b93508360208202850161103f85610e6a565b8060005b8581101561107b578484038952815161105c8582610ff7565b94506110678361100b565b925060208a01995050600181019050611043565b50829750879550505050505092915050565b600060208201905081810360008301526110a78184611018565b905092915050565b6110b881610ccf565b82525050565b60006020820190506110d360008301846110af565b92915050565b6000602082840312156110ef576110ee610ca5565b5b60006110fd84828501610cf8565b91505092915050565b61110f81610e7a565b811461111a57600080fd5b50565b60008135905061112c81611106565b92915050565b60006020828403121561114857611147610ca5565b5b60006111568482850161111d565b91505092915050565b600060e0830160008301516111776000860182610e84565b506020830151848203602086015261118f8282610eea565b915050604083015184820360408601526111a98282610eea565b91505060608301516111be6060860182610f23565b5060808301516111d16080860182610f23565b5060a083015184820360a08601526111e98282610eea565b91505060c08301516111fe60c0860182610f3e565b508091505092915050565b60006020820190508181036000830152611223818461115f565b905092915050565b600082825260208201905092915050565b7f4f6e6c79206272696467652063616e2072756e20746869732066756e6374696f60008201527f6e00000000000000000000000000000000000000000000000000000000000000602082015250565b600061129860218361122b565b91506112a38261123c565b604082019050919050565b600060208201905081810360008301526112c78161128b565b9050919050565b60006040820190506112e360008301856110af565b6112f060208301846110af565b9392505050565b600081519050919050565b600081905092915050565b6000611318826112f7565b6113228185611302565b9350611332818560208601610eaf565b80840191505092915050565b600061134a828461130d565b915081905092915050565b600061136082610caf565b9050919050565b61137081611355565b811461137b57600080fd5b50565b60008151905061138d81611367565b92915050565b6000602082840312156113a9576113a8610ca5565b5b60006113b78482850161137e565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000600282049050600182168061143657607f821691505b602082108103611449576114486113ef565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026114b17fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82611474565b6114bb8683611474565b95508019841693508086168417925050509392505050565b6000819050919050565b60006114f86114f36114ee84610e7a565b6114d3565b610e7a565b9050919050565b6000819050919050565b611512836114dd565b61152661151e826114ff565b848454611481565b825550505050565b600090565b61153b61152e565b611546818484611509565b505050565b5b8181101561156a5761155f600082611533565b60018101905061154c565b5050565b601f8211156115af576115808161144f565b61158984611464565b81016020851015611598578190505b6115ac6115a485611464565b83018261154b565b50505b505050565b600082821c905092915050565b60006115d2600019846008026115b4565b1980831691505092915050565b60006115eb83836115c1565b9150826002028217905092915050565b61160482610e93565b67ffffffffffffffff81111561161d5761161c6113c0565b5b611627825461141e565b61163282828561156e565b600060209050601f8311600181146116655760008415611653578287015190505b61165d85826115df565b8655506116c5565b601f1984166116738661144f565b60005b8281101561169b57848901518255600182019150602085019450602081019050611676565b868310156116b857848901516116b4601f8916826115c1565b8355505b6001600288020188555050505b505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061173682610e7a565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203611768576117676116fc565b5b60018201905091905056fea26469706673582212205045e961d17ec9aff46b365d61f4bde3d9e60426b259210190760cee0162e91364736f6c63430008130033608060405234801561001057600080fd5b5033600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555061099a806100616000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c80638c010a9e14610030575b600080fd5b61004a6004803603810190610045919061022f565b61004c565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146100dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100d3906102df565b60405180910390fd5b600081600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660405161010e906101bf565b61011992919061030e565b604051809103906000f080158015610135573d6000803e3d6000fd5b5090506000819050816000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050565b61062d8061033883390190565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006101fc826101d1565b9050919050565b61020c816101f1565b811461021757600080fd5b50565b60008135905061022981610203565b92915050565b600060208284031215610245576102446101cc565b5b60006102538482850161021a565b91505092915050565b600082825260208201905092915050565b7f4f6e6c792061646d696e2063616e2063616c6c20746869732066756e6374696f60008201527f6e00000000000000000000000000000000000000000000000000000000000000602082015250565b60006102c960218361025c565b91506102d48261026d565b604082019050919050565b600060208201905081810360008301526102f8816102bc565b9050919050565b610308816101f1565b82525050565b600060408201905061032360008301856102ff565b61033060208301846102ff565b939250505056fe608060405234801561001057600080fd5b5060405161062d38038061062d8339818101604052810190610032919061011d565b816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505061015d565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006100ea826100bf565b9050919050565b6100fa816100df565b811461010557600080fd5b50565b600081519050610117816100f1565b92915050565b60008060408385031215610134576101336100ba565b5b600061014285828601610108565b925050602061015385828601610108565b9150509250929050565b6104c18061016c6000396000f3fe6080604052600436106100345760003560e01c806398e1b41014610039578063b69ef8a814610043578063ed88c68e1461006e575b600080fd5b610041610078565b005b34801561004f57600080fd5b506100586101df565b60405161006591906102af565b60405180910390f35b6100766101e7565b005b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610108576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100ff9061034d565b60405180910390fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1661014a6101df565b6040516101569061039e565b60006040518083038185875af1925050503d8060008114610193576040519150601f19603f3d011682016040523d82523d6000602084013e610198565b606091505b50509050806101dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101d3906103ff565b60405180910390fd5b50565b600047905090565b60003073ffffffffffffffffffffffffffffffffffffffff163460405161020d9061039e565b60006040518083038185875af1925050503d806000811461024a576040519150601f19603f3d011682016040523d82523d6000602084013e61024f565b606091505b5050905080610293576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161028a9061046b565b60405180910390fd5b50565b6000819050919050565b6102a981610296565b82525050565b60006020820190506102c460008301846102a0565b92915050565b600082825260208201905092915050565b7f4f6e6c792061646d696e2063616e2063616c6c20746869732066756e6374696f60008201527f6e00000000000000000000000000000000000000000000000000000000000000602082015250565b60006103376021836102ca565b9150610342826102db565b604082019050919050565b600060208201905081810360008301526103668161032a565b9050919050565b600081905092915050565b50565b600061038860008361036d565b915061039382610378565b600082019050919050565b60006103a98261037b565b9150819050919050565b7f5061796d656e74206973206e6f742073756363657366756c6c00000000000000600082015250565b60006103e96019836102ca565b91506103f4826103b3565b602082019050919050565b60006020820190508181036000830152610418816103dc565b9050919050565b7f446f6e6174696f6e206973206e6f742073756363657366756c6c000000000000600082015250565b6000610455601a836102ca565b91506104608261041f565b602082019050919050565b6000602082019050818103600083015261048481610448565b905091905056fea26469706673582212209c1a3f763941294b9a8add7b20c3f1bc721bf9596516b9f3fd67fe5706eed48b64736f6c63430008130033a2646970667358221220afa2adba398265bdae97dfb0db7ac2e4e3526949bc9d42e2860e8a440715773664736f6c63430008130033",
}

// BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeMetaData.ABI instead.
var BridgeABI = BridgeMetaData.ABI

// BridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeMetaData.Bin instead.
var BridgeBin = BridgeMetaData.Bin

// DeployBridge deploys a new Ethereum contract, binding an instance of Bridge to it.
func DeployBridge(auth *bind.TransactOpts, backend bind.ContractBackend, _bridgeAddress common.Address) (common.Address, *types.Transaction, *Bridge, error) {
	parsed, err := BridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeBin), backend, _bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// DonationFactoryAddress is a free data retrieval call binding the contract method 0x5501e348.
//
// Solidity: function donationFactoryAddress() view returns(address)
func (_Bridge *BridgeCaller) DonationFactoryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "donationFactoryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DonationFactoryAddress is a free data retrieval call binding the contract method 0x5501e348.
//
// Solidity: function donationFactoryAddress() view returns(address)
func (_Bridge *BridgeSession) DonationFactoryAddress() (common.Address, error) {
	return _Bridge.Contract.DonationFactoryAddress(&_Bridge.CallOpts)
}

// DonationFactoryAddress is a free data retrieval call binding the contract method 0x5501e348.
//
// Solidity: function donationFactoryAddress() view returns(address)
func (_Bridge *BridgeCallerSession) DonationFactoryAddress() (common.Address, error) {
	return _Bridge.Contract.DonationFactoryAddress(&_Bridge.CallOpts)
}

// GetAllDonations is a free data retrieval call binding the contract method 0x38a59a07.
//
// Solidity: function getAllDonations() view returns((uint256,string,string,address,address,string,bool)[])
func (_Bridge *BridgeCaller) GetAllDonations(opts *bind.CallOpts) ([]DonationManagerDonationInfo, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getAllDonations")

	if err != nil {
		return *new([]DonationManagerDonationInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]DonationManagerDonationInfo)).(*[]DonationManagerDonationInfo)

	return out0, err

}

// GetAllDonations is a free data retrieval call binding the contract method 0x38a59a07.
//
// Solidity: function getAllDonations() view returns((uint256,string,string,address,address,string,bool)[])
func (_Bridge *BridgeSession) GetAllDonations() ([]DonationManagerDonationInfo, error) {
	return _Bridge.Contract.GetAllDonations(&_Bridge.CallOpts)
}

// GetAllDonations is a free data retrieval call binding the contract method 0x38a59a07.
//
// Solidity: function getAllDonations() view returns((uint256,string,string,address,address,string,bool)[])
func (_Bridge *BridgeCallerSession) GetAllDonations() ([]DonationManagerDonationInfo, error) {
	return _Bridge.Contract.GetAllDonations(&_Bridge.CallOpts)
}

// GetDonation is a free data retrieval call binding the contract method 0xef07a81f.
//
// Solidity: function getDonation(uint256 donationID) view returns((uint256,string,string,address,address,string,bool))
func (_Bridge *BridgeCaller) GetDonation(opts *bind.CallOpts, donationID *big.Int) (DonationManagerDonationInfo, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getDonation", donationID)

	if err != nil {
		return *new(DonationManagerDonationInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(DonationManagerDonationInfo)).(*DonationManagerDonationInfo)

	return out0, err

}

// GetDonation is a free data retrieval call binding the contract method 0xef07a81f.
//
// Solidity: function getDonation(uint256 donationID) view returns((uint256,string,string,address,address,string,bool))
func (_Bridge *BridgeSession) GetDonation(donationID *big.Int) (DonationManagerDonationInfo, error) {
	return _Bridge.Contract.GetDonation(&_Bridge.CallOpts, donationID)
}

// GetDonation is a free data retrieval call binding the contract method 0xef07a81f.
//
// Solidity: function getDonation(uint256 donationID) view returns((uint256,string,string,address,address,string,bool))
func (_Bridge *BridgeCallerSession) GetDonation(donationID *big.Int) (DonationManagerDonationInfo, error) {
	return _Bridge.Contract.GetDonation(&_Bridge.CallOpts, donationID)
}

// ApplyForDonation is a paid mutator transaction binding the contract method 0x8114b497.
//
// Solidity: function applyForDonation() returns()
func (_Bridge *BridgeTransactor) ApplyForDonation(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "applyForDonation")
}

// ApplyForDonation is a paid mutator transaction binding the contract method 0x8114b497.
//
// Solidity: function applyForDonation() returns()
func (_Bridge *BridgeSession) ApplyForDonation() (*types.Transaction, error) {
	return _Bridge.Contract.ApplyForDonation(&_Bridge.TransactOpts)
}

// ApplyForDonation is a paid mutator transaction binding the contract method 0x8114b497.
//
// Solidity: function applyForDonation() returns()
func (_Bridge *BridgeTransactorSession) ApplyForDonation() (*types.Transaction, error) {
	return _Bridge.Contract.ApplyForDonation(&_Bridge.TransactOpts)
}

// ApproveDonation is a paid mutator transaction binding the contract method 0x9eb7477c.
//
// Solidity: function approveDonation(address approvedAddress) returns()
func (_Bridge *BridgeTransactor) ApproveDonation(opts *bind.TransactOpts, approvedAddress common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "approveDonation", approvedAddress)
}

// ApproveDonation is a paid mutator transaction binding the contract method 0x9eb7477c.
//
// Solidity: function approveDonation(address approvedAddress) returns()
func (_Bridge *BridgeSession) ApproveDonation(approvedAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.ApproveDonation(&_Bridge.TransactOpts, approvedAddress)
}

// ApproveDonation is a paid mutator transaction binding the contract method 0x9eb7477c.
//
// Solidity: function approveDonation(address approvedAddress) returns()
func (_Bridge *BridgeTransactorSession) ApproveDonation(approvedAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.ApproveDonation(&_Bridge.TransactOpts, approvedAddress)
}

// StartDonation is a paid mutator transaction binding the contract method 0x32f3d5f9.
//
// Solidity: function startDonation(address destionationWallet, string title, string description, address owner, string coverImageCID) returns()
func (_Bridge *BridgeTransactor) StartDonation(opts *bind.TransactOpts, destionationWallet common.Address, title string, description string, owner common.Address, coverImageCID string) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "startDonation", destionationWallet, title, description, owner, coverImageCID)
}

// StartDonation is a paid mutator transaction binding the contract method 0x32f3d5f9.
//
// Solidity: function startDonation(address destionationWallet, string title, string description, address owner, string coverImageCID) returns()
func (_Bridge *BridgeSession) StartDonation(destionationWallet common.Address, title string, description string, owner common.Address, coverImageCID string) (*types.Transaction, error) {
	return _Bridge.Contract.StartDonation(&_Bridge.TransactOpts, destionationWallet, title, description, owner, coverImageCID)
}

// StartDonation is a paid mutator transaction binding the contract method 0x32f3d5f9.
//
// Solidity: function startDonation(address destionationWallet, string title, string description, address owner, string coverImageCID) returns()
func (_Bridge *BridgeTransactorSession) StartDonation(destionationWallet common.Address, title string, description string, owner common.Address, coverImageCID string) (*types.Transaction, error) {
	return _Bridge.Contract.StartDonation(&_Bridge.TransactOpts, destionationWallet, title, description, owner, coverImageCID)
}

// BridgeApplyForDonationIterator is returned from FilterApplyForDonation and is used to iterate over the raw logs and unpacked data for ApplyForDonation events raised by the Bridge contract.
type BridgeApplyForDonationIterator struct {
	Event *BridgeApplyForDonation // Event containing the contract specifics and raw log

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
func (it *BridgeApplyForDonationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeApplyForDonation)
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
		it.Event = new(BridgeApplyForDonation)
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
func (it *BridgeApplyForDonationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeApplyForDonationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeApplyForDonation represents a ApplyForDonation event raised by the Bridge contract.
type BridgeApplyForDonation struct {
	UserWalletAddress common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterApplyForDonation is a free log retrieval operation binding the contract event 0xff25dd54611daf0db1901cc5bad1d6e38fdc0f5d6cb1e18cf0342e9e2763d78d.
//
// Solidity: event ApplyForDonation(address userWalletAddress)
func (_Bridge *BridgeFilterer) FilterApplyForDonation(opts *bind.FilterOpts) (*BridgeApplyForDonationIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ApplyForDonation")
	if err != nil {
		return nil, err
	}
	return &BridgeApplyForDonationIterator{contract: _Bridge.contract, event: "ApplyForDonation", logs: logs, sub: sub}, nil
}

// WatchApplyForDonation is a free log subscription operation binding the contract event 0xff25dd54611daf0db1901cc5bad1d6e38fdc0f5d6cb1e18cf0342e9e2763d78d.
//
// Solidity: event ApplyForDonation(address userWalletAddress)
func (_Bridge *BridgeFilterer) WatchApplyForDonation(opts *bind.WatchOpts, sink chan<- *BridgeApplyForDonation) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "ApplyForDonation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeApplyForDonation)
				if err := _Bridge.contract.UnpackLog(event, "ApplyForDonation", log); err != nil {
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

// ParseApplyForDonation is a log parse operation binding the contract event 0xff25dd54611daf0db1901cc5bad1d6e38fdc0f5d6cb1e18cf0342e9e2763d78d.
//
// Solidity: event ApplyForDonation(address userWalletAddress)
func (_Bridge *BridgeFilterer) ParseApplyForDonation(log types.Log) (*BridgeApplyForDonation, error) {
	event := new(BridgeApplyForDonation)
	if err := _Bridge.contract.UnpackLog(event, "ApplyForDonation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

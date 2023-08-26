package bridge

import (
	bridge "bridge/contracts"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"strings"
)

type NetworkListener struct {
	url               string
	contractAddress   string
	crossChainUrl     string
	crossChainAddress string
	isFilecoin        bool
}

func (c *NetworkListener) listenNetwork() {
	client, err := ethclient.Dial(c.url)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(c.contractAddress)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			if c.isFilecoin && vLog.Topics[0] == common.HexToHash("0x2d4b597935f3cd67fb2eebf1db4debc934cee5c7baa7153f980fdbeb2e74084e") {
				fmt.Println("Topic", vLog.Topics)
				fmt.Println("Data", vLog.Data)
				fmt.Println("Adresss", vLog.Address)
				abiFile, err := ioutil.ReadFile("Lilypad.abi")
				contractAbi, err := abi.JSON(strings.NewReader(string(abiFile)))
				if err != nil {
					// Handle error
					fmt.Println(err)
				}
				event := struct {
					Wallet string
					Result string
				}{}
				err = contractAbi.UnpackIntoInterface(&event, "ResultForUser", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				// 0=>Wallet Address 1=>Amount 2=>ChainId
				fmt.Println("Event", event)
				c.sendRequestToDonationManager(common.HexToAddress(event.Wallet))
			}
			if c.isFilecoin && vLog.Topics[0] == common.HexToHash("0xb3ab19bd795fbabd419dbf34dda545d0dca86e54c45156d3bc156b184733cd88") {
				fmt.Println("Topic", vLog.Topics)
				fmt.Println("Data", vLog.Data)
				fmt.Println("Adresss", vLog.Address)
				abiFile, err := ioutil.ReadFile("DonationManager.abi")
				contractAbi, err := abi.JSON(strings.NewReader(string(abiFile)))
				if err != nil {
					// Handle error
					fmt.Println(err)
				}
				event := struct {
					Wallet common.Address
				}{}
				err = contractAbi.UnpackIntoInterface(&event, "ApplyForDonation", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}
				// 0=>Wallet Address 1=>Amount 2=>ChainId
				fmt.Println("Event", event)
				c.sendRequestToFilecoin(event.Wallet)
			}

		}
	}
}

func (c *NetworkListener) sendRequestToDonationManager(walletAddress common.Address) {
	err, contractInstance, _, res := CreateFunctionRequirementsForDonation(c.crossChainUrl, c.crossChainAddress, os.Getenv("private_key"))
	if err != nil {
		// Handle error
		fmt.Println(err)
	}
	fmt.Println("Gas Price", res.GasPrice)
	result, err2 := contractInstance.ApproveDonation(res, walletAddress)
	if err2 != nil {
		// Handle error
		fmt.Println("err", err2)
	}

	fmt.Println("result", result)
}

func (c *NetworkListener) sendRequestToFilecoin(walletAddress common.Address) {
	err, contractInstance, _, res := CreateFunctionRequirementsForFilecoin(c.crossChainUrl, c.crossChainAddress, os.Getenv("private_key"))
	if err != nil {
		// Handle error
		fmt.Println(err)
	}
	fmt.Println("Gas Price", res.GasPrice)
	result, err2 := contractInstance.Request(res, walletAddress.String())
	if err2 != nil {
		// Handle error
		fmt.Println("err", err2)
	}

	fmt.Println("result", result)
}
func CreateFunctionRequirementsForFilecoin(clientUrl string, lendingPoolAddress string, privateKey string) (error, *bridge.Lilypad, common.Address, *bind.TransactOpts) {
	err, client, _publicAddress, res := CreateFunctionRequirementsForControllers(
		clientUrl,
		"Lilypad.abi",
		lendingPoolAddress,
		privateKey)

	address := common.HexToAddress(lendingPoolAddress)
	contractInstance, err := bridge.NewLilypad(address, client)
	return err, contractInstance, _publicAddress, res
}

func CreateFunctionRequirementsForDonation(clientUrl string, lendingPoolAddress string, privateKey string) (error, *bridge.Bridge, common.Address, *bind.TransactOpts) {
	err, client, _publicAddress, res := CreateFunctionRequirementsForControllers(
		clientUrl,
		"Donation.abi",
		lendingPoolAddress,
		privateKey)

	address := common.HexToAddress(lendingPoolAddress)
	contractInstance, err := bridge.NewBridge(address, client)
	return err, contractInstance, _publicAddress, res
}

func CreateFunctionRequirementsForControllers(clientUrl string, walletAbiName string, oracleAddress string, privateKey string) (error, *ethclient.Client, common.Address, *bind.TransactOpts) {
	client, err := ethclient.Dial(clientUrl)
	if err != nil {
		// Handle error
	}

	address := common.HexToAddress(oracleAddress)
	abiFile, err := ioutil.ReadFile(walletAbiName)
	_, err = abi.JSON(strings.NewReader(string(abiFile)))
	if err != nil {
		// Handle error
		fmt.Println(err)
	}

	fmt.Println(address)
	//fmt.Println(client)
	//fmt.Println(contractAbi)

	if err != nil {
		// Handle error
	}

	_privateKey, _, _publicAddress, _ := GenerateKeypairFromPrivateKeyHex(privateKey)
	res, _ := BuildTransactionOptions(client, _publicAddress, _privateKey, 300000)
	return err, client, _publicAddress, res
}

func BuildTransactionOptions(client *ethclient.Client, fromAddress common.Address, prvKey *ecdsa.PrivateKey, gasLimit uint64) (*bind.TransactOpts, error) {

	// Retrieve the chainID
	chainID, cIDErr := client.ChainID(context.Background())

	if cIDErr != nil {
		return nil, cIDErr
	}

	// Retrieve suggested gas price
	gasPrice, gErr := client.SuggestGasPrice(context.Background())

	if gErr != nil {
		return nil, gErr
	}

	// Create empty options object
	txOptions, txOptErr := bind.NewKeyedTransactorWithChainID(prvKey, chainID)

	if txOptErr != nil {
		return nil, txOptErr
	}

	txOptions.Value = big.NewInt(0) // in wei
	txOptions.GasLimit = gasLimit   // in units
	txOptions.GasPrice = gasPrice

	return txOptions, nil
}

func GenerateKeypairFromPrivateKeyHex(privateKeyHex string) (*ecdsa.PrivateKey, *ecdsa.PublicKey, common.Address, error) {

	log.Println("Generating the keypair...")

	// If hex string has "0x" at the start discard it
	if privateKeyHex[:2] == "0x" {
		privateKeyHex = privateKeyHex[2:]
	}

	// Convert hex key to a private key object
	privateKey, privateKeyErr := crypto.HexToECDSA(privateKeyHex)

	if privateKeyErr != nil {
		return nil, nil, common.Address{}, privateKeyErr
	}

	// Generate the public key using the private key
	publicKey := privateKey.Public()

	// Cast crypto.Publickey object to the ecdsa.PublicKey object
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		return nil, nil, common.Address{}, errors.New("error casting public key to ECDSA")
	}

	// Convert publickey to a address
	pubkeyAsAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return privateKey, publicKeyECDSA, pubkeyAsAddress, nil
}

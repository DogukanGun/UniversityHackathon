package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func main() {
	// Replace this with your Ethereum RPC endpoint URL
	if err := godotenv.Load("server.env"); err != nil {
		log.Fatalf("Failed to load the env vars: %v", err)
	}
	avalancheUrl := os.Getenv("AVALANCHE_FUJI")
	walletAddress := os.Getenv("WALLET_ADDRESS")

	// Create a new RPC client
	ethereumClient, err := ethclient.Dial(avalancheUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Call the eth_getBalance method
	account := common.HexToAddress(walletAddress)
	balance, err := ethereumClient.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Convert the balance from wei to ether
	etherBalance := new(big.Float).Quo(new(big.Float).SetInt(balance), big.NewFloat(1e18))

	fmt.Printf("Balance of %s: %s ETH\n", walletAddress, etherBalance.String())

}

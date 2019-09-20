package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/viper"
)

func init() {
	initViper()
}

func main() {
	rootHashStr := "0x1a883108f6c929c5fe66e2f0df9726e54f5d106d9851711c4d413588d91e9d6f"
	leave1Str := "0xe58b11ef54b1cd83d5dcf3ab433827b907fdb23f1f16f8544b4ae290b4418654"
	leave2Str := "0xbee02ea41003d2f4e587698691f0be417c2c546e70c63648c59f93cbabc8f85f"

	rootHash := [32]byte{}
	copy(rootHash[:], []byte(rootHashStr))

	leave1 := [32]byte{}
	copy(leave1[:], []byte(leave1Str))

	leave2 := [32]byte{}
	copy(leave2[:], []byte(leave2Str))

	leaves := [][32]byte{leave1, leave2}

	set(rootHash, leaves)
	time.Sleep(30 * time.Second) // wait for mining
	get(rootHash)
}

func set(rootHash [32]byte, leaves [][32]byte) {
	client, err := ethclient.Dial(viper.GetString("RPC_URL"))
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(viper.GetString("PRIVATE_KEY"))
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)                   // in wei
	auth.GasLimit = viper.GetUint64("GAS_LIMIT") // in units
	auth.GasPrice = gasPrice

	contractAddress := common.HexToAddress(viper.GetString("CONTRACT_ADDRESS"))
	instance, err := NewMerkleRootHashOracle(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := instance.RegisterMerkleRootHash(auth, rootHash, leaves)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Tx created: %s\n", tx.Hash().Hex())
}

func get(mapKey [32]byte) {
	fmt.Println("Fetching leaves from map...")

	client, err := ethclient.Dial(viper.GetString("RPC_URL"))
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(viper.GetString("CONTRACT_ADDRESS"))
	instance, err := NewMerkleRootHashOracle(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	values, err := instance.GetRootHashContents(&bind.CallOpts{}, mapKey)
	if err != nil {
		log.Fatal(err)
	}

	for i, v := range values {
		humanPosition := i + 1
		fmt.Printf("Value %d: %s\n", humanPosition, v)
	}
}

func initViper() {
	// Declare config file
	viper.SetConfigFile("./config.yaml")

	// Searches for config file in given paths and read it
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// Confirm which config file is used
	fmt.Printf("Using config: %s\n", viper.ConfigFileUsed())
}

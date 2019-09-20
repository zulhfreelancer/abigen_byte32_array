package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/binary"
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
	// root and leaves input can be with or without 0x prefix
	rootHash := newHash("0x_MERKLE_ROOT_HASH")
	leave1 := newHash("0x_LEAVE_1_HASH")
	leave2 := newHash("0x_LEAVE_2_HASH")
	leaves := [][32]byte{leave1, leave2}

	set(rootHash, leaves)
	time.Sleep(30 * time.Second) // wait for mining
	get(rootHash)
}

func newHash(input string) common.Hash {
	if len(input) < 64 {
		log.Fatalf("Invalid input length for '%s'", input)
	}
	return common.HexToHash(input)
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

	log.Printf("Tx created: %s\n", tx.Hash().Hex())
}

func get(mapKey [32]byte) {
	log.Printf("Fetching leaves from map...\n")

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

	fmt.Println("-------")
	fmt.Println("Result:")
	fmt.Println("-------")
	for i, v := range values {
		humanCounter := i + 1
		fmt.Printf("Value %d: %s\n", humanCounter, byte32ArrayToString(v))
	}
}

func byte32ArrayToString(source [32]byte) string {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.LittleEndian, source)
	if err != nil {
		log.Fatal("binary.Write failed:", err)
	}

	valHash := common.BytesToHash(buf.Bytes())
	valHashStr := valHash.String()
	return valHashStr
}

func initViper() {
	// Declare config file
	viper.SetConfigFile("./config.yaml")

	// Searches for config file in given paths and read it
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}

	// Confirm which config file is used
	log.Printf("Using config: %s\n", viper.ConfigFileUsed())
}

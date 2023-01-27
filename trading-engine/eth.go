package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	store "MarketPlace"
)

func main(){
    client, err := ethclient.Dial("https://api.hyperspace.node.glif.io/") 
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("56bdc6f5be9d06f132398f70571312861d2edfb58ba0b03ed09618c843c18fa5")
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

	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
    auth.GasLimit = uint64(4086904)
	auth.GasFeeCap = gasPrice
	auth.GasTipCap = gasPrice

	address := common.HexToAddress("0xDB8d554C03EA59A08793Ee01746b2823DE2ED0d8")
	instance, err := store.NewApi(address, client)
	if err != nil {
		log.Fatal(err)
	}

	tx, err := instance.AddBid(auth, big.NewInt(65), big.NewInt(100))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("tx sent: %s", tx.Hash().Hex())
}
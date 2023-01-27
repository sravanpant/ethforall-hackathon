package main

import (
    "context"
    "fmt"
    "log"
    "math/big"
	"net/http"
	"bytes"
    "strings"
	"time"
    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/accounts/abi"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/ethereum/go-ethereum/ethclient"
    store "MarketPlace"
)

type NewAskEventStruct struct{
	Asker common.Address
	AskAmount *big.Int
	AskSize *big.Int
}

type NewBidEventStruct struct{
	Bidder common.Address
	BidAmount *big.Int
	BidSize *big.Int
}

var asks = []common.Address{}
var bids = []common.Address{} 
var lastBlock uint64 = 0 

func queryContract(client *ethclient.Client, contractAddress common.Address){
	latestBlock, err := client.BlockNumber(context.Background())
	if err != nil {
        log.Fatal(err)
    }
	if lastBlock == 0{
		lastBlock = latestBlock - 200
	}
	fmt.Println("Starting block :- ", lastBlock)
	fmt.Println("Ending Block :- ", latestBlock)
    query := ethereum.FilterQuery{
        FromBlock: big.NewInt(int64(lastBlock)),
        ToBlock:   big.NewInt(int64(latestBlock)),
        Addresses: []common.Address{
            contractAddress,
        },
    }
    logs, err := client.FilterLogs(context.Background(), query)
    if err != nil {
        log.Fatal(err)
    }
    contractAbi, err := abi.JSON(strings.NewReader(string(store.ApiMetaData.ABI)))
    if err != nil {
        log.Fatal(err)
    }
	logNewAskSig := []byte("NewAskEvent(address,uint256,uint256)")
	logNewBidSig := []byte("NewBidEvent(address,uint256,uint256)")
	logNewAskSigHash := crypto.Keccak256Hash(logNewAskSig)
	logNewBidSigHash := crypto.Keccak256Hash(logNewBidSig)
    for _, vLog := range logs {
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		switch vLog.Topics[0].Hex(){

		case logNewAskSigHash.Hex():
			var AskEventVar NewAskEventStruct 
			err := contractAbi.UnpackIntoInterface(&AskEventVar, "NewAskEvent", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			check := false
			for i:=0; i<len(asks); i++{
				if asks[i] == AskEventVar.Asker{
					check = true
				}
			}
			if !check{
				asks = append(asks, AskEventVar.Asker)
				fmt.Println("ASK EVENT :----------")
				fmt.Println(AskEventVar.Asker)
				fmt.Println(AskEventVar.AskAmount)
				fmt.Println(AskEventVar.AskSize) 
				posturl := "http://localhost:8000/addask"
				body := []byte(`{
					"id": "`+AskEventVar.Asker.Hex()+`",
					"amount":`+AskEventVar.AskAmount.String()+`,
					"size":`+AskEventVar.AskSize.String()+
				`}`)
				r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
				if err != nil {
					log.Fatal(err)
				}
				httpclient := &http.Client{}
				res, err := httpclient.Do(r)
				if err != nil {
					panic(err)
				}
				fmt.Println(res)
			}

		
		case logNewBidSigHash.Hex():

			var BidEventVar NewBidEventStruct
			err := contractAbi.UnpackIntoInterface(&BidEventVar, "NewBidEvent", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			check := false
			for i:=0; i<len(bids); i++{
				if bids[i] == BidEventVar.Bidder{
					check = true
				}
			}
			if !check{
				bids = append(bids, BidEventVar.Bidder)
				fmt.Println("BID EVENT :----------")
				fmt.Println(BidEventVar.Bidder)
				fmt.Println(BidEventVar.BidAmount)
				fmt.Println(BidEventVar.BidSize)
				posturl := "http://localhost:8000/addbid"
				body := []byte(`{
					"id": "`+BidEventVar.Bidder.Hex()+`",
					"amount":`+BidEventVar.BidAmount.String()+`,
					"size":`+BidEventVar.BidSize.String()+
				`}`)
				fmt.Println(bytes.NewBuffer(body))
				r, err := http.NewRequest("POST", posturl, bytes.NewBuffer(body))
				if err != nil {
					log.Fatal(err)
				}
				r.Header.Add("Content-Type", "application/json")
				httpclient := &http.Client{}
				res, err := httpclient.Do(r)
				if err != nil {
					panic(err)
				}
				fmt.Println(res)
			}
			 
		}

    }
	lastBlock = latestBlock+1
}

func main() {
	
    client, err := ethclient.Dial("https://api.hyperspace.node.glif.io/")
    if err != nil {
        log.Fatal(err)
    }
	contractAddress := common.HexToAddress("0xDB8d554C03EA59A08793Ee01746b2823DE2ED0d8")
	for{
		queryContract(client, contractAddress)
		time.Sleep(90*time.Second)
	}
	
}
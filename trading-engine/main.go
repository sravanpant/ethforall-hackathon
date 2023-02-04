package main

import(
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"fmt"
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	store "MarketPlace"
)

type Bid struct{
	ID string `json:"id"`
	Amount int `json:"amount"`
	Size int `json:"size"`
}

type Ask struct{
	ID string `json:"id"`
	Amount int `json:"amount"`
	Size int `json:"size"`
}

type ByAmountBid []Bid
type ByAmountAsk []Ask

func (a ByAmountBid) Len() int{ return len(a) }
func (a ByAmountBid) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAmountBid) Less(i, j int) bool { return a[i].Amount < a[j].Amount }
func (a ByAmountAsk) Len() int{ return len(a) }
func (a ByAmountAsk) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAmountAsk) Less(i, j int) bool { return a[i].Amount < a[j].Amount }

var bids = []Bid{}
var asks = []Ask{}

func orderBids(b Bid){
	bids = append(bids, b)
	sort.Slice(bids, func(i, j int) bool{
		return bids[i].Amount > bids[j].Amount
	})
}

func orderAsks(a Ask){
	asks = append(asks, a)
	sort.Sort(ByAmountAsk(asks))
}

func matchBidOrder(){
	bidLength := len(bids)
	for i:=0; i<bidLength; i++{
		b := bids[i]
		askLength := len(asks)
		for j:=0; j<askLength; j++{
			a := asks[j]
			if b.Amount >= a.Amount{
				if b.Size == a.Size{
					fmt.Println("Order created 00x")
					sendOrder(b.ID, a.ID, a.Amount, b.Size)
					asks = append(asks[:j], asks[j+1:]...)
					askLength--
					bids = append(bids[:i], bids[i+1:]...)
					bidLength--
					break
				}else if a.Size > b.Size{
					fmt.Println("Order created with greater ask size")
					sendOrder(b.ID, a.ID, a.Amount, b.Size)
					asks[j].Size = (a.Size - b.Size)
					bids = append(bids[:i], bids[i+1:]...)
					bidLength--
					break
				}
			}else{
				break
			}
			
		}
	}
}

func matchAskOrder(){
	askLength := len(asks)
	for i:=0;i<askLength;i++{
		bidLength := len(bids)
		a:=asks[i]
		for j:=0; j<bidLength; j++{
			b:=bids[j]
			if a.Amount <= b.Amount{
				if b.Size == a.Size{
					fmt.Println("Order created 00x")
					sendOrder(b.ID, a.ID, a.Amount, b.Size)
					asks = append(asks[:i], asks[i+1:]...)
					askLength--
					bids = append(bids[:j], bids[j+1:]...)
					bidLength--
					break
				}else if a.Size > b.Size{
					fmt.Println("Order created with greater ask size")
					sendOrder(b.ID, a.ID, a.Amount, b.Size)
					asks[i].Size = (a.Size - b.Size)
					bids = append(bids[:j], bids[j+1:]...)
					j--
					bidLength--
				}
			}else{
				break
			}
		}
	}
}

func getBids(context *gin.Context){
	context.IndentedJSON(http.StatusOK, bids)
}

func getAsks(context *gin.Context){
	context.IndentedJSON(http.StatusOK, asks)
}

func addBid(context *gin.Context){
	var newBid Bid
	if err:= context.BindJSON(&newBid); err!=nil{
		return
	}

	orderBids(newBid)
	matchBidOrder()
	context.IndentedJSON(http.StatusCreated, newBid)
}

func addAsk(context *gin.Context){
	var newAsk Ask
	if err:= context.BindJSON(&newAsk); err!=nil{
		return
	}

	orderAsks(newAsk)
	matchAskOrder()
	context.IndentedJSON(http.StatusCreated, newAsk)
}

func sendOrder(asker string, bidder string, amount int, size int){
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
    auth.GasLimit = uint64(6086904)
	auth.GasFeeCap = gasPrice
	auth.GasTipCap = gasPrice

	address := common.HexToAddress("0xDB8d554C03EA59A08793Ee01746b2823DE2ED0d8")
	instance, err := store.NewApi(address, client)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(asker+bidder, common.HexToAddress(bidder), common.HexToAddress(asker), big.NewInt(int64(amount)), big.NewInt(int64(size)))
	tx, err := instance.AddOrder(auth, asker+bidder, common.HexToAddress(bidder), common.HexToAddress(asker), big.NewInt(int64(amount)), big.NewInt(int64(size)))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("tx sent: %s", tx.Hash().Hex())
}

func main(){
	router:=gin.Default()
	router.GET("/bids", getBids)
	router.GET("/asks", getAsks)
	router.POST("addbid", addBid)
	router.POST("/addask", addAsk)
	router.Run("localhost:8000")
}
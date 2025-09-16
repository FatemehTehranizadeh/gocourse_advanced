package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

func main() {

	ch := make(chan float64, 2)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go gettingPricefromCoinpaprika(ch, &wg)
	go gettingPricefromCoingecko(ch, &wg)

	fmt.Println("The price of Bitcoin is: ", <- ch)
	

	wg.Wait()
	close(ch)

}

func gettingPricefromCoinpaprika(ch chan float64, wg *sync.WaitGroup) {
	resp, err := http.Get("https://api.coinpaprika.com/v1/tickers/btc-bitcoin")
	if err != nil {
		log.Fatal("Error while getting the data: ", err)
	}
	defer resp.Body.Close()

	respBuff, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error during reading: ", err)
	}

	var respData map[string]interface{}
	err = json.Unmarshal(respBuff, &respData)
	if err != nil {
		log.Fatal("Error while unmarshalling: ", err)
	}

	quotes := respData["quotes"].(map[string]interface{})
	price := quotes["USD"].(map[string]interface{})["price"]
	// fmt.Println("Price: ",price)
	ch <- price.(float64)
	defer wg.Done()
}


func gettingPricefromCoingecko(ch chan float64, wg *sync.WaitGroup) {
	resp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd")
	if err != nil {
		log.Fatal("Error while getting the data: ", err)
	}
	defer resp.Body.Close()

	respBuff, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error during reading: ", err)
	}

	var respData map[string]interface{}
	err = json.Unmarshal(respBuff, &respData)
	if err != nil {
		log.Fatal("Error while unmarshalling: ", err)
	}

	price := respData["bitcoin"].(map[string]interface{})["usd"]
	ch <- price.(float64)
	defer wg.Done()
}
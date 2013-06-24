package btce

import (
	"encoding/json"
	"github.com/mattomatic/go-bitcoin/common"
	"io/ioutil"
	"net/http"
	"strconv"
)

func getPrice(price string) common.Price {
	return common.Price(getFloat(price))
}

func getVolume(volume string) common.Volume {
	return common.Volume(getFloat(volume))
}

func getFloat(amount string) float64 {
	x, err := strconv.ParseFloat(amount, 64)

	if err != nil {
		panic(err.Error())
	}

	return x
}

func httpRequest(url string) []byte {
	resp, err := http.Get(url)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err.Error())
	}

	return body
}

func getOrderBook() *OrderBook {
	bytes := httpRequest(OrderBookUrl)
	book := &OrderBook{}

	err := json.Unmarshal(bytes, book)

	if err != nil {
		panic(err.Error())
	}

	return book
}

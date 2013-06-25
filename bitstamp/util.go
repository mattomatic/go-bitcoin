package bitstamp

import (
	"encoding/json"
	"fmt"
	"github.com/mattomatic/go-bitcoin/common"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func getPrice(price string) common.Price {
	return common.Price(getFloat(price))
}

func getVolume(volume string) common.Volume {
	return common.Volume(getFloat(volume))
}

func getTimestamp(amount string) time.Time {
	return time.Unix(int64(getInt(amount)), 0)
}

func getInt(amount string) int {
	x, err := strconv.Atoi(amount)

	if err != nil {
		panic(err.Error())
	}

	return x
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

func getTicker() *Ticker {
	bytes := httpRequest(TickerUrl)
	ticker := &Ticker{}

	err := json.Unmarshal(bytes, ticker)

	if err != nil {
		panic(err.Error())
	}

	return ticker
}

func getOrderBook() (book *OrderBook, err error) {
	defer func() {
		if r := recover(); r != nil {
			book, err = nil, fmt.Errorf("panic: %v", r)
		}
	}()

	bytes := httpRequest(OrderBookUrl)
	book = &OrderBook{}
	err = json.Unmarshal(bytes, book)

	if err != nil {
		panic(err.Error())
	}

	return book, nil
}

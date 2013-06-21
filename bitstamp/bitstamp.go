package bitstamp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	TickerUrl = "https://bitstamp.net/api/ticker"
)

func GetTicker() *Ticker {
	resp, err := http.Get(TickerUrl)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err.Error())
	}

	ticker := &Ticker{}
	err = json.Unmarshal(body, ticker)

	if err != nil {
		panic(err.Error())
	}

	return ticker
}

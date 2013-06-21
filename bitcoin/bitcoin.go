package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/mattomatic/go-bitcoin/bitstamp"
	"io/ioutil"
	"net/http"
)

func init() {
	flag.Parse()
}

func main() {
	resp, err := http.Get("https://www.bitstamp.net/api/ticker/")

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err.Error())
	}

	ticker := &bitstamp.Ticker{}
	json.Unmarshal(body, ticker)
	fmt.Println(ticker)
}

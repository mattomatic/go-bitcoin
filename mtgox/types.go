package mtgox

import (
	"time"
)

const (
	ExchangeId = "MTGOX"
)

type Feed struct {
	Type      string
	Timestamp time.Time
	Message   interface{}
}

type LoginReply struct {
	Op      string `json:"op"`
	Message string `json:"message"`
}

type Trade struct {
	AmountInt     string `json:"amount_int"`
	Date          int    `json:"date"`
	Item          string `json:"item"`
	PriceCurrency string `json:"price_currency"`
	PriceInt      string `json:"price_int"`
	Primary       string `json:"primary"`
	Properties    string `json:"properties"`
	Tid           string `json:"tid"`
	TradeType     string `json:"bid"`
	Type          string `json:"type"`
}

type TradeFeed struct {
	Header
	Trade `json:"trade"`
}

type Depth struct {
	Currency       string `json:"currency"`
	Item           string `json:"item"`
	Now            string `json:"now"`
	Price          string `json:"price"`
	PriceInt       string `json:"price_int"`
	TotalVolumeInt string `json:"total_volume_int"`
	Type           string `json:"type_str"`
	VolumeInt      string `json:"volume_int"`
}

type DepthFeed struct {
	Header
	Depth `json:"depth"`
}

type TickerField struct {
	currency string `json:"currency"`
	display  string `json:"display"`
	value    string `json:"value"`
	valueInt string `json:"value_int"`
}

type Ticker struct {
	average      TickerField `json:"avg"`
	bid          TickerField `json:"buy"`
	ask          TickerField `json:"sell"`
	last         TickerField `json:"last"`
	lastLocal    TickerField `json:"last_local"`
	lastOriginal TickerField `json:"last_orig"`
	high         TickerField `json:"high"`
	low          TickerField `json:"low"`
	volume       TickerField `json:"vol"`
	vwap         TickerField `json:"vwap"`
	timestamp    time.Time
}

type TickerFeed struct {
	Header
	Ticker `json:"ticker"`
}

type Header struct {
	Channel string `json:"channel"`
	Op      string `json:"op"`
	Origin  string `json:"origin"`
	Private string `json:"private"`
}

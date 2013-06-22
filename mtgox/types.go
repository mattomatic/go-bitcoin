package mtgox

import (
	"github.com/mattomatic/go-bitcoin/common"
)

type Feed common.Feed

type LoginReply struct {
	Op      string `json:"op"`
	Message string `json:"message"`
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

package mtgox

import (
	"github.com/mattomatic/go-bitcoin/common"
)

type TickerField struct {
	Currency string `json:"currency"`
	Display  string `json:"display"`
	Value    string `json:"value"`
	ValueInt string `json:"value_int"`
}

type Ticker struct {
	Average      TickerField `json:"avg"`
	Bid          TickerField `json:"buy"`
	Ask          TickerField `json:"sell"`
	Last         TickerField `json:"last"`
	LastLocal    TickerField `json:"last_local"`
	LastOriginal TickerField `json:"last_orig"`
	High         TickerField `json:"high"`
	Low          TickerField `json:"low"`
	Volume       TickerField `json:"vol"`
	Vwap         TickerField `json:"vwap"`
}

func (t *Ticker) GetExchange() common.Exchange { return common.Exchange(ExchangeId) }
func (t *Ticker) GetSymbol() common.Symbol     { return common.Symbol("BTC") }
func (t *Ticker) GetBid() common.Price         { return getPrice(t.Bid.ValueInt) }
func (t *Ticker) GetAsk() common.Price         { return getPrice(t.Ask.ValueInt) }
func (t *Ticker) GetHigh() common.Price        { return getPrice(t.High.ValueInt) }
func (t *Ticker) GetLow() common.Price         { return getPrice(t.Low.ValueInt) }
func (t *Ticker) GetVolume() common.Volume     { return getVolume(t.Volume.ValueInt) }

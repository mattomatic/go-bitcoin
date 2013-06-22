package bitstamp

import (
	"github.com/mattomatic/go-bitcoin/common"
	"time"
)

type Ticker struct {
	Timestamp string `json:"timestamp"`
	Bid       string `json:"bid"`
	Ask       string `json:"ask"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Last      string `json:"last"`
	Volume    string `json:"volume"`
}

func (t *Ticker) GetExchange() common.Exchange {
	return common.Exchange(ExchangeId)
}

func (t *Ticker) GetSymbol() common.Symbol {
	return common.Symbol("BTC")
}

func (t *Ticker) GetTimestamp() time.Time {
	return time.Unix(int64(getInt(t.Timestamp)), 0)
}

func (t *Ticker) GetBid() common.Price {
	return getPrice(t.Bid)
}

func (t *Ticker) GetAsk() common.Price {
	return getPrice(t.Ask)
}

func (t *Ticker) GetHigh() common.Price {
	return getPrice(t.High)
}

func (t *Ticker) GetLow() common.Price {
	return getPrice(t.Low)
}

func (t *Ticker) GetVolume() common.Volume {
	return getVolume(t.Volume)
}

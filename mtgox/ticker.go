package mtgox

import (
	"github.com/mattomatic/go-bitcoin/common"
	"time"
)

func (t *Ticker) Exchange() common.Exchange {
	return common.Exchange(ExchangeId)
}

func (t *Ticker) Symbol() common.Symbol {
	return common.Symbol("BTC")
}

func (t *Ticker) Timestamp() time.Time {
	return t.timestamp
}

func (t *Ticker) Bid() common.Price {
	return getPrice(t.bid.valueInt)
}

func (t *Ticker) Ask() common.Price {
	return getPrice(t.ask.valueInt)
}

func (t *Ticker) High() common.Price {
	return getPrice(t.high.valueInt)
}

func (t *Ticker) Low() common.Price {
	return getPrice(t.low.valueInt)
}

func (t *Ticker) Volume() common.Volume {
	return getVolume(t.volume.valueInt)
}

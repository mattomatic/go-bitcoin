package bitstamp

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
	return time.Unix(int64(getInt(t.timestamp)), 0)
}

func (t *Ticker) Bid() common.Price {
	return getPrice(t.bid)
}

func (t *Ticker) Ask() common.Price {
	return getPrice(t.ask)
}

func (t *Ticker) High() common.Price {
	return getPrice(t.high)
}

func (t *Ticker) Low() common.Price {
	return getPrice(t.low)
}

func (t *Ticker) Volume() common.Volume {
	return getVolume(t.volume)
}

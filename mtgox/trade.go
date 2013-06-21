package mtgox

import (
	"github.com/mattomatic/go-bitcoin/common"
	"time"
)

func (t *Trade) Exchange() common.Exchange {
	return common.Exchange(ExchangeId)
}

func (t *Trade) Symbol() common.Symbol {
	return common.Symbol(t.Item)
}

func (t *Trade) Timestamp() time.Time {
	return time.Unix(int64(t.Date), 0)
}

func (t *Trade) Volume() common.Volume {
	return getVolume(t.AmountInt)
}

func (t *Trade) Price() common.Price {
	return getPrice(t.PriceInt)
}

func (t *Trade) Currency() common.Currency {
	return common.Currency(t.PriceCurrency)
}

func (t *Trade) TradeId() common.TradeId {
	return common.TradeId(t.Tid)
}

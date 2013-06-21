package mtgox

import (
	"github.com/mattomatic/go-bitcoin/common"
	"strconv"
	"time"
)

const (
	Divisor = 1e8
)

func (t *Trade) Exchange() common.Exchange {
	return common.Exchange("MTGOX")
}

func (t *Trade) Symbol() common.Symbol {
	return common.Symbol(t.Item)
}

func (t *Trade) Timestamp() time.Time {
	return time.Unix(int64(t.Date), 0)
}

func (t *Trade) Volume() common.Volume {
	volume, err := strconv.Atoi(t.AmountInt)

	if err != nil {
		panic(err.Error())
	}

	return common.Volume(volume / Divisor)
}

func (t *Trade) Price() common.Price {
	price, err := strconv.Atoi(t.PriceInt)

	if err != nil {
		panic(err.Error())
	}

	return common.Price(price / Divisor)
}

func (t *Trade) Currency() common.Currency {
	return common.Currency(t.PriceCurrency)
}

func (t *Trade) TradeId() common.TradeId {
	return common.TradeId(t.Tid)
}

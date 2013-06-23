package bitstamp

import (
	"github.com/mattomatic/go-bitcoin/common"
)

type Order struct {
	Price  string
	Volume string
}

func (o *Order) GetExchange() common.Exchange {
	return common.Exchange(ExchangeId)
}

func (o *Order) GetSymbol() common.Symbol {
	return common.Symbol("BTC")
}

func (o *Order) GetPrice() common.Price {
	return getPrice(o.Price)
}

func (o *Order) GetVolume() common.Volume {
	return getVolume(o.Volume)
}

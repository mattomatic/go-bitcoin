package mtgox

import (
	"github.com/mattomatic/go-bitcoin/common"
	"github.com/petar/GoLLRB/llrb"
)

type Order struct {
	volume string
	price  string
}

func (o *Order) GetExchange() common.Exchange {
	return common.Exchange(ExchangeId)
}

func (o *Order) GetSymbol() common.Symbol {
	return common.Symbol("BTC")
}

func (o *Order) GetPrice() common.Price {
	return getPrice(o.price)
}

func (o *Order) GetVolume() common.Volume {
	return getVolume(o.volume)
}

// Define function to let red-black tree work for ordering orders :p
func (o *Order) Less(than llrb.Item) bool {
	return o.GetPrice() < than.(*Order).GetPrice()
}

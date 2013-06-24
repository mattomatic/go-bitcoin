package mtgox

import (
	"github.com/mattomatic/go-bitcoin/common"
	"github.com/petar/GoLLRB/llrb"
)

type Order struct {
	volume common.Volume
	price  common.Price
}

func (o *Order) GetExchange() common.Exchange { return common.Exchange(ExchangeId) }
func (o *Order) GetSymbol() common.Symbol     { return common.Symbol("BTC") }
func (o *Order) GetPrice() common.Price       { return o.price }
func (o *Order) GetVolume() common.Volume     { return o.volume }

// Define function to let red-black tree work for ordering orders :p
func (o *Order) Less(than llrb.Item) bool {
	return o.GetPrice() < than.(*Order).GetPrice()
}

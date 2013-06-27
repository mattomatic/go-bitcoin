package mtgox

import (
	"github.com/mattomatic/go-bitcoin/common"
)

type Depth struct {
	Currency    string `json:"currency"`
	Item        string `json:"item"`
	Now         string `json:"now"`
	Price       string `json:"price_int"`
	TotalVolume string `json:"total_volume_int"`
	Type        string `json:"type_str"`
}

func (g *Depth) GetExchange() common.Exchange { return ExchangeId }
func (g *Depth) GetSymbol() common.Symbol     { return "BTC" }
func (d *Depth) GetVolume() common.Volume     { return getVolume(d.TotalVolume) }
func (d *Depth) GetPrice() common.Price       { return getPrice(d.Price) }
func (d *Depth) GetFee() common.Fee           { return ExchangeFee }
func (d *Depth) GetSide() common.Side         { return getSide(d.Type) }

func getSide(side string) common.Side {
	switch side {
	case "bid":
		return common.Bid
	case "ask":
		return common.Ask
	default:
		panic("unrecognized side")
	}
}

package bitstamp

import (
	"encoding/json"
	"github.com/mattomatic/go-bitcoin/common"
)

type Order struct {
	price  common.Price
	volume common.Volume
	side   common.Side
}

func (o *Order) GetExchange() common.Exchange { return common.Exchange(ExchangeId) }
func (o *Order) GetSymbol() common.Symbol     { return common.Symbol("BTC") }
func (o *Order) GetPrice() common.Price       { return o.price }
func (o *Order) GetVolume() common.Volume     { return o.volume }
func (o *Order) GetSide() common.Side         { return o.side }

func (o *Order) UnmarshalJSON(bytes []byte) error {
	values := &[2]string{}
	err := json.Unmarshal(bytes, values)

	if err != nil {
		panic(err.Error())
	}

	o.price = getPrice(values[0])
	o.volume = getVolume(values[1])

	return err
}

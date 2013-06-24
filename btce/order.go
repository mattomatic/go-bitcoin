package btce

import (
	"encoding/json"
	"github.com/mattomatic/go-bitcoin/common"
)

type Order struct {
	price  common.Price
	volume common.Volume
}

func (o *Order) GetExchange() common.Exchange { return common.Exchange(ExchangeId) }
func (o *Order) GetSymbol() common.Symbol     { return common.Symbol("BTC") }
func (o *Order) GetPrice() common.Price       { return o.price }
func (o *Order) GetVolume() common.Volume     { return o.volume }

func (o *Order) UnmarshalJSON(bytes []byte) error {
	values := &[2]float64{}
	err := json.Unmarshal(bytes, values)

	if err != nil {
		panic(err.Error())
	}

	o.price = common.Price(values[0])
	o.volume = common.Volume(values[1])

	return err
}

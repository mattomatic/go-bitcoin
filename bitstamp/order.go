package bitstamp

import (
    "github.com/mattomatic/go-bitcoin/common"
)

type Order struct {
	Price  string
	Volume string
}

func (o *Order) GetPrice() common.Price {
    return getPrice(o.Price)
}

func (o *Order) GetVolume() common.Volume {
    return getVolume(o.Volume)
}
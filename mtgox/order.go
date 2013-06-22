package mtgox

import (
	"github.com/mattomatic/go-bitcoin/common"
	"github.com/mattomatic/go-heap/heap"
)

type Order struct {
	volume string
	price  string
}

func (o *Order) Less(than heap.Item) bool {
	return o.GetPrice() < than.(*Order).GetPrice()
}

func (o *Order) GetPrice() common.Price {
	return getPrice(o.price)
}

func (o *Order) GetVolume() common.Volume {
	return getVolume(o.volume)
}

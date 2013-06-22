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

func (d *Depth) GetVolume() common.Volume {
	return getVolume(d.TotalVolume)
}

func (d *Depth) GetPrice() common.Price {
	return getPrice(d.Price)
}

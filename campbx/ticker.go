package campbx

import (
	"github.com/mattomatic/go-bitcoin/common"
)

type Ticker struct {
	Last string `json:"Last Trade"`
	Bid  string `json:"Best Bid"`
	Ask  string `json:"Best Ask"`
}

func (t *Ticker) GetLast() common.Price { return getPrice(t.Last) }
func (t *Ticker) GetBid() common.Price  { return getPrice(t.Bid) }
func (t *Ticker) GetAsk() common.Price  { return getPrice(t.Ask) }

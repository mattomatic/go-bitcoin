package algo

import (
	"fmt"
	"github.com/mattomatic/go-bitcoin/common"
)

type Pair struct {
	Buy    common.Order
	Sell   common.Order
	Credit common.Amount
}

func (p *Pair) String() string {
	return fmt.Sprintf("%s -- %s Credit: $%v", common.OrderString(p.Buy), common.OrderString(p.Sell), p.Credit)
}

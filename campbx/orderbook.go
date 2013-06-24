package campbx

import (
	"github.com/mattomatic/go-bitcoin/common"
)

type OrderBook struct {
	Bids []Order `json:"Bids"`
	Asks []Order `json:"Asks"`
}

func (o *OrderBook) GetBids() <-chan common.Order {
	return getChannelForward(o.Bids)
}

func (o *OrderBook) GetAsks() <-chan common.Order {
	// campbx asks are sorted highest to lowest for some reason.
	return getChannelReverse(o.Asks)
}

func getChannelForward(orders []Order) chan common.Order {
	ch := make(chan common.Order)
	go func() {
		defer close(ch)
		for i := 0; i < len(orders); i++ {
			ch <- &orders[i]
		}
	}()
	return ch
}

func getChannelReverse(orders []Order) chan common.Order {
	ch := make(chan common.Order)
	go func() {
		defer close(ch)
		for i := len(orders) - 1; i >= 0; i-- {
			ch <- &orders[i]
		}
	}()
	return ch
}

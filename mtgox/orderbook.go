package mtgox

import (
	"github.com/mattomatic/go-bitcoin/common"
	"github.com/petar/GoLLRB/llrb"
)

type OrderBook struct {
	Bids *llrb.LLRB
	Asks *llrb.LLRB
}

func newOrderBook() *OrderBook {
	return &OrderBook{
		llrb.New(),
		llrb.New()}
}

func (o *OrderBook) GetBids() []common.Order {
	return make([]common.Order, 0)
}

func (o *OrderBook) GetAsks() []common.Order {
	return make([]common.Order, 0)
}

func (o *OrderBook) handleDepth(depth *Depth) {
	order := makeOrder(depth)

	switch depth.Type {
	case "bid":
		update(o.Bids, order)
	case "ask":
		update(o.Asks, order)
	default:
		panic("unrecognized depth side!")
	}
}

func update(tree *llrb.LLRB, order *Order) {
	if order.GetVolume() == 0 {
		tree.Delete(order)
	} else {
		tree.ReplaceOrInsert(order)
	}
}

func makeOrder(depth *Depth) *Order {
	return &Order{price: depth.Price, volume: depth.TotalVolume}
}

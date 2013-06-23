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

func (o *OrderBook) GetBids() chan common.Order {
	orders := make(chan common.Order)

	iterator := func(item llrb.Item) bool {
		orders <- item.(*Order)
	}

	go func() {
		defer close(orders)
		o.Bids.DescendLessOrEqual(o.Bids.Max(), iterator)
	}()

	return orders
}

func (o *OrderBook) GetAsks() chan common.Order {
	orders := make(chan common.Order)

	iterator := func(item llrb.Item) bool {
		orders <- item.(*Order)
	}

	go func() {
		defer close(orders)
		o.Asks.AscendGreaterOrEqual(o.Asks.Min(), adder)
	}()

	return orders
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

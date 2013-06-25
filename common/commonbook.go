package common

import (
	"github.com/petar/GoLLRB/llrb"
)

type CommonBook struct {
	bids *llrb.LLRB
	asks *llrb.LLRB
}

type Item struct {
	Diff
}

func (i *Item) Less(than llrb.Item) bool {
	that := than.(Diff)

	if i.GetPrice() < that.GetPrice() {
		return true
	}

	if i.GetPrice() > that.GetPrice() {
		return false
	}

	// same price -- break ties by exchange
	return i.GetExchange() < that.GetExchange()
}

func NewBook() *CommonBook {
	return &CommonBook{
		llrb.New(),
		llrb.New()}
}

func (c *CommonBook) GetBids() <-chan Order {
	orders := make(chan Order)

	iterator := func(item llrb.Item) bool {
		orders <- item.(Order)
		return true
	}

	go func() {
		defer close(orders)
		c.bids.DescendLessOrEqual(c.bids.Max(), iterator)
	}()

	return orders
}

func (c *CommonBook) GetAsks() <-chan Order {
	orders := make(chan Order)

	iterator := func(item llrb.Item) bool {
		orders <- item.(Order)
		return true
	}

	go func() {
		defer close(orders)
		c.asks.AscendGreaterOrEqual(c.asks.Min(), iterator)
	}()

	return orders
}

func (c *CommonBook) ApplyDiff(diff Diff) {
	switch diff.GetSide() {
	case Bid:
		updateTree(c.bids, diff)
	case Ask:
		updateTree(c.asks, diff)
	default:
		panic("unrecognized side!")
	}
}

func updateTree(tree *llrb.LLRB, diff Diff) {
	if diff.GetVolume() == 0 {
		tree.Delete(&Item{diff})
	} else {
		tree.ReplaceOrInsert(&Item{diff})
	}
}

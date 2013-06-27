package algo

import (
	"github.com/mattomatic/go-bitcoin/common"
)

// Generate a list of changes that occurred between the old book and the new book.
func GenerateDiffs(old, new common.OrderBook) <-chan common.Diff {
	diffs := make(chan common.Diff)
	go generate(diffs, old, new)
	return diffs
}

func generate(diffs chan common.Diff, old, new common.OrderBook) {
	defer close(diffs)
	walk(diffs, old.GetBids(), new.GetBids())
	walk(diffs, old.GetAsks(), new.GetAsks())
}

// walk down the book depth generating insert/delete/modify diffs
func walk(diffs chan common.Diff, oldOrders, newOrders <-chan common.Order) {
	old, oldOk := <-oldOrders
	new, newOk := <-newOrders

	for oldOk && newOk {
		if shouldInsertNew(old, new) {
			diffs <- insert(new)
			new, newOk = <-newOrders
		} else if shouldRemoveOld(old, new) {
			diffs <- remove(old)
			old, oldOk = <-oldOrders
		} else {
			if old.GetVolume() != new.GetVolume() {
				diffs <- insert(old)
			}

			new, newOk = <-newOrders
			old, oldOk = <-oldOrders
		}
	}

	for oldOk { // remove trailing oldies
		diffs <- remove(old)
		old, oldOk = <-oldOrders
	}

	for newOk { // insert trailing newbies
		diffs <- insert(new)
		new, newOk = <-newOrders
	}
}

func shouldInsertNew(old, new common.Order) bool {
	if old.GetSide() == common.Bid {
		return new.GetPrice() > old.GetPrice()
	} else {
		return new.GetPrice() < old.GetPrice()
	}
}

func shouldRemoveOld(old, new common.Order) bool {
	if old.GetSide() == common.Bid {
		return old.GetPrice() > new.GetPrice()
	} else {
		return old.GetPrice() < new.GetPrice()
	}
}

func remove(order common.Order) common.Diff {
	return makeDiff(order, 0)
}

func insert(order common.Order) common.Diff {
	return makeDiff(order, order.GetVolume())
}

func update(order common.Order) common.Diff {
	return makeDiff(order, order.GetVolume())
}

func makeDiff(order common.Order, volume common.Volume) common.Diff {
	return common.NewOrder(
		order.GetExchange(),
		order.GetSymbol(),
		volume,
		order.GetPrice(),
		order.GetFee(),
		order.GetSide())
}

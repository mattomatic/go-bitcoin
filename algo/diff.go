package algo

import (
	"github.com/mattomatic/go-bitcoin/common"
)

type diff struct {
	exchange common.Exchange
	symbol   common.Symbol
	volume   common.Volume
	price    common.Price
	side     common.Side
}

func (d *diff) GetExchange() common.Exchange { return d.exchange }
func (d *diff) GetSymbol() common.Symbol     { return d.symbol }
func (d *diff) GetVolume() common.Volume     { return d.volume }
func (d *diff) GetPrice() common.Price       { return d.price }
func (d *diff) GetSide() common.Side         { return d.side }

// Generate a list of changes that occurred between the old book and the new book.
func GenerateDiffs(old, new common.OrderBook) <-chan common.Diff {
	diffs := make(chan common.Diff)
	go generate(diffs, old, new)
	return diffs
}

func generate(diffs chan common.Diff, oldBook, newBook common.OrderBook) {
	defer close(diffs)
	walk(diffs, common.Bid, oldBook.GetBids(), newBook.GetBids())
	walk(diffs, common.Ask, oldBook.GetAsks(), newBook.GetAsks())
}

// walk down the book depth generating insert/delete/modify diffs
func walk(diffs chan common.Diff, side common.Side, oldOrders, newOrders <-chan common.Order) {
	old, oldOk := <-oldOrders
	new, newOk := <-newOrders

	for oldOk && newOk {
		if shouldInsertNew(side, old, new) {
			diffs <- insert(new, side)
			new, newOk = <-newOrders
		} else if shouldRemoveOld(side, old, new) {
			diffs <- remove(old, side)
			old, oldOk = <-oldOrders
		} else {
			if old.GetVolume() != new.GetVolume() {
				diffs <- insert(old, side)
			}

			new, newOk = <-newOrders
			old, oldOk = <-oldOrders
		}
	}

	for oldOk { // remove trailing oldies
		diffs <- remove(old, side)
		old, oldOk = <-oldOrders
	}

	for newOk { // insert trailing newbies
		diffs <- insert(new, side)
		new, newOk = <-newOrders
	}
}

func shouldInsertNew(side common.Side, old, new common.Order) bool {
	if side == common.Bid {
		return new.GetPrice() > old.GetPrice()
	} else {
		return new.GetPrice() < old.GetPrice()
	}
}

func shouldRemoveOld(side common.Side, old, new common.Order) bool {
	if side == common.Bid {
		return old.GetPrice() > new.GetPrice()
	} else {
		return old.GetPrice() < new.GetPrice()
	}
}

func remove(order common.Order, side common.Side) *diff {
	diff := makeDiff(order, side)
	diff.volume = 0
	return diff
}

func insert(order common.Order, side common.Side) *diff {
	return makeDiff(order, side)
}

func update(order common.Order, side common.Side) *diff {
	return makeDiff(order, side)
}

func makeDiff(order common.Order, side common.Side) *diff {
	return &diff{
		exchange: order.GetExchange(),
		symbol:   order.GetSymbol(),
		volume:   order.GetVolume(),
		price:    order.GetPrice(),
		side:     side,
	}
}

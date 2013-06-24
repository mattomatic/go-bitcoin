package common

type diff struct {
	exchange Exchange
	symbol   Symbol
	volume   Volume
	price    Price
	side     Side
}

func (d *diff) GetExchange() Exchange { return d.exchange }
func (d *diff) GetSymbol() Symbol     { return d.symbol }
func (d *diff) GetVolume() Volume     { return d.volume }
func (d *diff) GetPrice() Price       { return d.price }
func (d *diff) GetSide() Side         { return d.side }

// Generate a list of changes that occurred between the old book and the new book.
func GenerateDiffs(old, new OrderBook) <-chan DepthDiff {
	diffs := make(chan DepthDiff)
	go generate(diffs, old, new)
	return diffs
}

func generate(diffs chan DepthDiff, oldBook, newBook OrderBook) {
	defer close(diffs)
	walk(diffs, Bid, oldBook.GetBids(), newBook.GetBids())
	walk(diffs, Ask, oldBook.GetAsks(), newBook.GetAsks())
}

// walk down the book depth generating insert/delete/modify diffs
func walk(diffs chan DepthDiff, side Side, oldOrders, newOrders <-chan Order) {
	old, oldOk := <-oldOrders
	new, newOk := <-newOrders

	for oldOk && newOk {
		if shouldInsertNew(side, old, new) {
			diffs <- insert(new, side)
			new, newOk = <-newOrders
		} else if shouldRemoveOld(side, old, new) {
			diffs <- remove(old, Bid)
			old, oldOk = <-oldOrders
		} else {
			if old.GetVolume() != new.GetVolume() {
				diffs <- insert(old, Bid)
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

func shouldInsertNew(side Side, old, new Order) bool {
	if side == Bid {
		return new.GetPrice() > old.GetPrice()
	} else {
		return new.GetPrice() < old.GetPrice()
	}
}

func shouldRemoveOld(side Side, old, new Order) bool {
	if side == Bid {
		return old.GetPrice() > new.GetPrice()
	} else {
		return old.GetPrice() < new.GetPrice()
	}
}

func remove(order Order, side Side) *diff {
	diff := makeDiff(order, side)
	diff.volume = 0
	return diff
}

func insert(order Order, side Side) *diff {
	return makeDiff(order, side)
}

func update(order Order, side Side) *diff {
	return makeDiff(order, side)
}

func makeDiff(order Order, side Side) *diff {
	return &diff{
		exchange: order.GetExchange(),
		symbol:   order.GetSymbol(),
		volume:   order.GetVolume(),
		price:    order.GetPrice(),
		side:     side,
	}
}

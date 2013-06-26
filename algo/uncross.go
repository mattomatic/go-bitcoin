package algo

import (
	"container/list"
	"github.com/mattomatic/go-bitcoin/common"
)

func Uncross(bids, asks <-chan common.Order) []Pair {
	pairs := list.New()

	bid, bidOk, bidLimit := next(bids)
	ask, askOk, askLimit := next(asks)

	for bidOk && askOk && bid.GetPrice() > ask.GetPrice() {
		volume := min(bidLimit, askLimit)
		bidLimit -= volume
		askLimit -= volume

		addPairs(pairs, volume, bid, ask)

		if bidLimit == 0 {
			bid, bidOk, bidLimit = next(bids)
		}

		if askLimit == 0 {
			ask, askOk, askLimit = next(asks)
		}
	}

	// spin off rest of items in the channel
	for bidOk || askOk {
		bid, bidOk, bidLimit = next(bids)
		ask, askOk, askLimit = next(asks)
	}

	return toSlice(pairs)
}

func toSlice(pairs *list.List) []Pair {
	slice := make([]Pair, pairs.Len())
	index := 0
	for o := pairs.Front(); o != nil; o = o.Next() {
		slice[index] = o.Value.(Pair)
		index++
	}
	return slice
}

func addPairs(pairs *list.List, limit common.Volume, bid, ask common.Order) {
	buy := getOpposingOrder(ask, limit)
	sell := getOpposingOrder(bid, limit)
	credit := getCredit(buy, sell)
	pair := Pair{Buy: buy, Sell: sell, Credit: credit}
	pairs.PushBack(pair)
}

func getCredit(buy, sell common.Order) common.Amount {
	volume := common.Amount(buy.GetVolume())
	credit := common.Amount(sell.GetPrice() - buy.GetPrice())
	return volume * credit
}

func getOpposingOrder(order common.Order, volume common.Volume) common.Order {
	return common.NewOrder(order.GetExchange(), order.GetSymbol(), volume, order.GetPrice(), getOpposingSide(order.GetSide()))
}

func getOpposingSide(side common.Side) common.Side {
	if side == common.Bid {
		return common.Ask
	}
	return common.Bid
}

func min(lhs, rhs common.Volume) common.Volume {
	if lhs < rhs {
		return lhs
	}

	return rhs
}

func next(orders <-chan common.Order) (common.Order, bool, common.Volume) {
	order, ok := <-orders

	if !ok { // no more orders on this side
		return order, ok, common.Volume(0)
	}

	return order, ok, order.GetVolume()
}

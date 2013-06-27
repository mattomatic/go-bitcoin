package algo

import (
	"github.com/mattomatic/go-bitcoin/common"
)

func Uncross(bids, asks <-chan common.Order) []Pair {
	pairs := make([]Pair, 0)

	bid, bidOk, bidLimit := next(bids)
	ask, askOk, askLimit := next(asks)

	for bidOk && askOk && common.GetFeeAdjustedPrice(bid) > common.GetFeeAdjustedPrice(ask) {
		volume := min(bidLimit, askLimit)
		bidLimit -= volume
		askLimit -= volume

		pairs = addPairs(pairs, volume, bid, ask)

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

	return pairs
}

func addPairs(pairs []Pair, limit common.Volume, bid, ask common.Order) []Pair {
	credit := getCredit(bid, ask, limit)
	cost := getCost(bid, ask, limit)
	roi := credit / cost
	buy := getOpposingOrder(ask, limit)
	sell := getOpposingOrder(bid, limit)
	pair := Pair{Buy: buy, Sell: sell, Credit: credit, Cost: cost, Roi: roi}
	return append(pairs, pair)
}

// Calculate the expense for hitting the bid and lifting the offer
func getCost(bid, ask common.Order, limit common.Volume) common.Amount {
	return common.Amount(limit) * common.Amount(common.GetFeeAdjustedPrice(ask))
}

// Calculate credit for hitting the bid and lifting the offer for the limit amount
func getCredit(bid, ask common.Order, limit common.Volume) common.Amount {
	credit := common.Amount(common.GetFeeAdjustedPrice(bid) - common.GetFeeAdjustedPrice(ask))
	return common.Amount(limit) * credit
}

func getOpposingOrder(order common.Order, volume common.Volume) common.Order {
	return common.NewOrder(
		order.GetExchange(),
		order.GetSymbol(),
		volume,
		order.GetPrice(),
		order.GetFee(),
		getOpposingSide(order.GetSide()))
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

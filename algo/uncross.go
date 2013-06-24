package algo

import (
	"github.com/mattomatic/go-bitcoin/common"
)

func Uncross(book common.OrderBook) float64 {
	return uncross(book.GetBids(), book.GetAsks())
}

func uncross(bids, asks <-chan common.Order) float64 {
	bid, bidOk := <-bids
	ask, askOk := <-asks
	credit := 0.0

	for {
		if !bidOk || !askOk {
			break
		}
		if bid.GetPrice() < ask.GetPrice() {
			break
		}

		if bid.GetVolume() < ask.GetVolume() {
			// buy bidvolume at offer
			// sell bidvolume at bid
			// collect difference
			// move to next bid
			credit += float64(bid.GetVolume()) * float64(bid.GetPrice()-ask.GetPrice())
			bid, bidOk = <-bids
		} else {
			credit += float64(ask.GetVolume()) * float64(bid.GetPrice()-ask.GetPrice())
			ask, askOk = <-asks
		}
	}

	// spin off rest of items in the channel
	for bidOk || askOk {
		bid, bidOk = <-bids
		ask, askOk = <-asks
	}

	return credit
}

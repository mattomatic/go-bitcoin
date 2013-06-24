package main

import (
	"fmt"
	"github.com/mattomatic/go-bitcoin/bitstamp"
	"github.com/mattomatic/go-bitcoin/btce"
	"github.com/mattomatic/go-bitcoin/common"
	"github.com/mattomatic/go-bitcoin/mtgox"
	"time"
)

func multiplex(output chan<- common.DepthDiff, input <-chan common.DepthDiff) {
	for diff := range input {
		output <- diff
	}
}

func printBook(book common.OrderBook, depth int) {
	bids := book.GetBids()
	asks := book.GetAsks()
	bid, bidOk := <-bids
	ask, askOk := <-asks

	for depth >= 0 && bidOk && askOk {
		fmt.Println(common.OrderString(bid), "---", common.OrderString(ask))
		bid, bidOk = <-bids
		ask, askOk = <-asks
		depth--
	}

	for bidOk || askOk {
		bid, bidOk = <-bids
		ask, askOk = <-asks
	}
}

func main() {
	diffs := make(chan common.DepthDiff)
	go multiplex(diffs, btce.GetDepthDiffChannel())
	go multiplex(diffs, mtgox.GetDepthDiffChannel())
	go multiplex(diffs, bitstamp.GetDepthDiffChannel())

	book := common.NewBook()

	for diff := range diffs {
		book.ApplyDiff(diff)
		printBook(book, 20)
		fmt.Println(time.Now(), "---------------->", common.DepthDiffString(diff))
	}
}

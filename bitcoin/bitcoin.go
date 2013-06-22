package main

import (
	"fmt"
	"github.com/mattomatic/go-bitcoin/common"
	"github.com/mattomatic/go-bitcoin/mtgox"
)

func main() {
	var gx common.Client = mtgox.NewClient()
	gx.ToggleOrderBookFeeds()
	gx.ToggleTickerFeeds()
	gx.ToggleAsync()
	gxchan := gx.Channel()

	for {
		select {
		case gxfeed := <-gxchan:
			if gxfeed.GetType() == common.TickerFeed {
				fmt.Println("ticker", common.TickerString(gxfeed.GetMessage().(common.Ticker)))
			}
			if gxfeed.GetType() == common.OrderBookFeed {
				book := gxfeed.GetMessage().(common.OrderBook)
				for _, bid := range book.GetBids() {
					fmt.Println("bid", common.OrderString(bid))
				}

				for _, ask := range book.GetAsks() {
					fmt.Println("ask", common.OrderString(ask))
				}
			}

			fmt.Println("----------------------")
		}
	}
}

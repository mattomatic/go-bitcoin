package main

import (
	"fmt"
	"github.com/mattomatic/go-bitcoin/bitstamp"
	"github.com/mattomatic/go-bitcoin/common"
)

func print(book common.OrderBook) {
	for bid := range book.GetBids() {
		fmt.Println("bid", common.OrderString(bid))
	}

	for ask := range book.GetAsks() {
		fmt.Println("ask", common.OrderString(ask))
	}

}

func main() {
	client := bitstamp.NewClient()
	client.ToggleOrderBookFeeds()
	client.ToggleAsync()
	feeds := client.Channel()

	for {
		select {
		case feed := <-feeds:
			if feed.GetType() == common.OrderBookFeed {
				book := feed.GetMessage().(common.OrderBook)
				print(book)
			}

			fmt.Println("----------------------")
		}
	}
}

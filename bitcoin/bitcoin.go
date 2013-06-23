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
	books := bitstamp.GetOrderBookChannel()

	for book := range books {
		print(book)
		fmt.Println("----------")
	}
}

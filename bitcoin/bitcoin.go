package main

import (
	"flag"
	"fmt"
	"github.com/mattomatic/go-bitcoin/mtgox"
	"github.com/mattomatic/go-bitcoin/common"
)

func init() {
	flag.Parse()
}

func main() {
	a := mtgox.NewClient("ws://websocket.mtgox.com:80")
	a.ToggleTradeFeeds()
	afeeds := a.Feeds()

	b := mtgox.NewClient("ws://websocket.mtgox.com:80")
	b.ToggleTradeFeeds()
	bfeeds := b.Feeds()

	var afeed, bfeed *mtgox.Feed

	for {
		select {
		case afeed = <-afeeds:
		    trade := afeed.Message.(common.Trade)
		    fmt.Println("trade:", trade)
			fmt.Println("afeed:", afeed.Message)
		case bfeed = <-bfeeds:
		    trade := bfeed.Message.(common.Trade)
		    fmt.Println("trade:", trade)
			fmt.Println("bfeed:", bfeed.Message)
		}
	}
}

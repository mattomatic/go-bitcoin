package main

import (
	"flag"
	"fmt"
	"github.com/mattomatic/go-bitcoin/mtgox"
)

func init() {
	flag.Parse()
}

//=======================

//=======================

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
			fmt.Println("afeed:", afeed.Message)
		case bfeed = <-bfeeds:
			fmt.Println("bfeed:", bfeed.Message)
		}
	}
}

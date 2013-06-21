package main

import (
	"flag"
	"fmt"
	"github.com/mattomatic/go-bitcoin/mtgox"
	"github.com/mattomatic/go-bitcoin/common"
	"time"
)

func init() {
	flag.Parse()
}

func main() {
	a := mtgox.NewClient("ws://websocket.mtgox.com:80")
	a.ToggleTradeFeeds()
	//a.ToggleDepthFeeds()
	//a.ToggleTickerFeeds()
	afeeds := a.Feeds()

	b := mtgox.NewClient("ws://websocket.mtgox.com:80")
	b.ToggleTradeFeeds()
	//b.ToggleDepthFeeds()
	//b.ToggleTickerFeeds()
	bfeeds := b.Feeds()

	var afeed, bfeed *mtgox.Feed

	for {
		select {
		case afeed = <-afeeds:
			fmt.Println("diff:", time.Now().Sub(afeed.Timestamp))
			fmt.Println("trade:", common.TradeString(afeed.Message.(common.Trade)))
		case bfeed = <-bfeeds:
			fmt.Println("diff:", time.Now().Sub(bfeed.Timestamp))
			fmt.Println("bfeed:", bfeed.Timestamp.String(), bfeed.Message)
			fmt.Println("trade:", common.TradeString(bfeed.Message.(common.Trade)))
		}
	}
}

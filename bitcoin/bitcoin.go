package main

import (
	"flag"
	"fmt"
	"github.com/mattomatic/go-bitcoin/bitstamp"
	"github.com/mattomatic/go-bitcoin/mtgox"
)

func init() {
	flag.Parse()
}

func main() {
	bs := bitstamp.NewClient()
	bs.ToggleTickerFeeds()
	bs.ToggleAsync()

	gx := mtgox.NewClient()
	gx.ToggleTickerFeeds()
	gx.ToggleAsync()

	bsfeeds := bs.Channel()
	gxfeeds := gx.Channel()

	for {
		select {
		case bsfeed := <-bsfeeds:
			fmt.Println("bs", bsfeed)
		case gxfeed := <-gxfeeds:
			fmt.Println("gx", gxfeed)
		}
	}
}

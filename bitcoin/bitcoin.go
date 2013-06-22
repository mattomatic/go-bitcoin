package main

import (
	"fmt"
	"github.com/mattomatic/go-bitcoin/bitstamp"
	"github.com/mattomatic/go-bitcoin/common"
)

func main() {
    var bs common.Client = bitstamp.NewClient()
    bs.ToggleOrderBookFeeds()
    bs.ToggleAsync()
    
    for feed := range bs.Channel() {
        fmt.Println(feed)
    }
}

package main

import (
	"fmt"
	"github.com/mattomatic/go-bitcoin/bitstamp"
	"github.com/mattomatic/go-bitcoin/campbx"
	"github.com/mattomatic/go-bitcoin/common"
	"github.com/mattomatic/go-bitcoin/mtgox"
)

func main() {
	campbxDiffs := campbx.GetDepthDiffChannel()
	bitstampDiffs := bitstamp.GetDepthDiffChannel()
	mtgoxDiffs := mtgox.GetDepthDiffChannel()

	for {
		select {
		case a := <-campbxDiffs:
			fmt.Println(common.DepthDiffString(a))
		case b := <-bitstampDiffs:
			fmt.Println(common.DepthDiffString(b))
		case c := <-mtgoxDiffs:
			fmt.Println(common.DepthDiffString(c))
		}
	}
}

package main

import (
	"fmt"
	"github.com/mattomatic/go-bitcoin/bitstamp"
	"github.com/mattomatic/go-bitcoin/btce"
	"github.com/mattomatic/go-bitcoin/campbx"
	"github.com/mattomatic/go-bitcoin/common"
	"github.com/mattomatic/go-bitcoin/mtgox"
)

func multiplex(output chan<- common.DepthDiff, input <-chan common.DepthDiff) {
	for diff := range input {
		output <- diff
	}
}

func main() {
	diffs := make(chan common.DepthDiff)
	go multiplex(diffs, btce.GetDepthDiffChannel())
	go multiplex(diffs, mtgox.GetDepthDiffChannel())
	go multiplex(diffs, campbx.GetDepthDiffChannel())
	go multiplex(diffs, bitstamp.GetDepthDiffChannel())

	for diff := range diffs {
		fmt.Println(common.DepthDiffString(diff))
	}
}

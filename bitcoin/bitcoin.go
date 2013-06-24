package main

import (
	"fmt"
	"github.com/mattomatic/go-bitcoin/btce"
	"github.com/mattomatic/go-bitcoin/common"
)

func main() {
	diffs := btce.GetDepthDiffChannel()

	for {
		select {
		case diff := <-diffs:
			fmt.Println(common.DepthDiffString(diff))
		}
	}
}

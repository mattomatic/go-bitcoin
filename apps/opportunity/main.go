// this program displays a list of opportunities up to a specified ROI
package main

import (
	"flag"
	"fmt"
	"github.com/mattomatic/go-bitcoin/common"
	"github.com/mattomatic/go-bitcoin/strategy"
	"time"
)

func main() {
	roi := flag.Float64("roi", 0.00, "the return on investment to display")
	max := flag.Int("max", 200, "the maximum number of opportunities to display")
	flag.Parse()

	pairs := strategy.GetFilteredPairsChannel(common.Amount(*roi))

	for pair := range pairs {
		for i, p := range pair {
		    if i < *max {
			    fmt.Println(p.String())
			}
		}
		fmt.Println(">>>", time.Now())
	}
}

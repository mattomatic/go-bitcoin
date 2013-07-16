// This program displays the combined order book for multiple exchanges
// up to a specified depth limit
package main

import (
	"flag"
	"fmt"
	"github.com/mattomatic/go-bitcoin/bbo"
	"github.com/mattomatic/go-bitcoin/common"
	"time"
)

func main() {
	depth := flag.Int("depth", 15, "order book depth to display")
	flag.Parse()

	books := bbo.GetCombinedBookChannel()

	for book := range books {
		common.PrintBook(book, *depth)
		fmt.Println(">>>", time.Now())
	}
}

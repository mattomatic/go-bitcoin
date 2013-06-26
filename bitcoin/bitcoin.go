package main

import (
	"fmt"
	"github.com/mattomatic/go-bitcoin/algo"
	"github.com/mattomatic/go-bitcoin/bbo"
	"github.com/mattomatic/go-bitcoin/common"
)

func main() {
	books := bbo.GetCombinedBookChannel()

	for book := range books {
		common.PrintBook(book, 20)
		pairs := algo.Uncross(book.GetBids(), book.GetAsks())

		fmt.Println()
		for _, pair := range pairs {
			fmt.Println(pair.String())
		}

		fmt.Println("--------------------")

	}
}

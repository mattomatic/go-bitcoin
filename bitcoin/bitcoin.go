package main

import (
	"fmt"
	"github.com/mattomatic/go-bitcoin/bitstamp"
)

func main() {
	book := bitstamp.GetOrderBook()
	fmt.Println(book)
}

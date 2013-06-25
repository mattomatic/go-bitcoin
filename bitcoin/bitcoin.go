package main

import (
	"fmt"
	"github.com/mattomatic/go-bitcoin/bbo"
	"github.com/mattomatic/go-bitcoin/common"
	"github.com/mattomatic/go-bitcoin/algo"
	"time"
)

func main() {
    books := bbo.GetCombinedBookChannel()
    
    for book := range books {
        common.PrintBook(book, 20)
        fmt.Println(time.Now(), "------------>", algo.Uncross(book))
    }
}

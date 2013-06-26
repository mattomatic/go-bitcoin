package main

import (
	"fmt"
	"github.com/mattomatic/go-bitcoin/algo"
	"github.com/mattomatic/go-bitcoin/bbo"
	"time"
)

func main() {
	books := bbo.GetCombinedBookChannel()
    ticks := time.Tick(time.Second)
    
    book := <-books
    now := <-ticks
    
    for {
        select {
        case book = <-books:
        case now = <-ticks:
            pairs := algo.Uncross(book.GetBids(), book.GetAsks())
            
            for i, pair := range pairs {
                fmt.Println(i, pair.String())
            }           
            fmt.Println("----------->", now)
        }
    }
}

func min(a, b int) int {
    if a < b {
        return a
    }
    
    return b
}
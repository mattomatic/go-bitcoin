package bbo

import (
	"github.com/mattomatic/go-bitcoin/bitstamp"
	"github.com/mattomatic/go-bitcoin/btce"
	"github.com/mattomatic/go-bitcoin/campbx"
	"github.com/mattomatic/go-bitcoin/common"
	"github.com/mattomatic/go-bitcoin/mtgox"
)

// Get a channel that contains diffs from all the exchanges
func GetCombinedDiffsChannel() <-chan common.DepthDiff {
    diffs := make(chan common.DepthDiff)
    go pipe(btce.GetDepthDiffChannel(), diffs)
	go pipe(mtgox.GetDepthDiffChannel(), diffs)
	go pipe(bitstamp.GetDepthDiffChannel(), diffs)
	go pipe(campbx.GetDepthDiffChannel(), diffs)
	return diffs
}

// Get a channel that contains the combined book of all the exchanges
func GetCombinedBookChannel() <-chan common.OrderBook {
    books := make(chan common.OrderBook)
    diffs := GetCombinedDiffsChannel()
    book := common.NewBook()
    
    go func() {
        defer close(books)
        
        for diff := range diffs {
            book.ApplyDiff(diff)
            books <- book
        }
    }()
    
    return books
}

// pipe everything from the input channel to the output channel
func pipe(input <-chan common.DepthDiff, output chan<- common.DepthDiff) {
	for diff := range input {
		output <- diff
	}
}
package bbo

import (
	"github.com/mattomatic/go-bitcoin/bitstamp"
	"github.com/mattomatic/go-bitcoin/btce"
	"github.com/mattomatic/go-bitcoin/campbx"
	"github.com/mattomatic/go-bitcoin/common"
	"github.com/mattomatic/go-bitcoin/mtgox"
)

// Get a channel that contains diffs from all the exchanges
func GetCombinedDiffsChannel() <-chan common.Diff {
	diffs := make(chan common.Diff)
	go pipe(btce.GetDiffChannel(), diffs)
	go pipe(mtgox.GetDiffChannel(), diffs)
	go pipe(bitstamp.GetDiffChannel(), diffs)
	go pipe(campbx.GetDiffChannel(), diffs)
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
func pipe(input <-chan common.Diff, output chan<- common.Diff) {
	for diff := range input {
		output <- diff
	}
}

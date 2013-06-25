package common

import "fmt"

func TradeString(t Trade) string {
	return fmt.Sprintf("%s:%s %f@%f", t.GetExchange(), t.GetSymbol(), t.GetVolume(), t.GetPrice())
}

func TickerString(t Ticker) string {
	return fmt.Sprintf("%s:%s %f -- %f", t.GetExchange(), t.GetSymbol(), t.GetBid(), t.GetAsk())
}

func OrderString(o Order) string {
	return fmt.Sprintf("%s:%s %f @ %f", o.GetExchange(), o.GetSymbol(), o.GetVolume(), o.GetPrice())
}

func DiffString(d Diff) string {
	return fmt.Sprintf("%s:%s %s %f @ %f", d.GetExchange(), d.GetSymbol(), d.GetSide(), d.GetVolume(), d.GetPrice())
}

// Print a representation of the book to a certain depth
func PrintBook(book OrderBook, depth int) {
	bids := book.GetBids()
	asks := book.GetAsks()

	bid, bidOk := <-bids
	ask, askOk := <-asks

	for depth >= 0 && bidOk && askOk {
		fmt.Println(OrderString(bid), "---", OrderString(ask))
		bid, bidOk = <-bids
		ask, askOk = <-asks
		depth--
	}

	// clean out the rest of the channel
	for bidOk || askOk {
		bid, bidOk = <-bids
		ask, askOk = <-asks
	}
}

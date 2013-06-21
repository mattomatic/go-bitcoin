package common

import "fmt"

func TradeString(t Trade) string {
	return fmt.Sprintf("%s:%s %f@%f", t.GetExchange(), t.GetSymbol(), t.GetVolume(), t.GetPrice())
}

func TickerString(t Ticker) string {
	return fmt.Sprintf("%s:%s %f -- %f", t.GetExchange(), t.GetSymbol(), t.GetBid(), t.GetAsk())
}

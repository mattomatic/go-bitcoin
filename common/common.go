package common

import "fmt"

func TradeString(t Trade) string {
	return fmt.Sprintf("%s:%s %f@%f", t.Exchange(), t.Symbol(), t.Volume(), t.Price())
}

func TickerString(t Ticker) string {
	return fmt.Sprintf("%s:%s %f -- %f", t.Exchange(), t.Symbol(), t.Bid(), t.Ask())
}

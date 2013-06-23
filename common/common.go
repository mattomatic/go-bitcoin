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

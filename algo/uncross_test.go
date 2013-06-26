package algo

import (
	"github.com/mattomatic/go-bitcoin/common"
	"testing"
)

func TestUncrossNoOrders(t *testing.T) {
	bids := getBids(0, "FOO")
	asks := getAsks(0, "BAR")

	pairs := Uncross(bids, asks)

	if len(pairs) != 0 {
		t.Error()
	}
}

func TestUncrossOnce(t *testing.T) {
	bids := getBids(1, "FOO")
	asks := getAsks(1, "BAR")

	pairs := Uncross(bids, asks)

	// bids and asks here both have the same price, so nothing should
	// be in cross.
	if len(pairs) != 0 {
		t.Error()
	}
}

func TestUncrossTwice(t *testing.T) {
	bids := getBids(2, "FOO")
	asks := getAsks(2, "BAR")

	pairs := Uncross(bids, asks)

	if len(pairs) != 1 {
		t.Error()
	}

	pair := pairs[0]

	if pair.Buy.GetSide() != common.Bid {
		t.Error()
	}

	if pair.Sell.GetSide() != common.Ask {
		t.Error()
	}

	if pair.Buy.GetExchange() != "BAR" || pair.Buy.GetPrice() != 1 {
		t.Error()
	}

	if pair.Sell.GetExchange() != "FOO" || pair.Sell.GetPrice() != 2 {
		t.Error()
	}

	if pair.Credit != 100 { // 100 volume * $1 difference in price
		t.Error()
	}
}

func getBids(n int, exchange common.Exchange) chan common.Order {
	orders := make(chan common.Order, n)
	for i := n; i > 0; i-- {
		orders <- common.NewOrder(exchange, "BTC", 100, common.Price(i), common.Bid)
	}
	close(orders)
	return orders
}

func getAsks(n int, exchange common.Exchange) chan common.Order {
	orders := make(chan common.Order, n)
	for i := 1; i <= n; i++ {
		orders <- common.NewOrder(exchange, "BTC", 100, common.Price(i), common.Ask)
	}
	close(orders)
	return orders
}

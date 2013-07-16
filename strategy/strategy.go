package strategy

import (
	"fmt"
	"github.com/mattomatic/go-bitcoin/algo"
	"github.com/mattomatic/go-bitcoin/bbo"
	"github.com/mattomatic/go-bitcoin/common"
	"time"
)

func GetPairsChannel() <-chan []algo.Pair {
	pairs := make(chan []algo.Pair)

	go func() {
		defer close(pairs)
		for book := range bbo.GetCombinedBookChannel() {
			pairs <- algo.Uncross(book.GetBids(), book.GetAsks())
		}
	}()

	return pairs
}

func GetFilteredPairsChannel(roi common.Amount) <-chan []algo.Pair {
	pairs := make(chan []algo.Pair)

	go func() {
		defer close(pairs)

		for pair := range GetPairsChannel() {
			pairs <- FilterPairs(pair, roi)
		}
	}()

	return pairs
}

func FilterPairs(pairs []algo.Pair, roi common.Amount) []algo.Pair {
	filtered := make([]algo.Pair, 0)
	for _, p := range pairs {
		if p.Roi >= roi {
			filtered = append(filtered, p)
		}
	}
	return filtered
}

func GetExchangeHaveChannel() <-chan common.ExchangeBudget {
	budgets := make(chan common.ExchangeBudget)

	go func() {
		defer close(budgets)

		for {
			time.Sleep(1 * time.Second)
			budgets <- common.NewExchangeBudget()
		}
	}()

	return budgets
}

func GetExchangeWantChannel(roi common.Amount) <-chan common.ExchangeBudget {
	budgets := make(chan common.ExchangeBudget)

	go func() {
		defer close(budgets)

		for pair := range GetFilteredPairsChannel(roi) {
			budgets <- GetExchangeBudget(pair)
		}
	}()

	return budgets
}

func GetExchangeBudget(pairs []algo.Pair) common.ExchangeBudget {
	budget := common.NewExchangeBudget()

	for _, pair := range pairs {
		// add the buy order cost to the USD budget for that exchange
		budget.Get(pair.Buy.GetExchange()).USD += common.Volume(pair.Cost)

		// add the sell order volume to the BTC budget for that exchange
		budget.Get(pair.Sell.GetExchange()).BTC += pair.Sell.GetVolume()
	}

	return budget
}

func GetExchangeBudgetDiffChannel(roi common.Amount) <-chan common.ExchangeBudget {
	diffs := make(chan common.ExchangeBudget)
	go getExchangeBudgetDiffs(roi, diffs)
	return diffs
}

func getExchangeBudgetDiffs(roi common.Amount, diffs chan common.ExchangeBudget) {
	defer close(diffs)

	haves := GetExchangeHaveChannel()
	wants := GetExchangeWantChannel(roi)

	have := <-haves
	want := <-wants

	for {
		select {
		case have = <-haves:
			diffs <- want.Subtract(have)
		case want = <-wants:
			diffs <- want.Subtract(have)
		}
	}
}

func Run(roi common.Amount) {
	pairs := GetFilteredPairsChannel(roi)

	for pair := range pairs {
		for _, p := range pair {
		        fmt.Println(p.String())
		}
	fmt.Println("-------", time.Now())
	}
}

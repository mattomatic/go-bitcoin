package common

type ExchangeBudget map[Exchange]*Budget

func NewExchangeBudget() ExchangeBudget {
	return make(ExchangeBudget)
}

// Get the budget for a particular exchange or make a new one if none exist
func (e ExchangeBudget) Get(ex Exchange) (b *Budget) {
	if _, ok := e[ex]; !ok {
		e[ex] = &Budget{}
	}

	return e[ex]
}

// The budget type represents either:
// 1. the amount of each currency we have (in an account)
// 2. the amount of each currency we need (to execute our trades)
type Budget struct {
	USD Volume
	BTC Volume
}

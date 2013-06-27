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

func (e ExchangeBudget) Subtract(o ExchangeBudget) ExchangeBudget {
    exchangeBudget := NewExchangeBudget()
    
    for _, exchange := range unionKeys(e, o) {
        budget := exchangeBudget.Get(exchange)
        budget.USD = e.Get(exchange).USD - o.Get(exchange).USD
        budget.BTC = e.Get(exchange).BTC - o.Get(exchange).BTC
    }
    
    return exchangeBudget
}

func unionKeys(a, b ExchangeBudget) []Exchange {
    m := make(map[Exchange]bool)
    
    for k, _ := range a {
        m[k] = true
    }
    
    for k, _ := range b {
        m[k] = true
    }
    
    keys := make([]Exchange, 0)
    
    for k, _ := range m {
        keys = append(keys, k)
    }
    
    return keys
}

// The budget type represents either:
// 1. the amount of each currency we have (in an account)
// 2. the amount of each currency we need (to execute our trades)
type Budget struct {
	USD Volume
	BTC Volume
}

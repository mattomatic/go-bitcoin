package mtgox

import (
	"github.com/mattomatic/go-bitcoin/common"
	"time"
)

type Trade struct {
	AmountInt     string `json:"amount_int"`
	Date          int    `json:"date"`
	Item          string `json:"item"`
	PriceCurrency string `json:"price_currency"`
	PriceInt      string `json:"price_int"`
	Primary       string `json:"primary"`
	Properties    string `json:"properties"`
	Tid           string `json:"tid"`
	TradeType     string `json:"bid"`
	Type          string `json:"type"`
}

func (t *Trade) GetExchange() common.Exchange { return common.Exchange(ExchangeId) }
func (t *Trade) GetSymbol() common.Symbol     { return common.Symbol(t.Item) }
func (t *Trade) GetTimestamp() time.Time      { return time.Unix(int64(t.Date), 0) }
func (t *Trade) GetVolume() common.Volume     { return getVolume(t.AmountInt) }
func (t *Trade) GetPrice() common.Price       { return getPrice(t.PriceInt) }
func (t *Trade) GetCurrency() common.Currency { return common.Currency(t.PriceCurrency) }
func (t *Trade) GetTradeId() common.TradeId   { return common.TradeId(t.Tid) }

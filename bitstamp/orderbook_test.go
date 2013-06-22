package bitstamp

import (
	"github.com/mattomatic/go-bitcoin/common"
	"testing"
)

func TestOrderBookInterface(t *testing.T) {
	var _ common.OrderBook = &OrderBook{}
}

package bitstamp

import (
	"github.com/mattomatic/go-bitcoin/common"
	"testing"
)

func TestTickerInterface(t *testing.T) {
	var _ common.Ticker = &Ticker{}
}

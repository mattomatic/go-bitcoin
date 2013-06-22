package bitstamp

import (
	"github.com/mattomatic/go-bitcoin/common"
	"testing"
)

func TestOrderInterface(t *testing.T) {
	var _ common.Order = &Order{}
}

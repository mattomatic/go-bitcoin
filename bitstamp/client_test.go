package bitstamp

import (
	"github.com/mattomatic/go-bitcoin/common"
	"testing"
)

func TestClientInterface(t *testing.T) {
	var _ common.Client = &Client{}
}

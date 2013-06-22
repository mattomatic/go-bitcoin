package mtgox

import (
	"github.com/mattomatic/go-bitcoin/common"
	"testing"
)

func TestFeedInterface(t *testing.T) {
	var _ common.Feed = &Feed{}
}

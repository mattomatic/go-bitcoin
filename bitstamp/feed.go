package bitstamp

import (
	"github.com/mattomatic/go-bitcoin/common"
	"time"
)

type Feed struct {
	Type      common.FeedType
	Timestamp time.Time
	Message   interface{}
}

func (f *Feed) GetType() common.FeedType {
	return f.Type
}

func (f *Feed) GetTimestamp() time.Time {
	return f.Timestamp
}

func (f *Feed) GetMessage() interface{} {
	return f.Message
}

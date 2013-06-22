package mtgox

import (
	"encoding/json"
	"github.com/mattomatic/go-bitcoin/common"
	"time"
)

func (feed *Feed) UnmarshalJSON(bytes []byte) error {
	feed.Timestamp = time.Now()

	header := &Header{}
	err := json.Unmarshal(bytes, header)

	if err != nil {
		panic(err.Error())
	}

	feed.Type = getFeedType(header.Private)

	switch feed.Type {
	case common.DepthFeed:
		msg := &DepthFeed{}
		err = json.Unmarshal(bytes, msg)
		feed.Message = msg

	case common.TickerFeed:
		msg := &TickerFeed{}
		msg.Timestamp = feed.Timestamp
		err = json.Unmarshal(bytes, msg)
		feed.Message = msg

	case common.TradeFeed:
		msg := &TradeFeed{}
		err = json.Unmarshal(bytes, msg)
		feed.Message = msg

	default:
		panic("unrecognized feed type!")
	}

	if err != nil {
		panic(err.Error())
	}

	return err
}

package mtgox

import (
	"encoding/json"
	"time"
)

func (feed *Feed) UnmarshalJSON(bytes []byte) (err error) {
	feed.Timestamp = time.Now()

	header := &Header{}
	err = json.Unmarshal(bytes, header)

	if err != nil {
		panic(err.Error())
	}

	feed.Type = header.Private

	switch feed.Type {
	case "depth":
		msg := &DepthFeed{}
		err = json.Unmarshal(bytes, msg)
		feed.Message = msg
	case "ticker":
		msg := &TickerFeed{}
		err = json.Unmarshal(bytes, msg)
		feed.Message = msg
	case "trade":
		msg := &TradeFeed{}
		err = json.Unmarshal(bytes, msg)
		feed.Message = msg
	default:
		panic("unrecognized feed type!")
	}

	return err
}

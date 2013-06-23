package mtgox

import (
	"encoding/json"
)

func (feed *Feed) UnmarshalJSON(bytes []byte) error {
	feed.kind = getKind(bytes)
	feed.message = getMessage(bytes, feed.kind)
	return nil
}

func getKind(bytes []byte) string {
	header := &Header{}
	err := json.Unmarshal(bytes, header)

	if err != nil {
		panic(err.Error())
	}

	return header.Private
}

func getMessage(bytes []byte, kind string) interface{} {
	var msg interface{}
	var err error

	switch kind {
	case "depth":
		msg = &DepthFeed{}
		err = json.Unmarshal(bytes, msg)
	case "ticker":
		msg = &TickerFeed{}
		err = json.Unmarshal(bytes, msg)
	case "trade":
		msg = &TradeFeed{}
		err = json.Unmarshal(bytes, msg)
	default:
		panic("unrecognized feed type!")
	}

	if err != nil {
		panic(err.Error())
	}

	return msg
}

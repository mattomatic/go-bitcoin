package mtgox

type LoginReply struct {
	Op      string `json:"op"`
	Message string `json:"message"`
}

type TradeFeed struct {
	Header
	Trade `json:"trade"`
}

type DepthFeed struct {
	Header
	Depth `json:"depth"`
}

type TickerFeed struct {
	Header
	Ticker `json:"ticker"`
}

type Header struct {
	Channel string `json:"channel"`
	Op      string `json:"op"`
	Origin  string `json:"origin"`
	Private string `json:"private"`
}

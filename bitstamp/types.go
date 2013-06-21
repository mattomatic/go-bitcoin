package bitstamp

const (
	ExchangeId = "BITSTAMP"
)

type Ticker struct {
	Timestamp string `json:"timestamp"`
	Bid       string `json:"bid"`
	Ask       string `json:"ask"`
	High      string `json:"high"`
	Low       string `json:"low"`
	Last      string `json:"last"`
	Volume    string `json:"volume"`
}

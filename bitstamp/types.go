package bitstamp

const (
	ExchangeId = "BITSTAMP"
)

type Ticker struct {
	timestamp string `json:"timestamp"`
	bid       string `json:"bid"`
	ask       string `json:"ask"`
	high      string `json:"high"`
	low       string `json:"low"`
	last      string `json:"last"`
	volume    string `json:"volume"`
}

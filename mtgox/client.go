package mtgox

import (
	"code.google.com/p/go.net/websocket"
	"encoding/json"
	"github.com/mattomatic/go-bitcoin/common"
)

const (
	MtGoxUrl = "ws://websocket.mtgox.com:80"
	Protocol = "ws"
	LocalUrl = "ws://localhost"
)

type Client struct {
	conn    *websocket.Conn
	encoder *json.Encoder
	decoder *json.Decoder
	feeds   chan common.Feed
	book    *OrderBook
}

func NewClient() *Client {
	conn, err := websocket.Dial(MtGoxUrl, Protocol, LocalUrl)

	if err != nil {
		panic(err.Error())
	}

	encoder := json.NewEncoder(conn)
	decoder := json.NewDecoder(conn)
	feeds := make(chan common.Feed)
	book := newOrderBook()

	client := &Client{conn, encoder, decoder, feeds, book}

	reply := &LoginReply{}
	err = decoder.Decode(reply)

	if err != nil {
		panic(err.Error())
	}

    // Go async now
    go client.async()

	return client
}

func (c *Client) Channel() <-chan common.Feed {
	return c.feeds
}

func (c *Client) ToggleTickerFeeds() {
	c.toggle("ticker")
}

func (c *Client) ToggleTradeFeeds() {
	c.toggle("trades")
}

func (c *Client) ToggleOrderBookFeeds() {
	c.toggle("depth")
}

func (c *Client) toggle(feed string) {
	msg := map[string]string{"op": "mtgox.subscribe", "type": feed}
	err := c.encoder.Encode(&msg)

	if err != nil {
		panic(err.Error())
	}
}

func (c *Client) convertDepthToBookFeed(feed *Feed) {
	if feed.Type == common.DepthFeed {
		feed.Type = common.OrderBookFeed
		depth := &feed.Message.(*DepthFeed).Depth
		c.book.handleDepth(depth)
		feed.Message = c.book
	}
}

func (c *Client) async() {
	feed := &Feed{}

	for {
		err := c.decoder.Decode(&feed)

		if err != nil {
			panic(err.Error())
		}

		c.convertDepthToBookFeed(feed)
		c.feeds <- feed
	}
}

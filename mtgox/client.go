package mtgox

import (
	"code.google.com/p/go.net/websocket"
	"encoding/json"
)

type Client struct {
	conn    *websocket.Conn
	encoder *json.Encoder
	decoder *json.Decoder
	feeds   chan *Feed
}

func NewClient(url string) *Client {
	conn, err := websocket.Dial(url, "ws", "ws://localhost")

	if err != nil {
		panic(err.Error())
	}

	encoder := json.NewEncoder(conn)
	decoder := json.NewDecoder(conn)
	feeds := make(chan *Feed)

	client := &Client{conn, encoder, decoder, feeds}

	reply := &LoginReply{}
	err = decoder.Decode(reply)

	if err != nil {
		panic(err.Error())
	}

	go client.loop()

	return client
}

func (c *Client) Feeds() chan *Feed {
	return c.feeds
}

func (c *Client) ToggleTickerFeeds() {
	c.toggle("ticker")
}

func (c *Client) ToggleTradeFeeds() {
	c.toggle("trades")
}

func (c *Client) ToggleDepthFeeds() {
	c.toggle("depth")
}

func (c *Client) toggle(feed string) {
	msg := map[string]string{"op": "mtgox.subscribe", "type": feed}
	err := c.encoder.Encode(&msg)

	if err != nil {
		panic(err.Error())
	}
}

func (c *Client) loop() {
	feed := &Feed{}

	for {
		err := c.decoder.Decode(&feed)

		if err != nil {
			panic(err.Error())
		}

		c.feeds <- feed
	}
}

package mtgox

import (
	"code.google.com/p/go.net/websocket"
	"encoding/json"
)

const (
	MtGoxUrl = "ws://websocket.mtgox.com:80"
	Protocol = "ws"
	LocalUrl = "ws://localhost"
)

type Connection struct {
	encoder *json.Encoder
	decoder *json.Decoder
}

type Feed struct {
	kind    string
	message interface{}
}

func connect() *Connection {
	conn, err := websocket.Dial(MtGoxUrl, Protocol, LocalUrl)

	if err != nil {
		panic(err.Error())
	}

	encoder := json.NewEncoder(conn)
	decoder := json.NewDecoder(conn)

	reply := &LoginReply{}
	err = decoder.Decode(reply)

	if err != nil {
		panic(err.Error())
	}

	return &Connection{encoder, decoder}
}

func (c *Connection) subscribe(feed string) {
	msg := map[string]string{"op": "mtgox.subscribe", "type": feed}
	err := c.encoder.Encode(&msg)

	if err != nil {
		panic(err.Error())
	}
}

func (c *Connection) poll() chan *Feed {
	feeds := make(chan *Feed)

	go func() {
		feed := &Feed{}

		for {
			err := c.decoder.Decode(&feed)

			if err != nil {
				panic(err.Error())
			}

			feeds <- feed
		}
	}()

	return feeds
}

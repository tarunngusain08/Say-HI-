package contracts

import "github.com/gorilla/websocket"

type Message struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
}

type Client struct {
	Username string
	Conn     *websocket.Conn
}

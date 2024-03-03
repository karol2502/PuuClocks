package sockets

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Client struct {
	Socket *websocket.Conn

	Receive chan []byte

	Lobby *Lobby
}

func (c *Client) Read() {
	defer c.Socket.Close()
	for {
		_, msg, err := c.Socket.ReadMessage()
		fmt.Println("Received message: ", string(msg))
		if err != nil {
			_ = fmt.Errorf("there was a error when reading message for Client: %w", err)
			return
		}
		c.Lobby.Forward <- msg
	}
}

func (c *Client) Write() {
	defer c.Socket.Close()
	for msg := range c.Receive {
		fmt.Println("Writting message: ", string(msg))
		err := c.Socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}

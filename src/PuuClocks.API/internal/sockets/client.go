package sockets

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client interface {
	ReceiveMessage([]byte)
	Close()
}

type client struct {
	ID uuid.UUID
	Socket *websocket.Conn

	Receive chan []byte

	Lobby Lobby
}

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func NewClient(conn *websocket.Conn, l Lobby) Client {
	c := &client{
		ID: uuid.New(),
		Socket:  conn,
		Receive: make(chan []byte, Upgrader.ReadBufferSize),
		Lobby:   l,
	}

	l.JoinLobby(c)

	defer func() {
		l.LeaveLobby(c)
	}()

	go c.Write()
	c.Read()

	return c
}

func (c *client) Read() {
	defer c.Socket.Close()
	for {
		_, msg, err := c.Socket.ReadMessage()
		fmt.Println("Received message: ", string(msg))
		if err != nil {
			_ = fmt.Errorf("there was a error when reading message for Client: %w", err)
			return
		}
		c.Lobby.ForwardMessage(Message{
			SocketID: c.ID,
			Data: string(msg),
		})
	}
}

func (c *client) Write() {
	defer c.Socket.Close()
	for msg := range c.Receive {
		fmt.Println("Writing message: ", string(msg))
		err := c.Socket.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			return
		}
	}
}

func (c *client) ReceiveMessage(msg []byte) {
	c.Receive <- msg
}

func (c *client) Close() {
	close(c.Receive)
}

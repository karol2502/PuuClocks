package sockets

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Lobby struct {
	Forward chan ([]byte)
	Join    chan *Client
	Leave   chan *Client

	Clients map[*Client]bool
}

func NewLobby() *Lobby {
	return &Lobby{
		Forward: make(chan []byte),
		Join:    make(chan *Client),
		Leave:   make(chan *Client),
		Clients: make(map[*Client]bool),
	}
}

func (l *Lobby) Run() {
	for {
		select {
		case client := <-l.Join:
			l.Clients[client] = true
		case client := <-l.Leave:
			delete(l.Clients, client)
			close(client.Receive)
		case msg := <-l.Forward:
			for c := range l.Clients {
				c.Receive <- msg
			}
		}
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (l *Lobby) JoinLobby(w http.ResponseWriter, r *http.Request){
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	c := &Client{
		Socket: conn,
		Receive: make(chan []byte, upgrader.ReadBufferSize),
		Lobby: l,
	}
	
	l.Join <- c

	defer conn.Close()
	defer func() {
		l.Leave <- c
	}()

	go c.Write()
	c.Read()
}
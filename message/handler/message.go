package handler

import (
	"Say-Hi/message/contracts"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

type MessageHandler struct {
	clients  map[*contracts.Client]bool
	mu       sync.Mutex
	upgrader websocket.Upgrader
}

func NewMessageHandler(upgrader websocket.Upgrader) *MessageHandler {
	return &MessageHandler{
		clients:  make(map[*contracts.Client]bool),
		mu:       sync.Mutex{},
		upgrader: upgrader,
	}
}

func (m *MessageHandler) SendMessage(w http.ResponseWriter, r *http.Request) {
	conn, err := m.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	user1 := r.URL.Query().Get("user1")
	user2 := r.URL.Query().Get("user2")
	client1 := &contracts.Client{Username: user1, Conn: conn}
	client2 := &contracts.Client{Username: user2, Conn: conn}

	m.mu.Lock()
	m.clients[client1] = true
	m.clients[client2] = true
	m.mu.Unlock()

	defer func() {
		m.mu.Lock()
		delete(m.clients, client1)
		delete(m.clients, client2)
		m.mu.Unlock()
	}()

	for {
		var msg contracts.Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
			return
		}

		m.mu.Lock()
		for c := range m.clients {
			if c.Username == msg.Recipient {
				if err := client1.Conn.WriteJSON(msg); err != nil {
					log.Println(err)
					return
				}
			}
		}
		m.mu.Unlock()
	}
}

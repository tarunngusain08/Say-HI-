package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// Message represents a chat message.
type Message struct {
	Sender    string `json:"sender"`
	Recipient string `json:"recipient"`
	Content   string `json:"content"`
	Timestamp string `json:"timestamp"`
}

// Client represents a connected client.
type Client struct {
	Username string
	Conn     *websocket.Conn
}

var (
	clients   = make(map[*Client]bool)
	clientsMu sync.Mutex
	upgrader  = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	username := r.URL.Query().Get("username")
	client := &Client{Username: username, Conn: conn}

	clientsMu.Lock()
	clients[client] = true
	clientsMu.Unlock()

	defer func() {
		clientsMu.Lock()
		delete(clients, client)
		clientsMu.Unlock()
	}()

	for {
		var msg Message
		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Println(err)
			return
		}

		// Handle your message logic here (e.g., store in database, send to recipient)
		// For simplicity, we'll broadcast the message to all connected clients.

		clientsMu.Lock()
		for c := range clients {
			if err := c.Conn.WriteJSON(msg); err != nil {
				log.Println(err)
				return
			}
		}
		clientsMu.Unlock()
	}
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	port := 8080
	fmt.Printf("Server is running on :%d\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

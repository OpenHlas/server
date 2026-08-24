package controllers

import (
	"encoding/json"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/gouef/web-project/models"
)

type NodeHub struct {
	mu      sync.RWMutex
	clients map[*websocket.Conn]struct{}
}

func NewNodeHub() *NodeHub {
	return &NodeHub{clients: make(map[*websocket.Conn]struct{})}
}

func (hub *NodeHub) Add(client *websocket.Conn) {
	hub.mu.Lock()
	defer hub.mu.Unlock()
	hub.clients[client] = struct{}{}
}

func (hub *NodeHub) Remove(client *websocket.Conn) {
	hub.mu.Lock()
	defer hub.mu.Unlock()
	delete(hub.clients, client)
}

func (hub *NodeHub) Broadcast(payload []byte) {
	hub.mu.RLock()
	clients := make([]*websocket.Conn, 0, len(hub.clients))
	for client := range hub.clients {
		clients = append(clients, client)
	}
	hub.mu.RUnlock()

	for _, client := range clients {
		if err := client.WriteMessage(websocket.TextMessage, payload); err != nil {
			hub.Remove(client)
			_ = client.Close()
		}
	}
}

type nodeEvent struct {
	Event string          `json:"event"`
	Data  json.RawMessage `json:"data"`
}

func (hub *NodeHub) HandleConnection(writer http.ResponseWriter, request *http.Request) {
	if request.Header.Get("Authorization") == "" {
		http.Error(writer, "Authorization is required.", http.StatusUnauthorized)
		return
	}

	upgrader := websocket.Upgrader{
		CheckOrigin: func(*http.Request) bool { return true },
	}
	client, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		return
	}
	defer client.Close()

	hub.Add(client)
	defer hub.Remove(client)

	for {
		_, payload, err := client.ReadMessage()
		if err != nil {
			return
		}

		var event nodeEvent
		if err := json.Unmarshal(payload, &event); err != nil || event.Event != "message_create" {
			continue
		}

		var message models.MessageCreatedData
		if err := json.Unmarshal(event.Data, &message); err != nil {
			continue
		}

		response, err := models.NewMessageCreatedEvent(message)
		if err != nil {
			continue
		}
		hub.Broadcast(response)
	}
}

package hub

import "github.com/wanyuqin/etcd-ui/backend/logger"

type Hub struct {
	Clients map[string]*Client

	Broadcast chan []byte

	Register chan *Client

	Unregister chan *Client
}

var H *Hub

func NewHub() *Hub {
	hub := &Hub{
		Broadcast:  make(chan []byte),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[string]*Client),
	}
	if H != nil {
		return H
	}
	H = hub
	return hub
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			if client.ID == "" {
				logger.Errorf("websocket client id is null")
				return
			}
			if _, ok := h.Clients[client.ID]; ok {
				logger.Errorf(" %s client register failed ID duplicate:", client.ID)
				return
			}

			h.Clients[client.ID] = client

		case client := <-h.Unregister:
			if _, ok := h.Clients[client.ID]; ok {
				delete(h.Clients, client.ID)
				close(client.Send)
			}
		case message := <-h.Broadcast:
			for id, client := range h.Clients {
				select {
				case client.Send <- message:
				default:
					close(client.Send)
					delete(h.Clients, id)
				}
			}

		}
	}
}

func (h *Hub) OnlineClientCount() int {
	return len(h.Clients)
}

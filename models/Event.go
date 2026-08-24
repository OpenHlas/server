package models

import "encoding/json"

type Event struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

type MessageCreatedData struct {
	ID        string `json:"id"`
	ServerID  string `json:"server_id"`
	ChannelID string `json:"channel_id"`
	AuthorID  string `json:"author_id"`
	Author    string `json:"author_name"`
	Content   string `json:"content"`
	Timestamp int64  `json:"timestamp"`
}

func NewMessageCreatedEvent(message MessageCreatedData) ([]byte, error) {
	return json.Marshal(Event{
		Event: "message_created",
		Data:  message,
	})
}

package models
package models

import (
	"encoding/json"
	"testing"
)

func TestNewMessageCreatedEventMatchesClientEnvelope(t *testing.T) {
	payload, err := NewMessageCreatedEvent(MessageCreatedData{
		ID:        "msg_1",
		ServerID:  "srv_1",
		ChannelID: "general",
		AuthorID:  "usr_1",
		Author:    "Jan",
		Content:   "Hello",
		Timestamp: 123,
	})
	if err != nil {
		t.Fatalf("marshal event: %v", err)
	}

	var envelope struct {
		Event string                 `json:"event"`
		Data  MessageCreatedData `json:"data"`
	}
	if err := json.Unmarshal(payload, &envelope); err != nil {
		t.Fatalf("decode event: %v", err)
	}
	if envelope.Event != "message_created" {
		t.Fatalf("event = %q, want message_created", envelope.Event)
	}
	if envelope.Data.ChannelID != "general" || envelope.Data.Content != "Hello" {
		t.Fatalf("unexpected data: %+v", envelope.Data)
	}
}

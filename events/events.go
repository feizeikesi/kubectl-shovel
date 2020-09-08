package events

import (
	"encoding/json"
	"fmt"
)

// Event is struct for events
type Event struct {
	Type    EventType `json:"type"`
	Message string    `json:"message"`
}

// EventType is type for further parsing
type EventType string

const (
	Error  EventType = "error"
	Status EventType = "status"
	Result EventType = "result"
)

// NewEvent is used to publish new event
func NewEvent(eventType EventType, message string) error {
	data, err := json.Marshal(Event{
		Type:    eventType,
		Message: message,
	})
	if err != nil {
		return err
	}
	fmt.Println(string(data))

	return nil
}

// GetEvent is used to read published event
func GetEvent(data string) (*Event, error) {
	event := &Event{}
	if err := json.Unmarshal([]byte(data), event); err != nil {
		return nil, err
	}

	return event, nil
}
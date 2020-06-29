package govent

import (
	"fmt"
)

// EventHandler is an event handler interface
type EventHandler interface {
	Handle(Event) error
}

type EventData map[string]interface{}

// Event is an event
type Event interface {
	String() string
	GetData() EventData
}

// We keep a slice of EventHandlers for each event names
// There is no restrictions on the event names
// When a new one is registered it will be added as a new key
var registry = map[string][]EventHandler{}

// Listen is registering a new EventHandler
func Listen(name string, eh EventHandler) error {
	if _, ok := registry[name]; !ok {
		registry[name] = []EventHandler{}
	}
	registry[name] = append(registry[name], eh)

	return nil
}

// Publish runs all the EventHandlers for said event
func Publish(ev Event) error {
	ehs, ok := registry[ev.String()]
	if !ok {
		return fmt.Errorf("Event %s not found", ev)
	}
	for _, eh := range ehs {
		err := eh.Handle(ev)
		if err != nil {
			return err
		}
	}

	return nil
}

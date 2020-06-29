package govent

import (
	"fmt"
	"testing"
)

type TestEvent struct{}

func (ev TestEvent) String() string {
	return "event.one"
}

func (ev TestEvent) GetData() EventData {
	return map[string]interface{}{
		"test": "test",
	}
}

type LogListener struct{}

func (log LogListener) Handle(ev Event) error {
	result = fmt.Sprintf("Logging %s with %s", ev, ev.GetData()["test"])
	return nil
}

var result string

func TestOneEvent(t *testing.T) {
	logListener := LogListener{}
	testEvent := TestEvent{}

	err := Listen(testEvent.String(), logListener)
	if err != nil {
		t.Fatal("Unexpected error")
	}

	err = Publish(testEvent)
	if err != nil {
		t.Fatal("Unexpected error")
	}

	expected := "Logging event.one with test"
	if result != expected {
		t.Fatalf("Error asserting that\n%s\nis equal to\n%s", result, expected)
	}
}

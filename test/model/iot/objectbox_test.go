package iot

import (
	. "github.com/objectbox/objectbox-go/test/model/iot/object"
	"testing"
)

func TestObjectBoxEvents(t *testing.T) {
	objectBox := CreateObjectBox()
	box := objectBox.Box(1)
	event := Event{
		Device: "my device",
	}
	objectId, err := box.Put(&event)
	if err != nil {
		t.Fatalf("Could not add event: %v", err)
	}
	t.Logf("Added object ID %v", objectId)

	event.Device = "2nd device"
	objectId, err = box.Put(&event)
	if err != nil {
		t.Fatalf("Could not add 2nd event: %v", err)
	}
	t.Logf("Added 2nd object ID %v", objectId)

	objectRead, err := box.Get(objectId)
	if err != nil {
		t.Fatalf("Could not get 2nd event by ID: %v", err)
	}
	eventRead := objectRead.(*Event)
	if objectId != eventRead.Id || event.Device != eventRead.Device {
		t.Fatalf("Event data error: %v vs. %v", event, eventRead)
	}
}

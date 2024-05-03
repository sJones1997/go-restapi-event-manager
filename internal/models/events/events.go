package events

import (
	"time"
)

type Event struct {
	ID          int
	Uuid        string
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int
}

var events []Event

func (e *Event) Save() {
	// Soon to be added to DB
	events = append(events, *e)
}

func GetAllEvents() []Event {
	return events
}

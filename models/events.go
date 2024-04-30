package models

import (
	"time"
)

type Event struct {
	ID          int
	uuid        string
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UserID      int
}

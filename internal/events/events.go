package events

import (
	"database/sql"
	"errors"
	"github.com/sJones1997/go-restapi-event-manager/db"
	"github.com/sJones1997/go-restapi-event-manager/internal/utils/HTTPError"
	"log"
	"net/http"
	"time"
)

type Event struct {
	ID          int
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	CreatedAt   time.Time
	UserID      int
}

func (e *Event) Save() (Event, error) {
	// Soon to be added to DB
	query := `
		INSERT INTO events(name, description, location, user_id) 
		VALUES (?, ?, ?, ?)`
	stmt, err := db.CONN.Prepare(query)

	if err != nil {
		return Event{}, err
	}

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.UserID)

	if err != nil {
		return Event{}, err
	}

	resultID, err := result.LastInsertId()

	event, err := GetEvent(resultID)
	if err != nil {
		return Event{}, err
	}

	return event, nil
}

func GetEvent(eventId int64) (Event, error) {
	var event Event

	query := `
		SELECT id, name, description, location, created_at, user_id 
		from events 
		where id = ?`

	err := db.CONN.QueryRow(query, eventId).Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.CreatedAt, &event.UserID)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return Event{}, HTTPError.New(http.StatusNotFound, "Event not found")
		}

		return Event{}, err
	}

	return event, nil
}

func GetAllEvents() ([]Event, error) {

	rows, err := db.CONN.Query("SELECT id, name, description, location, created_at, user_id FROM events")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.CreatedAt, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return events, nil

}

func DeleteEvent(eventId int64) error {

	_, err := GetEvent(eventId)
	if err != nil {
		return err
	}

	query := `DELETE FROM events WHERE id = ?`

	stmt, err := db.CONN.Prepare(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(eventId)

	if err != nil {
		return err
	}

	return nil

}

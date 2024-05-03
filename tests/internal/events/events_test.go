package events

import (
	"bytes"
	"encoding/json"
	"github.com/sJones1997/go-restapi-event-manager/internal/events"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestCreateEvent(t *testing.T) {

	router := events.NewServer()

	w := httptest.NewRecorder()

	event := events.Event{
		Name:        "test",
		Description: "test description",
		Location:    "test location",
		DateTime:    time.Now(),
	}

	payload, _ := json.Marshal(event)

	req, _ := http.NewRequest("POST", "/events", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/json")

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

}

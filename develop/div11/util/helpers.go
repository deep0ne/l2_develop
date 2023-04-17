package util

import (
	"div11/domain"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const timeLayout = "2006-01-02"

func EventParser(r *http.Request) (int, domain.Event, error) {
	err := r.ParseForm()

	if err != nil {
		return 0, domain.Event{}, err
	}
	userId := r.Form.Get("user_id")
	ID, err := strconv.Atoi(userId)
	if err != nil {
		return 0, domain.Event{}, errors.New("Невалидный ID")
	}
	eventName := r.Form.Get("event_name")
	eventDate := r.Form.Get("date")
	t, err := time.Parse(timeLayout, eventDate)
	if err != nil {
		return 0, domain.Event{}, errors.New("Невалидная дата")
	}

	event := domain.Event{
		Name: eventName,
		Date: t,
	}

	return ID, event, nil
}

func JSONError(w http.ResponseWriter, err error, code int) {
	errJSON := make(map[string]string)
	errJSON["error"] = err.Error()
	bytes, _ := json.MarshalIndent(errJSON, "", "\t")

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)

	fmt.Fprintln(w, string(bytes))
}

func JSONWriter(w http.ResponseWriter, events []domain.Event, statusCode int) {
	resp := make(map[string][]domain.Event)
	resp["result"] = events

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)

	bytes, _ := json.MarshalIndent(resp, "", "\t")
	w.Write(bytes)
}

func GetEventsByDate(events []domain.Event, date string, days int) ([]domain.Event, error) {
	eventsByDate := make([]domain.Event, 0)
	t, err := time.Parse(timeLayout, date)
	if err != nil {
		return nil, err
	}
	for _, event := range events {
		switch days {
		case 1:
			if t.Equal(event.Date) {
				eventsByDate = append(eventsByDate, event)
			}
		case 7:
			if event.Date.Sub(t).Hours()/24 <= 7 {
				eventsByDate = append(eventsByDate, event)
			}
		case 30:
			if event.Date.Sub(t).Hours()/24 <= 30 {
				eventsByDate = append(eventsByDate, event)
			}
		}
	}
	return eventsByDate, nil
}

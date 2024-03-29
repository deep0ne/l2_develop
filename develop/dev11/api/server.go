package api

import (
	"div11/domain"
	"div11/middleware"
	"div11/util"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/sirupsen/logrus"
)

const (
	DAY   = 1
	WEEK  = 7
	MONTH = 30
)

type Server struct {
	Config util.Config
	Store  domain.UserInfo
	Logger *logrus.Logger
}

func NewServer(config util.Config) *Server {
	return &Server{
		Config: config,
		Store:  make(domain.UserInfo),
		Logger: util.NewLogger(),
	}
}

func (server *Server) NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/create_event", middleware.LoggingMiddleware(server.Logger, http.HandlerFunc(server.createEvent)))
	mux.Handle("/update_event", middleware.LoggingMiddleware(server.Logger, http.HandlerFunc(server.updateEvent)))
	mux.Handle("/delete_event", middleware.LoggingMiddleware(server.Logger, http.HandlerFunc(server.deleteEvent)))
	mux.Handle("/events_for_day", middleware.LoggingMiddleware(server.Logger, http.HandlerFunc(server.getEvents)))
	mux.Handle("/events_for_week", middleware.LoggingMiddleware(server.Logger, http.HandlerFunc(server.getEvents)))
	mux.Handle("/events_for_month", middleware.LoggingMiddleware(server.Logger, http.HandlerFunc(server.getEvents)))

	return mux
}

func (server *Server) createEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		util.JSONError(w, errors.New("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	ID, event, err := util.EventParser(r)
	if err != nil {
		util.JSONError(w, err, http.StatusBadRequest)
		return
	}

	server.Store[ID] = append(server.Store[ID], event)

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]string{"result": fmt.Sprintf("Event Юзера %d создан успешно", ID)})
	if err != nil {
		util.JSONError(w, err, http.StatusInternalServerError)
	}
}

func (server *Server) updateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		util.JSONError(w, errors.New("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}
	ID, event, err := util.EventParser(r)
	if err != nil {
		util.JSONError(w, err, http.StatusBadRequest)
		return
	}

	events, ok := server.Store[ID]
	if !ok {
		util.JSONError(w, errors.New("Юзера с таким ID не существует"), http.StatusInternalServerError)
		return
	}

	for i, e := range events {
		if e.Name == event.Name {
			events[i].Date = event.Date
		}
	}

	err = json.NewEncoder(w).Encode(map[string]string{"result": fmt.Sprintf("Время встречи '%s' у юзера %d успешно обновлено", event.Name, ID)})
	if err != nil {
		util.JSONError(w, err, http.StatusInternalServerError)
	}
}

func (server *Server) deleteEvent(w http.ResponseWriter, r *http.Request) {
	var deleted bool
	if r.Method != "POST" {
		util.JSONError(w, errors.New("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	ID, event, err := util.EventParser(r)
	if err != nil {
		util.JSONError(w, err, http.StatusBadRequest)
		return
	}
	events, ok := server.Store[ID]
	if !ok {
		util.JSONError(w, errors.New("Юзера с таким ID не существует"), http.StatusInternalServerError)
		return
	}
	for i, e := range events {
		if e.Name == event.Name {
			events = append(events[:i], events[i+1:]...)
			deleted = true
		}
	}

	if !deleted {
		util.JSONError(w, errors.New(fmt.Sprintf("Встречи с именем %s не существует", event.Name)), http.StatusBadRequest)
		return
	}

	if len(events) == 0 {
		delete(server.Store, ID)
	}

	err = json.NewEncoder(w).Encode(map[string]string{"result": fmt.Sprintf("Встреча '%s' у юзера %d успешно удалена", event.Name, ID)})
	if err != nil {
		util.JSONError(w, err, http.StatusInternalServerError)
	}
}

func (server *Server) getEvents(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		util.JSONError(w, errors.New("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}
	req := r.URL.Query()
	id, err := strconv.Atoi(req.Get("user_id"))

	if err != nil {
		util.JSONError(w, errors.New("Невалидный ID"), http.StatusBadRequest)
		return
	}

	event, ok := server.Store[id]
	if !ok {
		util.JSONError(w, errors.New("Юзера с таким ID не существует"), http.StatusInternalServerError)
		return
	}

	date := req.Get("date")
	events := make([]domain.Event, 0)

	switch r.URL.Path {
	case "/events_for_day":
		events, err = util.GetEventsByDate(event, date, DAY)
	case "/events_for_week":
		events, err = util.GetEventsByDate(event, date, WEEK)
	case "/events_for_month":
		events, err = util.GetEventsByDate(event, date, MONTH)
	}
	if err != nil {
		util.JSONError(w, err, http.StatusInternalServerError)
	}

	util.JSONWriter(w, events, http.StatusOK)
}

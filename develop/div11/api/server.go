package api

import (
	"div11/domain"
	"div11/util"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

const (
	DAY   = 1
	WEEK  = 7
	MONTH = 30
)

type Server struct {
	Config util.Config
	Store  domain.UserInfo
}

func NewServer(config util.Config) *Server {
	return &Server{
		Config: config,
		Store:  make(domain.UserInfo),
	}
}

func (server *Server) NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/create_event", http.HandlerFunc(server.createEvent))
	mux.Handle("/update_event", http.HandlerFunc(server.updateEvent))
	mux.Handle("/delete_event", http.HandlerFunc(server.deleteEvent))
	mux.Handle("/events_for_day", http.HandlerFunc(server.getEvents))
	mux.Handle("/events_for_week", http.HandlerFunc(server.getEvents))
	mux.Handle("/events_for_month", http.HandlerFunc(server.getEvents))

	return mux
}

func (server *Server) createEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		util.JSONError(w, errors.New("Method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	ID, event, err := util.CreateEventParser(r)
	if err != nil {
		util.JSONError(w, err, http.StatusInternalServerError)
		return
	}

	server.Store[ID] = append(server.Store[ID], event)

	json.NewEncoder(w).Encode(map[string]string{"result": fmt.Sprintf("Event Юзера %d создан успешно", ID)})
}

func (server *Server) updateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (server *Server) deleteEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
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
		util.JSONError(w, errors.New("Юзера с таким ID не существует"), http.StatusBadRequest)
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

	util.JSONWriter(w, events, http.StatusOK)
}

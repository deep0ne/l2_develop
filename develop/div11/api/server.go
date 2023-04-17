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

type Server struct {
	Config util.Config
	Store  map[int][]domain.Event
}

func NewServer(config util.Config) *Server {
	return &Server{
		Config: config,
		Store:  make(map[int][]domain.Event),
	}
}

func (server *Server) NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("/create_event", http.HandlerFunc(server.createEvent))
	mux.Handle("/update_event", http.HandlerFunc(server.updateEvent))
	mux.Handle("/delete_event", http.HandlerFunc(server.deleteEvent))
	mux.Handle("/events_for_day", http.HandlerFunc(server.getEventsForDay))
	mux.Handle("/events_for_week", http.HandlerFunc(server.getEventsForWeek))
	mux.Handle("/events_for_month", http.HandlerFunc(server.getEventsForMonth))

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

func (server *Server) getEventsForDay(w http.ResponseWriter, r *http.Request) {
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
	util.JSONEventsResponse(w, event, date, 1)
}

func (server *Server) getEventsForWeek(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func (server *Server) getEventsForMonth(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

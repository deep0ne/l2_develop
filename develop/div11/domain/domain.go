package domain

import (
	"time"
)

type Event struct {
	Name string    `json:"Имя встречи"`
	Date time.Time `json:"Время встречи"`
}

type User struct {
	ID     int
	Events []Event
}

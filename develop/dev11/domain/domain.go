package domain

import (
	"time"
)

type Event struct {
	Name string    `json:"Название встречи"`
	Date time.Time `json:"Время встречи"`
}

type UserInfo map[int][]Event

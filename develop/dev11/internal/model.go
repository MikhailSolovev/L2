package internal

import "time"

type Event struct {
	ID     int       `json:"id"`
	UserID int       `json:"user_id"`
	Event  string    `json:"event"`
	Date   time.Time `json:"date"`
}

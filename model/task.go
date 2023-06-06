package model

import "time"

type Task struct {
	ID    int       `db:"id"`
	Title string    `db:"title"`
	Task  string    `db:"task"`
	Data  time.Time `db:"date_created"`
}

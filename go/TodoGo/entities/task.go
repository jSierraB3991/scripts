package entities

import "time"

type Task struct {
	Id         int64
	Name       string
	Done       bool
	CreateAt   time.Time
	CompleteAt time.Time
}

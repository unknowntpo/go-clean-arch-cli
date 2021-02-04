package main

import (
	"time"
)

type Task struct {
	Id       int       // task id
	Desc     string    // task description
	CreateTS time.Time // when the task was created
	DoneTS   time.Time // when the task was done
	IsDone   bool      // task is done or not
}

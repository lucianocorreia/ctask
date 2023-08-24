package db

import "time"

// Task is a task
type Task struct {
	ID      uint
	Name    string
	Project string
	status  string
	Created time.Time
}

// FilterValue returns the value to be used for filtering
func (t Task) FilterValue() string {
	return t.Name
}

// Title returns the title of the task
func (t Task) Title() string {
	return t.Name
}

// Description returns the description of the task
func (t Task) Description() string {
	return t.Project
}

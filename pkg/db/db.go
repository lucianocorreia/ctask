package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// OpenDB opens the database, and creates the table if it doesn't exist
func OpenDB(path string) (*TaskDB, error) {
	db, err := sql.Open("sqlite3", fmt.Sprintf("%s/tasks.db", path))
	if err != nil {
		return nil, err
	}

	t := TaskDB{db, path}

	if !t.tableExists("tasks") {
		err := t.createTableTasks()
		if err != nil {
			defer db.Close()
			return nil, err
		}
	}

	return &t, nil
}

// Insert inserts a new task into the database
func (t *TaskDB) Insert(name, project string) error {
	_, err := t.DB.Exec(
		"INSERT INTO tasks(name, project, status, created) VALUES(?, ?, ?, ?)",
		name,
		project,
		todo.String(),
		time.Now())

	return err
}

func (t *TaskDB) GetTasks() ([]Task, error) {
	var tasks []Task
	rows, err := t.DB.Query("SELECT * FROM tasks")
	if err != nil {
		return tasks, fmt.Errorf("error getting tasks: %w", err)
	}

	for rows.Next() {
		var task Task
		err := rows.Scan(
			&task.ID,
			&task.Name,
			&task.Project,
			&task.Status,
			&task.Created,
		)

		if err != nil {
			return tasks, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

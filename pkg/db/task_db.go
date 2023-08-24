package db

import (
	"database/sql"
	"fmt"
)

// TaskDB is the database for tasks
type TaskDB struct {
	DB      *sql.DB
	DataDir string
}

func (t *TaskDB) tableExists(name string) bool {
	if _, err := t.DB.Exec(fmt.Sprintf("SELECT * FROM %s", name)); err == nil {
		return true
	}

	return false
}

func (t *TaskDB) createTableTasks() error {
	_, err := t.DB.Exec(`CREATE TABLE "tasks" ( "id" INTEGER, "name" TEXT NOT NULL, "project" TEXT, "status" TEXT, "created" DATETIME, PRIMARY KEY("id" AUTOINCREMENT))`)

	return err
}

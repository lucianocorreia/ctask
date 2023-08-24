package db

import "database/sql"

type TaskDB struct {
	db      *sql.DB
	dataDir string
}

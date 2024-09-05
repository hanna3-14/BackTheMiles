package db

import (
	"database/sql"

	"github.com/hanna3-14/BackTheMiles/pkg/helpers"
)

type SQLDBAdapter struct {
	Database *sql.DB
}

func NewSQLDBAdapter(databaseFile string) (*SQLDBAdapter, error) {
	path := helpers.SafeGetEnv("PATH_TO_VOLUME")
	db, err := sql.Open("sqlite3", path+databaseFile)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	return &SQLDBAdapter{Database: db}, nil
}

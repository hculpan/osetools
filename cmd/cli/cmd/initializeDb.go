package cmd

import (
	"database/sql"
	"errors"
	"fmt"
	"os"

	"github.com/hculpan/osetools/internal/db"
	"github.com/joho/godotenv"
)

var dbase *sql.DB

func initializeDb() (*db.Queries, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("missing .env file: %w", err)
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		return nil, errors.New("missing DB_PATH in configuration file")
	}

	dbase, err = sql.Open("sqlite", dbPath)
	if err != nil {
		return nil, err
	}

	return db.New(dbase), nil
}

func closeDb() {
	if dbase != nil {
		dbase.Close()
	}
}

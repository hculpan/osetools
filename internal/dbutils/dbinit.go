package dbutils

import (
	_ "embed"
	"log"
	"os"

	"github.com/hculpan/osetools/internal/db"
	"github.com/joho/godotenv"

	"database/sql"

	_ "github.com/glebarez/go-sqlite"
)

var dbase *sql.DB
var queries *db.Queries

//go:embed scripts/schema.sql
var ddl string

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("missing .env file: %s", err.Error())
	}

	// Update this whatever variables you need to retrieve to
	// initialize the db
	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		log.Fatal("missing DB_PATH in configuration file")
	}

	dbase, err = sql.Open("sqlite", dbPath)
	if err != nil {
		log.Fatal(err)
	}

	queries = db.New(dbase)
}

func Close() {
	if dbase != nil {
		dbase.Close()
	}
}

func GetDdl() string {
	return ddl
}

func GetQueries() *db.Queries {
	return queries
}

package dbutils

import (
	_ "embed"
	"fmt"
	"log"
	"os"

	"github.com/hculpan/osetools/internal/db"
	"github.com/joho/godotenv"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
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

	dbServer := os.Getenv("DB_SERVER")
	if dbServer == "" {
		log.Fatal("missing DB_SERVER in configuration file")
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		log.Fatal("missing DB_NAME in configuration file")
	}

	dbUser := os.Getenv("DB_USER")
	if dbUser == "" {
		log.Fatal("missing DB_USER in configuration file")
	}

	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		log.Fatal("missing DB_PASSWORD in configuration file")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", dbUser, dbPassword, dbServer, dbName)

	dbase, err = sql.Open("mysql", dsn)
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

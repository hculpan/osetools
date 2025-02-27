package cmd

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/hculpan/osetools/internal/db"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

var dbase *sql.DB

func initializeDb() (*db.Queries, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("missing .env file: %w", err)
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
		return nil, err
	}

	if err := dbase.Ping(); err != nil {
		return nil, fmt.Errorf("database connection failed: %w", err)
	}

	return db.New(dbase), nil
}

func closeDb() {
	if dbase != nil {
		dbase.Close()
	}
}

package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/hculpan/osetools/internal/dbutils"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("missing .env file: %s", err.Error())
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("missing PORT in configuration file")
	}

	defer dbutils.Close()

	router := routes(chi.NewRouter())

	slog.Info("starting server on port " + port)
	http.ListenAndServe(":"+port, router)
}

package main

import (
    "database/sql"
    "log"
    "net/http"
    "tender-service/internal/database"
    "tender-service/internal/handlers"
    "tender-service/internal/repository"
    "tender-service/internal/service"
    "tender-service/pkg/config"
    "tender-service/pkg/logger"
	"github.com/joho/godotenv"

    _ "github.com/lib/pq"
)

func main() {
    logger.Init()

	err := godotenv.Load()
    if err != nil {
        logger.Error.Fatalf("Error loading .env file")
    }

    cfg := config.LoadConfig()

    db, err := sql.Open("postgres", cfg.PostgresConn)
    if err != nil {
        logger.Error.Fatalf("Unable to connect to the database: %v\n", err)
    }

    database.InitDB(db)

    tenderRepo := repository.NewTenderRepository(db)
    tenderService := service.NewTenderService(tenderRepo)

    http.HandleFunc("/api/ping", handler.PingHandler)
    http.HandleFunc("/api/tenders/new", handler.CreateTenderHandler(tenderService))
    http.HandleFunc("/api/tenders", handler.ListTendersHandler(tenderService))

    logger.Info.Printf("Starting server at %s\n", cfg.ServerAddress)
    log.Fatal(http.ListenAndServe(cfg.ServerAddress, nil))
}
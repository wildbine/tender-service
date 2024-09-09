package main

import (
    "database/sql"
    "fmt"
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

    connStr := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable",
        cfg.PostgresUsername, cfg.PostgresPassword, cfg.PostgresHost, cfg.PostgresDatabase)
    
    fmt.Println("Connecting to database with:", connStr)

    db, err := sql.Open("postgres", connStr)
    if err != nil {
        logger.Error.Fatalf("Unable to connect to the database: %v\n", err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        logger.Error.Fatalf("Unable to ping the database: %v\n", err)
    }

    database.InitDB(db)

    tenderRepo := repository.NewTenderRepository(db)
    tenderService := service.NewTenderService(tenderRepo)
    tenderHandler := handler.NewTenderHandler(tenderService)

    http.HandleFunc("/api/ping", handler.PingHandler)
    http.HandleFunc("/api/tenders/new", tenderHandler.CreateTenderHandler)
    http.HandleFunc("/api/tenders", tenderHandler.GetTenderHandler)
    
    logger.Info.Printf("Starting server at %s\n", cfg.ServerAddress)
    log.Fatal(http.ListenAndServe(cfg.ServerAddress, nil))
}
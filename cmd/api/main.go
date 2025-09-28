package main

import (
	"database/sql"
	"log"
	"rest-api-in-gin/internal/database"
	"rest-api-in-gin/internal/env"
	"rest-api-in-gin/internal/logger"

	_ "github.com/joho/godotenv/autoload" // Automatically loads environment variables
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

type application struct {
	port      int
	jwtSecret string
	models    database.Models
	logger    *logrus.Logger
}

func main() {

	// Initialize with Debug level so you see all logs during development
	logger.InitLogger("logs/app.log", logrus.DebugLevel)

	logger.AppLogger.Info("Application started")
	logger.AppLogger.Debug("Debugging enabled")

	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	models := database.NewModels(db)

	app := &application{

		port:      env.GetEnvInt("PORT", 8080),
		jwtSecret: env.GetEnvString("JWT_SECRET", "some-secret-1213123"),
		models:    models,
	}

	if err := app.serve(); err != nil {
		log.Fatal(err)
	}

}

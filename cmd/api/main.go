package main

import (
	"database/sql"
	"log"
	_ "rest-api-in-gin/docs"
	"rest-api-in-gin/internal/database"
	"rest-api-in-gin/internal/env"
	"rest-api-in-gin/internal/logger"

	_ "github.com/joho/godotenv/autoload" // Automatically loads environment variables
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

// @title Go Gin Rest API
// @version 1.0
// @description A rest API in Go using Gin framework.
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter your bearer token in the format **Bearer &lt;token&gt;**

// Apply the security definition to your endpoints
// @security BearerAuth

type application struct {
	port      int
	jwtSecret string
	models    database.Models
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

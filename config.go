package main

import (
	"database/sql"
	"github.com/CodyBrunson/groceryList/internal/database"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

type Config struct {
	port         string
	filePathRoot string
	platform     string
	db           *database.Queries
}

func NewConfig() *Config {
	godotenv.Load()
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	FILEPATHROOT := os.Getenv("FILEPATH_ROOT")
	if FILEPATHROOT == "" {
		log.Fatal("FILEPATH_ROOT is not set")
		return nil
	}

	DBSTRING := os.Getenv("DB_URL")
	if DBSTRING == "" {
		log.Fatal("DBSTRING is not set")
	}
	PLATFORM := os.Getenv("PLATFORM")
	switch PLATFORM {
	case "DEV":
		log.Println("Application is running in DEV mode.  Accessing dev database.")
		DBSTRING = strings.Replace(DBSTRING, "###", "dev", 1)
	case "PROD":
		log.Println("Application is running in PROD mode.  Accessing public database.")
		DBSTRING = strings.Replace(DBSTRING, "###", "public", 1)
	case "":
		log.Fatal("PLATFORM is not set")
		return nil
	}

	db, err := sql.Open("postgres", DBSTRING)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
		return nil
	}

	queries := database.New(db)
	newConfig := &Config{
		port:         PORT,
		filePathRoot: FILEPATHROOT,
		platform:     PLATFORM,
		db:           queries,
	}

	return newConfig

}

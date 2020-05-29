package settings

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

const (
	// VERSION defines project version
	VERSION = "0.1.0"
)

var (
	// DB is global database instance
	DB *gorm.DB
	// Router is global router instance
	Router *gin.Engine
	// Env is global configs
	Env map[string]string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
	Env, err = godotenv.Read()
	if err != nil {
		log.Println("Error reading .env file")
	}

	if Env["API_PORT"] == "" {
		Env["API_PORT"] = os.Getenv("PORT")
	}
	if Env["API_SECRET"] == "" {
		Env["API_SECRET"] = os.Getenv("API_SECRET")
	}
	if Env["DB_DRIVER"] == "" {
		Env["DB_DRIVER"] = os.Getenv("DB_DRIVER")
	}
	if Env["DATABASE_URL"] == "" {
		Env["DATABASE_URL"] = os.Getenv("DATABASE_URL")
	}
}

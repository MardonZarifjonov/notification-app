package database

import (
	"fmt"
	"log"
	"notification-app/api/models"
	s "notification-app/api/settings"

	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// _ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func init() {
	var err error

	dbDriver := s.Env["DB_DRIVER"]
	dbUsername := s.Env["DB_USER"]
	dbPassword := s.Env["DB_PASSWORD"]
	dbHost := s.Env["DB_HOST"]
	dbPort := s.Env["DB_PORT"]
	dbName := s.Env["DB_NAME"]

	if len(s.Env["DATABASE_URL"]) > 10 {
		db, err = gorm.Open(dbDriver, s.Env["DATABASE_URL"])
	} else if dbDriver == "mysql" {
		DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbName)
		db, err = gorm.Open(dbDriver, DBURL)
	} else if dbDriver == "postgres" {
		DBURL := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", dbHost, dbPort, dbName, dbUsername, dbPassword)
		db, err = gorm.Open(dbDriver, DBURL)
	} else if dbDriver == "sqlite3" {
		db, err = gorm.Open("sqlite3", dbHost)
	} else {
		log.Fatal("Unknown Database Driver")
	}

	if err != nil {
		fmt.Printf("Cannot connect to %s database\n", dbDriver)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database\n", dbDriver)
	}

	//Printing query
	db.LogMode(true)

	//Automatically create migration as per model
	db.Debug().AutoMigrate(
		&models.Order{},
		&models.User{},
		&models.Notification{},
	)
}

//GetDB function return the instance of db
func GetDB() *gorm.DB {
	return db
}

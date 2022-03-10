package database

import (
	"JWT-GoFiber/models"
	"log"
	"os"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB
var err error

type DbInstance struct {
	Db *gorm.DB
	// FirstName string `json:"firstname"`
	// LastName  string  `json:"lastname"`
	// Email     string  `json:"email"`
}
var Database DbInstance

func ConnectDb(){
	db , err := gorm.Open(sqlite.Open("jwtgo.db"), &gorm.Config{}) 

	if err != nil{
		log.Fatal("Failed to open database")
		os.Exit(2)
	}
	log.Println("Database Connected succefully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")
	// add migrations
	db.AutoMigrate(& models.User{})
	Database = DbInstance{Db: db}
}
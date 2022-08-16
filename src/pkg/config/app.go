package config

import (
	"fmt"
	"log"
	"os"

	"github.com/felipelaptrin/go-api/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error

func Connect() {
	user := os.Getenv("DATABASE_USER")
	password := os.Getenv("DATABASE_PASSWORD")
	database := os.Getenv("DATABASE_NAME")
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")

	connection_string := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		database,
	)
	Db, err = gorm.Open(mysql.Open(connection_string), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic("Failed to connect database")
	} else {
		fmt.Println("Connection to database was a success!")
	}
}

func Migrate() {
	Db.AutoMigrate(&models.User{})
	log.Println("Database Migration Completed.")
}

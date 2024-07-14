package cmd

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() {
	pg_host := os.Getenv("POSTGRES_HOST")
	pg_user := os.Getenv("POSTGRES_USER")
	pg_password := os.Getenv("POSTGRES_PASSWORD")
	pg_db := os.Getenv("POSTGRES_DB")
	pg_port := os.Getenv("POSTGRES_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Cairo", pg_host, pg_user, pg_password, pg_db, pg_port)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("couldn't connect to database:", err)
	}
	log.Println("connected to database")
	db = database
}

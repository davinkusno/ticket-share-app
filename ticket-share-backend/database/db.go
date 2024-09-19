package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Credential struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         int
}

// Connect initializes the database connection
func Connect(cred Credential) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", 
		cred.Host, cred.Username, cred.Password, cred.DatabaseName, cred.Port)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
}

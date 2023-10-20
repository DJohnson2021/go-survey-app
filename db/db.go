package db

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	// "github.com/joho/godotenv"
)

var DB *gorm.DB

func InitDatabase() {
	var err error

	dsn := os.Getenv("POSTGRES_STRING_DEV")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: &s", err.Error())
	}
}

func CloseDatabase() {
    sqlDB, err := DB.DB()
    if err != nil {
        log.Fatalf("Failed to get underlying database from GORM: %s", err.Error())
    }
    sqlDB.Close()
}

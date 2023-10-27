package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var (
	DB *gorm.DB
)

func OpenDatabaseConnection() {
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_URL")), &gorm.Config{})
	if err != nil {
		panic("failed to connect entities")
	}
	DB = db
}

func CloseDatabaseConnection() {
	db, err := DB.DB()
	if err != nil {
		fmt.Println("Error on closing connection:", err.Error())
	}
	db.Close()
}

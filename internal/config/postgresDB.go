package config

import (
	"os"

	"github.com/Kawar1mi/golang-crud-gin-gorm/internal/helper"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DatabaseConnection() *gorm.DB {

	// DSN in format - "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	if err != nil {
		helper.FatalIfError(err)
	}

	PingDB(db)

	helper.InfoMsg("connected to database")

	return db
}

func PingDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		helper.FatalIfError(err)
	}

	err = sqlDB.Ping()
	if err != nil {
		helper.FatalIfError(err)
	}
}

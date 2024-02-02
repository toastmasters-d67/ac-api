package helpers

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func InitDb() *gorm.DB {
	HOST := os.Getenv("DB_HOST")
	PORT := os.Getenv("DB_PORT")
	DATABASE := os.Getenv("DB_DATABASE")
	USER := os.Getenv("DB_USER")
	PASSWORD := os.Getenv("DB_PASSWORD")
	SSL := os.Getenv("DB_SSL")
	TIMEZONE := os.Getenv("DB_TIMEZONE")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
		HOST, PORT, USER, PASSWORD, DATABASE, SSL, TIMEZONE)

	gormDB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
	})
	if err != nil {
		panic("open gorm db error")
	}

	return gormDB
}

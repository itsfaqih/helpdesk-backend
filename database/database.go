package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func Init() *gorm.DB {
	var err error
	dsn := "host=localhost user=postgres password=postgres dbname=helpdesk port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect database. Error: " + err.Error())
	}

	fmt.Println("Connection Opened to Database")

	return db
}

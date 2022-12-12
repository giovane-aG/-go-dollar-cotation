package create_connection

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func CreateConnection() *gorm.DB {
	var db *gorm.DB
	var err error

	dsn := "user=postgres host=localhost dbname=postgres password=postgres port=5435 sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	return db
}

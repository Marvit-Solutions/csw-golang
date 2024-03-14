package database

import (
	"fmt"

	"github.com/Marvit-Solutions/csw-golang/library/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDBSQL initializes the PostgreSQL database connection.
func InitDBSQL(env config.Config, parentKey string) (db *gorm.DB, err error) {
	dbUser := env.GetString(parentKey + `.user`)
	dbPass := env.GetString(parentKey + `.pass`)
	dbName := env.GetString(parentKey + `.database`)
	dbHost := env.GetString(parentKey + `.address`)
	dbPort := env.GetString(parentKey + `.port`)

	// Construct the connection string.
	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		dbUser,
		dbPass,
		dbName,
		dbHost,
		dbPort,
	)

	// Open a connection to the database.
	db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("failed connect to database")
	}

	return db, err
}

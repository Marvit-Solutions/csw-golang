package database

import (
	"csw-golang/library/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDBSQL(env config.Config, parentKey string) (db *gorm.DB, err error) {
	dbUser := env.GetString(parentKey + `.user`)
	dbPass := env.GetString(parentKey + `.pass`)
	dbName := env.GetString(parentKey + `.database`)
	dbHost := env.GetString(parentKey + `.address`)
	dbPort := env.GetString(parentKey + `.port`)

	connectionString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		dbUser,
		dbPass,
		dbName,
		dbHost,
		dbPort,
	)

	db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("failed connect to database")
	}

	return db, err
}

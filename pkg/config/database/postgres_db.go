package database

import (
	"poc-recognition/pkg/config"
	"poc-recognition/pkg/logger"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB struct {
	db *gorm.DB
}

func NewPostgresDB() *PostgresDB {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  config.DB_POSTGRES_URI,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		logger.Error("failed to connect database", err.Error())
		return nil
	}

	return &PostgresDB{
		db: db,
	}
}

package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	// Initialize Sqlite
	db, err = InitilizeSQLite()

	if err != nil {
		return fmt.Errorf("Error initilizing sqlite %v", err)
	}

	return nil
}

func GetSLite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}

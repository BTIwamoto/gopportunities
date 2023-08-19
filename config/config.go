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

	//Initialize SQL Server
	db, err = InitializeSQLServer()
	if err != nil {
		return fmt.Errorf("error initializing sqlServer: %v", err)
	}

	return nil
}

func GetSQLServer() *gorm.DB {
	return db
}

func GetLogger(prefix string) *Logger {
	// Initialize Logger
	logger = NewLogger(prefix)
	return logger
}

package config

import (
	"github.com/BTIwamoto/gopportunities/schemas"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func InitializeSQLServer() (*gorm.DB, error) {
	logger := GetLogger("sqlServer")
	connString := "sqlserver://sa:Hmnioq@183@localhost:1433?database=Opportunities&connection+timeout=30"

	//Create Database and Connect
	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		logger.Errorf("database opening error: %v", err)
		return nil, err
	}

	// Migrate the Schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("database automigration error: %v", err)
		return nil, err
	}

	// Return Database
	return db, nil
}

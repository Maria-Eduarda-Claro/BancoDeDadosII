package database

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func GetDatabaseConnection() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(DATABASE_NAME))

}

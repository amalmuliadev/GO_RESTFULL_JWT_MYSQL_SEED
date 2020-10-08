package database

import (
	"../../config"
	// mysql is driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Connect to the DATABASE
func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(config.DBDRIVER, config.DBURL)
	if err != nil {
		return nil, err
	}
	return db, nil
}

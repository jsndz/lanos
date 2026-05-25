package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(url), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

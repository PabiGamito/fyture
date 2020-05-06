package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // Sqlite Driver for gorm
)

func SetupModels() *gorm.DB {
	db, err := gorm.Open("sqlite3", "test.db")

	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Feature{}, &User{})

	return db
}

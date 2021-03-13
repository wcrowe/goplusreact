package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Has to be uppercase
func Connect() {
	dsn := "root:password@tcp(localhost:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

}

package config

import (
	"latihanotp/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateDatabase() *gorm.DB {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/otpgmail"))
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{})
	return db
}

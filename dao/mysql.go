package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() (err error) {
	dsn := "root:688@tcp(127.0.0.1:3306)/gormtest?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return
}

func Close() (err error) {
	sqlDb, err := DB.DB()
	sqlDb.Close()
	return
}

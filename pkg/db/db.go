package db

import (
	"proj1/models"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

func DBConfig() (*gorm.DB, error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/crud?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&models.Student{}, &models.User{})

	return db, nil
}

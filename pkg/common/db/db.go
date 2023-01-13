package db

import (
	"fmt"
	"log"

	"shortlink/pkg/common/config"
	"shortlink/pkg/common/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",config.GetEnv("DB_USER"), config.GetEnv("DB_PASSWORD"),config.GetEnv("HOST"),config.GetEnv("PORT"),config.GetEnv("DB_NAME") )
	db, e := gorm.Open(mysql.Open(url), &gorm.Config{})

	if e != nil {
		log.Fatalln(e)
	}

	db.AutoMigrate(&models.User{}, &models.Link{}, &models.VisitLink{})

	return db
}
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
	// load  env
	config, err := config.InitConfig()

	if err != nil {
		log.Fatalln(err)
	}

	url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",config.DB_USER, config.DB_PASSWORD,config.HOST,config.PORT,config.DB_NAME )
	db, e := gorm.Open(mysql.Open(url), &gorm.Config{})

	if e != nil {
		log.Fatalln(e)
	}

	db.AutoMigrate(&models.User{})

	return db
}
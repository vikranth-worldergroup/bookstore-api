package config

import (
	"log"
	"bookstore/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Connect(){
	dsn:="root:@Vikranth08@tcp(127.0.0.1:3306)/bookstore?charset=utf8mb4&parseTime=True&loc=Local"
	db,err:=gorm.Open(mysql.Open(dsn),&gorm.Config{})

	if err!=nil{
		log.Fatal("failed to connect db",err)
	}
	Db=db
	Db.AutoMigrate(&models.Book{},&models.Author{})
}
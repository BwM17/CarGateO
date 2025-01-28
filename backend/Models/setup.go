package models

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{}) 
	if err != nil {
		log.Fatalln("failed to connect to Database") 
	} 

	err = database.AutoMigrate(&User{}, &Car{}, &Log{}) 
	if err != nil {  
		log.Fatalln("failed to migrate Data set") 
	} 

	DB = database
}
package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"task_go/model"
)

var db *gorm.DB
var port string

func SetupModels() {
	connect := "host=localhost user=postgres password=12345678 dbname=task_go port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(connect), &gorm.Config{})
	if err != nil {
		return
	}

	db.AutoMigrate(&model.User{})

	SetUpDBConnection(db)
	SetPortConnection(":8080")

}

func SetUpDBConnection(DB *gorm.DB) {
	db = DB
}

func GetDBConnection() *gorm.DB {
	return db
}

func SetPortConnection(Port string) {
	port = Port
}

func GetPortConnection() string {
	return port
}
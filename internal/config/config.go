package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

const DB_USER = "swapnesh"
const DB_PASSWORD = "Swapnesh@1998"
const DB_NAME = "suryansh"

var db *gorm.DB

func Connect() {
	var db_args = DB_USER + ":" + DB_PASSWORD + "@/" + DB_NAME + "?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open("mysql", db_args)
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}

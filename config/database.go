package config

import (
	"github.com/jinzhu/gorm"
	"os"
)

var Db *gorm.DB
var err error

/**
this function to connect to database
*/
func ConnectToDatabase() {
	Db, err = gorm.Open("mysql", os.Getenv("DATABASE_USERNAME")+":"+os.Getenv("DATABASE_PASSWORD")+"@tcp("+os.Getenv("DATABASE_HOST")+":"+
		os.Getenv("DATABASE_PORT")+")/"+os.Getenv("DATABASE_NAME")+"?charset=utf8&parseTime=True&loc=Local&character_set_server=utf8")
	Db.LogMode(true)
}


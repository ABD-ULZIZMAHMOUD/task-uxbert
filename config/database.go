package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"os"
)

var Db *gorm.DB
var err error

/**
* this function to connect to database
 */
func ConnectToDatabase() {
	Db, err = gorm.Open("mysql", os.Getenv("DATABASE_USERNAME")+":"+os.Getenv("DATABASE_PASSWORD")+"@tcp("+os.Getenv("DATABASE_HOST")+":"+
		os.Getenv("DATABASE_PORT")+")/"+os.Getenv("DATABASE_NAME")+"?charset=utf8&parseTime=True&loc=Local&character_set_server=utf8")
	if err != nil {
		log.Fatal(err)
	}
	Db.LogMode(true)
}

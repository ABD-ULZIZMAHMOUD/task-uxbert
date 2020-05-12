package provider

import (
	"task-uxbert/config"
	"task-uxbert/models"
)

func Migrate() {
	config.Db.AutoMigrate(&models.User{}, models.Room{}, models.Message{})
}

func truncate(tables ...string) {
	for _, tc := range tables {
		s := "TRUNCATE TABLE " + tc
		config.Db.Exec(s)
	}
}

func Drop(tables ...string) {
	for _, tc := range tables {
		s := "DROP TABLE " + tc
		config.Db.Exec(s)
	}
}

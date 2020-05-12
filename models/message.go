package models

import "github.com/jinzhu/gorm"

type Message struct {
	gorm.Model
	RoomID   int    `json:"room_id" `
	Content  string `sql:"type:text;" json:"content"`
	Receiver int    `json:"receiver"`
	Sender   int    `json:"sender"`
	IsRead   int    `json:"is_read"`
}

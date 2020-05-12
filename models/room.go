package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

/***
* why i add full_name i room struct because this is  customer chat service so
* probability of user change his data is low so this will increase performance to get rooms with any joins
 */
type Room struct {
	gorm.Model
	User1       int       `json:"user_1"`
	User2       int       `json:"user_2"`
	LastMessage time.Time `json:"last_message"`
	Messages    Message   `json:"message" binding:"nostructlevel"`
	FullName1   string    `json:"first_name_1" `
	FullName2   string    `json:"first_name_2" `
}

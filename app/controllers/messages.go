package controllers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"task-uxbert/app/validations"
	"task-uxbert/config"
	helpers "task-uxbert/helper"
	"task-uxbert/models"
	"time"
)

func StoreMessageINRoom(g *gin.Context) {
	user := models.GetUserFromHeader(g)
	// init visitor User struct to validate request
	message := new(models.Message)
	/**
	* get request and parse it to validation
	* if there any error will return with message
	 */
	err := validations.MassageValidate(g, message)
	/***
	* return response if there an error if true you
	* this mean you have errors so we will return and bind data
	 */
	if helpers.ReturnNotValidRequest(err, g) {
		return
	}
	message.Sender = int(user.ID)
	message.IsRead = 2
	userRecevier := models.GetUserBYID(message.Receiver)
	if userRecevier.ID != 0 || user.ID != 0 {
		helpers.ReturnNotFound(g, "Not found this user!")
		return
	}
	var oldRoom models.Room
	config.Db.Where("user1 = ? and user2 = ?", message.Sender, message.Receiver).
		Or("user1 = ? and user2 = ?", message.Receiver, message.Sender).First(&oldRoom)
	if oldRoom.ID == 0 && user.Type == 2 {
		oldRoom.LastMessage = time.Now()
		oldRoom.FullName1 = user.FullName
		oldRoom.FullName2 = userRecevier.FullName

		config.Db.Create(&oldRoom)
		// write in chan new room from this user to both user
		writeToClientChatRoom(oldRoom, int(user.ID), int(userRecevier.ID))
		writeToClientChatRoom(oldRoom, int(userRecevier.ID), int(user.ID))

	} else if oldRoom.ID == 0 && user.Type != 2 {
		helpers.ReturnNotFound(g, "Not found this user!")
		return
	} else {
		oldRoom.LastMessage = time.Now()
		config.Db.Save(&oldRoom)
	}
	message.RoomID = int(oldRoom.ID)
	config.Db.Create(&message)
	writeToClientChat(*message)
	helpers.OkResponse(g, "room", message)
	return
}

func MessageInRoom(g *gin.Context) {
	user := models.GetUserFromHeader(g)
	room_id, _ := strconv.Atoi(g.Param("id"))
	var room models.Room
	config.Db.Where("user1 = ? or user2 = ?", user.ID, user.ID).First(&room, room_id)
	if room.ID == 0 {
		helpers.ReturnNotFound(g, "this room not exist!")
	}
	limit_stirng := g.DefaultQuery("limit", "10")
	offset_string := g.DefaultQuery("offset", "0")
	limit, _ := strconv.Atoi(limit_stirng)
	offset, _ := strconv.Atoi(offset_string)
	if limit > 10 || limit == 0 {
		limit = 10
	}
	var messages []models.Message
	config.Db.Where("room_id = ?", room_id).Offset(offset).Limit(limit).Order("id desc").Find(&messages)
	helpers.OkResponse(g, "messages in room ", messages)
	return
}

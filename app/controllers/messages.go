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

/***
* store new message to normal user from admin
 */
func StoreMessageINRoomAdmin(g *gin.Context) {
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
	// get users and check if exist
	userRecevier := models.GetUserBYID(message.Receiver)
	if userRecevier.ID == 0 || user.ID == 0 {
		helpers.ReturnNotFound(g, "Not found this user!")
		return
	}
	// get room and check this user can chat or not
	var oldRoom models.Room
	config.Db.Where("user1 = ? and user2 = ?", message.Sender, message.Receiver).
		Or("user1 = ? and user2 = ?", message.Receiver, message.Sender).First(&oldRoom)
	if oldRoom.ID == 0 {
		helpers.ReturnNotFound(g, "Not found this room!")
		return

	} else {
		oldRoom.LastMessage = time.Now()
		config.Db.Save(&oldRoom)
	}
	// store message in database
	message.RoomID = int(oldRoom.ID)
	config.Db.Create(&message)
	// write message in chan to send to user
	writeToClientChat(*message)
	helpers.OkResponse(g, "room", message)
	return
}

/***
* get messages in room to user with limit and offset
 */
func MessageInRoom(g *gin.Context) {
	user := models.GetUserFromHeader(g)
	room_id, _ := strconv.Atoi(g.Param("id"))
	var room models.Room
	config.Db.Where("user1 = ? or user2 = ?", user.ID, user.ID).First(&room, room_id)
	if room.ID == 0 {
		helpers.ReturnNotFound(g, "this room not exist!")
		return
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

/***
* store new message to admin
 */
func StoreMessageINRoomNormal(g *gin.Context) {
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
	// get users and check if exist
	userRecevier := models.GetUserBYID(message.Receiver)
	if userRecevier.ID == 0 || user.ID == 0 {
		helpers.ReturnNotFound(g, "Not found this user!")
		return
	}
	// get room and check this user can chat or not
	var oldRoom models.Room
	config.Db.Where("user1 = ? and user2 = ?", message.Sender, message.Receiver).
		Or("user1 = ? and user2 = ?", message.Receiver, message.Sender).First(&oldRoom)
	// to unsure that user chat with admins only
	if oldRoom.ID == 0 && user.Type == 1 {
		// create new room and start start chat
		oldRoom.LastMessage = time.Now()
		oldRoom.FullName1 = user.FullName
		oldRoom.FullName2 = userRecevier.FullName

		config.Db.Create(&oldRoom)
		// write in chan new room from this user to both user
		writeToClientChatRoom(oldRoom, int(user.ID), int(userRecevier.ID))
		writeToClientChatRoom(oldRoom, int(userRecevier.ID), int(user.ID))

	} else if oldRoom.ID == 0 && user.Type != 1 {
		// want to chat with normal user
		helpers.ReturnNotFound(g, "can't chat with this user!")
		return
	} else {
		oldRoom.LastMessage = time.Now()
		config.Db.Save(&oldRoom)
	}
	// store message in database
	message.RoomID = int(oldRoom.ID)
	config.Db.Create(&message)
	// write message in chan to send to user
	writeToClientChat(*message)
	helpers.OkResponse(g, "room", message)
	return
}

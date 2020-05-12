package controllers

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"task-uxbert/config"
	helpers "task-uxbert/helper"
	"task-uxbert/models"
	"time"
)

/***
* start chat with with user and if you normal user can create new room
 */
func OpenRoom(g *gin.Context) {
	// get user from header
	user1 := models.GetUserFromHeader(g)
	// get user how you want to chat with him
	user_id, _ := strconv.Atoi(g.Param("id"))
	user2 := models.GetUserBYID(user_id)

	var room models.Room
	room.User1 = int(user1.ID)
	room.User2 = int(user2.ID)
	// check this user found or not
	if room.User2 == 0 || room.User1 == 0 {
		helpers.ReturnNotFound(g, "Not found this user!")
		return
	}
	result := make(map[string]interface{})
	var messages []models.Message
	var oldRoom models.Room
	// check if user have oldRoom with this user or not
	config.Db.Where("user1 = ? and user2 = ?", room.User1, room.User2).Or("user1 = ? and user2 = ?", room.User2, room.User1).
		Preload("Messages").First(&oldRoom)
	// if not found and this user is normal create new room
	if oldRoom.ID == 0 && user1.Type == 2 {
		room.LastMessage = time.Now()
		room.FullName1 = user1.FullName
		room.FullName2 = user2.FullName

		config.Db.Create(&room)
		// write in chan new room from this user to both user
		writeToClientChatRoom(room, int(user1.ID), int(user2.ID))
		writeToClientChatRoom(room, int(user2.ID), int(user1.ID))
		result["room"] = room
		result["message"] = []int{}
		helpers.OkResponse(g, "done create room", result)
		return
	}
	// get latest message to room and return it
	config.Db.Where("room_id = ?", oldRoom.ID).Limit(50).Order("id desc").Find(&messages)
	// make all massage to user how want to chat read
	go func() {
		config.Db.Table("messages").Where("room_id = ?", oldRoom.ID).
			Where("receiver = ?", user1.ID).UpdateColumn(map[string]interface{}{"is_read": 1})
	}()
	result["room"] = oldRoom
	result["messages"] = messages
	helpers.OkResponse(g, "room", result)
	return
}

func MYRooms(g *gin.Context) {
	var rooms []models.Room
	// get user from header
	user := models.GetUserFromHeader(g)
	// make limit and offset from query param
	limit_stirng := g.DefaultQuery("limit", "30")
	offset_string := g.DefaultQuery("offset", "0")
	limit, _ := strconv.Atoi(limit_stirng)
	offset, _ := strconv.Atoi(offset_string)
	if limit > 30 || limit == 0 {
		limit = 30
	}
	// get rooms and latest massage in each room
	config.Db.Where("user1 = ? or user2 = ?", user.ID, user.ID).
		Preload("Messages").Limit(limit).Offset(offset).Order("last_message desc").Find(&rooms)

	// return response
	helpers.OkResponse(g, "room", rooms)
	return
}

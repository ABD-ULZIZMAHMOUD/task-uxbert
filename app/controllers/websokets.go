package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"task-uxbert/config"
	"task-uxbert/models"
	"time"
)

// hubs to all user connection
var hubsChat = make(map[int]map[int]*websocket.Conn)

// to get user latest connection
var UserConnection = make(map[int]int)

// make chan to every connection
var UserChan = make(map[int]map[int]chan interface{})

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Data struct {
	Type        int         `json:"type"`
	MessageData interface{} `json:"message_data"`
}

/**
* this function open connection socket and put token of client to hubs
* to send all messages in it
 */
func WshandlerChat(w http.ResponseWriter, r *http.Request, token string) {
	// this time to unsure no two connection made in same time
	time.Sleep(3 * time.Second)
	// upgrade connection to websocket
	wsupgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := wsupgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Failed to set websocket upgrade", err)
		return
	}
	// check if user login or  not 
	user := models.GetUserBYToken(token)
	user_id := int(user.ID)
	if user_id == 0 {
		return
	}
	// to make user can open multiple connection to same user like many taps or mobile 
	index := 1
	// check if this first in hub or not
	if _, ok := hubsChat[user_id]; ok {
		index = UserConnection[user_id] + 1
	} else {
		// make new connection and put it in hub
		Connection := make(map[int]*websocket.Conn)
		chanelMap := make(map[int]chan interface{})
		hubsChat[user_id] = Connection
		UserChan[user_id] = chanelMap
	}
	// update latest connection
	UserConnection[user_id] = index
	// add connection to hub
	hubsChat[user_id][index] = conn
	// make chan and add to user chans
	chanle := make(chan interface{})
	UserChan[user_id][index] = chanle

	go func() {
		// start ping
		writeInChanPing(user_id, index)
	}()
	go func() {
		// start write to socket
		writeToClientFromChan(user_id, index)
	}()
	go func() {
		// start read pong from socket
		ReadFromClientSocket(user_id, index)
	}()
	return
}

func writeToClientChat(message models.Message) {
	if conne, ok := UserChan[message.Receiver]; ok {
		for _, conn := range conne {
			var data Data
			data.Type = 2
			data.MessageData = message
			conn <- data
		}
	}
	if conne, ok := UserChan[message.Sender]; ok {
		for _, conn := range conne {
			var data Data
			data.Type = 2
			data.MessageData = message
			conn <- data
		}
	}
	return
}

func writeToClientChatRoom(room models.Room, id int, user_id int) {
	user := models.GetUserBYID(user_id)
	object := make(map[string]interface{})
	object["user"] = user
	object["room"] = room
	if _, ok := hubsChat[user_id]; ok {
		object["online"] = 1
	} else {
		object["online"] = 2
	}
	if conne, ok := UserChan[id]; ok {
		for _, conn := range conne {
			var data Data
			data.Type = 5
			data.MessageData = object
			conn <- data
		}
	}
	return
}

/***
* add ping in chan every 4 seconds
 */
func writeInChanPing(user_id int, index int) {
	for {
		if _, ok := hubsChat[user_id]; ok {
			var data Data
			data.Type = 1
			data.MessageData = "ping"
			UserChan[user_id][index] <- data
		}
		time.Sleep(4 * time.Second)
	}
}

/***
* write any thing in chan in socket
 */
func writeToClientFromChan(id int, index int) {
	for {
		if conn, ok := hubsChat[id][index]; ok {
			conn.UnderlyingConn()
			client, ok := <-UserChan[id][index]
			if ok {
				err := conn.WriteJSON(client)
				if err != nil {
					RemoveFromHubsChatAndSendTOAllUserTest(id, index)
					return
				}
			}
		} else {
			break
		}
	}
	return
}

/***
* read from socket ping or any message else
 */
func ReadFromClientSocket(id int, index int) {
	count := 0
	timeOut := 12
	// check user connection by time out
	go func() {
		for timeOut > count {
			time.Sleep(1 * time.Second)
			count = count + 1
		}
		// if user lost connection by timeout remove form hub
		RemoveFromHubsChatAndSendTOAllUserTest(id, index)
		return
	}()
	for {
		if conn, ok := hubsChat[id][index]; ok {
			conn.UnderlyingConn()
			_, m, err := conn.ReadMessage()
			s := string(m)
			// if error when read message this connection is lost and remove from hub
			if err != nil {
				RemoveFromHubsChatAndSendTOAllUserTest(id, index)
				count = 13
				return
			} else {
				var data Data
				err = json.Unmarshal([]byte(s), &data)
				// this receive pong from socket
				if data.Type == 1 {
					count = 0
				}
				if data.Type == 2 {
					// if user read message receive message and make it read
					message_id := int(data.MessageData.(float64))
					config.Db.Table("messages").Where("id = ?", message_id).
						UpdateColumn(map[string]interface{}{"is_read": 1})
				}
			}
		} else {
			// if connection lost for any reason and not found in hub remove it
			RemoveFromHubsChatAndSendTOAllUserTest(id, index)
			count = 13
			return
		}
	}
}

// delete from hubs if connection lost
func RemoveFromHubsChatAndSendTOAllUserTest(id int, index int) {
	// this last connection to this user delete from hubs by index
	delete(hubsChat[id], index)
	delete(UserChan[id], index)
	// if this last connection to this user delete from hubs all connections
	if len(hubsChat[id]) == 0 {
		delete(hubsChat, id)
	}
}

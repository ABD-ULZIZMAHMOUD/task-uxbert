package models

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"task-uxbert/config"
)

type User struct {
	gorm.Model
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
	// customer= 2  and admin= 1 values
	Type int `json:"type"`
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

/**
* generate token based on user data
 */
func GenerateToken(user *User) {
	user.Token = uuid.New().String()
	config.Db.Save(&user)
	user.Password = ""
	return
}

/***
* get user by token
 */
func GetUserBYToken(token string) User {
	var user User
	config.Db.Where("token = ?", token).First(&user)
	return user
}

/***
* get user by ID
 */
func GetUserBYID(id int) User {
	var user User
	config.Db.Where("id = ?", id).First(&user)
	return user
}

/***
* get user from header
 */
func GetUserFromHeader(c *gin.Context) User {
	var user User
	_ = json.Unmarshal([]byte(c.GetHeader("user")), &user)
	return user
}

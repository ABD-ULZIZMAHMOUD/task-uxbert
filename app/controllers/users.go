package controllers

import (
	"github.com/gin-gonic/gin"
	"task-uxbert/app/validations"
	"task-uxbert/config"
	helpers "task-uxbert/helper"
	"task-uxbert/models"
)

func Login(g *gin.Context) {
	// init user login struct to validate request
	login := new(models.Login)
	/**
	* get request and parse it to validation
	* if there any error will return with message
	 */
	err := validations.LoginValidate(g, login)
	/***
	* return response if there an error if true you
	* this mean you have errors so we will return and bind data
	 */
	if helpers.ReturnNotValidRequest(err, g) {
		return
	}
	/**
	* check if user exists
	* check if user not blocked
	 */
	var user models.User
	config.Db.Where("email = ?", login.Email).First(&user)
	if user.ID == 0 {
		helpers.ReturnNotFound(g, "Not found this user!")
		return
	}
	/**
	* now check if password are valid
	* if user password is not valid we will return invalid email
	* or password
	 */
	check := helpers.CheckPasswordHash(login.Password, user.Password)
	if !check {
		helpers.ReturnNotFound(g, "your email or your password are not valid")
		return
	}
	// update token then return with the new data
	models.GenerateToken(&user)
	// now user is login we can return his info
	helpers.OkResponse(g, "you are login now", user)
}

/**
* Register new user on system
 */
func Register(g *gin.Context) {
	// init visitor User struct to validate request
	user := new(models.User)
	/**
	* get request and parse it to validation
	* if there any error will return with message
	 */
	err := validations.RegisterValidate(g, user)
	/***
	* return response if there an error if true you
	* this mean you have errors so we will return and bind data
	 */
	if helpers.ReturnNotValidRequest(err, g) {
		return
	}
	/**
	* check if this email exists database
	* if this email found will return
	 */
	config.Db.Find(&user, "email = ? ", user.Email)
	if user.ID != 0 {
		helpers.ReturnResponseWithMessageAndStatus(g, 400, "this email is exist!", false)
		return
	}
	//set type 2
	user.Type = 2
	user.Password, _ = helpers.HashPassword(user.Password)
	// create new user based on register struct
	config.Db.Create(&user)
	// now user is login we can return his info
	helpers.OkResponse(g, "Thank you for register in our system you can login now!", user)
}

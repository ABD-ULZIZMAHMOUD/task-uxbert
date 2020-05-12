package validations

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"task-uxbert/models"
)

/**
* validate login user request
 */
func LoginValidate(g *gin.Context, login *models.Login) *govalidator.Validator {
	/// Validation rules
	rules := govalidator.MapData{
		"email":    []string{"required", "min:6", "max:50", "email"},
		"password": []string{"required", "min:6", "max:50"},
	}
	opts := govalidator.Options{
		Request: g.Request, // request object
		Rules:   rules,     // rules map
		Data:    login,
	}
	return govalidator.New(opts)
}

/**
* validate Register user request
 */
func RegisterValidate(g *gin.Context, user *models.User) *govalidator.Validator {
	/// Validation rules
	rules := govalidator.MapData{
		"email":     []string{"required", "min:6", "max:50", "email"},
		"password":  []string{"required", "min:6", "max:50"},
		"full_name": []string{"required", "min:6", "max:50"},
	}
	opts := govalidator.Options{
		Request: g.Request, // request object
		Rules:   rules,     // rules map
		Data:    user,
	}
	return govalidator.New(opts)
}

package validations

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
	"task-uxbert/models"
)

/**
* validate Register user request
 */
func MassageValidate(g *gin.Context, message *models.Message) *govalidator.Validator {
	/// Validation rules
	rules := govalidator.MapData{
		"content":  []string{"required", "min:1"},
		"receiver": []string{"required", "min:1"},
	}
	opts := govalidator.Options{
		Request: g.Request, // request object
		Rules:   rules,     // rules map
		Data:    message,
	}
	return govalidator.New(opts)
}

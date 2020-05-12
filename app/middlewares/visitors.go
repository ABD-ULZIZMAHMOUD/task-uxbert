package middlewares

import (
	"github.com/gin-gonic/gin"
	"task-uxbert/config"
	"task-uxbert/models"
)

/**
* middleware to make sure that Authorization
* @ param   handler func
* @ return  func
 */
func UserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User
		config.Db.Table("users").Where("token = ?", c.GetHeader("Authorization")).Scan(&user)
		if user.ID == 0 {
			{
				RespondWithError(c, 401, "Invalid API toke")
				return
			}
		}
		c.Next()
		return
	}
}

package middlewares

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	helpers "task-uxbert/helper"
	"task-uxbert/models"
)

/***
* middleware to make sure that Authorization and user is normal
 */
func VisitorMiddleware() gin.HandlerFunc {
	return func(g *gin.Context) {
		if g.GetHeader("Authorization") == "" {
			helpers.ReturnForbidden(g, "Invalid API token!")
			return
		}

		user := models.GetUserBYToken(g.GetHeader("Authorization"))
		if user.ID == 0 || user.Type != 2 {
			helpers.ReturnForbidden(g, "Invalid API token!")
			g.Abort()
			return
		}

		userJson, _ := json.Marshal(&user)
		g.Request.Header.Set("user", string(userJson))
		g.Next()
		return
	}
}

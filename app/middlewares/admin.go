package middlewares

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	helpers "task-uxbert/helper"
	"task-uxbert/models"
)

/***
* middleware to make sure that Authorization and user is admin
 */
func AdminMiddleware() gin.HandlerFunc {
	return func(g *gin.Context) {
		models.GetUserBYToken(g.GetHeader("Authorization"))
		user := models.GetUserBYToken(g.GetHeader("Authorization"))
		if user.ID == 0 || user.Type != 1 {
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

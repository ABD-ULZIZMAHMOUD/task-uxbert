package routes

import (
	"github.com/gin-gonic/gin"
	"task-uxbert/app/controllers"
	"task-uxbert/config"
)

func AuthRoutes() {
	auth := config.Router.Group("users")
	{
		auth.POST("login", controllers.Login)
		auth.POST("register", controllers.Register)
		auth.GET("chats/ws/:token", func(g *gin.Context) {
			token := g.Param("token")
			controllers.WshandlerChat(g.Writer, g.Request, token)
		})
	}

}

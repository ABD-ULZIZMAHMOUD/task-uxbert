package routes

import (
	"task-uxbert/app/controllers"
	"task-uxbert/app/middlewares"
	"task-uxbert/config"
)

func VisitorRoutes() {
	normal := config.Router.Group("normal").Use(middlewares.VisitorMiddleware())
	{
		// id reference to user how you want to chat with him
		normal.GET("open-room/:id", controllers.OpenRoomNormal)
		normal.POST("store-message", controllers.StoreMessageINRoomNormal)
		// id reference to room_id
		normal.GET("messages/:id", controllers.MessageInRoom)
	}

}

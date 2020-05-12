package routes

import (
	"task-uxbert/app/controllers"
	"task-uxbert/app/middlewares"
	"task-uxbert/config"
)

func AdminRoutes() {
	admin := config.Router.Group("admin").Use(middlewares.AdminMiddleware())
	{
		// id reference to user how you want to chat with him
		admin.GET("open-room/:id", controllers.OpenRoomAdmin)
		admin.GET("my-rooms", controllers.MYRooms)
		// id reference to room_id
		admin.GET("messages/:id", controllers.MessageInRoom)
		admin.POST("store-message", controllers.StoreMessageINRoomAdmin)
	}

}

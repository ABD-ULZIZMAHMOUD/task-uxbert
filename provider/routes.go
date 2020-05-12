package provider

import (
	"os"
	"task-uxbert/app/middlewares"
	"task-uxbert/config"
	"task-uxbert/routes"
)

/***
* set all routes and middlewares
 */
func SetRoutes() {
	config.Router = config.SetupRouter()
	// set cors origin
	config.Router.Use(middlewares.CORSMiddleware())

	routes.AuthRoutes()
	routes.VisitorRoutes()
	routes.AdminRoutes()
	// run server
	_ = config.Router.Run(os.Getenv("PORT"))
}

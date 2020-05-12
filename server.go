package main

import (
	"github.com/subosito/gotenv"
	"task-uxbert/config"
	"task-uxbert/provider"
)

func main() {
	_ = gotenv.Load()
	config.ConnectToDatabase()
	provider.Migrate()
	provider.Seed()
	provider.SetRoutes()
}

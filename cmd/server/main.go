package main

import (
	"github.com/edaywalid/chat-app/internal/app"
	"github.com/edaywalid/chat-app/internal/router"
)

func main() {
	app, err := app.NewApp(".")
	if err != nil {
		panic(err)
	}

	r := router.SetupRoutes(app)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}

package router

import (
	"fmt"

	"github.com/edaywalid/chat-app/internal/app"
	"github.com/edaywalid/chat-app/internal/routes"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *app.App) *gin.Engine {
	router := gin.Default()
	router.Use(app.Middlewares.CorsMiddleware.CORSMiddleware())
	fmt.Println("is router nil?", router == nil)
	routes.NewAuthRoutes(app).Setup(router)
	routes.NewChatRoutes(app).Setup(router)
	return router
}

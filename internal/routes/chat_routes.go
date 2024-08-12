package routes

import (
	"github.com/edaywalid/chat-app/internal/app"
	"github.com/gin-gonic/gin"
)

type ChatRoutes struct {
	app *app.App
}

func NewChatRoutes(app *app.App) *ChatRoutes {
	return &ChatRoutes{app}
}

func (r *ChatRoutes) Setup(router *gin.Engine) {
	router.GET("/ws/message", r.app.Middlewares.AuthMiddleware.AuthMiddleWare(), r.app.Handlers.WebSocketHandler.SendOneOnOneMessage)
	router.GET("/ws/group", r.app.Middlewares.AuthMiddleware.AuthMiddleWare(), r.app.Handlers.WebSocketHandler.SendGroupMessage)
	router.GET("/ws/broadcast", r.app.Middlewares.AuthMiddleware.AuthMiddleWare(), r.app.Handlers.WebSocketHandler.SendBroadcastMessage)
}

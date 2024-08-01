package routes

import (
	"github.com/edaywalid/chat-app/internal/app"
	"github.com/gin-gonic/gin"
)

type AuthRoutes struct {
	app *app.App
}

func NewAuthRoutes(app *app.App) *AuthRoutes {
	return &AuthRoutes{app}
}

func (r *AuthRoutes) Setup(router *gin.Engine) {
	router.POST("/register", r.app.Handlers.AuthHandler.Register)
}

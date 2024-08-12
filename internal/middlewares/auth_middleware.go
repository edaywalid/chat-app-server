package middlewares

import (
	"github.com/edaywalid/chat-app/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type AuthMiddleware struct {
	jwtService *services.JwtService
}

func NewAuthMiddleware(jwtService *services.JwtService) *AuthMiddleware {
	return &AuthMiddleware{jwtService}
}

func (m *AuthMiddleware) AuthMiddleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken, err := ctx.Cookie("access_token")
		if err != nil {
			ctx.JSON(401, gin.H{"error": "unauthorized"})
			ctx.Abort()
			return
		}

		userID, err := m.jwtService.ValidateToken(accessToken)
		if err != nil {
			if err.(*jwt.ValidationError).Errors&jwt.ValidationErrorExpired == 0 {
				refreshToken, err := ctx.Cookie("refresh_token")
				if err != nil {
					ctx.JSON(401, gin.H{"error": "unauthorized"})
					ctx.Abort()
					return
				}

				newAccessToken, err := m.jwtService.RefreshToken(refreshToken)
				if err != nil {
					ctx.JSON(401, gin.H{"error": "unauthorized"})
					ctx.Abort()
					return
				}

				ctx.SetCookie("access_token", newAccessToken, 60*15, "/", "localhost", false, true)

				userID, err = m.jwtService.ValidateToken(newAccessToken)
				if err != nil {
					ctx.JSON(401, gin.H{"error": "unauthorized"})
					ctx.Abort()
					return
				}
			} else {
				ctx.JSON(401, gin.H{"error": "unauthorized"})
				ctx.Abort()
				return
			}
		}

		ctx.Set("user_id", userID)
		ctx.Next()
	}
}

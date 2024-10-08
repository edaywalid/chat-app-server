package handlers

import (
	"net/http"

	"github.com/edaywalid/chat-app/internal/services"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService}
}

func (h *AuthHandler) Register(c *gin.Context) {
	var input struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.authService.Register(input.Username, input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User registered successfully. Please check your email for verification code.",
	})
}

func (h *AuthHandler) Login(c *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tokenPair, err := h.authService.Login(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("access_token", tokenPair.AccessToken, 60*15, "/", "localhost", false, true)
	c.SetCookie("refresh_token", tokenPair.RefreshToken, 60*60*24*7, "/", "localhost", false, true)
	c.String(http.StatusOK, "Logged in successfully")
}

func (h *AuthHandler) ConfirmEmail(c *gin.Context) {
	var input struct {
		Email string `json:"email" binding:"required"`
		Code  string `json:"code" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.authService.ConfirmEmail(input.Email, input.Code); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Account verified successfully"})
}

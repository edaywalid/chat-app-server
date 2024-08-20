package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/edaywalid/chat-app/internal/models"
	"github.com/edaywalid/chat-app/internal/services"
	internal_ws "github.com/edaywalid/chat-app/internal/websocket"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WebSocketHandler struct {
	wsManager   *internal_ws.Manager
	chatService *services.ChatService
}

func NewWebSocketHandler(wsManager *internal_ws.Manager, chatService *services.ChatService) *WebSocketHandler {
	return &WebSocketHandler{
		wsManager:   wsManager,
		chatService: chatService,
	}
}

func (h *WebSocketHandler) SendOneOnOneMessage(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	extracted_userid, exists := c.Get("user_id")

	if !exists {
		return
	}
	userID := extracted_userid.(uuid.UUID)
	h.wsManager.AddClient(userID, conn)
	defer h.wsManager.RemoveClient(userID)
	go h.chatService.ListenForDirectMessage(userID)
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			panic(err)
			return
		}
		messageModel := &models.Message{}
		err = json.Unmarshal(message, &messageModel)
		if err != nil {
			panic(err)
			return
		}

		messageModel.SenderID = userID.String()
		messageModel.CreatedAt = time.Now()

		if messageModel.RecipientID != "" {
			err = h.chatService.SendDirectMessage(messageModel)
		}
	}
}

func (h *WebSocketHandler) SendGroupMessage(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(err)
		return
	}

	extracted_userid, exists := c.Get("user_id")
	if !exists {
		panic(err)
		return
	}
	userID := extracted_userid.(uuid.UUID)
	h.wsManager.AddClient(userID, conn)
	defer h.wsManager.RemoveClient(userID)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			panic(err)
			break
		}
		messageModel := &models.GroupMessage{}
		err = json.Unmarshal(message, &messageModel)
		if err != nil {
			panic(err)
			break
		}
		messageModel.SenderID = userID.String()

		if messageModel.GroupID == "" {
			err = h.chatService.SendGroupMessage(messageModel)
		}
	}
}

func (h *WebSocketHandler) SendBroadcastMessage(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		panic(err)
		return
	}

	extracted_userid, exists := c.Get("user_id")
	if !exists {
		panic(err)
		return
	}
	userID := extracted_userid.(uuid.UUID)
	h.wsManager.AddClient(userID, conn)
	defer h.wsManager.RemoveClient(userID)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			panic(err)
			break
		}
		messageModel := &models.Message{}
		err = json.Unmarshal(message, &messageModel)
		if err != nil {
			panic(err)
			break
		}
		messageModel.SenderID = userID.String()
		err = h.chatService.BroadcastMessage(messageModel)
	}
}

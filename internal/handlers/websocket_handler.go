package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/edaywalid/chat-app/internal/models"
	"github.com/edaywalid/chat-app/internal/services"
	internal_ws "github.com/edaywalid/chat-app/internal/websocket"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
		log.Printf("error upgrading to websocket: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error upgrading to websocket"})
		return
	}

	extracted_userid, exists := c.Get("user_id")

	if !exists {
		log.Printf("error getting user_id from context: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting user_id from context"})
		return
	}
	userID := extracted_userid.(uuid.UUID)
	h.wsManager.AddClient(userID, conn)

	defer func() {
		h.wsManager.RemoveClient(userID)
		conn.Close()
	}()

	go h.chatService.ListenForDirectMessage(userID)
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("error reading message: %v", err)
			break
		}
		messageModel := &models.Message{}
		err = json.Unmarshal(message, &messageModel)
		if err != nil {
			log.Printf("error unmarshalling message: %v", err)
			conn.WriteMessage(websocket.TextMessage, []byte("error unmarshalling message"))
			continue
		}

		messageModel.ID = primitive.NewObjectID()
		messageModel.SenderID = userID.String()
		messageModel.CreatedAt = time.Now()

		if messageModel.RecipientID != "" {
			err = h.chatService.SendDirectMessage(messageModel)
			if err != nil {
				log.Printf("error sending direct message: %v", err)
				conn.WriteMessage(websocket.TextMessage, []byte("error sending direct message"))
			}
		}
	}
}

func (h *WebSocketHandler) SendGroupMessage(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("error upgrading to websocket: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error upgrading to websocket"})
		return
	}

	extracted_userid, exists := c.Get("user_id")
	if !exists {
		log.Printf("error getting user_id from context: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting user_id from context"})
		return
	}
	userID := extracted_userid.(uuid.UUID)
	h.wsManager.AddClient(userID, conn)
	defer func() {
		h.wsManager.RemoveClient(userID)
		conn.Close()
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("error reading message: %v", err)
			break
		}
		messageModel := &models.GroupMessage{}
		err = json.Unmarshal(message, &messageModel)
		if err != nil {
			log.Printf("error unmarshalling message: %v", err)
			conn.WriteMessage(websocket.TextMessage, []byte("error unmarshalling message"))
			continue
		}

		messageModel.ID = primitive.NewObjectID()
		messageModel.SenderID = userID.String()

		if messageModel.GroupID == "" {
			err = h.chatService.SendGroupMessage(messageModel)
			if err != nil {
				log.Printf("error sending group message: %v", err)
				conn.WriteMessage(websocket.TextMessage, []byte("error sending group message"))
			}
		}
	}
}

func (h *WebSocketHandler) SendBroadcastMessage(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("error upgrading to websocket: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error upgrading to websocket"})
		return
	}

	extracted_userid, exists := c.Get("user_id")
	if !exists {
		log.Printf("error getting user_id from context: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error getting user_id from context"})
		return
	}
	userID := extracted_userid.(uuid.UUID)
	h.wsManager.AddClient(userID, conn)

	defer func() {
		h.wsManager.RemoveClient(userID)
		conn.Close()
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("error reading message: %v", err)
			break
		}
		messageModel := &models.Message{}
		err = json.Unmarshal(message, &messageModel)
		if err != nil {
			log.Printf("error unmarshalling message: %v", err)
			conn.WriteMessage(websocket.TextMessage, []byte("error unmarshalling message"))
			continue
		}

		messageModel.ID = primitive.NewObjectID()
		messageModel.SenderID = userID.String()

		err = h.chatService.BroadcastMessage(messageModel)
		if err != nil {
			log.Printf("error sending broadcast message: %v", err)
			conn.WriteMessage(websocket.TextMessage, []byte("error sending broadcast message"))
		}
	}
}

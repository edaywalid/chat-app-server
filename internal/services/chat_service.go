package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/edaywalid/chat-app/internal/models"
	"github.com/edaywalid/chat-app/internal/websocket"
	"github.com/google/uuid"
)

type ChatService struct {
	wsManager    *websocket.Manager
	redisService *RedisService
}

func NewChatService(wsManager *websocket.Manager, redisService *RedisService) *ChatService {
	return &ChatService{
		wsManager:    wsManager,
		redisService: redisService,
	}
}

func (s *ChatService) SendDirectMessage(message *models.Message) error {
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		return err
	}
	recipientID := message.RecipientID
	channel := fmt.Sprintf("direct:%s", recipientID)
	return s.redisService.Publish(channel, jsonMessage)
}

func (s *ChatService) SendGroupMessage(message *models.GroupMessage) error {
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		return err
	}
	groupID := message.GroupID
	channel := fmt.Sprintf("group:%d", groupID)
	return s.redisService.Publish(channel, jsonMessage)
}

func (s *ChatService) BroadcastMessage(message *models.Message) error {
	jsonMessage, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return s.redisService.Publish("broadcast", jsonMessage)
}

func (s *ChatService) ListenForDirectMessage(userID uuid.UUID) {
	channel := fmt.Sprintf("direct:%d", userID)
	pubsub := s.redisService.Subscribe(channel)
	defer pubsub.Close()

	for {
		msg, err := pubsub.ReceiveMessage(context.Background())
		if err != nil {
			log.Printf("Error receiving message: %v", err)
			continue
		}

		var message models.Message
		if err := json.Unmarshal([]byte(msg.Payload), &message); err != nil {
			log.Printf("Error unmarshalling message: %v", err)
			continue
		}

		client := s.wsManager.GetClient(userID)
		if client == nil {
			log.Printf("Client not found for user %s", userID)
			continue
		}

		if err := client.Conn.WriteJSON(message); err != nil {
			log.Printf("Error writing message to client: %v", err)
		}

	}
}

func (s *ChatService) ListenForGroupMessage(userID, groupID uuid.UUID) {
	channel := fmt.Sprintf("group:%d", groupID)
	pubsub := s.redisService.Subscribe(channel)
	defer pubsub.Close()

	for {
		msg, err := pubsub.ReceiveMessage(context.Background())
		if err != nil {
			log.Printf("Error receiving message: %v", err)
			continue
		}

		var message models.Message
		if err := json.Unmarshal([]byte(msg.Payload), &message); err != nil {
			log.Printf("Error unmarchilling message : %v", err)
		}

		client := s.wsManager.GetClient(userID)
		if client == nil {
			log.Printf("Client not found for user %s", userID)
			continue
		}

		if err := client.Conn.WriteJSON(message); err != nil {
			log.Printf("Error writing message to client: %v", err)
		}
	}
}

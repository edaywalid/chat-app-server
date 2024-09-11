package repositories

import (
	"fmt"

	"github.com/edaywalid/chat-app/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type ChatRepository struct {
	mongo_client *mongo.Client
}

func NewChatRepository(client *mongo.Client) *ChatRepository {
	return &ChatRepository{client}
}

func (r *ChatRepository) SaveMessage(message models.Message) error {
	// Save message to database
	collection := r.mongo_client.Database("chat").Collection("messages")
	res, err := collection.InsertOne(nil, message)
	if err != nil {
		return err
	}
	fmt.Println("Inserted message with ID: ", res.InsertedID)
	return nil
}
func (r *ChatRepository) SaveGroupMessage(message models.GroupMessage) error {

	// Save message to database
	collection := r.mongo_client.Database("chat").Collection("group_messages")
	res, err := collection.InsertOne(nil, message)
	if err != nil {
		return err
	}
	fmt.Println("Inserted message with ID: ", res.InsertedID)
	return nil
}

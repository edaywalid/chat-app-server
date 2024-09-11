package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	SenderID    string             `json:"sender_id" bson:"sender_id"`
	RecipientID string             `json:"recipient_id" bson:"recipient_id"`
	Content     string             `json:"content" bson:"content"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
}

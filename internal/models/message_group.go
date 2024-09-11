package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type GroupMessage struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	SenderID  string             `json:"sender_id" bson:"sender_id"`
	GroupID   string             `json:"group_id" bson:"group_id"`
	Content   string             `json:"content" bson:"content"`
	CreatedAt string             `json:"created_at" bson:"created_at"`
}

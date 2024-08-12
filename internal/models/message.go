package models

type Message struct {
	ID          string `json:"id" bson:"_id"`
	SenderID    string `json:"sender_id" bson:"sender_id"`
	RecipientID string `json:"recipient_id" bson:"recipient_id"`
	Content     string `json:"content" bson:"content"`
	CreatedAt   string `json:"created_at" bson:"created_at"`
}

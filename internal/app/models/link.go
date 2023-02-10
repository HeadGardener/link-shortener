package models

import "time"

type Link struct {
	URL        string    `json:"url" bson:"url"`
	ShortURL   string    `json:"short_url" bson:"short_url"`
	Identifier string    `json:"identifier" bson:"identifier"`
	ID         uint32    `json:"id" bson:"id"`
	UserID     string    `json:"user_id" bson:"user_id"`
	CreatedAt  time.Time `json:"created_at" bson:"created_at"`
}

type InputLink struct {
	URL       string `json:"url" bson:"url"`
	CustomURL string `json:"custom_url" bson:"custom_url"`
}

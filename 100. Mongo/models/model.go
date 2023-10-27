package model;

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogPost struct{
	ObjectId primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string `json:"title"`
	Content string `json:"content"`
	Genre []string	`json:"genre"`
	IsRead bool `json:"isRead"`
}
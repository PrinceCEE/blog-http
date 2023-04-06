package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostReactions struct {
	User primitive.ObjectID `json:"user,omitempty" bson:"user,omitempty"`
	Time primitive.DateTime
}

type Post struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Content string             `json:"content,omitempty" bson:"content,omitempty"`
}

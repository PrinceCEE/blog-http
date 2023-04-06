package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PostReactions struct {
	User primitive.ObjectID `json:"user,omitempty" bson:"user,omitempty"`
	Time primitive.DateTime `json:"time,omitempty" bson:"time,omitempty"`
}

type Post struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Content   string             `json:"content,omitempty" bson:"content,omitempty"`
	User      primitive.ObjectID `json:"user,omitempty" bson:"user,omitempty"`
	Likes     []PostReactions    `json:"likes,omitempty" bson:"likes,omitempty"`
	UnLikes   []PostReactions    `json:"unLikes,omitempty" bson:"unLikes,omitempty"`
	CreatedAt primitive.DateTime `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt primitive.DateTime `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

var PostSchemaValidation = bson.M{
	"$jsonSchema": bson.M{
		"bsonType":             "object",
		"required":             []string{"_id", "content", "user"},
		"additionalProperties": false,
		"properties": bson.M{
			"_id": bson.M{"bsonType": "objectId"},
			"content": bson.M{
				"bsonType":    "string",
				"description": "must be string and is required",
			},
			"user": bson.M{
				"bsonType":    "objectId",
				"description": "must be an objectID and is required",
			},
			"likes": bson.M{
				"bsonType": "array",
				"items": bson.M{
					"bsonType": "object",
					"required": []string{"user", "time"},
					"properties": bson.M{
						"user": bson.M{
							"bsonType":    "objectId",
							"description": "`user` in `likes` is required",
						},
						"time": bson.M{
							"bsonType":    "date",
							"description": "`time` in `likes` is required",
						},
					},
				},
			},
			"unLikes": bson.M{
				"bsonType": "array",
				"items": bson.M{
					"bsonType": "object",
					"required": []string{"user", "time"},
					"properties": bson.M{
						"user": bson.M{
							"bsonType":    "objectId",
							"description": "`user` in `likes` is required",
						},
						"time": bson.M{
							"bsonType":    "date",
							"description": "`time` in `likes` is required",
						},
					},
				},
			},
			"createdAt": bson.M{
				"bsonType":    "date",
				"description": "createdAt is required",
			},
			"updatedAt": bson.M{
				"bsonType": "date",
			},
		},
	},
}

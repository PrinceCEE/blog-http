package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentReactions struct {
	User primitive.ObjectID `json:"user,omitempty" bson:"user,omitempty"`
	Time primitive.DateTime `json:"time,omitempty" bson:"time,omitempty"`
}

type Comment struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Post      primitive.ObjectID `json:"post,omitempty" bson:"post,omitempty"`
	Content   string             `json:"content,omitempty" bson:"content,omitempty"`
	User      primitive.ObjectID `json:"user,omitempty" bson:"user,omitempty"`
	Likes     []PostReactions    `json:"likes,omitempty" bson:"likes,omitempty"`
	UnLikes   []PostReactions    `json:"unLikes,omitempty" bson:"unLikes,omitempty"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

func (c *Comment) MarshalBSON() ([]byte, error) {
	if c.CreatedAt.IsZero() {
		c.CreatedAt = time.Now()
	}
	c.UpdatedAt = time.Now()

	type _c Comment
	return bson.Marshal((*_c)(c))
}

var CommentSchemaValidation = bson.M{
	"$jsonSchema": bson.M{
		"bsonType":             "object",
		"required":             []string{"_id", "content", "user", "post"},
		"additionalProperties": false,
		"properties": bson.M{
			"_id": bson.M{"bsonType": "objectId"},
			"post": bson.M{
				"bsonType":    "objectId",
				"description": "posot is required",
			},
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

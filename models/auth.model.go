package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Auth struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	User     primitive.ObjectID `json:"user,omitempty" bson:"user,omitempty"`
	Password string             `json:"password" bson:"password"`
}

var AuthSchemaValidation = bson.M{
	"$jsonSchema": bson.M{
		"bsonType":             "object",
		"required":             []string{"_id", "content", "user"},
		"additionalProperties": false,
		"properties": bson.M{
			"_id": bson.M{"bsonType": "objectId"},
			"user": bson.M{
				"bsonType":    "objectId",
				"description": "must be an objectId and is required",
			},
			"password": bson.M{
				"bsonType":    "string",
				"description": "must be a string and is required",
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

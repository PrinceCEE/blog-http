package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Code struct {
	ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	User   primitive.ObjectID `json:"user,omitempty" bson:"user,omitempty"`
	Code   string             `json:"code" bson:"code"`
	IsUsed bool               `json:"isUsed" bson:"isUsed"`
}

var CodeSchemaValidation = bson.M{
	"$jsonSchema": bson.M{
		"bsonType":             "object",
		"required":             []string{"_id", "code", "user", "isUsed"},
		"additionalProperties": false,
		"properties": bson.M{
			"_id": bson.M{"bsonType": "objectId"},
			"user": bson.M{
				"bsonType":    "objectId",
				"description": "must be an objectId and is required",
			},
			"code": bson.M{
				"bsonType":    "string",
				"description": "must be a string and is required",
			},
			"isUsed": bson.M{
				"bsonType":    "bool",
				"description": "must be a boolean and is required",
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

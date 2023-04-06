package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName  string             `json:"lastName,omitempty" bson:"lastName,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
	CreatedAt primitive.DateTime `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt primitive.DateTime `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

var (
	strDesc = "must be a string and is required"
	strName = "string"
)

var UserSchemaValidation = bson.M{
	"$jsonSchema": bson.M{
		"bsonType":             "object",
		"additionalProperties": false,
		"required":             []string{"_id", "firstName", "lastName", "email"},
		"properties": bson.M{
			"_id": bson.M{
				"bsonType": "objectId",
			},
			"firstName": bson.M{
				"bsonType":    strName,
				"description": strDesc,
			},
			"lastName": bson.M{
				"bsonType":    strName,
				"description": strDesc,
			},
			"email": bson.M{
				"bsonType":    strName,
				"description": strDesc,
			},
		},
	},
}

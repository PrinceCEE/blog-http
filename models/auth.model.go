package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Auth struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	User      primitive.ObjectID `json:"user,omitempty" bson:"user,omitempty"`
	Password  string             `json:"password" bson:"password"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

func (a *Auth) MarshalBSON() ([]byte, error) {
	if a.CreatedAt.IsZero() {
		a.CreatedAt = time.Now()
	}
	a.UpdatedAt = time.Now()

	type _a Auth
	return bson.Marshal((*_a)(a))
}

var AuthSchemaValidation = bson.M{
	"$jsonSchema": bson.M{
		"bsonType":             "object",
		"required":             []string{"_id", "password", "user"},
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

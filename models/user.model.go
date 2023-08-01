package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	FirstName string             `json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName  string             `json:"lastName,omitempty" bson:"lastName,omitempty"`
	Email     string             `json:"email,omitempty" bson:"email,omitempty"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}

func (u *User) MarshalBSON() ([]byte, error) {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	u.UpdatedAt = time.Now()

	type _u User
	return bson.Marshal((*_u)(u))
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
			"createdAt": bson.M{
				"bsonType": "date",
			},
			"updatedAt": bson.M{
				"bsonType": "date",
			},
		},
	},
}

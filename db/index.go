package db

import (
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	UserCollection    mongo.Collection
	PostCollection    mongo.Collection
	CommentCollection mongo.Collection
	AuthCollection    mongo.Collection
)

/*
* 1. Connect to Mongo
* 2. Create collections with the validators and export them
 */

func Connect() {
	// connect to mongoDB
	// create the collections and assign them
}

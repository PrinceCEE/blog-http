package db

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	UserCollection    *mongo.Collection
	PostCollection    *mongo.Collection
	CommentCollection *mongo.Collection
	AuthCollection    *mongo.Collection
)

const (
	DB_NAME                 = "blog-http"
	USER_COLLECTION_NAME    = "users"
	POST_COLLECTION_NAME    = "posts"
	COMMENT_COLLECTION_NAME = "comments"
	AUTH_COLLECTION_NAME    = "auths"
)

// Connect to the DB
// And initiliase the Collections
func Connect() error {
	dbUrl := os.Getenv("DB_URL")
	opts := options.Client().ApplyURI(dbUrl)
	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return err
	}

	if err = client.Ping(context.Background(), nil); err != nil {
		return err
	}

	UserCollection = client.Database(DB_NAME).Collection(USER_COLLECTION_NAME)
	PostCollection = client.Database(DB_NAME).Collection(POST_COLLECTION_NAME)
	CommentCollection = client.Database(DB_NAME).Collection(COMMENT_COLLECTION_NAME)
	AuthCollection = client.Database(DB_NAME).Collection(AUTH_COLLECTION_NAME)

	return nil
}

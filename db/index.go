package db

import (
	"blog-http/models"
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/exp/slices"
)

var (
	UserCollection    *mongo.Collection = &mongo.Collection{}
	PostCollection    *mongo.Collection = &mongo.Collection{}
	CommentCollection *mongo.Collection = &mongo.Collection{}
	AuthCollection    *mongo.Collection = &mongo.Collection{}
	CodeCollection    *mongo.Collection = &mongo.Collection{}
)

const (
	DB_NAME                 = "blog-http"
	USER_COLLECTION_NAME    = "users"
	POST_COLLECTION_NAME    = "posts"
	COMMENT_COLLECTION_NAME = "comments"
	AUTH_COLLECTION_NAME    = "auths"
	CODE_COLLECTION_NAME    = "codes"
)

type collectionConfig struct {
	name   string
	coll   *mongo.Collection
	schema bson.M
}

// Connect to the DB
// And initiliase the Collections
func Connect() error {
	todoCtx := context.TODO()
	dbUrl := os.Getenv("DB_URL")
	opts := options.Client().ApplyURI(dbUrl)
	client, err := mongo.Connect(todoCtx, opts)
	if err != nil {
		return err
	}

	if err = client.Ping(todoCtx, nil); err != nil {
		return err
	}

	db := client.Database(DB_NAME)
	collectionsConfig := []collectionConfig{
		{USER_COLLECTION_NAME, UserCollection, models.UserSchemaValidation},
		{AUTH_COLLECTION_NAME, AuthCollection, models.AuthSchemaValidation},
		{POST_COLLECTION_NAME, PostCollection, models.PostSchemaValidation},
		{COMMENT_COLLECTION_NAME, CommentCollection, models.CommentSchemaValidation},
		{CODE_COLLECTION_NAME, CodeCollection, models.CodeSchemaValidation},
	}
	collectionNames, err := db.ListCollectionNames(todoCtx, bson.M{})
	if err != nil {
		return err
	}

	for _, v := range collectionsConfig {
		v := v
		index := slices.IndexFunc(collectionNames, func(n string) bool {
			return n == v.name
		})

		if index == -1 {
			err = db.CreateCollection(todoCtx, v.name, options.CreateCollection().SetValidator(v.schema))
			if err != nil {
				return err
			}
		}

		*v.coll = *db.Collection(v.name)
	}
	return nil
}

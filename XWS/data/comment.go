package data

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	author  User
	content string
}

type Comments []*Comment

func getCommentByPost(id string) (Comments, error) {
	postCollection := Client.Database("xws").Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	docID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		fmt.Println("[ERROR] can't convert string to ObjectID", err)
		return nil, ErrPostNotFound
	}

	var targetPost Post
	err = postCollection.FindOne(ctx, bson.M{"_id": docID}).Decode(&targetPost)
	if err != nil {
		fmt.Println("[ERROR] FindOne() ObjectIDFromHex :", err)
		return nil, ErrPostNotFound
	}

	return targetPost.Comments, nil
}

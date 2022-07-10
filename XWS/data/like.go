package data

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Like struct {
	Author  string `json:"username,omitempty"`
	Content bool   `json:"boolean,omitempty"`
}

type Likes []*Like

func getLikeCountByPost(id string) (int, int, error) {
	postCollection := Client.Database("xws").Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	docID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		fmt.Println("[ERROR] can't convert string to ObjectID", err)
		return 0, 0, ErrPostNotFound
	}

	var targetPost Post
	err = postCollection.FindOne(ctx, bson.M{"_id": docID}).Decode(&targetPost)
	if err != nil {
		fmt.Println("[ERROR] FindOne() ObjectIDFromHex :", err)
		return 0, 0, ErrPostNotFound
	}

	likeCount := 0
	dislikeCount := 0

	for i := 1; i < len(targetPost.Likes); i++ {
		if targetPost.Likes[i].Content == true {
			likeCount++
		} else {
			dislikeCount++
		}
	}
	return likeCount, dislikeCount, nil
}

package data

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	Username   string             `json:"username,omitempty"`
	TxtContent string             `json:"text,omitempty"`
	Hyperlink  string             `json:"link,omitempty"`
	Comments   Comments           `json:"-"`
	Likes      Likes              `json:"-"`
}

type Posts []*Post

var ErrPostNotFound = fmt.Errorf("Post not found")

func AddPostToDB(post Post) primitive.ObjectID {
	fmt.Println("[DEBUG] entered adding post to db")
	postCollection := Client.Database("xws").Collection("posts")
	doc, err := bson.Marshal(post)
	if err != nil {
		fmt.Println("[ERROR] marshaling to bson.d")
	}
	result, err := postCollection.InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Println("[ERROR] inserting into database")
	}
	postID := result.InsertedID.(primitive.ObjectID)
	fmt.Println(postID)
	return postID
}

func AddCommentToPost(post Post, comment Comment) primitive.ObjectID {
	fmt.Println("[DEBUG] entered adding post to db")
	postCollection := Client.Database("xws").Collection("posts")
	post.Comments = append(post.Comments, &comment)
	doc, err := bson.Marshal(post)
	if err != nil {
		fmt.Println("[ERROR] marshaling to bson.d")
	}
	result, err := postCollection.UpdateOne(context.TODO(), bson.M{"ID": post.ID}, doc)
	if err != nil {
		fmt.Println("[ERROR] inserting into database")
	}
	postID := result.UpsertedID.(primitive.ObjectID)
	fmt.Println(postID)
	return postID
}

func AddLikeToPost(post Post, like Like) primitive.ObjectID {
	fmt.Println("[DEBUG] entered adding post to db")
	postCollection := Client.Database("xws").Collection("posts")
	post.Likes = append(post.Likes, &like)
	doc, err := bson.Marshal(post)
	if err != nil {
		fmt.Println("[ERROR] marshaling to bson.d")
	}
	result, err := postCollection.UpdateOne(context.TODO(), bson.M{"ID": post.ID}, doc)
	if err != nil {
		fmt.Println("[ERROR] inserting into database")
	}
	postID := result.UpsertedID.(primitive.ObjectID)
	fmt.Println(postID)
	return postID
}

package data

import (
	"context"
	"fmt"
	"time"

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

func GetPostsUser(username string) (Posts, []primitive.ObjectID) {
	fmt.Printf("[DEBUG] entered getting posts for user: %s\n", username)
	var posts Posts
	postsCollection := Client.Database("xws").Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	filter := bson.D{{"username", username}}
	cursor, err := postsCollection.Find(ctx, filter)
	if err != nil {
		fmt.Println("[ERROR] reading from db")
	}

	cursor.All(ctx, &posts)
	postIDs := make([]primitive.ObjectID, len(posts))
	for i := 0; i < len(posts); i++ {
		postIDs[i] = posts[i].ID
	}
	return posts, postIDs
}

func AddCommentToPost(id string, comment Comment) primitive.ObjectID {
	postCollection := Client.Database("xws").Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	docID, err := primitive.ObjectIDFromHex(id)

	var foundPost Post
	err = postCollection.FindOne(ctx, bson.M{"_id": docID}).Decode(&foundPost)

	foundPost.Comments = append(foundPost.Comments, &comment)

	if err != nil {
		fmt.Println("[ERROR] marshaling to bson.d")
	}

	result, err := postCollection.UpdateOne(context.TODO(), bson.M{"_id": foundPost.ID}, foundPost)
	if err != nil {
		fmt.Println("[ERROR] inserting into database")
	}

	postID := result.UpsertedID.(primitive.ObjectID)
	fmt.Println(postID)
	return postID
}

func AddLikeToPost(id string, like Like) primitive.ObjectID {
	postCollection := Client.Database("xws").Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	docID, err := primitive.ObjectIDFromHex(id)

	var foundPost Post
	err = postCollection.FindOne(ctx, bson.M{"_id": docID}).Decode(&foundPost)

	foundPost.Likes = append(foundPost.Likes, &like)
	if err != nil {
		fmt.Println("[ERROR] marshaling to bson.d")
	}
	result, err := postCollection.UpdateOne(context.TODO(), bson.M{"_id": foundPost.ID}, foundPost)
	if err != nil {
		fmt.Println("[ERROR] inserting into database")
	}
	postID := result.UpsertedID.(primitive.ObjectID)
	fmt.Println(postID)
	return postID
}

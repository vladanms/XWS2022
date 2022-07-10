package data

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username   string             `json:"username,omitempty"`
	TxtContent string             `json:"text,omitempty"`
	Hyperlink  string             `json:"link,omitempty"`
	Comments   Comments           `json:"comments,omitempty"`
	Likes      Likes              `json:"likes,omitempty"`
	ImagePath  string             `json:"imagePath,omitempty"`
	ImageName  string             `json:"imageName,omitempty"`
}

type PostNotification struct {
	PostNotifID primitive.ObjectID `bson:"_id,omitempty"`
	PostID      primitive.ObjectID
	Recipient   string
}

type Posts []*Post
type PostNotifications []*PostNotification

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

func AddCommentToPost(id primitive.ObjectID, comment Comment) primitive.ObjectID {
	fmt.Println("[DEBUG] entered adding comment to post")
	postCollection := Client.Database("xws").Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	var foundPost Post
	err := postCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&foundPost)

	foundPost.Comments = append(foundPost.Comments, &comment)

	if err != nil {
		fmt.Println("[ERROR] marshaling to bson.d")
	}

	update := bson.M{"$set": bson.M{"comments": foundPost.Comments}}
	_, err = postCollection.UpdateOne(ctx, bson.M{"_id": foundPost.ID}, update)
	if err != nil {
		fmt.Println("[ERROR] inserting into database")
		fmt.Println(err)
	}
	fmt.Println("[DEBUG] added comment to post")
	return primitive.NilObjectID
}

func AddLikeToPost(id primitive.ObjectID, like Like) primitive.ObjectID {
	fmt.Println("[DEBUG] entered adding like to database")
	postCollection := Client.Database("xws").Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	fmt.Println(id)
	fmt.Println(like)
	var foundPost Post
	err := postCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&foundPost)
	//if user already liked this post
	for _, current := range foundPost.Likes {
		if current.Author == like.Author {
			return primitive.NilObjectID
		}
	}
	foundPost.Likes = append(foundPost.Likes, &like)
	if err != nil {
		fmt.Println("[ERROR] marshaling to bson.d")
		return primitive.NilObjectID
	}
	update := bson.M{"$set": bson.M{"likes": foundPost.Likes}}
	_, err = postCollection.UpdateOne(ctx, bson.M{"_id": foundPost.ID}, update)
	if err != nil {
		fmt.Println("[ERROR] inserting into database")
		fmt.Println(err)
	}

	return primitive.NilObjectID
}
func GetSinglePost(postID primitive.ObjectID) *Post {
	var post Post
	postsCollection := Client.Database("xws").Collection("posts")
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	filter := bson.D{{"_id", postID}}
	err := postsCollection.FindOne(ctx, filter).Decode(&post)
	if err != nil {
		return nil
	}
	return &post

}

func AddPostNotificationsToDB(postNotifications PostNotifications) {
	fmt.Println("[DEBUG] entered adding post notifs to db")
	postNotificationsCollection := Client.Database("xws").Collection("postNotifications")
	docs := make([]interface{}, 0)
	for i := 0; i < len(postNotifications); i++ {
		doc, err := bson.Marshal(postNotifications[i])
		if err != nil {
			fmt.Println("[ERROR] marshaling to bson.d")
		}
		docs = append(docs, doc)
	}
	_, err := postNotificationsCollection.InsertMany(context.TODO(), docs)
	if err != nil && err != bson.ErrNilReader {
		fmt.Println("[ERROR] inserting into database")
	}

}
func GetPostNotificationsForUser(username string) PostNotifications {
	postNotificationsCollection := Client.Database("xws").Collection("postNotifications")
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	filter := bson.M{"recipient": username}

	cur, err := postNotificationsCollection.Find(ctx, filter)
	if err != nil {
		fmt.Println("[ERROR] reading from database")
		return nil
	}
	var postNotifications PostNotifications
	err = cur.All(ctx, &postNotifications)
	if err != nil {
		fmt.Println("[ERROR] decoding results")
		return nil
	}
	err = DeletePostNotifications(postNotifications)
	if err != nil {
		fmt.Println("[ERROR] deleting notifications")
		return nil
	}
	return postNotifications
}

func DeletePostNotifications(postNotifications PostNotifications) error {
	postNotificationsCollection := Client.Database("xws").Collection("postNotifications")
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	for i := 0; i < len(postNotifications); i++ {
		_, err := postNotificationsCollection.DeleteOne(ctx, bson.M{"_id": postNotifications[i].PostNotifID})
		if err != nil {
			return err
		}
	}
	return nil
}

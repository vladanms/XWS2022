package server

import (
	"context"
	"fmt"
	"posts_service/database"
	protos "posts_service/protos/posts"

	"github.com/hashicorp/go-hclog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Posts struct {
	log hclog.Logger
	protos.UnimplementedPostsServer
}

func NewPosts(l hclog.Logger, u protos.UnimplementedPostsServer) *Posts {
	return &Posts{l, u}
}

func (p *Posts) GetAllPostsFromUser(ctx context.Context, pr *protos.PostsRequest) (*protos.PostsResponse, error) {
	p.log.Info("Get posts from user handle")
	PostsCollection := database.Client.Database("xws").Collection("posts")
	p.log.Info(pr.Username)
	cursor, err := PostsCollection.Find(ctx, bson.D{{"username", pr.Username}})
	if err != nil {
		p.log.Error("retrieving posts from database")
	}
	allPosts := protos.PostsResponse{}
	err = cursor.All(ctx, &allPosts.Results)
	if err != nil {
		fmt.Println(err)
		p.log.Error("decoding posts from bson to struct")
	}

	return &allPosts, err
}

func (p *Posts) GetNotificationPosts(ctx context.Context, npr *protos.NotificationPostsRequest) (*protos.PostsResponse, error) {
	p.log.Info("Get Notification posts for user")
	PostNotificationsCollection := database.Client.Database("xws").Collection("postNotifications")
	cursor, err := PostNotificationsCollection.Find(ctx, bson.M{"recipient": npr.Username})
	if err != nil {
		p.log.Error("retrieving post notifs from database")
		return nil, err
	}
	if cursor.RemainingBatchLength() == 0 {
		return nil, fmt.Errorf("no notifs")
	}
	postNotifications := protos.PostNotifications{}
	err = cursor.All(ctx, &postNotifications.PostNotifs)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	DeletePostNotifications(ctx, &postNotifications)
	var posts protos.PostsResponse
	for _, notif := range postNotifications.PostNotifs {
		req := protos.SinglePostRequest{ID: notif.PostID}
		post, err := GetSinglePost(ctx, &req)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		posts.Results = append(posts.Results, post)
	}

	return &posts, nil
}
func DeletePostNotifications(ctx context.Context, postNotifs *protos.PostNotifications) error {
	postNotificationsCollection := database.Client.Database("xws").Collection("postNotifications")
	for i := 0; i < len(postNotifs.PostNotifs); i++ {
		id, err := primitive.ObjectIDFromHex(postNotifs.PostNotifs[i].NotifID)
		if err != nil {
			fmt.Println(err)
			return err
		}
		_, err = postNotificationsCollection.DeleteOne(ctx, bson.M{"_id": id})
		if err != nil {
			return err
		}
	}
	return nil
}

func GetSinglePost(ctx context.Context, npr *protos.SinglePostRequest) (*protos.PostResponse, error) {
	PostsCollection := database.Client.Database("xws").Collection("posts")
	var post protos.PostResponse
	postID, err := primitive.ObjectIDFromHex(npr.ID)
	if err != nil {
		return nil, err
	}
	err = PostsCollection.FindOne(ctx, bson.M{"_id": postID}).Decode(&post)
	if err != nil {
		return nil, err
	}

	return &post, nil
}
func (p *Posts) CreatePost(ctx context.Context, npr *protos.CreateRequest) (*protos.CreateResponse, error) {
	PostsCollection := database.Client.Database("xws").Collection("posts")
	doc, err := bson.Marshal(npr)
	if err != nil {
		fmt.Println("[ERROR] marshaling to bson")
		fmt.Println(err)
	}
	res, err := PostsCollection.InsertOne(ctx, doc)
	if err != nil {
		p.log.Error("inserting post into database")
		return nil, err
	}
	postID := res.InsertedID.(primitive.ObjectID)
	fmt.Println(postID)

	return &protos.CreateResponse{PostID: postID.Hex()}, nil
}

func (p *Posts) AddCommentToPost(ctx context.Context, npr *protos.CommentRequest) (*protos.CommentResponse, error) {
	postCollection := database.Client.Database("xws").Collection("posts")
	var foundPost protos.PostResponse
	fmt.Println(npr.PostID)
	postID, err := primitive.ObjectIDFromHex(npr.PostID)
	if err != nil {
		p.log.Error("converting from string to ObjectID")
		return nil, err
	}
	fmt.Println(postID)
	err = postCollection.FindOne(ctx, bson.M{"_id": postID}).Decode(&foundPost)
	comment := protos.Comment{Author: npr.Author, Content: npr.Content}
	foundPost.Comments = append(foundPost.Comments, &comment)
	fmt.Println(foundPost.Comments)
	if err != nil {
		fmt.Println("[ERROR] marshaling to bson.d")
		fmt.Println(err)
		return nil, err
	}
	update := bson.M{"$set": bson.M{"comments": foundPost.Comments}}
	fmt.Println(update)
	_, err = postCollection.UpdateOne(ctx, bson.M{"_id": postID}, update)
	if err != nil {
		fmt.Println("[ERROR] inserting into database")
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("[DEBUG] added comment to post")
	return &protos.CommentResponse{}, nil
}

func (p *Posts) AddLikeToPost(ctx context.Context, npr *protos.LikeRequest) (*protos.LikeResponse, error) {
	postCollection := database.Client.Database("xws").Collection("posts")
	var foundPost protos.PostResponse
	like := protos.Like{Author: npr.Author, Content: npr.Content}
	postID, err := primitive.ObjectIDFromHex(npr.PostID)
	if err != nil {
		p.log.Error("converting from string to ObjectID")
		return nil, err
	}
	err = postCollection.FindOne(ctx, bson.M{"_id": postID}).Decode(&foundPost)
	if err != nil {
		fmt.Println("[ERROR] marshaling to bson")
		fmt.Println(err)
		return nil, err
	}
	flag := false
	for _, current := range foundPost.Likes {
		if current.Author == like.Author {
			current.Content = like.Content
			flag = true
		}
	}
	if !flag {
		foundPost.Likes = append(foundPost.Likes, &like)
	}
	update := bson.M{"$set": bson.M{"likes": foundPost.Likes}}
	_, err = postCollection.UpdateOne(ctx, bson.M{"_id": postID}, update)
	if err != nil {
		fmt.Println("[ERROR] inserting into database")
		fmt.Println(err)
		return nil, err
	}
	fmt.Println("[DEBUG] added comment to post")
	return &protos.LikeResponse{}, nil

}

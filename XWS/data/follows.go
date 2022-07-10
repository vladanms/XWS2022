package data

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Follow struct {
	FollowID primitive.ObjectID `bson:"_id,omitempty"`
	Follower string             `json:"follower"`
	Followee string             `json:"followee"`
}

type FollowRequest struct {
	FollowRequestID primitive.ObjectID `bson:"_id,omitempty"`
	Requester       string             `json:"requester"`
	Requestee       string             `json:"requestee"`
}

type Follows []*Follow
type FollowRequests []*FollowRequest

func AddFollowToDB(follow Follow) {
	followsCollection := Client.Database("xws").Collection("follows")

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	filter := bson.D{{"follower", follow.Follower}, {"followee", follow.Followee}}
	var result Follow
	err := followsCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil && err != mongo.ErrNoDocuments {
		return
	}
	if result.FollowID != primitive.NilObjectID {
		fmt.Println("[ERROR] already following")
		return
	}

	doc, err := bson.Marshal(follow)
	if err != nil {
		fmt.Println("[ERROR] marshaling to bson.d")
	}
	_, err = followsCollection.InsertOne(ctx, doc)
	if err != nil {
		fmt.Println("[ERROR] inserting into database")
	}
}

func AddFollowRequestToDB(followRequest FollowRequest) {
	followRequestsCollection := Client.Database("xws").Collection("followRequests")

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	filter := bson.D{{"requester", followRequest.Requester}, {"requestee", followRequest.Requestee}}
	var result Follow
	err := followRequestsCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil && err != mongo.ErrNoDocuments {
		return
	}
	if result.FollowID != primitive.NilObjectID {
		fmt.Println("[ERROR] already requested to follow")
		return
	}

	doc, err := bson.Marshal(followRequest)
	if err != nil {
		fmt.Println("[ERROR] marshaling to bson.d")
	}
	_, err = followRequestsCollection.InsertOne(ctx, doc)
	if err != nil {
		fmt.Println("[ERROR] inserting into database")
	}
}
func GetFollow(follower, followee string) (*Follow, error) {
	followsCollection := Client.Database("xws").Collection("follows")

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	var result Follow
	filter := bson.D{{"follower", follower}, {"followee", followee}}

	sRes := followsCollection.FindOne(ctx, filter)
	if sRes.Err() != nil {
		fmt.Println("[DEBUG] no follow")
		return nil, fmt.Errorf("follow not found")
	}
	err := sRes.Decode(&result)
	if err != nil {
		fmt.Println("[ERROR] could not decode result")
		return nil, err
	}
	return &result, nil
}

func GetFollowRequest(requester, requestee string) (*FollowRequest, error) {
	followRequestsCollection := Client.Database("xws").Collection("followRequests")

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	var result FollowRequest
	filter := bson.D{{"requester", requester}, {"requestee", requestee}}

	sRes := followRequestsCollection.FindOne(ctx, filter)
	if sRes.Err() != nil {
		fmt.Println("[ERROR] no request")
		return nil, fmt.Errorf("request not found")
	}
	err := sRes.Decode(&result)
	if err != nil {
		fmt.Println("[ERROR] could not decode result")
		return nil, err
	}
	return &result, nil
}
func GetAllFollowRequests(requestee string) (FollowRequests, error) {
	followRequestsCollection := Client.Database("xws").Collection("followRequests")

	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	fmt.Println(requestee)
	var result FollowRequests
	filter := bson.M{"requestee": requestee}
	cursor, err := followRequestsCollection.Find(ctx, filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}
	err = cursor.All(ctx, &result)
	if err != nil {
		fmt.Println("[ERROR] decoding result")
		return nil, err
	}
	fmt.Println("[DEBUG] finished getting requests successfully")
	return result, nil
}

func DeleteFollowRequest(followRequestID primitive.ObjectID) error {
	followRequestsCollection := Client.Database("xws").Collection("followRequests")
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)

	_, err := followRequestsCollection.DeleteOne(ctx, bson.M{"_id": followRequestID})
	if err != nil {
		return err
	}
	return nil
}

func GetAllFollowers(username string) []string {
	fmt.Println("[DEBUG] entered getting followers")
	followsCollection := Client.Database("xws").Collection("follows")
	ctx, _ := context.WithTimeout(context.Background(), 15*time.Second)
	var follows Follows
	filter := bson.D{{"followee", username}}
	cursor, err := followsCollection.Find(ctx, filter)
	if err != nil {
		fmt.Println("[ERROR] retrieving follows")
		return nil
	}
	cursor.All(ctx, &follows)
	fmt.Printf("Number of followers in get all followers %d\n", len(follows))
	followers := make([]string, len(follows))
	for i := 0; i < len(follows); i++ {
		followers[i] = follows[i].Follower
	}
	return followers
}

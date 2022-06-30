package server

import (
	"context"
	"fmt"
	"follows_service/database"
	protos "follows_service/protos/follows"

	"github.com/hashicorp/go-hclog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Follows struct {
	log hclog.Logger
	protos.UnimplementedFollowsServer
}

func NewFollows(l hclog.Logger, u protos.UnimplementedFollowsServer) *Follows {
	return &Follows{l, u}
}

func (f *Follows) AddFollowToDB(ctx context.Context, fr *protos.Follow) (*protos.EmptyResponse, error) {
	f.log.Info("Adding Follow to Database")
	followsCollection := database.Client.Database("xws").Collection("follows")
	filter := bson.D{{"follower", fr.Follower}, {"followee", fr.Followee}}
	var result protos.Follow
	err := followsCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil && err != mongo.ErrNoDocuments {
		return &protos.EmptyResponse{}, err
	}
	if result.FollowID != "" {
		fmt.Println("[ERROR] already following")
		return &protos.EmptyResponse{}, fmt.Errorf("aleready following")
	}

	doc, err := bson.Marshal(fr)
	if err != nil {
		fmt.Println("[ERROR] marshaling to bson.d")
		return &protos.EmptyResponse{}, err
	}
	_, err = followsCollection.InsertOne(ctx, doc)
	if err != nil {
		fmt.Println("[ERROR] inserting into database")
		return &protos.EmptyResponse{}, err
	}
	return &protos.EmptyResponse{}, nil
}

func (f *Follows) AddFollowRequestToDB(ctx context.Context, fr *protos.FollowRequest) (*protos.EmptyResponse, error) {
	followRequestsCollection := database.Client.Database("xws").Collection("followRequests")

	filter := bson.D{{"requester", fr.Requester}, {"requestee", fr.Requestee}}
	var result protos.FollowRequest
	err := followRequestsCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil && err != mongo.ErrNoDocuments {
		return &protos.EmptyResponse{}, err
	}
	if result.RequestID != "" {
		fmt.Println("[ERROR] already requested to follow")
		return &protos.EmptyResponse{}, err
	}

	doc, err := bson.Marshal(fr)
	if err != nil {
		fmt.Println("[ERROR] marshaling to bson.d")
		return &protos.EmptyResponse{}, err
	}
	_, err = followRequestsCollection.InsertOne(ctx, doc)
	if err != nil {
		fmt.Println("[ERROR] inserting into database")
		return &protos.EmptyResponse{}, err
	}
	return &protos.EmptyResponse{}, nil
}

func (f *Follows) GetFollowRequests(ctx context.Context, fr *protos.GetFollowRRequest) (*protos.FollowRequests, error) {
	followRequestsCollection := database.Client.Database("xws").Collection("followRequests")
	var result protos.FollowRequests
	filter := bson.M{"requestee": fr.Username}
	cursor, err := followRequestsCollection.Find(ctx, filter)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	if cursor.RemainingBatchLength() == 0 {
		return nil, fmt.Errorf("no requests")
	}
	err = cursor.All(ctx, &result.Results)
	if err != nil {
		fmt.Println("[ERROR] decoding result")
		return nil, err
	}
	fmt.Println("[DEBUG] finished getting requests successfully")
	return &result, nil
}

func (f *Follows) DeleteFollowRequest(ctx context.Context, fr *protos.DeleteFollowRRequest) (*protos.EmptyResponse, error) {
	followRequestsCollection := database.Client.Database("xws").Collection("followRequests")
	frID, err := primitive.ObjectIDFromHex(fr.FollowRequestID)
	if err != nil {
		f.log.Error("converting to ObjectID from string")
		return nil, err
	}
	_, err = followRequestsCollection.DeleteOne(ctx, bson.M{"_id": frID})
	if err != nil {
		return &protos.EmptyResponse{}, err
	}
	return &protos.EmptyResponse{}, nil
}
func (f *Follows) GetFollow(ctx context.Context, fr *protos.Follow) (*protos.Follow, error) {
	followsCollection := database.Client.Database("xws").Collection("follows")
	var result protos.Follow
	filter := bson.D{{"follower", fr.Follower}, {"followee", fr.Followee}}

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
func (f *Follows) GetFollowRequest(ctx context.Context, fr *protos.FollowRequest) (*protos.FollowRequest, error) {
	f.log.Info("get follow request service")
	followsCollection := database.Client.Database("xws").Collection("followRequests")
	var result protos.FollowRequest
	filter := bson.D{{"requester", fr.Requester}, {"requestee", fr.Requestee}}

	sRes := followsCollection.FindOne(ctx, filter)
	if sRes.Err() != nil {
		fmt.Println(sRes.Err())
		fmt.Println("[DEBUG] no follow request")
		return nil, fmt.Errorf("follow request not found")
	}
	err := sRes.Decode(&result)
	if err != nil {
		fmt.Println(err)
		fmt.Println("[ERROR] could not decode result")
		return nil, err
	}
	return &result, nil
}

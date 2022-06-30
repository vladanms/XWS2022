package server

import (
	"context"
	"fmt"
	"users_service/database"
	protos "users_service/protos/user"

	"github.com/hashicorp/go-hclog"
	"go.mongodb.org/mongo-driver/bson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Users struct {
	log hclog.Logger
	protos.UnimplementedUserServer
}

func NewUsers(l hclog.Logger, u protos.UnimplementedUserServer) *Users {
	return &Users{l, u}
}

func (u *Users) GetUsers(ctx context.Context, ur *protos.UsersRequest) (*protos.UsersResponse, error) {
	u.log.Info("Get users handle")
	usersCollection := database.Client.Database("xws").Collection("users")
	cursor, err := usersCollection.Find(ctx, bson.D{})
	if err != nil {
		u.log.Error("retrieving users from database")
	}
	var allUsers protos.UsersResponse
	err = cursor.All(ctx, &allUsers.Results)
	if err != nil {
		u.log.Error("decoding users from bson to struct")
	}

	return &allUsers, err
}

func (u *Users) GetPublicUsers(ctx context.Context, ur *protos.UsersRequest) (*protos.UsersResponse, error) {
	u.log.Info("Get public users handle")
	usersCollection := database.Client.Database("xws").Collection("users")
	cursor, err := usersCollection.Find(ctx, bson.D{})
	if err != nil {
		u.log.Error("[ERROR] retrieving users from database")
	}
	var allUsers protos.UsersResponse
	err = cursor.All(ctx, &allUsers.Results)
	if err != nil {
		u.log.Error("[ERROR] decoding users form bson to struct")
	}
	var publicUsers protos.UsersResponse
	for _, user := range allUsers.Results {
		if user.Public {
			publicUsers.Results = append(publicUsers.Results, user)
		}
	}
	return &publicUsers, err
}

func (u *Users) GetUserByUsername(ctx context.Context, ur *protos.UserByUsernameRequest) (*protos.UserResponse, error) {
	u.log.Info("Get user by username handle")
	usersCollection := database.Client.Database("xws").Collection("users")

	var result protos.UserResponse
	fmt.Println(ur.Username)
	filter := bson.D{{"username", ur.Username}}

	sRes := usersCollection.FindOne(ctx, filter)
	if sRes.Err() != nil {
		fmt.Println(sRes.Err())
		u.log.Error("no user with that username")
		return &protos.UserResponse{}, sRes.Err()
	}
	err := sRes.Decode(&result)
	if err != nil {
		fmt.Println(err)
		u.log.Error("could not decode result")
		return nil, err
	}
	user := &result
	return user, err
}
func (u *Users) CreateUser(ctx context.Context, ur *protos.UserResponse) (*protos.UsersResponse, error) {
	userCollection := database.Client.Database("xws").Collection("users")
	if ur.DateOfBirth == nil {
		ur.DateOfBirth = &timestamppb.Timestamp{}
		fmt.Println(ur.DateOfBirth)
	}
	doc, err := bson.Marshal(ur)
	if err != nil {
		fmt.Println("[ERROR] marshaling to bson.d")
		return nil, err
	}
	result, err := userCollection.InsertOne(ctx, doc)
	if err != nil {
		fmt.Println("[ERROR] inserting into database")
		return nil, err
	}
	fmt.Println(result.InsertedID)
	return &protos.UsersResponse{}, nil
}

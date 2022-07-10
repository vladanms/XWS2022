package server

import (
	"context"
	"fmt"
	"job_offers_service/database"
	protos "job_offers_service/protos/joboffers"

	"github.com/hashicorp/go-hclog"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type JobOffers struct {
	log hclog.Logger
	protos.UnimplementedJobOffersServer
}

func NewJobOffers(l hclog.Logger, u protos.UnimplementedJobOffersServer) *JobOffers {
	return &JobOffers{l, u}
}

func (jo *JobOffers) GetJobOffers(ctx context.Context, jor *protos.JobOffersRequest) (*protos.JobOffersResponse, error) {
	jo.log.Info("Get job offers handle")
	jobOffersCollection := database.Client.Database("xws").Collection("job_offers")
	cursor, err := jobOffersCollection.Find(ctx, bson.D{})
	if err != nil {
		jo.log.Error("retrieving job offers from database")
	}
	var allJobOffers protos.JobOffersResponse
	err = cursor.All(ctx, &allJobOffers.Results)
	if err != nil {
		jo.log.Error("decoding job offers from bson to struct")
	}

	return &allJobOffers, err
}

func (jo *JobOffers) CreateJobOffer(ctx context.Context, jor *protos.JobOffer) (*protos.JobOffersResponse, error) {
	jobOfferCollection := database.Client.Database("xws").Collection("job_offers")
	doc, err := bson.Marshal(jor)
	if err != nil {
		fmt.Println("[ERROR] marshaling to bson.d")
		return nil, err
	}
	result, err := jobOfferCollection.InsertOne(ctx, doc)
	if err != nil {
		fmt.Println("[ERROR] inserting into database")
		return nil, err
	}
	fmt.Println(result.InsertedID)
	return &protos.JobOffersResponse{}, nil
}

func (jo *JobOffers) RemoveJobOffer(ctx context.Context, jor *protos.JobOffer) (*protos.JobOffersRequest, error) {
	jo.log.Info("removing job offer")
	jobOfferCollection := database.Client.Database("xws").Collection("job_offers")
	objectId, err := primitive.ObjectIDFromHex(jor.ID)
	if err != nil {
		fmt.Println("[ERROR] can't convert string to ObjectID", err)
		return nil, err
	}
	fmt.Println(objectId)
	jor.ID = ""
	filter := bson.M{"_id": objectId}
	_, err = jobOfferCollection.DeleteOne(ctx, filter)
	if err != nil {
		fmt.Println(err)
		fmt.Println("[ERROR] removing job offer")
	}
	jo.log.Info("finished updating user")
	return &protos.JobOffersRequest{}, nil
}

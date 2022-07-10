package handlers

import (
	jobOfferProtos "job_offers_service/protos/joboffers"
	"log"
)

// Users handler for getting and updating users
type JobOffers struct {
	l   *log.Logger
	joc jobOfferProtos.JobOffersClient
}

// NewUsers returns a new products handler with the given logger
func NewJobOffers(l *log.Logger, joc jobOfferProtos.JobOffersClient) *JobOffers {
	return &JobOffers{l, joc}
}

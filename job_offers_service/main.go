package main

import (
	protos "job_offers_service/protos/joboffers"
	"job_offers_service/server"
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

func main() {
	log := hclog.Default()
	gs := grpc.NewServer()
	uus := protos.UnimplementedJobOffersServer{}
	us := server.NewJobOffers(log, uus)
	protos.RegisterJobOffersServer(gs, us)

	l, err := net.Listen("tcp", ":9095")
	if err != nil {
		log.Error("Unable to listen")
		os.Exit(1)
	}
	gs.Serve(l)
}

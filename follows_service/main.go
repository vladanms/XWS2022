package main

import (
	protos "follows_service/protos/follows"
	"follows_service/server"
	"net"
	"os"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

func main() {
	log := hclog.Default()
	gs := grpc.NewServer()
	uus := protos.UnimplementedFollowsServer{}
	us := server.NewFollows(log, uus)
	protos.RegisterFollowsServer(gs, us)

	l, err := net.Listen("tcp", ":9094")
	if err != nil {
		log.Error("Unable to listen")
		os.Exit(1)
	}
	gs.Serve(l)
}

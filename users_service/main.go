package main

import (
	"net"
	"os"
	protos "users_service/protos/user"
	"users_service/server"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

func main() {
	log := hclog.Default()
	gs := grpc.NewServer()
	uus := protos.UnimplementedUserServer{}
	us := server.NewUsers(log, uus)
	protos.RegisterUserServer(gs, us)

	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Unable to listen")
		os.Exit(1)
	}
	gs.Serve(l)
}

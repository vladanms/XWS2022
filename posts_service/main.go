package main

import (
	"net"
	"os"
	protos "posts_service/protos/posts"
	"posts_service/server"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
)

func main() {
	log := hclog.Default()
	gs := grpc.NewServer()
	uus := protos.UnimplementedPostsServer{}
	us := server.NewPosts(log, uus)
	protos.RegisterPostsServer(gs, us)

	l, err := net.Listen("tcp", ":9093")
	if err != nil {
		log.Error("Unable to listen")
		os.Exit(1)
	}
	gs.Serve(l)
}

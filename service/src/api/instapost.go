package api

import (
	"context"
	"log"

	pb "helloworld/gen/proto"

	"google.golang.org/grpc"
)

// server is used to implement helloworld.GreeterServer.
type instaServer struct {
	pb.UnimplementedInstaPostServer
}

// SayHello implements helloworld.GreeterServer
func (s *instaServer) GetPosts(ctx context.Context, in *pb.InstaPostRequest) (*pb.InstaPostResponse, error) {
	log.Printf("Received: %v", in.GetId())
	return &pb.InstaPostResponse{Stats: in.GetId() + " = 10 likes; 5 comments"}, nil
}

func RegisterInstaPostServer(s grpc.ServiceRegistrar) {
	pb.RegisterInstaPostServer(s, &instaServer{})
}

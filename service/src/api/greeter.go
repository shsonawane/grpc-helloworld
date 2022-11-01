package api

import (
	"context"
	"log"

	pb "helloworld/gen/proto"

	"google.golang.org/grpc"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Greeter 9 Server says 'Hello " + in.GetName() + "!'"}, nil
}

func RegisterGreeterServer(s grpc.ServiceRegistrar) {
	pb.RegisterGreeterServer(s, &server{})
}

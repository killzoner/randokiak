package main

import (
	"fmt"
	"log"
	"net"

	"github.com/Pallinder/go-randomdata"
	"golang.org/x/net/context"

	pb "rdk.io/protocols"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// server is used to implement greeter.GreeterServer
type server struct {
}

func (s server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println(randomdata.GenerateProfile(randomdata.RandomGender))
	log.Printf("Answering to msg '%v'", in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	// Register reflection service on grpc server
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

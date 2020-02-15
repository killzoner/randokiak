package main

import (
	"fmt"
	"log"
	"net"

	randomdata "github.com/Pallinder/go-randomdata"
	"golang.org/x/net/context"

	pb "rdk.io/protocols"
	"rdk.io/utils"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	defaultPort = "50051"
)

// server is used to implement greeter.GreeterServer
type server struct {
}

func (s server) GetProfile(ctx context.Context, in *pb.ProfileRequest) (*pb.ProfileReply, error) {
	rand := randomdata.GenerateProfile(randomdata.RandomGender)
	fmt.Println(rand)
	p := pb.Profile{
		Gender: rand.Gender,
		Name: &pb.ProfileName{
			First: rand.Name.First,
			Last:  rand.Name.Last,
			Title: rand.Name.Title,
		},
		Location: &pb.ProfileLocation{
			Street:   rand.Location.Street,
			City:     rand.Location.City,
			State:    rand.Location.State,
			Postcode: int64(rand.Location.Postcode),
		},
		Email: rand.Email,
		Login: &pb.ProfileLogin{
			Username: rand.Login.Username,
			Password: rand.Login.Password,
			Salt:     rand.Login.Salt,
			Md5:      rand.Login.Md5,
			Sha1:     rand.Login.Sha1,
			Sha256:   rand.Login.Sha256,
		},
		Dob:        rand.Dob,
		Registered: rand.Registered,
		Phone:      rand.Phone,
		Cell:       rand.Cell,
		Id: &pb.ProfileId{
			Name: rand.ID.Name,
		},
		Picture: &pb.ProfilePicture{
			Large:     rand.Picture.Large,
			Medium:    rand.Picture.Medium,
			Thumbnail: rand.Picture.Thumbnail,
		},
		Nat: rand.Nat,
	}

	return &pb.ProfileReply{Profile: &p}, nil
}

func main() {
	port := fmt.Sprintf(":%s", utils.GetEnv("PORT", defaultPort))
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterRandomizerServer(s, &server{})
	// Register reflection service on grpc server
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

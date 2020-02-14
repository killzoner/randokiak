package controller

import (
	"net/http"

	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	pb "rdk.io/protocols"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

// Status godoc
// @Summary Ping
// @Description Ping
// @Success 200
// @Router /dumb [get]
func (controller *Controller) Status(ctx *gin.Context) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print it's response
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %v", r.Message)

	ctx.JSON(http.StatusOK, gin.H{
		"message": r.Message,
	})
}

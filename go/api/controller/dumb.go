package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/gin-gonic/gin"

	"google.golang.org/grpc"
	"rdk.io/protocols"
	pb "rdk.io/protocols"
	"rdk.io/utils"
)

const (
	defaultGenHost         = "localhost"
	defaultGenPort         = "50051"
	defaultPulsarHost      = "localhost"
	defaultPulsarPort      = "6650"
	defaultTopic           = "randokiak-topic"
	defaultTnbFetchProfile = "50"
)

// Status godoc
// @Summary Ping
// @Description Ping
// @Success 200
// @Router /dumb [get]
func (controller *Controller) Status(ctx *gin.Context) {
	genHost := utils.GetEnv("GEN_HOST", defaultGenHost)
	genPort := utils.GetEnv("GEN_PORT", defaultGenPort)
	genAddress := fmt.Sprintf("%s:%s", genHost, genPort)

	conn, err := grpc.Dial(genAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	grpcClient := pb.NewRandomizerClient(conn)

	pulsarHost := utils.GetEnv("PULSAR_HOST", defaultPulsarHost)
	pulsarPort := utils.GetEnv("PULSAR_PORT", defaultPulsarPort)
	pulsarAdress := fmt.Sprintf("pulsar://%s:%s", pulsarHost, pulsarPort)

	pulsarClient, err := pulsar.NewClient(pulsar.ClientOptions{
		URL: pulsarAdress,
	})
	defer pulsarClient.Close()

	topic := utils.GetEnv("TOPIC", defaultTopic)
	pulsarProducer, err := pulsarClient.CreateProducer(pulsar.ProducerOptions{
		Topic: topic,
	})
	defer pulsarProducer.Close()

	var wg sync.WaitGroup

	/***
	 * Warn, a send or receive to a nil channel
	 * hangs forever, be sure to initialize size!!!
	 */
	nbFetchProfile, _ := strconv.ParseInt(utils.GetEnv("NB_FETCH_PROFILE", defaultTnbFetchProfile), 10, 64)
	ch := make(chan *pb.ProfileReply, nbFetchProfile)
	profiles := make([]pb.ProfileReply, 0, nbFetchProfile)

	for i := 0; i < int(nbFetchProfile); i++ {
		wg.Add(1)
		go getProfileAndPushToPulsar(ch, &grpcClient, &pulsarProducer, &wg)
	}

	wg.Wait()

	/*
	 * close channel as we know all tasks are done
	 * if we don't close, range hangs forever too
	 */
	close(ch)

	for ret := range ch {
		profiles = append(profiles, *ret)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"profiles": profiles,
	})
}

func getProfileAndPushToPulsar(
	chanProfiles chan *pb.ProfileReply,
	grpcClient *pb.RandomizerClient,
	pulsarProducer *pulsar.Producer,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	r, err := (*grpcClient).GetProfile(context.Background(), &pb.ProfileRequest{})
	if err != nil {
		log.Fatalf("could not get profile: %v", err)
	}
	log.Printf("Got 1 profil")

	pushToPulsar(pulsarProducer, r)

	//send response to channel
	chanProfiles <- r
}

func pushToPulsar(pulsarProducer *pulsar.Producer, message *protocols.ProfileReply) {

	json, err := json.Marshal(message)

	_, err = (*pulsarProducer).Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte(json),
	})

	if err != nil {
		fmt.Println("Failed to publish message", err)
	}
	fmt.Println("Published message")
}

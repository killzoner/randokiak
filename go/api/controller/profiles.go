package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/gin-gonic/gin"

	"rdk.io/protocols"
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

// AskMoreProfiles godoc
// @Summary Profiles
// @Description Get more profiles
// @Success 200
// @Router /profiles [get]
func (cc *Controller) AskMoreProfiles(ctx *gin.Context, connectors *IConnectors) {

	conn := (*connectors).getGrpcConnection()
	defer conn.Close()
	grpcClient := (*connectors).getGrpcClient(&conn)

	pulsarClient := (*connectors).getPulsarClient()
	defer pulsarClient.Close()

	pulsarProducer := (*connectors).getPulsarProducer(&pulsarClient)
	defer pulsarProducer.Close()

	var wg sync.WaitGroup

	/***
	 * Warn, a send or receive to a nil channel
	 * hangs forever, be sure to initialize size!!!
	 */
	nbFetchProfile, _ := strconv.ParseInt(utils.GetEnv("NB_FETCH_PROFILE", defaultTnbFetchProfile), 10, 64)
	ch := make(chan *protocols.ProfileReply, nbFetchProfile)
	profiles := make([]protocols.ProfileReply, 0, nbFetchProfile)

	for i := 0; i < int(nbFetchProfile); i++ {
		wg.Add(1)
		go cc.getProfileAndPushToPulsar(ch, &grpcClient, &pulsarProducer, &wg)
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

func (cc *Controller) getProfileAndPushToPulsar(
	chanProfiles chan *protocols.ProfileReply,
	grpcClient *protocols.RandomizerClient,
	pulsarProducer *pulsar.Producer,
	wg *sync.WaitGroup,
) {
	defer wg.Done()

	r, err := (*grpcClient).GetProfile(context.Background(), &protocols.ProfileRequest{})
	if err != nil {
		log.Fatalf("could not get profile: %v", err)
	}
	log.Printf("Got 1 profil")

	cc.pushToPulsar(pulsarProducer, r)

	//send response to channel
	chanProfiles <- r
}

func (cc *Controller) pushToPulsar(
	pulsarProducer *pulsar.Producer,
	message *protocols.ProfileReply,
) pulsar.MessageID {

	json, _ := json.Marshal(message)

	MessageID, err := (*pulsarProducer).Send(context.Background(), &pulsar.ProducerMessage{
		Payload: []byte(json),
	})

	if err != nil {
		log.Println("Failed to publish message", err)
	}
	log.Println("Published message")

	return MessageID
}

func (cc *Controller) unexportedFunction() int {
	return 8
}

// ExportedFunction godoc
func (cc *Controller) ExportedFunction() int {
	return 8
}

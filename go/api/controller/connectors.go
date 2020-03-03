package controller

import (
	"fmt"
	"log"

	"github.com/apache/pulsar-client-go/pulsar"
	"google.golang.org/grpc"
	"rdk.io/protocols"
	"rdk.io/utils"
)

// Connectors godoc
type Connectors struct {
	conn           *grpc.ClientConn
	grpcClient     *protocols.RandomizerClient
	pulsarClient   *pulsar.Client
	pulsarProducer *pulsar.Producer
}

// IConnectors godoc
type IConnectors interface {
	getGrpcConnection() *grpc.ClientConn
	getPulsarClient() *pulsar.Client
	getPulsarProducer() *pulsar.Producer
}

func (cc *Connectors) getGrpcConnection() *grpc.ClientConn {
	genHost := utils.GetEnv("GEN_HOST", defaultGenHost)
	genPort := utils.GetEnv("GEN_PORT", defaultGenPort)
	genAddress := fmt.Sprintf("%s:%s", genHost, genPort)

	conn, err := grpc.Dial(genAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return conn
}

func (cc *Connectors) getPulsarClient() pulsar.Client {
	pulsarHost := utils.GetEnv("PULSAR_HOST", defaultPulsarHost)
	pulsarPort := utils.GetEnv("PULSAR_PORT", defaultPulsarPort)
	pulsarAdress := fmt.Sprintf("pulsar://%s:%s", pulsarHost, pulsarPort)

	pulsarClient, _ := pulsar.NewClient(pulsar.ClientOptions{
		URL: pulsarAdress,
	})
	return pulsarClient
}

func (cc *Connectors) getPulsarProducer(pulsarClient *pulsar.Client) pulsar.Producer {
	topic := utils.GetEnv("TOPIC", defaultTopic)
	pulsarProducer, _ := (*pulsarClient).CreateProducer(pulsar.ProducerOptions{
		Topic: topic,
	})
	return pulsarProducer
}

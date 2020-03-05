package controller

import (
	"fmt"
	"log"

	"github.com/apache/pulsar-client-go/pulsar"
	"google.golang.org/grpc"
	"rdk.io/protocols"
	"rdk.io/utils"
)

// GrpcClientConnectionWrp godoc
type GrpcClientConnectionWrp struct {
	conn *grpc.ClientConn
}

// IGrpcClientConnectionWrp godoc
type IGrpcClientConnectionWrp interface {
	Close() error
	GetConnection() *grpc.ClientConn
}

// Close godoc
func (cc *GrpcClientConnectionWrp) Close() error {
	return cc.conn.Close()
}

// GetConnection godoc
func (cc *GrpcClientConnectionWrp) GetConnection() *grpc.ClientConn {
	return cc.conn
}

// Connectors godoc
type Connectors struct {
	conn           *IGrpcClientConnectionWrp
	grpcClient     *protocols.RandomizerClient
	pulsarClient   *pulsar.Client
	pulsarProducer *pulsar.Producer
}

// IConnectors godoc
type IConnectors interface {
	getGrpcConnection() IGrpcClientConnectionWrp
	getGrpcClient(conn *IGrpcClientConnectionWrp) protocols.RandomizerClient
	getPulsarClient() pulsar.Client
	getPulsarProducer(pulsarClient *pulsar.Client) pulsar.Producer
}

func (cc *Connectors) getGrpcConnection() IGrpcClientConnectionWrp {
	genHost := utils.GetEnv("GEN_HOST", defaultGenHost)
	genPort := utils.GetEnv("GEN_PORT", defaultGenPort)
	genAddress := fmt.Sprintf("%s:%s", genHost, genPort)

	conn, err := grpc.Dial(genAddress, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	return &GrpcClientConnectionWrp{conn: conn}
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

func (cc *Connectors) getGrpcClient(conn *IGrpcClientConnectionWrp) protocols.RandomizerClient {
	return protocols.NewRandomizerClient((*conn).GetConnection())
}

// NewConnectors godoc
func NewConnectors() IConnectors {
	return &Connectors{}
}

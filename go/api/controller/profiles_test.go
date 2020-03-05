package controller

/**
 * Test private methods
 * with same package name
 * Run with go test -v ./...
 *
 * mockgen annotations allow generating mock files with `go generate` command
 */

//go:generate mockgen -source=controller.go -destination=mock_controller.go -package controller
//goo:generate mockgen -source=../../protocols/random.pb.go -package controller -destination=mock_randomizerclient.go

//for save reflection mode
//goo:generate mockgen -destination=mock/mock_randomizerclient2.go -package mock rdk.io/protocols RandomizerClient

import (
	"context"
	"fmt"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/apache/pulsar-client-go/pulsar"
	"github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	"google.golang.org/grpc"
	"rdk.io/protocols"
)

func TestUnexportedFunction_(t *testing.T) {
	t.Logf("Here IController.unexportedFunction is visible")
	c := Controller{}
	want := 8
	got := c.unexportedFunction()
	if got != want {
		t.Fatalf(`Fail! Wanted '%v', got '%v'`, want, got)
	}
}

// implements pulsar.Producer
type MockPulsarProducer struct{}

var sendmock func() (pulsar.MessageID, error)

/**
 * Manual mock
 * Generated with vs code => Go: Generate interface stubs
 */
func (cc MockPulsarProducer) Topic() string {
	panic("not implemented") // TODO: Implement
}
func (cc MockPulsarProducer) Name() string {
	panic("not implemented") // TODO: Implement
}
func (cc MockPulsarProducer) Send(_ context.Context, _ *pulsar.ProducerMessage) (pulsar.MessageID, error) {
	return sendmock()
}
func (cc MockPulsarProducer) SendAsync(_ context.Context, _ *pulsar.ProducerMessage, _ func(pulsar.MessageID, *pulsar.ProducerMessage, error)) {
	panic("not implemented") // TODO: Implement
}
func (cc MockPulsarProducer) LastSequenceID() int64 {
	panic("not implemented") // TODO: Implement
}
func (cc MockPulsarProducer) Flush() error {
	panic("not implemented") // TODO: Implement
}
func (cc MockPulsarProducer) Close() {
}

//---------------------

// implements pulsar.Client
type MockPulsarClient struct{}

func (cc MockPulsarClient) CreateProducer(_ pulsar.ProducerOptions) (pulsar.Producer, error) {
	panic("not implemented") // TODO: Implement
}
func (cc MockPulsarClient) Subscribe(_ pulsar.ConsumerOptions) (pulsar.Consumer, error) {
	panic("not implemented") // TODO: Implement
}
func (cc MockPulsarClient) CreateReader(_ pulsar.ReaderOptions) (pulsar.Reader, error) {
	panic("not implemented") // TODO: Implement
}
func (cc MockPulsarClient) TopicPartitions(topic string) ([]string, error) {
	panic("not implemented") // TODO: Implement
}
func (cc MockPulsarClient) Close() {
}

//---------------------

type MockRandomizerClient struct {
	calls int
}

var getProfile func() (*protocols.ProfileReply, error)

func (cc *MockRandomizerClient) GetProfile(ctx context.Context, in *protocols.ProfileRequest, opts ...grpc.CallOption) (*protocols.ProfileReply, error) {
	cc.calls++
	// println(cc.calls)
	return getProfile()
}

//---------------------

type MockGrpcClientConnectionWrp struct {
}

// IGrpcClientConnectionWrp godoc
func (cc *MockGrpcClientConnectionWrp) Close() error {
	return nil
}

// IGrpcClientConnectionWrp godoc
func (cc *MockGrpcClientConnectionWrp) GetConnection() *grpc.ClientConn {
	return nil
}

//---------------------

func getProducer() pulsar.Producer {
	return &MockPulsarProducer{}
}

func mockAsIConnector(m *MockIConnectors) IConnectors {
	return m
}

func TestPushToPulsar(t *testing.T) {
	c := Controller{}
	reply := protocols.ProfileReply{}

	// "cast" MockPulsarProducer as pulsar.Producer interface
	producer := getProducer()

	// set return value for mock
	sendmock = func() (pulsar.MessageID, error) {
		return pulsar.LatestMessageID(), nil
	}
	id, _ := sendmock()

	want := fmt.Sprintf("%s", id)
	got := fmt.Sprintf("%s", c.pushToPulsar(&producer, &reply))
	// we compare string representations as it's different object references
	if got != want {
		t.Fatalf(`Fail! Wanted '%v', got '%v'`, want, got)
	}

}

func TestAskMoreProfiles(t *testing.T) {
	c := Controller{}
	gin.SetMode(gin.TestMode)
	resp := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(resp)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockIConnectors(ctrl)
	impl := mockAsIConnector(m)

	conn := &MockGrpcClientConnectionWrp{}
	pulsarClient := &MockPulsarClient{}
	randomizerClient := &MockRandomizerClient{}
	producer := &MockPulsarProducer{}

	// set return value for mock
	sendmock = func() (pulsar.MessageID, error) {
		return pulsar.LatestMessageID(), nil
	}

	// set return value for mock
	getProfile = func() (*protocols.ProfileReply, error) {
		return &protocols.ProfileReply{Profile: &protocols.Profile{}}, nil
	}

	m.
		EXPECT().
		getGrpcConnection().
		Return(conn).
		AnyTimes()

	m.
		EXPECT().
		getGrpcClient(gomock.Any()).
		Return(randomizerClient).
		AnyTimes()

	m.
		EXPECT().
		getPulsarClient().
		Return(pulsarClient).
		AnyTimes()

	m.
		EXPECT().
		getPulsarProducer(gomock.Any()).
		Return(producer).
		AnyTimes()

	// run as background
	go c.AskMoreProfiles(ctx, &impl)

	want := 50
	var got int
	deadline := time.Now().Add(5 * time.Second)
	for {
		got = randomizerClient.calls
		if got == want || time.Now().After(deadline) {
			break
		}
	}
	if got != want {
		t.Fatalf(`Fail! Wanted '%v', got '%v'`, want, got)
	}
}

package controller

/**
 * Test private methods
 * with same package name
 * Run with go test -v ./...
 */

import (
	"context"
	"fmt"
	"testing"

	"github.com/apache/pulsar-client-go/pulsar"
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
	panic("not implemented") // TODO: Implement
}

func getProducer() pulsar.Producer {
	return &MockPulsarProducer{}
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

	/**
	 * TODO: make tests more accurate,
	 * but it's a quick demo on how we mock
	 * a do internal (private testing)
	 */

	want := fmt.Sprintf("%s", id)
	got := fmt.Sprintf("%s", c.pushToPulsar(&producer, &reply))
	// we compare string representations as it's different object references
	if got != want {
		t.Fatalf(`Fail! Wanted '%v', got '%v'`, want, got)
	}

}

package server

import (
	"context"
	"fmt"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"log"
	proto "rk-grpc/api/gen/v1"
)

func StartReceiver() {
	// The default client is HTTP.
	c, err := cloudevents.NewClientHTTP()
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}

	log.Fatal(c.StartReceiver(context.Background(), receive))
}

func receive(event cloudevents.Event) {
	// do something with event.
	fmt.Printf("%s", event)
}

// GreeterServer Implementation of GreeterServer.
type GreeterServer struct{}

func NewGreeterServer() *GreeterServer {
	return &GreeterServer{}
}

// Greeter Handle Greeter method.
func (server *GreeterServer) Greeter(context.Context, *proto.GreeterRequest) (*proto.GreeterResponse, error) {
	return &proto.GreeterResponse{}, nil
}

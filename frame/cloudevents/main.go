package main

import (
	"context"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"log"
)

func main() {
	c, err := cloudevents.NewClientHTTP()
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}

	// Create an Event.
	event := cloudevents.NewEvent()
	event.SetID("A234-1234-1234")
	event.SetSource("cloud-hcce/api/v1/deployments")
	event.SetType("cloud-hcce.deployments.create")
	event.SetData(cloudevents.ApplicationJSON, map[string]string{"hello": "world"})

	// optional
	event.SetExtension("idc", "wowjoy")
	event.SetExtension("env", "dev")
	event.SetExtension("namespace", "dev")
	event.SetExtension("service", "cloud-hcce")
	// Set a target.
	ctx := cloudevents.ContextWithTarget(context.Background(), "http://localhost:8080/")
	// Send that Event.
	if result := c.Send(ctx, event); cloudevents.IsUndelivered(result) {
		log.Fatalf("failed to send, %v", result)
	} else {
		log.Printf("sent: %v", event)
		log.Printf("result: %v", result)
	}
}

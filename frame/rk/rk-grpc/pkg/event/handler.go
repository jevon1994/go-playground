package main

import (
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type Handler interface {
	Handle(event *cloudevents.Event) error
}

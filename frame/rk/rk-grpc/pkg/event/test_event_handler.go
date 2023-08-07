package event

import (
	"fmt"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type TestEventCallBackHandler struct {
}

func (th *TestEventCallBackHandler) Handle(event *cloudevents.Event) error {
	fmt.Printf("this is event callback,%v \n", event)
	return nil
}

type TestEventCallBackResultHandler struct{}

func (terh *TestEventCallBackResultHandler) Handle(event *cloudevents.Event) error {
	fmt.Printf("this is event callback result,%v \n", event)
	return nil
}

func init() {
	ReqBus.AddHandler(GetCallBackType("cloud-hcce", "deployments.create"), &TestEventCallBackHandler{})
	ReqBus.AddHandler(GetCallBackResultType("cloud-hcce", "deployments.create"), &TestEventCallBackResultHandler{})
}

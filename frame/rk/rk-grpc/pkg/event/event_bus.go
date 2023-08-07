package event

import (
	"fmt"
	"github.com/asaskevich/EventBus"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

var HandlerMap = make(map[string]Handler)
var ReqBus *RequestEventBus

type RequestEventBus struct {
	Bus EventBus.Bus
}

func NewRequestEventBus() *RequestEventBus {
	return &RequestEventBus{Bus: EventBus.New()}
}

func InitRequestEventBus() *RequestEventBus {
	eventBus := NewRequestEventBus()

	for topic, h := range HandlerMap {
		eventBus.Bus.Subscribe(topic, h.Handle)
	}
	return eventBus
}

func (eb *RequestEventBus) AddHandler(topic string, handler Handler) {
	HandlerMap[topic] = handler
}

func (re *RequestEventBus) Dispatch(event *cloudevents.Event) error {
	ReqBus.Bus.Publish(event.Type(), event)
	return nil
}

func GetCallBackType(service, event string) string {
	return fmt.Sprintf("req:%v.%v", service, event)
}

func GetCallBackResultType(service, event string) string {
	return fmt.Sprintf("r:%v.%v", service, event)
}

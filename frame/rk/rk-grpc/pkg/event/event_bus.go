package main

import (
	"errors"
	"github.com/asaskevich/EventBus"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"reflect"
)

type InProcBus struct {
	handlers map[string]Handler
}

func NewInProcBus() *InProcBus {
	var handlerMap = make(map[string]Handler)
	return &InProcBus{handlers: handlerMap}
}

func (b *InProcBus) AddHandler(handler Handler) {
	handlerType := reflect.TypeOf(handler)
	queryTypeName := handlerType.In(0).Elem().Name() // 获取函数第一个参数的名称，在上面例子里面就是GetAlertByIdQuery
	b.handlers[queryTypeName] = handler
}

// ErrHandlerNotFound defines an error if a handler is not found.
var ErrHandlerNotFound = errors.New("handler not found")

func (re *RequestEventBus) Dispatch(event *cloudevents.Event) error {
	eBus.Bus.Publish(event.Type(), event)
	//var msgName = reflect.TypeOf(event).Elem().Name()
	//
	//handler := handlerMap[event.Type()]
	//if handler == nil {
	//	return ErrHandlerNotFound
	//}
	//var params = []reflect.Value{}
	//params = append(params, reflect.ValueOf(event))
	//
	//ret := reflect.ValueOf(handler).Call(params) // 通过反射机制调用函数
	//err := ret[0].Interface()
	//if err == nil {
	//	return nil
	//}
	return nil
}

var HandlerMap = make(map[string]Handler)
var eBus *RequestEventBus

type RequestEventBus struct {
	Bus EventBus.Bus
}

func NewRequestEventBus() *RequestEventBus {
	return &RequestEventBus{Bus: EventBus.New()}
}

func (eb *RequestEventBus) AddHandler(topic string, handler Handler) {
	HandlerMap[topic] = handler
}

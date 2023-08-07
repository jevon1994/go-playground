package main

//
//// 事件结构体
//import "fmt"
//
//type Event struct {
//	Name string
//	Data interface{}
//}
//
//// 事件处理器接口
//type EventHandler interface {
//	Handle(event Event)
//}
//
//// 消息队列接口
//type MessageQueue interface {
//	Consume(queueName string, handler EventHandler)
//	Publish(queueName string, event Event)
//}
//
//// 事件处理器实现
//type MyEventHandler struct{}
//
//func (h *MyEventHandler) Handle(event Event) {
//	// 实现对事件的响应处理逻辑
//	fmt.Printf("handling event %v\n", event.Name)
//}
//
//// 消息队列实现
//type MyMessageQueue struct {
//	handlers map[string][]EventHandler
//}
//
//func (mq *MyMessageQueue) Consume(queueName string, handler EventHandler) {
//	if _, ok := mq.handlers[queueName]; !ok {
//		mq.handlers[queueName] = []EventHandler{}
//	}
//	mq.handlers[queueName] = append(mq.handlers[queueName], handler)
//}
//
//func (mq *MyMessageQueue) Publish(queueName string, event Event) {
//	handlers := mq.handlers[queueName]
//	for _, handler := range handlers {
//		handler.Handle(event)
//	}
//}
//
//// 示例代码
//func main() {
//	mq := &MyMessageQueue{
//		handlers: make(map[string][]EventHandler),
//	}
//	handler := &MyEventHandler{}
//	mq.Consume("test", handler)
//	mq.Publish("test", Event{Name: "test", Data: "hello"})
//}

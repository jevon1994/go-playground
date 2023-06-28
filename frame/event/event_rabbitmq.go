// 事件结构体
type Event struct {
	Name string
	Data interface{}
}

// 事件处理器接口
type EventHandler interface {
	Handle(event Event)
}

// 事件总线结构体
type EventBus struct {
	handlers map[string][]EventHandler
}

// 注册事件处理器
func (bus *EventBus) RegisterHandler(eventType string, handler EventHandler) {
	if _, ok := bus.handlers[eventType]; !ok {
		bus.handlers[eventType] = []EventHandler{}
	}
	bus.handlers[eventType] = append(bus.handlers[eventType], handler)
}

// 分发事件
func (bus *EventBus) Publish(event Event) {
	handlers := bus.handlers[event.Name]
	for _, handler := range handlers {
		handler.Handle(event)
	}
}

// 事件处理器实现
type MyEventHandler struct{}

func (h *MyEventHandler) Handle(event Event) {
	// 实现对事件的响应处理逻辑
	fmt.Printf("handling event %v\n", event.Name)
}

// 示例代码
func main() {
	bus := &EventBus{
		handlers: make(map[string][]EventHandler),
	}
	handler := &MyEventHandler{}
	bus.RegisterHandler("test", handler)
	bus.Publish(Event{Name: "test", Data: "hello"})
}

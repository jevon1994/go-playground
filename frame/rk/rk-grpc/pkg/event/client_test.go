package main

import (
	"context"
	"fmt"
	_ "github.com/asaskevich/EventBus"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"net/http"
	"testing"
)

func TestClient(t *testing.T) {
	doPush()

	eBus = initBus()
	InitTask()

	client := NewStreamServiceClient()
	initGin(client)
}

func EngineIinitialize() {
	eBus = initBus()
}

func doPush() {
	bytes := make([]byte, 10)
	request := StreamPushRequest{
		EventData: map[string]interface{}{"eventdata": bytes},
	}
	client := NewStreamServiceClient()
	push, _ := client.Push(context.Background(), &request)
	fmt.Println(push)
}

func initBus() *RequestEventBus {
	eventBus := NewRequestEventBus()
	eventBus.AddHandler("cloud-hcce.deployments.create", &TestEventHandler{})
	for topic, handler := range HandlerMap {
		eventBus.Bus.Subscribe(topic, handler.Handle)
	}
	return eventBus
}

func initGin(client *StreamServiceClient) {
	// Create a new boot instance.
	engine := gin.Default()
	engine.Use(ParseEvent())
	engine.POST("/api/v1/stream/handler", client.Handle)
	engine.Run()
}

// Handle
// @Summary Handle
// @Id 1
// @Tags Handle
// @version 1.0
// @Param name query string true "name"
// @produce application/json
// @Success 200 {object} GreeterResponse
// @Router /api/v1/stream/handler [post]
func (client *StreamServiceClient) Handle(ctx *gin.Context) {
	event, _ := cloudevents.NewEventFromHTTPRequest(ctx.Request)
	//eBus.Dispatch(event)
	eBus.Bus.Publish(event.Type(), event)
	ctx.JSON(http.StatusOK, &Response{
		Message: "call back success",
		Code:    "00000",
		Data:    "",
	})
}

func InitTask() {
	c := cron.New()
	c.AddFunc("@every 2s", doPush)
	c.Start()
}

type TestEventHandler struct {
}

func (th *TestEventHandler) Handle(event *cloudevents.Event) error {
	fmt.Println(event)
	return nil
}

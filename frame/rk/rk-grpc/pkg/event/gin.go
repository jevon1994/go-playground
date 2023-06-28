package client

import (
	"context"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/client"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	se "rk-grpc/pkg/event"
)

type StreamGinServiceClient struct {
	*se.RequestEventBus
	client.Client
}

func NewStreamGinServiceClient() *StreamGinServiceClient {

	c, err := cloudevents.NewClientHTTP()
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}

	se.ReqBus = se.InitRequestEventBus()
	return &StreamGinServiceClient{se.ReqBus,
		c}
}

func (client *StreamGinServiceClient) InitEngineofGin(engine *gin.Engine) {
	engine.Use(se.ParseEvent())
	engine.POST(STREAM_HANDLER_URI, client.Handle)
}

func (client *StreamGinServiceClient) Poll(ctx context.Context, out *StreamPushRequest) (Response, error) {
	panic("11")
}

func (client *StreamGinServiceClient) Push(ctx context.Context, out *StreamPushRequest) (Response, error) {
	// Create an Event.
	event, cetx := se.NewEventFromStreamPushRequest(ctx, out)
	event.SetData(cloudevents.ApplicationJSON, out.Bytes)
	// Send that Event.
	if result := client.Send(cetx, event); cloudevents.IsUndelivered(result) {
		log.Fatalf("failed to send, %v", result)
	} else {
		log.Printf("sent: %v", event)
	}
	return Response{
		Code:    "00000",
		Message: "send event success",
		Data:    event.ID(),
	}, nil
}

//${biz}/api/v1/stream/handler
//func (client *StreamServiceClient) Handle(ctx context.Context, in *StreamCallBackRequest) (Response, error) {
//	panic("111")
//}

// Handle
// @Summary Handle
// @Tags Handle
// @version 1.0
// @produce application/json
// @Success 200 {object} Response
// @Router /api/v1/stream/handler [post]
func (client *StreamGinServiceClient) Handle(ctx *gin.Context) {
	event, _ := cloudevents.NewEventFromHTTPRequest(ctx.Request)
	se.ReqBus.Bus.Publish(event.Type(), event)
	ctx.JSON(http.StatusOK, &Response{
		Message: "call back success",
		Code:    "00000",
		Data:    "",
	})
}

package event

import (
	"context"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/cloudevents/sdk-go/v2/client"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type StreamGinServiceClient struct {
	*AbstractStreamServiceClient
	client.Client
}

func NewStreamGinServiceClient() *StreamGinServiceClient {
	c, err := cloudevents.NewClientHTTP()
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}

	ReqBus = InitRequestEventBus()
	return &StreamGinServiceClient{&AbstractStreamServiceClient{ReqBus},
		c}
}

func (client *StreamGinServiceClient) InitEngineofGin(engine *gin.Engine) {
	engine.Use(ParseEvent())
	engine.POST(STREAM_HANDLER_URI, client.Handle)
}

func (client *StreamGinServiceClient) Push(ctx context.Context, out *StreamPushRequest) (Response, error) {
	// Create an Event.
	reqEvent, cetx, err := NewEventFromStreamPushRequest(ctx, out)
	reqEvent.SetData(cloudevents.ApplicationJSON, map[string][]byte{"evetdata": out.Data})
	// Send that Event.
	if result := client.Send(cetx, reqEvent); cloudevents.IsACK(result) {
		log.Printf("send success, %v", result)
	} else {
		log.Fatalf("failed to send, %v", result)
	}
	return Response{
		Code:    "00000",
		Message: "send event success",
		Data:    reqEvent.ID(),
	}, err
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
	reqEvent, _ := cloudevents.NewEventFromHTTPRequest(ctx.Request)
	ReqBus.Bus.Publish(reqEvent.Type(), reqEvent)
	ctx.JSON(http.StatusOK, &Response{
		Message: "call back success",
		Code:    "00000",
		Data:    "",
	})
}

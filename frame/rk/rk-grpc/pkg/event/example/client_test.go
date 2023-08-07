package example

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"rk-grpc/pkg/event"
	"testing"
)

var client *event.StreamGinServiceClient

func TestClient(t *testing.T) {
	//test
	InitTask()
	client = event.NewStreamGinServiceClient()

	initGin(client.InitEngineofGin)
}

func doPush() {
	data := make(map[string]interface{})
	data["test_record_id"] = "A111-122323"
	marshal, _ := json.Marshal(data)

	m := make(map[string]interface{})
	m["idc"] = "unicom"
	m["env"] = "dev"
	m["namespace"] = "unicom"
	m["service"] = "cloud-hcce"

	request := event.StreamPushRequest{
		Type:       event.GetCallBackType("cloud-hcce", "deploments.create"),
		Target:     event.STREAM_SERVICE_URL,
		Extensions: m,
		Data:       marshal,
	}
	push, _ := client.Push(context.Background(), &request)
	fmt.Println(push)
}

func initGin(fn func(engine *gin.Engine)) {
	// Create a new boot instance.
	engine := gin.Default()
	fn(engine)
	engine.Run()
}

func InitTask() {
	c := cron.New()
	c.AddFunc("@every 2s", doPush)
	c.Start()
}

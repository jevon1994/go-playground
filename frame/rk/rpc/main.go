// Copyright (c) 2021 rookie-ninja
//
// Use of this source code is governed by an Apache-style
// license that can be found in the LICENSE file.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/robfig/cron/v3"
	"github.com/rookie-ninja/rk-boot/v2"
	"github.com/rookie-ninja/rk-demo/api/gen/v1"
	"github.com/rookie-ninja/rk-grpc/v2/boot"
	"google.golang.org/grpc"
	"log"
)

func main() {
	boot := rkboot.NewBoot()
	InitTask()
	// register grpc
	entry := rkgrpc.GetGrpcEntry("rpc")
	entry.AddRegFuncGrpc(registerGreeter)
	entry.AddRegFuncGw(greeter.RegisterGreeterHandlerFromEndpoint)

	// Bootstrap
	boot.Bootstrap(context.TODO())

	// Wait for shutdown sig
	boot.WaitForShutdownSig(context.TODO())
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
	client, _ := cloudevents.NewClientHTTP()
	event := cloudevents.NewEvent()
	event.SetSource("111")
	event.SetType("ccc.ccc")
	cetx := cloudevents.ContextWithTarget(context.Background(), "http://localhost:8080/v1/hello")
	event.SetData(cloudevents.ApplicationJSON, marshal)
	// Send that Event.
	if result := client.Send(cetx, event); cloudevents.IsACK(result) {
		log.Printf("send success, %v", result)
	} else {
		log.Fatalf("failed to send, %v", result)
	}
}

func InitTask() {
	c := cron.New()
	c.AddFunc("@every 2s", doPush)
	c.Start()
}

func registerGreeter(server *grpc.Server) {
	greeter.RegisterGreeterServer(server, &GreeterServer{})
}

//GreeterServer GreeterServer struct
type GreeterServer struct{}

// Hello response with hello message
func (server *GreeterServer) Test(ctx context.Context, in *greeter.CloudEvent) (*greeter.GreeterResponse, error) {
	fmt.Println(in)
	return &greeter.GreeterResponse{}, nil
}

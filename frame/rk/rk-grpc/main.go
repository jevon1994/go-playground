// Copyright (c) 2021 rookie-ninja
//
// Use of this source code is governed by an Apache-style
// license that can be found in the LICENSE file.
package main

import (
	"context"
	"embed"
	_ "embed"
	"fmt"
	_ "github.com/rookie-ninja/rk-boot/gf"
	rkboot "github.com/rookie-ninja/rk-boot/v2"
	rkentry "github.com/rookie-ninja/rk-entry/v2/entry"
	"os"
	proto "rk-grpc/api/gen/v1"
)

//go:embed boot.yaml
var boot []byte

//go:embed api/gen/v1
var docsFS embed.FS

//go:embed api/gen/v1
var staticFS embed.FS

func init() {
	rkentry.GlobalAppCtx.AddEmbedFS(rkentry.DocsEntryType, "greeter", &docsFS)
	rkentry.GlobalAppCtx.AddEmbedFS(rkentry.SWEntryType, "greeter", &docsFS)
	rkentry.GlobalAppCtx.AddEmbedFS(rkentry.StaticFileHandlerEntryType, "greeter", &staticFS)
}

func main() {
	// Bootstrap basic entries from rk-boot config.
	//rkentry.BootstrapPluginEntryFromYAML(boot)
	// Set REGION=beijing
	os.Setenv("REGION", "test")

	// Create a new boot instance.
	boot := rkboot.NewBoot()

	config := rkentry.GlobalAppCtx.GetConfigEntry("demo").Viper.GetString("demo")
	// Load config which is config/beijing.yaml
	fmt.Println(config)

	// Bootstrap
	boot.Bootstrap(context.Background())

	// Wait for shutdown sig
	boot.WaitForShutdownSig(context.Background())

	//// Bootstrap rk-grpc entry from rk-boot config
	//res := rkgrpc.RegisterGrpcEntryYAML(boot)
	//
	//// Get GrpcEntry
	//grpcEntry := res["greeter"].(*rkgrpc.GrpcEntry)
	//// Register gRPC server
	//grpcEntry.AddRegFuncGrpc(func(server *grpc.Server) {
	//	proto.RegisterGreeterServer(server, &GreeterServer{})
	//})
	//// Register rk-grpc-gateway func
	//grpcEntry.AddRegFuncGw(proto.RegisterGreeterHandlerFromEndpoint)
	//
	//// Bootstrap rk-grpc entry
	//grpcEntry.Bootstrap(context.Background())
	//
	//// Wait for shutdown signal
	//rkentry.GlobalAppCtx.WaitForShutdownSig()
	//
	//// Interrupt rk-gin entry
	//grpcEntry.Interrupt(context.Background())
}

// GreeterServer Implementation of GreeterServer.
type GreeterServer struct{}

// Greeter Handle Greeter method.
func (server *GreeterServer) Greeter(context.Context, *proto.GreeterRequest) (*proto.GreeterResponse, error) {
	return &proto.GreeterResponse{}, nil
}

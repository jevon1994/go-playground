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
	rkentry "github.com/rookie-ninja/rk-entry/v2/entry"
	rkgrpc "github.com/rookie-ninja/rk-grpc/v2/boot"
	"net/http"
	"os"
	"rk-grpc/cmd"
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
	cmd.InitZmq()
	cmd.InitTask()
	bootstrap()
}

func bootstrap() {
	// Bootstrap basic entries from rk-boot config.
	//rkentry.BootstrapPluginEntryFromYAML(boot)
	// Set REGION=beijing
	os.Setenv("REGION", "test")

	// Create a new boot instance.
	//boot := rkboot.NewBoot()
	//
	//config := rkentry.GlobalAppCtx.GetConfigEntry("demo").Viper.GetString("demo")
	//// Load config which is config/beijing.yaml
	//fmt.Println(config)
	//
	//// Bootstrap
	//boot.Bootstrap(context.Background())
	//
	//// Wait for shutdown sig
	//boot.WaitForShutdownSig(context.Background())

	//// Bootstrap rk-grpc entry from rk-boot config
	res := rkgrpc.RegisterGrpcEntryYAML(boot)

	// Get GrpcEntry
	grpcEntry := res["greeter"].(*rkgrpc.GrpcEntry)
	// Register gRPC server
	m := &MyHandler{}
	grpcEntry.HttpMux.Handle("/v2/ss", m)

	cmd.InitGrpc(grpcEntry)
	// Bootstrap rk-grpc entry
	grpcEntry.Bootstrap(context.Background())
	// Wait for shutdown signal
	rkentry.GlobalAppCtx.WaitForShutdownSig()

	//// Interrupt rk-web entry
	//grpcEntry.Interrupt(context.Background())
}

// 自定义的处理器
type MyHandler struct{}

// 实现 http.Handler 接口的 ServeHTTP 方法
func (h MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

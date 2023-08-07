package infra

import (
	"google.golang.org/grpc"
	greeter "rk-grpc/api/gen/v1"
	"rk-grpc/internal/global"
	"rk-grpc/internal/server"
)

func registerGreeter(g *grpc.Server) {
	greeter.RegisterGreeterServer(g, server.NewGreeterServer())
}

func init() {
	global.AddRegFuncGrpc(registerGreeter)
	global.AddRegFuncGw(greeter.RegisterGreeterHandlerFromEndpoint)
}

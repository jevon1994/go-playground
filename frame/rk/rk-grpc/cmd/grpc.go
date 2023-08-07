package cmd

import (
	rkgrpc "github.com/rookie-ninja/rk-grpc/v2/boot"
	"rk-grpc/internal/global"
	_ "rk-grpc/internal/infra"
)

func InitGrpc(grpcEntry *rkgrpc.GrpcEntry) {
	grpcEntry.AddGwMuxOptions(global.GetGwMuxOptions()...)
	grpcEntry.AddUnaryInterceptors(global.GetInterceptos()...)
	grpcEntry.AddRegFuncGw(global.GetGwRegFs()...)
	grpcEntry.AddRegFuncGrpc(global.GetGrpcRegFs()...)
}

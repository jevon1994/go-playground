package global

import (
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	rkgrpc "github.com/rookie-ninja/rk-grpc/v2/boot"
	"google.golang.org/grpc"
)

var Entry = &rkgrpc.GrpcEntry{}

func AddInterceptor(ipt ...grpc.UnaryServerInterceptor) {
	Entry.AddUnaryInterceptors(ipt...)
}

func AddRegFuncGw(f ...rkgrpc.GwRegFunc) {
	Entry.AddRegFuncGw(f...)
}

func AddRegFuncGrpc(f ...rkgrpc.GrpcRegFunc) {
	Entry.AddRegFuncGrpc(f...)
}

func AddGwMuxOption(opt ...runtime.ServeMuxOption) {
	Entry.AddGwMuxOptions(opt...)
}

func GetInterceptos() []grpc.UnaryServerInterceptor {
	return Entry.UnaryInterceptors
}

func GetGwMuxOptions() []runtime.ServeMuxOption {
	return Entry.GwMuxOptions
}

func GetGrpcRegFs() []rkgrpc.GrpcRegFunc {
	return Entry.GrpcRegF
}

func GetGwRegFs() []rkgrpc.GwRegFunc {
	return Entry.GwRegF
}

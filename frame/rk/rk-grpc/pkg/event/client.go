package event

import (
	"context"
)

const (
	STREAM_SERVICE_URI        = "/api/v1/stream"
	STREAM_SERVICE_URL        = "http://localhost:8080/api/v1/stream/handler"
	STREAM_HANDLER_URI        = "/api/v1/stream/handler"
	STREAM_HANDLER_HOST_LOCAL = "localhost:8080"
)

type StreamService interface {
	Push(ctx context.Context, out *StreamPushRequest) (Response, error)
	Poll(ctx context.Context, out *StreamPushRequest) (Response, error)
}

func GetStreamHandlerURL(host string) string {
	return host + STREAM_HANDLER_URI
}

func GetStreamHandlerURLofLocalService() string {
	return GetStreamHandlerURL(STREAM_HANDLER_HOST_LOCAL + STREAM_HANDLER_URI)
}

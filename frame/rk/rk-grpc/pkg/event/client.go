package client

import (
	"context"
)

const (
	STREAM_HANDLER_URI = "/api/v1/stream/handler"
)

type StreamPushRequest struct {
	Url        string                 `json:"url"`
	Type       string                 `json:"type"`
	Source     string                 `json:"source"`
	Extensions map[string]interface{} `json:"extensions"`
	Bytes      []byte                 `json:"data"`
}

//type StreamCallBackRequest struct {
//	Idc       string `json:"idc"`
//	Env       string `json:"env"`
//	Namespace string `json:"namespace"`
//}

type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
}

type StreamService interface {
	Push(ctx context.Context, out *StreamPushRequest) (Response, error)
	Poll(ctx context.Context, out *StreamPushRequest) (Response, error)
}

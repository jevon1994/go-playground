package event

import "context"

type AbstractStreamServiceClient struct {
	*RequestEventBus
}

func (ss *AbstractStreamServiceClient) Push(ctx context.Context, out *StreamPushRequest) (Response, error) {
	//TODO implement me
	panic("implement me")
}

func (ss *AbstractStreamServiceClient) Poll(ctx context.Context, out *StreamPushRequest) (Response, error) {
	//TODO implement me
	panic("implement me")
}

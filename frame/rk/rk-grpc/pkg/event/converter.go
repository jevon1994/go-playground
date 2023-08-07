package event

import (
	"context"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	"github.com/google/uuid"
)

func NewEventFromStreamPushRequest(ctx context.Context, out *StreamPushRequest) (cloudevents.Event, context.Context, error) {
	event := cloudevents.NewEvent()
	event.SetType(out.Type)
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return event, ctx, err
	}
	event.SetID(newUUID.String())
	if out.Source == "" {
		event.SetSource(STREAM_SERVICE_URI)
	} else {
		event.SetSource(out.Source)
	}
	event.SetType(out.Type)
	for k, v := range out.Extensions {
		event.SetExtension(k, v)
	}
	event.SetExtension("target", out.Target)
	// Set a target.
	cetx := cloudevents.ContextWithTarget(ctx, STREAM_SERVICE_URL)
	return event, cetx, nil
}

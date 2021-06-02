package main

import (
	"context"
	"fmt"
	"testing"
)

func TestHandler(t *testing.T) {
	e := new(EventHandler)
	ctx := context.Background()
	handle1(e, ctx)
	handle2(e.Handle, ctx)
}

type interceptor func(ctx context.Context, h handler)

type handler func(ctx context.Context)

func TestHandle(t *testing.T) {
	h := func(ctx context.Context) {
		fmt.Println("go to the supermarket...")
	}
	inter1 := func(ctx context.Context, h handler) {
		fmt.Println("clean the floor...")
		h(ctx)
	}
	ctx := context.Background()
	inter1(ctx, h)
}

func handle1(h IHandler, ctx context.Context) {
	h.Handle(ctx)
}

func handle2(f EventHandlerFunc, ctx context.Context) {
	f.HandleEvent(ctx)
}

type IHandler interface {
	Handle(ctx context.Context)
}

type EventHandler struct{}

func (e EventHandler) Handle(ctx context.Context) {
	fmt.Println(ctx)
}

type EventHandlerFunc func(ctx context.Context)

func (h EventHandlerFunc) HandleEvent(ctx context.Context) {
	h(ctx)
}

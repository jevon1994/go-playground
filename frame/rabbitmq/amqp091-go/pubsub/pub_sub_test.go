package main

import (
	"testing"
)

func TestPubSub(t *testing.T) {
	producer.pub()
	consumer.sub()
}

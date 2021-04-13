package main

import (
	"reflect"
	"testing"
	"time"
)

type TestStrc struct {
	a string
	b string
	c uint
	d uint8
	e []byte
	f time.Time
}

func TestReflect(t *testing.T) {
	t2 := TestStrc{
		a: "",
		b: nil,
	}
	of := reflect.ValueOf(t2)
	for i := 0; i < of.NumField(); i++ {

	}
}

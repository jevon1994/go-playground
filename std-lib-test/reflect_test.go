package main

import (
	"fmt"
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
		b: "",
	}
	of := reflect.ValueOf(t2)
	for i := 0; i < of.NumField(); i++ {
	}
}

func TestReflectNew(t *testing.T) {
	t3 := new(TestStrc)
	of := reflect.TypeOf(t3)
	elem := of.Elem()
	value := reflect.New(elem)
	fmt.Println(value)
}

func TestReflectSlice(t *testing.T) {
	var v *int
	of := reflect.TypeOf(v)
	sliceType := reflect.SliceOf(of)
	emptySlice := reflect.MakeSlice(sliceType, 1, 1)
	ints := emptySlice.Interface().([]*int)
	for i := 0; i < 2; i++ {
		ints = append(ints, &i)
	}
	fmt.Println(ints)
}

type Options struct {
	t *time.Time
}

type TestDto struct {
	Options
}

type Option interface {
	apply(*Options)
}

func (e emptyOption) apply(options *Options) {
	e.apply(options)
}

type emptyOption struct{}

type funcOption struct {
	f func(*Options)
}

func NewFuncOption(f func(*Options)) *funcOption {
	return &funcOption{
		f: f,
	}
}

func TestValueDate(t *testing.T) {
	t2 := new(TestStrc)
	of := reflect.ValueOf(t2)
	fmt.Println(of)
}

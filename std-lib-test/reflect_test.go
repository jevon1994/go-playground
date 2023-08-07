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
		a: "1",
		b: "2",
	}
	of := reflect.ValueOf(t2)
	for i := 0; i < of.NumField(); i++ {
	}
	if of.IsZero() {
		return
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
	strc := TestStrc{a: "1"}
	bytes := make([]byte, 0)
	of := reflect.ValueOf(bytes)
	fmt.Println(of.Type().String())
	personType := reflect.TypeOf(strc)
	filedName, isOk := personType.FieldByName("Name")
	if isOk {
		fmt.Println("FiledName =", filedName.Name, "Type =", filedName.Type, "Tag =", filedName.Tag)
	} else {
		fmt.Println("Filed Name not exist")
	}
}

type Item struct {
	Name string
	Age  int
}

func TestDeepEqual(t *testing.T) {
	m := &Item{
		Name: "1",
		Age:  0,
	}
	n := &Item{
		Name: "1",
		Age:  0,
	}

	fmt.Println(reflect.DeepEqual(m, n))
}

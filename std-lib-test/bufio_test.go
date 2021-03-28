package main

import (
	"fmt"
	"log"
	"strings"
	"testing"
)

type Function func(a int)

func TestReader(t *testing.T) {
	c := []string{"1", "2"}
	join := strings.Join(c, "?")
	fmt.Printf("%s \n", join)
	//b := []byte(`{"Name": "Wednesday", "Age": 6, "Parents": ["Gomez", "Morticia"]}`)
	//var f interface{}
	//err := json.Unmarshal(b, &f)
	//if err == nil{
	//	if m := f.(map[string]interface{}); true{
	//		fmt.Print(m)
	//
	//	}
	//}
	defer func() {
		if err := recover(); err != nil {
			fmt.Print("====================")
		}
	}()
	aPanic(1)

}
func aPanic(a int) {
	panic(a)
}

func logHandler(function Function) Function {
	return func(a int) {
		defer func() {
			if err := recover(); err != nil {
				log.Fatal(err)
			}
		}()
		function(a)
	}

}

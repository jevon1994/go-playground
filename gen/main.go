package main

import "fmt"

var TestDeploy map[string]map[string]byte

func main() {
	//m := make(map[string]bool, 10)
	test := make(map[string]map[string]byte)
	//test["test"] = m
	TestDeploy = test
	fmt.Println(1 & 1)

}

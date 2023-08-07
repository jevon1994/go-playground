package main

import (
	"embed"
	"fmt"
	"testing"
)

//go:embed res
var f embed.FS

func TestEmbed(t *testing.T) {
	data, err := f.ReadFile("res/hello.go")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data)) // hello world!
}

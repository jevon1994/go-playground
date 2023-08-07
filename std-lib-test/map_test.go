package main

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	m := make(map[string]int)

	delete(m, "1")

	fmt.Println(m)
}

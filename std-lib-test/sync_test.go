package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestSync(t *testing.T) {
	once()
}

func once() {

	once := &sync.Once{}
	for i := 0; i < 4; i++ {
		i := i
		go func() {
			once.Do(func() {
				fmt.Printf("first %d\n", i)
			})
		}()
	}
}

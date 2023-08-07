package bit

import (
	"fmt"
	"testing"
)

func TestBit(t *testing.T) {
	a := 4
	b := 5
	fmt.Println((a ^ b) < 0)
}

func TestOnce(t *testing.T) {
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
	fmt.Println(add(5, 6))
	fmt.Println(minus(6, 5))
	fmt.Println(multiply(5, 6))
	fmt.Println(divide(30, 6))
}

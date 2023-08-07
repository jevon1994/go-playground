package algorithms

import (
	"fmt"
	"testing"
)

func TestFact(t *testing.T) {
	Fact(5)
	fmt.Println(add(10, 5))

}

func add(a int, b int) int {
	var s, c int
	for b != 0 {
		s = a ^ b
		c = (a & b) << 1
		a = s
		b = c
	}
	return a
}

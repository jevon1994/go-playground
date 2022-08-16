package algorithms

import "fmt"

func Fact(n int) int {
	if n < 1 {
		fmt.Println("start...")
		return 1
	}
	fmt.Printf("f(%d) = %d * f(%d)\n", n, n, n-1)
	z := n * Fact(n-1)
	fmt.Printf("f(%d) = %d \n", n, z)
	return z
}

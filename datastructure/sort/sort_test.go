package sort

import (
	"fmt"
	"testing"
)

var arr = []int{2, 4, 5, 9, 3, 1, 0, 6}

func Test_Bubble(t *testing.T) {
	Bubble(arr)
	fmt.Println(arr)
}

func Test_insertion(t *testing.T) {
	Insertion(arr)
	fmt.Println(arr)
}

func Test_shell(t *testing.T) {
	Shell(arr)
	fmt.Println(arr)
}

func Test_Msort(t *testing.T) {
	ints := make([]int, len(arr))
	Msort(arr, ints, 0, len(arr)-1)
	fmt.Println(arr)
}

func Test_quick(t *testing.T) {
	Quick(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

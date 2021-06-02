package sort

import (
	"fmt"
	"testing"
)

func TestBubble(t *testing.T) {
	arr := []int{2, 4, 5, 9, 3, 1, 0, 6}
	bubble(arr)
	fmt.Println(arr)
}

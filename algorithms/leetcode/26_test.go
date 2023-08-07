package main

import (
	"fmt"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{0, 0, 1, 1, 2, 2, 3, 4, 5, 5}
	duplicates := removeDuplicates(nums)
	fmt.Println(duplicates)
}

func removeDuplicates(nums []int) int {
	length := len(nums)
	if len(nums) == 0 {
		return 0
	}
	var slow, fast int
	for fast < length {
		if nums[slow] != nums[fast] {
			slow = slow + 1
			nums[slow] = nums[fast]
		}
		fast = fast + 1
	}
	return slow + 1
}

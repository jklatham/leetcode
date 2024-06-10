package main

import (
	"fmt"
)

func countPairs(nums []int, target int) int {
	var numPairs int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			fmt.Println(nums[i], " plus ", nums[j], "equals:", (nums[i] + nums[j]))
			if nums[i]+nums[j] < target {
				fmt.Println("Sum is less than", target, "incrementing numPairs")
				numPairs += 1
			}
		}
	}
	return numPairs
}

func main() {

	arr1 := []int{-1, 1, 2, 3, 1}
	arr2 := []int{-6, 2, 5, -2, -7, -1, 3}

	fmt.Println(countPairs(arr1, 2))
	fmt.Println(countPairs(arr2, -2))
}

// 17/21 are returning true
// should be 10
// Issue was j := 1. Every time i incremented, j was reset to 1 resulting in extra iterations. Changed to j := i + 1;

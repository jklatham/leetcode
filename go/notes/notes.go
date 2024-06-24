package notes

import "fmt"

func TryLoops(nums []int) int {
	for i := range nums {
		fmt.Println("Range:", nums[i])
	}

	for i := range len(nums) {
		fmt.Println("Range(len:", nums[i])
	}

	return 1
}

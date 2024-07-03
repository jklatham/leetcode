// 1509 - Minimum Difference Between Largest and Smallest Value in Three Moves

package problems

import "slices"

func MinDifference(nums []int) int {
	n := len(nums)
	if n <= 4 {
		return 0
	}

	slices.Sort(nums)

	min_diff := min(
		nums[n-1]-nums[3],
		nums[n-2]-nums[2],
		nums[n-3]-nums[1],
		nums[n-4]-nums[0])

	return min_diff
}

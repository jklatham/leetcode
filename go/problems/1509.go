// 1509 - Minimum Difference Between Largest and Smallest Value in Three Moves

package problems

import "fmt"

func threeMoves(nums []int) []int {
	count := 3
	avg := getArrAvg(nums)
	var left, right []int

	for i := 0; i < len(nums); i++ {
		if nums[i] < avg {
			left = append(left, nums[i])
		} else {
			right = append(right, nums[i])
		}
	}

	if len(left) > len(right) {
		for count > 0 {
			threeMovesDown(nums)
			fmt.Println("Down", nums, count)
			count--
		}
	} else {
		for count > 0 {
			threeMovesUp(nums)
			fmt.Println("Up", nums, count)
			count--
		}
	}

	return nums
}

func threeMovesDown(nums []int) []int {
	newHi := getHighestIndex(nums)
	newLo := getLowestIndex(nums)

	nums[newHi] = nums[newLo]

	return nums
}

func threeMovesUp(nums []int) []int {
	newHi := getHighestIndex(nums)
	newLo := getLowestIndex(nums)

	nums[newLo] = nums[newHi]

	return nums
}

func getMovedDiff(nums []int) int {
	if len(nums) <= 3 {
		return 0
	} else {
		newNums := threeMoves(nums)
		return (newNums[getHighestIndex(nums)] - newNums[getLowestIndex(nums)])
	}
}

func getArrDiff(nums []int) int {
	return nums[getHighestIndex(nums)] - nums[getLowestIndex(nums)]
}

func getArrAvg(nums []int) int {
	sum := 0

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}

	return sum / len(nums)
}

func getLowestIndex(nums []int) int {
	lowest := 0
	i := 0
	for i < len(nums) {
		if nums[i] < nums[lowest] {
			lowest = i
		}
		i++
	}
	return lowest
}

func getHighestIndex(nums []int) int {
	highest := 0
	i := 0
	for i < len(nums) {
		if nums[i] > nums[highest] {
			highest = i
		}
		i++
	}
	return highest
}

func MinDifference(nums []int) int {
	return getMovedDiff(nums)
}

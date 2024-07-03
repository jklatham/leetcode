// 1509 - Minimum Difference Between Largest and Smallest Value in Three Moves

package problems

func threeMoves(nums []int) []int {
	tempdownArr := append([]int{}, nums...)
	tempupArr := append([]int{}, nums...)

	up := getArrDiff(threeMovesUp(tempupArr))
	down := getArrDiff(threeMovesDown(tempdownArr))

	if up > down {
		return tempdownArr
	} else {
		return tempupArr
	}
}

func threeMovesDown(nums []int) []int {
	for i := 0; i < 3; i++ {
		newHi := getHighestIndex(nums)
		newLo := getLowestIndex(nums)

		nums[newHi] = nums[newLo]
	}

	return nums
}

func threeMovesUp(nums []int) []int {
	for i := 0; i < 3; i++ {
		newHi := getHighestIndex(nums)
		newLo := getLowestIndex(nums)

		nums[newLo] = nums[newHi]
	}
	return nums
}

func getMovedDiff(nums []int) int {
	if len(nums) <= 3 {
		return 0
	} else {
		nums = threeMoves(nums)
		return (nums[getHighestIndex(nums)] - nums[getLowestIndex(nums)])
	}
}

func getArrDiff(nums []int) int {
	return nums[getHighestIndex(nums)] - nums[getLowestIndex(nums)]
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

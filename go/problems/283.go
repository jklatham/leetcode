package problems

func MoveZeroes(nums []int) {
	count := 0

	for i := range nums {
		if nums[i] != 0 {
			nums[count] = nums[i]
			count += 1
		}
	}

	for count < len(nums) {
		nums[count] = 0
		count += 1
	}
}

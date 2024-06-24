package problems

func MinKBitFlips(nums []int, k int) int {
	currentFlips := 0
	result := 0

	for i := range nums {
		if i-k >= 0 && nums[i-k] == 2 {
			currentFlips -= 1
		}
		if (currentFlips+nums[i])%2 == 0 {
			if i+k > len(nums) {
				return -1
			}
			nums[i] = 2
			currentFlips += 1
			result += 1
		}
	}

	return result
}

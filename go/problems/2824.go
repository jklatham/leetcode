package problems

func CountPairs(nums []int, target int) int {
	var numPairs int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] < target {
				numPairs += 1
			}
		}
	}
	return numPairs
}

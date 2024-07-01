package problems

func ThreeConsecutiveOdds(arr []int) bool {
	count := 0
	i := 0
	for i < len(arr) {
		if arr[i]%2 != 0 {
			count += 1
		} else if count < 3 {
			count = 0
		}
		i++
	}

	if count >= 3 {
		return true
	} else {
		return false
	}
}

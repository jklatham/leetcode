package problems

import "fmt"

func ThreeConsecutiveOdds(arr []int) bool {
	count := 0
	i := 0
	for i < len(arr) {
		if arr[i]%2 != 0 {
			count += 1
		} else if count > 0 && count < 3 {
			count -= 1
		}
		fmt.Println(count)
		i++
	}

	if count >= 3 {
		return true
	} else {
		return false
	}
}

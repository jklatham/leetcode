// Minimum Number of Moves to Seat Everyone

package problems

import (
	"slices"
)

func MinMovesToSeat(seats []int, students []int) int {
	var res int

	slices.Sort(seats)
	slices.Sort(students)

	for i := range seats {
		val := seats[i] - students[i]
		if val < 0 {
			res += -val
		} else {
			res += val
		}
	}
	return res
}

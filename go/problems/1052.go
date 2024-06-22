package problems

import "fmt"

// since customers.length == grumpy.length - hash map?
// if i >= len(grumpy)-minutes { grumpy[i] == 0 }

// HiddenTech can be used in any consecutive <minutes> not just the last <minutes>
// Find the MAXIMUM number of satisfied customers.

func MaxSatisfied(customers []int, grumpy []int, minutes int) int {
	low := 0
	high := 0
	maxSat := 0
	m1 := make(map[int][]int)

	for i := 0; i < len(customers); i++ {
		m1[i] = append(m1[i], customers[i])
	}
	fmt.Println(m1)

	for high < (len(customers)) {
		if (high-low+1) < minutes || grumpy[high] == 0 {
			maxSat += customers[high]
			//fmt.Println(customers[high], customers[low], maxSat)
			high += 1

		} else {
			maxSat += customers[high]
			maxSat -= customers[low]
			//fmt.Println(customers[high], customers[low], maxSat)
			high += 1
			low += 1

		}
	}

	return maxSat
}

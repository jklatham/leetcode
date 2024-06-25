package problems

// since customers.length == grumpy.length - hash map?
// if i >= len(grumpy)-minutes { grumpy[i] == 0 }

// HiddenTech can be used in any consecutive <minutes> not just the last <minutes>
// Find the MAXIMUM number of satisfied customers.

func MaxSatisfied(customers []int, grumpy []int, minutes int) int {
	low := 0
	high := 0

	// init
	for i := range customers {
		if grumpy[i] == 0 {
			low += customers[i]
		}
	}

	// additional satisfied customers if hidden tech used
	for i := 0; i < minutes; i++ {
		if grumpy[i] == 1 {
			high += customers[i]
		}
	}

	sat := high

	// Sliding Window to find maximum additional customers
	for i := minutes; i < len(customers); i++ {
		if grumpy[i-minutes] == 1 {
			high -= customers[i-minutes]
		}
		if grumpy[i] == 1 {
			high += customers[i]
		}
		if high > sat {
			sat = high
		}
	}

	return low + sat
}

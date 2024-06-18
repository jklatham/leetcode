package problems

// Intuition
/*
- Each worker can only have one job. (If job = True)
- Worker should choose the job that matches their skill level. (Worker <= Difficulty)
- More difficult jobs do not equate to more value. (Difficulty != Profit)
- (if worker[j] <= difficulty[i] && profit[i] >= workerProfit) job = True && profit += profit[i]
*/

/*
Time Complexity: O(m*n) (653ms - Beats 12%.)
Space Complexity: O(N) (6.98MB - Beats 92%.)
*/

func MaxProfitAssignment(difficulty []int, profit []int, worker []int) int {

	var total int
	var bestJob = make([]int, (len(worker)))

	for j := range worker {
		for i := range difficulty {
			if worker[j] >= difficulty[i] && profit[i] >= bestJob[j] {
				bestJob[j] = profit[i]
			}
		}
		total += bestJob[j]
	}

	return total
}

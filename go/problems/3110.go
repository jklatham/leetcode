package problems

func ScoreOfStrings(s string) int {
	var score int

	for i := 1; i < len(s); i++ {
		res := (int(s[i-1]) - int(s[i]))

		if res < 0 {
			score += -res
		} else {
			score += res
		}
	}

	return int(score)
}

package main

import (
	"fmt"
	"leetgode/problems"
)

func main() {

	arr1 := []int{2, 6, 4, 1}

	arr2 := []int{1, 2, 34, 3, 4, 5, 7, 23, 12}

	fmt.Println(problems.ThreeConsecutiveOdds(arr1))
	fmt.Println(problems.ThreeConsecutiveOdds(arr2))
}

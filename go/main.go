package main

import (
	"fmt"
	"leetgode/problems"
)

func main() {

	n1 := 8
	eL1 := [][]int{{0, 3}, {0, 4}, {1, 3}, {2, 4}, {2, 7}, {3, 5}, {3, 6}, {3, 7}, {4, 6}}

	fmt.Println(problems.GetAncestors(n1, eL1))
}

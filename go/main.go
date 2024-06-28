package main

import (
	"fmt"
	"leetgode/problems"
)

func main() {
	numCities := 5
	roads := [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 2}, {1, 3}, {2, 4}}
	fmt.Println(problems.MaximumImportance(numCities, roads))
}

package main

import (
	"fmt"
	"leetgode/problems"
)

func main() {

	arr1 := []int{-1, 1, 2, 3, 1}
	arr2 := []int{-6, 2, 5, -2, -7, -1, 3}

	fmt.Println(problems.CountPairs(arr1, 2))
	fmt.Println(problems.CountPairs(arr2, -2))

}

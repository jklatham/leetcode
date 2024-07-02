package main

import (
	"fmt"
	"leetgode/problems"
)

func main() {

	arr1 := []int{1, 2, 2, 1}
	arr2 := []int{2, 2}
	arr3 := []int{4, 9, 5}
	arr4 := []int{9, 4, 9, 8, 4}

	fmt.Println(problems.Intersect(arr1, arr2))
	fmt.Println(problems.Intersect(arr3, arr4))
}

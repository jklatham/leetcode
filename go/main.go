package main

import (
	"fmt"
	"leetgode/problems"
)

func main() {
	arr1 := []int{2, 3, 1, 3, 2, 4, 6, 7, 9, 2, 19}
	arr2 := []int{2, 1, 4, 3, 9, 6}
	arr3 := []int{28, 6, 22, 8, 44, 17}
	arr4 := []int{22, 28, 8, 6}

	fmt.Println(problems.RelativeSortArray(arr1, arr2))
	fmt.Println(problems.RelativeSortArray(arr3, arr4))

}

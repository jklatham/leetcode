package main

import (
	"fmt"
	"leetgode/problems"
)

func main() {

	diff1 := []int{2, 4, 6, 8, 10}
	prof1 := []int{10, 20, 30, 40, 50}
	worker1 := []int{4, 5, 6, 7}

	diff2 := []int{85, 47, 57}
	prof2 := []int{24, 66, 99}
	worker2 := []int{40, 25, 25}

	diff3 := []int{66, 1, 28, 73, 53, 35, 45, 60, 100, 44, 59, 94, 27, 88, 7, 18, 83, 18, 72, 63}
	prof3 := []int{66, 20, 84, 81, 56, 40, 37, 82, 53, 45, 43, 96, 67, 27, 12, 54, 98, 19, 47, 77}
	worker3 := []int{61, 33, 68, 38, 63, 45, 1, 10, 53, 23, 66, 70, 14, 51, 94, 18, 28, 78, 100, 16}

	fmt.Println("Total:", problems.MaxProfitAssignment(diff1, prof1, worker1))
	fmt.Println("Total:", problems.MaxProfitAssignment(diff2, prof2, worker2))
	fmt.Println("Total:", problems.MaxProfitAssignment(diff3, prof3, worker3))

}

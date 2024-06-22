package main

import (
	"fmt"
	"leetgode/problems"
)

func main() {

	customers := []int{1, 0, 1, 2, 1, 1, 7, 5}
	grumpy := []int{0, 1, 0, 1, 0, 1, 0, 1}
	minutes := 3

	c2 := []int{4, 10, 10}
	g2 := []int{1, 1, 0}
	m2 := 2

	c3 := []int{5, 8}
	g3 := []int{0, 1}
	m3 := 1

	fmt.Println("Total:", problems.MaxSatisfied(customers, grumpy, minutes))
	fmt.Println("Total:", problems.MaxSatisfied(c2, g2, m2))
	fmt.Println("Total:", problems.MaxSatisfied(c3, g3, m3))

}

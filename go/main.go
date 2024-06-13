package main

import (
	"fmt"
	"leetgode/problems"
)

func main() {

	seats := []int{3, 1, 5}
	students := []int{2, 7, 4}

	seats2 := []int{2, 2, 6, 6}
	students2 := []int{1, 3, 2, 6}

	seats3 := []int{4, 1, 5, 9}
	students3 := []int{1, 3, 2, 6}

	fmt.Println(problems.MinMovesToSeat(seats, students))
	fmt.Println(problems.MinMovesToSeat(seats2, students2))
	fmt.Println(problems.MinMovesToSeat(seats3, students3))

}

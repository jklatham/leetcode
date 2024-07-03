package main

import (
	"fmt"
	"leetgode/problems"
)

func main() {

	nums1 := []int{5, 3, 2, 4}
	nums2 := []int{1, 5, 0, 10, 14}
	nums3 := []int{3, 100, 20}
	nums4 := []int{6, 6, 0, 1, 1, 4, 6} // expect 2 got 4
	nums5 := []int{1, 5, 6, 14, 15}     // expect 1 got 4
	nums6 := []int{82, 81, 95, 75, 20}  // expect 1 got 13

	fmt.Println(problems.MinDifference(nums1))
	fmt.Println(problems.MinDifference(nums2))
	fmt.Println(problems.MinDifference(nums3))
	fmt.Println(problems.MinDifference(nums4))
	fmt.Println(problems.MinDifference(nums5))
	fmt.Println(problems.MinDifference(nums6))
}

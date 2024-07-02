// 350 - Intersection of Two Arrays II

// Naive

package problems

import (
	"slices"
)

func sortedOrNot(arr []int, n int) bool {
	if n == 0 || n == 1 {
		return true
	}

	return (arr[n-1] >= arr[n-2]) && sortedOrNot(arr, n-1)
}

func Intersect(nums1 []int, nums2 []int) []int {
	sub := []int{}
	n1 := len(nums1)
	n2 := len(nums2)
	if !sortedOrNot(nums1, n1) {
		slices.Sort(nums1)
	}
	if !sortedOrNot(nums2, n2) {
		slices.Sort(nums2)
	}

	i, j := 0, 0

	for i < n1 && j < n2 {
		if nums1[i] < nums2[j] {
			i++
		} else if nums2[j] < nums1[i] {
			j++
		} else {
			sub = append(sub, nums2[j])
			i++
			j++
		}
	}

	return sub
}

// Follow ups
// What if the given array is already sorted? How would you optimize?
/* sortedOrNot */
// What if num1's size is small compared to num2's size? Which algorithm is better?
// What if the elements of nums2 are stored on disk, and the memory is limited such that you cannot
// load all elements into the memory at once?

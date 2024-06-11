package problems

import (
	"slices"
)

func RelativeSortArray(arr1 []int, arr2 []int) []int {
	var res, end []int

	arr2_set := map[int]int{}
	arr1_cnt := map[int]int{}

	for i := 0; i < len(arr2); i += 1 {
		arr2_set[arr2[i]] = i
	}

	for i := 0; i < len(arr1); i += 1 {
		if _, ok := arr2_set[arr1[i]]; !ok {
			end = append(end, arr1[i])
		}
		arr1_cnt[arr1[i]] += 1
	}

	slices.Sort(end)

	for _, i := range arr2 {
		for range arr1_cnt[i] {
			res = append(res, i)
		}
	}

	res = append(res, end...)

	return res

}

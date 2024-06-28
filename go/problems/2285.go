// 2285 - Maximum Total Importance of Roads
/*
You are given an integer n denoting the number of cities in a country.
The cities are numbered from 0 to n - 1.

You are also given a 2D integer array roads where roads[i] = [ai, bi]
denotes that there exists a bidirectional road connecting cities ai and bi.

You need to assign each city with an integer value from 1 to n, where each value can only be used once.
The importance of a road is then defined as the sum of the values of the two cities it connects.

Return the maximum total importance of all roads possible after assigning the values optimally.
*/

// map value 1..n to city
// for each road return sum of value of connected cities
package problems

import (
	"fmt"
)

/* // returns map of cities:value
func getCityValueMap(n int) map[int]int {
	cityValue := make(map[int]int)
	for i := range n {
		cityValue[i] = i
	}
	return cityValue
} */

func MaximumImportance(n int, roads [][]int) int64 {
	cityValue := make(map[int]int)
	//cityValue := make([]int, n)
	fmt.Println(cityValue)
	// Count connections in roads
	for _, a := range roads {
		for _, b := range a {
			cityValue[b] += 1
		}
	}
	fmt.Println(cityValue)

	// Increment duplicates
	/* for i := 0; i <= len(cityValue); i++ {
		for j := 1; j < len(cityValue); j++ {
			if i != j && cityValue[i] == cityValue[j] {
				cityValue[j] += 1
				j = 1
			}
		}
	} */

	i := 0

	for i < len(cityValue) {
		for next := range cityValue {
			if cityValue[i] == cityValue[next] {
				cityValue[next] += 1
			}
			i++
		}
	}
	fmt.Println(cityValue)

	// find importance
	total := 0
	for i := 0; i < len(roads); i++ {
		total += (cityValue[roads[i][0]] + cityValue[roads[i][1]])
	}
	fmt.Println(total)
	maximum := int64(total)

	return maximum
}

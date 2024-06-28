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

import "sort"

func MaximumImportance(n int, roads [][]int) int64 {
	cityValue := make(map[int]int)
	// map each city to its value (number of connections)
	// increment city value for each city in road
	for _, a := range roads {
		cityValue[a[0]]++
		cityValue[a[1]]++
	}

	// Crate slice containing all nodes
	keys := make([]int, 0, len(cityValue))
	for k := range cityValue {
		keys = append(keys, k)
	}
	// Sort slice descending order based on city value
	sort.Slice(keys, func(i, j int) bool {
		return cityValue[keys[i]] > cityValue[keys[j]]
	})

	// assign city map value from highest to lowest
	cityMap := make(map[int]int)

	for i, k := range keys {
		cityMap[k] = n - i
	}

	// Iterate the roads and sum the value of connected nodes
	total := 0
	for _, road := range roads {
		total += cityMap[road[0]] + cityMap[road[1]]
	}

	return int64(total)
}

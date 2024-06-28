// Find the index of first occurrence in a String

package problems

import "strings"

func StrStr(haystack string, needle string) int {
	var firstOccurrence int
	if strings.Contains(haystack, needle) {
		firstOccurrence = strings.Index(haystack, needle)
	} else {
		firstOccurrence = -1
	}

	return firstOccurrence
}

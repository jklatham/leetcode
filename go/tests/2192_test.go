package problems

import (
	"leetgode/problems"
	"reflect"
	"testing"
)

func TestGetAncestors(t *testing.T) {
	t.Run("Case 1", func(t *testing.T) {
		nodes := 8
		edgelist := [][]int{{0, 3}, {0, 4}, {1, 3}, {2, 4}, {2, 7}, {3, 5}, {3, 6}, {3, 7}, {4, 6}}
		result := problems.GetAncestors(nodes, edgelist)
		expected := [][]int{{}, {}, {}, {0, 1}, {0, 2}, {0, 1, 3}, {0, 1, 2, 3, 4}, {0, 1, 2, 3}}

		getGetAncestors(t, result, expected)
	})
}

// New function "DeepEqual" - not typesafe
func getGetAncestors(t testing.TB, got, want [][]int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %q want %q", got, want)
	}
	/* 	if got != want {
		t.Errorf("got %q want %q", got, want)
	} */
}

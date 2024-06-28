package problems

import (
	"leetgode/problems"
	"testing"
)

/* func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output:	6
} */

func TestMaximumImportance(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		numCities := 5
		roads := [][]int{{0, 1}, {1, 2}, {2, 3}, {0, 2}, {1, 3}, {2, 4}}
		result := problems.MaximumImportance(numCities, roads)
		var expected int64 = 43
		getMaximumImportance(t, result, expected)
	})

	t.Run("Case 2", func(t *testing.T) {
		numCities := 5
		roads := [][]int{{0, 3}, {2, 4}, {1, 3}}
		result := problems.MaximumImportance(numCities, roads)
		var expected int64 = 20
		getMaximumImportance(t, result, expected)
	})
}

func getMaximumImportance(t testing.TB, got, want int64) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

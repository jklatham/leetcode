package problems

import (
	"leetgode/problems"
	"testing"
)

func TestStrStr(t *testing.T) {

	t.Run("Case 1", func(t *testing.T) {
		string1 := "sadbutsad"
		string2 := "sad"
		result := problems.StrStr(string1, string2)
		expected := 0
		getFirstOccurrence(t, result, expected)
	})

}

func getFirstOccurrence(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

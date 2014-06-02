package levenshtein

import "testing"

type LevenshteinTest struct {
	s1       string
	s2       string
	expected int
}

var LevenshteinTests = []LevenshteinTest{
	{
		"kitten",
		"sitting",
		3,
	},
}

func TestContains(t *testing.T) {
	for _, lt := range LevenshteinTests {
		actual := LevenshteinDitance(lt.s1, lt.s2)
		if lt.expected != actual {
			t.Errorf("Levenshtein distance of %s and %s: want %d got %d", lt.s1, lt.s2, lt.expected, actual)
		}
	}
}

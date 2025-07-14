package task2

import (
	"testing"
)


func TestWordCount(t *testing.T) {
	input := "foo bar FOO BAR Foo Bar fOO bAR"
	expected := map[string]int{
		"foo": 4,
		"bar": 4,
	} 

	got := CountWordFreq(input)

	for word, count := range got {
		if expected[word] != count {
			t.Fatalf("Expected %v, got %v instead", expected, got)
		}
	}
}

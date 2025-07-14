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

func TestIsPalindrome(t *testing.T) {
	tests := map[string]bool {
		"madam": true,
		"racecar": true,
		"": true,
		"a": true,
		"Madam": true,
		"RaceCar": true,
		"A man, a plan, a canal, Panama": true,
		"Was it a car or a cat I saw?": true,
		"No 'x' in Nixon": true,
	} 


	for input, output := range tests {
		if IsPalindrome(input) != output {
			t.Errorf("IsPalindrome(%s) = %t, got IsPalindrome(%s)= %t instead", input, output, input, IsPalindrome(input))
		}
	}
}

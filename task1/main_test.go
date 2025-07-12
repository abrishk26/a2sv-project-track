package main

import (
	"testing"
)

func TestParseGrades(t *testing.T) {
	inputs := "80.6 90.8 99 100"

	grades, err := parseGrades(inputs)

	if err != nil {
		t.Fatal(err)
	}

	expected := []float64{80.6, 90.8, 99.0, 100.0}

	for i := range grades {
		if grades[i] != expected[i] {
			t.Fatalf("Expected %f, got %f instead", expected[i], grades[i])
		}
	}
}

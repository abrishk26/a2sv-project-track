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

func TestCalculateAvg(t *testing.T) {
	input := []float64{80.3, 70.4, 90.1, 67.4}

	var expected float64 = 77.05

	res := calculateAvg(input)

	if expected != res {
		t.Fatalf("Expected %f, got %f instead", expected, res)
	}
}

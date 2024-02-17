package main

import (
	"math"
	"testing"
)

const tolerance = 1e-6 // Define a tolerance for floating-point comparisons

func TestLinearRegression(t *testing.T) {
	// Define test datasets
	testData := []struct {
		name     string
		dataset  []DataPoint
		expected struct {
			a float64
			b float64
		}
	}{
		{
			name: "Dataset 1",
			dataset: []DataPoint{
				{10, 8.04}, {8, 6.95}, {13, 7.58}, {9, 8.81}, {11, 8.33},
				{14, 9.96}, {6, 7.24}, {4, 4.26}, {12, 10.84}, {7, 4.82}, {5, 5.68},
			},
			expected: struct {
				a float64
				b float64
			}{3.0000909090909094, 0.500090909090909},
		},
		// Define other test datasets and their expected results similarly
	}

	// Run tests for each dataset
	for _, testData := range testData {
		t.Run(testData.name, func(t *testing.T) {
			a, b := linearRegression(testData.dataset)
			if !equalWithinTolerance(a, testData.expected.a, tolerance) || !equalWithinTolerance(b, testData.expected.b, tolerance) {
				t.Errorf("For %s, expected a = %f, b = %f, but got a = %f, b = %f", testData.name, testData.expected.a, testData.expected.b, a, b)
			}
		})
	}
}

func equalWithinTolerance(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

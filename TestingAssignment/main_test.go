package main

import (
	"testing"

	"github.com/montanaflynn/stats"
)

// Creating tests for each dataset
func TestLinRegCoef1(t *testing.T) {
	//Defining X and Y coordinates
	type Coordinate struct {
		X float64
		Y float64
	}
	//Test 1
	d1Coef := []Coordinate{
		{X: 10.0, Y: 8.04},
		{X: 8.0, Y: 6.95},
		{X: 13.0, Y: 7.58},
		{X: 9.0, Y: 8.81},
		{X: 11.0, Y: 8.33},
		{X: 14.0, Y: 9.96},
		{X: 6.0, Y: 7.24},
		{X: 4.0, Y: 4.26},
		{X: 12.0, Y: 10.84},
		{X: 7.0, Y: 4.82},
		{X: 5.0, Y: 5.68},
	}
	//Expected coefficient stated and running test resulted coefficient
	expectedCoeff := 0.5001

	statsCoef := make([]stats.Coordinate, len(d1Coef))
	for i, coord := range d1Coef {
		statsCoef[i] = stats.Coordinate{X: coord.X, Y: coord.Y}
	}

	resultCoeff := LinRegCoef(statsCoef)

	if resultCoeff != expectedCoeff {
		t.Errorf("Expected coefficient: %.4f, got: %.4f", expectedCoeff, resultCoeff)
	}
}

func TestLinRegCoef2(t *testing.T) {
	type Coordinate struct {
		X float64
		Y float64
	}
	//Test 2
	d2Coef := []Coordinate{
		{X: 10.0, Y: 9.14},
		{X: 8.0, Y: 8.14},
		{X: 13.0, Y: 8.74},
		{X: 9.0, Y: 8.77},
		{X: 11.0, Y: 9.26},
		{X: 14.0, Y: 8.10},
		{X: 6.0, Y: 6.13},
		{X: 4.0, Y: 3.10},
		{X: 12.0, Y: 9.13},
		{X: 7.0, Y: 7.26},
		{X: 5.0, Y: 4.74},
	}

	expectedCoeff := 0.5007

	statsCoef := make([]stats.Coordinate, len(d2Coef))
	for i, coord := range d2Coef {
		statsCoef[i] = stats.Coordinate{X: coord.X, Y: coord.Y}
	}

	resultCoeff := LinRegCoef(statsCoef)

	if resultCoeff != expectedCoeff {
		t.Errorf("Expected coefficient: %.4f, got: %.4f", expectedCoeff, resultCoeff)
	}
}

func TestLinRegCoef3(t *testing.T) {
	type Coordinate struct {
		X float64
		Y float64
	}
	//Test 3
	d3Coef := []Coordinate{
		{X: 10.0, Y: 7.46},
		{X: 8.0, Y: 6.77},
		{X: 13.0, Y: 12.74},
		{X: 9.0, Y: 7.11},
		{X: 11.0, Y: 7.81},
		{X: 14.0, Y: 8.84},
		{X: 6.0, Y: 6.08},
		{X: 4.0, Y: 5.39},
		{X: 12.0, Y: 8.15},
		{X: 7.0, Y: 6.42},
		{X: 5.0, Y: 5.73},
	}

	expectedCoeff := 0.4999

	statsCoef := make([]stats.Coordinate, len(d3Coef))
	for i, coord := range d3Coef {
		statsCoef[i] = stats.Coordinate{X: coord.X, Y: coord.Y}
	}

	resultCoeff := LinRegCoef(statsCoef)

	if resultCoeff != expectedCoeff {
		t.Errorf("Expected coefficient: %.4f, got: %.4f", expectedCoeff, resultCoeff)
	}
}

func TestLinRegCoef4(t *testing.T) {
	type Coordinate struct {
		X float64
		Y float64
	}
	//Test 4
	d4Coef := []Coordinate{
		{X: 8.0, Y: 6.58},
		{X: 8.0, Y: 5.76},
		{X: 8.0, Y: 7.71},
		{X: 8.0, Y: 8.84},
		{X: 8.0, Y: 8.47},
		{X: 8.0, Y: 7.04},
		{X: 8.0, Y: 5.25},
		{X: 19.0, Y: 12.50},
		{X: 8.0, Y: 5.56},
		{X: 8.0, Y: 7.91},
		{X: 8.0, Y: 6.89},
	}

	expectedCoeff := 0.4999

	statsCoef := make([]stats.Coordinate, len(d4Coef))
	for i, coord := range d4Coef {
		statsCoef[i] = stats.Coordinate{X: coord.X, Y: coord.Y}
	}

	resultCoeff := LinRegCoef(statsCoef)

	if resultCoeff != expectedCoeff {
		t.Errorf("Expected coefficient: %.4f, got: %.4f", expectedCoeff, resultCoeff)
	}
}

//Time to Benchmark our Functions

func BenchmarkLinRegCoef1(b *testing.B) {
	type Coordinate struct {
		X float64
		Y float64
	}
	//Test 1
	d1Coef := []Coordinate{
		{X: 10.0, Y: 8.04},
		{X: 8.0, Y: 6.95},
		{X: 13.0, Y: 7.58},
		{X: 9.0, Y: 8.81},
		{X: 11.0, Y: 8.33},
		{X: 14.0, Y: 9.96},
		{X: 6.0, Y: 7.24},
		{X: 4.0, Y: 4.26},
		{X: 12.0, Y: 10.84},
		{X: 7.0, Y: 4.82},
		{X: 5.0, Y: 5.68},
	}

	statsCoef := make([]stats.Coordinate, len(d1Coef))
	for i, coord := range d1Coef {
		statsCoef[i] = stats.Coordinate{X: coord.X, Y: coord.Y}
	}

	for i := 0; i < b.N; i++ {
		LinRegCoef(statsCoef)
	}
}

func BenchmarkLinRegCoef2(b *testing.B) {
	type Coordinate struct {
		X float64
		Y float64
	}
	//Test 2
	d2Coef := []Coordinate{
		{X: 10.0, Y: 9.14},
		{X: 8.0, Y: 8.14},
		{X: 13.0, Y: 8.74},
		{X: 9.0, Y: 8.77},
		{X: 11.0, Y: 9.26},
		{X: 14.0, Y: 8.10},
		{X: 6.0, Y: 6.13},
		{X: 4.0, Y: 3.10},
		{X: 12.0, Y: 9.13},
		{X: 7.0, Y: 7.26},
		{X: 5.0, Y: 4.74},
	}

	statsCoef := make([]stats.Coordinate, len(d2Coef))
	for i, coord := range d2Coef {
		statsCoef[i] = stats.Coordinate{X: coord.X, Y: coord.Y}
	}

	for i := 0; i < b.N; i++ {
		LinRegCoef(statsCoef)
	}
}

func BenchmarkLinRegCoef3(b *testing.B) {
	type Coordinate struct {
		X float64
		Y float64
	}
	//Test 3
	d3Coef := []Coordinate{
		{X: 10.0, Y: 7.46},
		{X: 8.0, Y: 6.77},
		{X: 13.0, Y: 12.74},
		{X: 9.0, Y: 7.11},
		{X: 11.0, Y: 7.81},
		{X: 14.0, Y: 8.84},
		{X: 6.0, Y: 6.08},
		{X: 4.0, Y: 5.39},
		{X: 12.0, Y: 8.15},
		{X: 7.0, Y: 6.42},
		{X: 5.0, Y: 5.73},
	}

	statsCoef := make([]stats.Coordinate, len(d3Coef))
	for i, coord := range d3Coef {
		statsCoef[i] = stats.Coordinate{X: coord.X, Y: coord.Y}
	}

	for i := 0; i < b.N; i++ {
		LinRegCoef(statsCoef)
	}
}

func BenchmarkLinRegCoef4(b *testing.B) {
	type Coordinate struct {
		X float64
		Y float64
	}
	//Test 4
	d4Coef := []Coordinate{
		{X: 8.0, Y: 6.58},
		{X: 8.0, Y: 5.76},
		{X: 8.0, Y: 7.71},
		{X: 8.0, Y: 8.84},
		{X: 8.0, Y: 8.47},
		{X: 8.0, Y: 7.04},
		{X: 8.0, Y: 5.25},
		{X: 19.0, Y: 12.50},
		{X: 8.0, Y: 5.56},
		{X: 8.0, Y: 7.91},
		{X: 8.0, Y: 6.89},
	}

	statsCoef := make([]stats.Coordinate, len(d4Coef))
	for i, coord := range d4Coef {
		statsCoef[i] = stats.Coordinate{X: coord.X, Y: coord.Y}
	}

	for i := 0; i < b.N; i++ {
		LinRegCoef(statsCoef)
	}
}

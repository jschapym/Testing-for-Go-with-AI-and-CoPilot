package main

import (
	"fmt"
)

// Define a struct to represent a dataset point
type DataPoint struct {
	X float64
	Y float64
}

// Define a function to perform linear regression
func linearRegression(data []DataPoint) (float64, float64) {
	n := len(data)
	sumX, sumY, sumXY, sumXX := 0.0, 0.0, 0.0, 0.0

	for _, point := range data {
		sumX += point.X
		sumY += point.Y
		sumXY += point.X * point.Y
		sumXX += point.X * point.X
	}

	b := (float64(n)*sumXY - sumX*sumY) / (float64(n)*sumXX - sumX*sumX)
	a := (sumY - b*sumX) / float64(n)

	return a, b
}

func main() {
	// Define datasets from the Anscombe Quartet
	dataset1 := []DataPoint{{10, 8.04}, {8, 6.95}, {13, 7.58}, {9, 8.81}, {11, 8.33}, {14, 9.96}, {6, 7.24}, {4, 4.26}, {12, 10.84}, {7, 4.82}, {5, 5.68}}
	dataset2 := []DataPoint{{10, 9.14}, {8, 8.14}, {13, 8.74}, {9, 8.77}, {11, 9.26}, {14, 8.10}, {6, 6.13}, {4, 3.10}, {12, 9.13}, {7, 7.26}, {5, 4.74}}
	dataset3 := []DataPoint{{10, 7.46}, {8, 6.77}, {13, 12.74}, {9, 7.11}, {11, 7.81}, {14, 8.84}, {6, 6.08}, {4, 5.39}, {12, 8.15}, {7, 6.42}, {5, 5.73}}
	dataset4 := []DataPoint{{8, 6.58}, {8, 5.76}, {8, 7.71}, {8, 8.84}, {8, 8.47}, {8, 7.04}, {8, 5.25}, {19, 12.50}, {8, 5.56}, {8, 7.91}, {8, 6.89}}

	// Perform linear regression on each dataset
	a1, b1 := linearRegression(dataset1)
	a2, b2 := linearRegression(dataset2)
	a3, b3 := linearRegression(dataset3)
	a4, b4 := linearRegression(dataset4)

	// Output the regression coefficients
	fmt.Println("Dataset 1: a =", a1, "b =", b1)
	fmt.Println("Dataset 2: a =", a2, "b =", b2)
	fmt.Println("Dataset 3: a =", a3, "b =", b3)
	fmt.Println("Dataset 4: a =", a4, "b =", b4)
}

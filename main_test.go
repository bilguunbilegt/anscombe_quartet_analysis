package main

import (
	"math"
	"testing"

	"github.com/montanaflynn/stats"
)

func TestCheckDataQuality(t *testing.T) {
	tests := []struct {
		x, y   []float64
		hasErr bool
	}{
		{[]float64{1, 2, 3}, []float64{4, 5, 6}, false},
		{[]float64{}, []float64{4, 5, 6}, true},
		{[]float64{1, 2, 3}, []float64{}, true},
		{[]float64{1, 2, 3}, []float64{4, 5}, true},
		{[]float64{1, 2, 3, 4}, []float64{4, 5, 6, 7}, false},
	}

	for _, test := range tests {
		err := checkDataQuality(test.x, test.y)
		if (err != nil) != test.hasErr {
			t.Errorf("checkDataQuality(%v, %v) = %v; want error: %v", test.x, test.y, err, test.hasErr)
		}
	}
}

func TestLinearRegression(t *testing.T) {
	x := []float64{1, 2, 3, 4, 5}
	y := []float64{2, 4, 6, 8, 10}

	series, err := linearRegression(x, y)
	if err != nil {
		t.Fatalf("linearRegression(%v, %v) returned error: %v", x, y, err)
	}

	// Check the first and last points for intercept and slope
	if series[0].Y != 2*x[0] || series[len(series)-1].Y != 2*x[len(x)-1] {
		t.Errorf("Expected first and last points to match intercept and slope. Got first point: %.2f, last point: %.2f", series[0].Y, series[len(series)-1].Y)
	}
}

func TestCalculateRSquared(t *testing.T) {
	x := []float64{1, 2, 3, 4, 5}
	y := []float64{2, 4, 6, 8, 10}
	intercept := 0.0
	slope := 2.0

	rSquared, err := calculateRSquared(x, y, intercept, slope)
	if err != nil {
		t.Fatalf("calculateRSquared(%v, %v, %.2f, %.2f) returned error: %v", x, y, intercept, slope, err)
	}

	expectedRSquared := 1.0
	if rSquared != expectedRSquared {
		t.Errorf("Expected R-squared %.2f, got %.2f", expectedRSquared, rSquared)
	}
}

func TestCorrelation(t *testing.T) {
	x := []float64{1, 2, 3, 4, 5}
	y := []float64{2, 4, 6, 8, 10}

	correlation, err := stats.Correlation(x, y)
	if err != nil {
		t.Fatalf("stats.Correlation(%v, %v) returned error: %v", x, y, err)
	}

	expectedCorrelation := 1.0
	if !almostEqual(correlation, expectedCorrelation, 1e-9) {
		t.Errorf("Expected correlation %.2f, got %.2f", expectedCorrelation, correlation)
	}
}

// Helper function to compare floating-point numbers with tolerance
func almostEqual(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

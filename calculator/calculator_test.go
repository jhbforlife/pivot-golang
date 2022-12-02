package calculator_test

import (
	"fmt"
	"testing"

	"github.com/jhbforlife/pivot-golang/calculator"
)

func TestCalculator(t *testing.T) {
	testCases := []struct {
		a        int
		b        int
		operator string
		want     int
	}{
		{123, 321, "+", 444},
		{900, 99, "+", 999},
		{2, 3, "+", 5},
		{321, 123, "-", 198},
		{24, 6, "-", 18},
		{1, 2, "-", -1},
		{10, 100, "*", 1000},
		{42, 0, "*", 0},
		{31, 2, "*", 62},
		{121, 11, "/", 11},
		{81, 9, "/", 9},
		{100, 0, "/", 0},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d%s%d", tc.a, tc.operator, tc.b), func(t *testing.T) {
			switch tc.operator {
			case "+":
				if got := calculator.Add(tc.a, tc.b); got != tc.want {
					t.Errorf("got: %d - want: %d", got, tc.want)
				}
			case "-":
				if got := calculator.Subtract(tc.a, tc.b); got != tc.want {
					t.Errorf("got: %d - want: %d", got, tc.want)
				}
			case "*":
				if got := calculator.Multiply(tc.a, tc.b); got != tc.want {
					t.Errorf("got: %d - want: %d", got, tc.want)
				}
			case "/":
				if got, err := calculator.Divide(tc.a, tc.b); err != nil {
					if tc.b != 0 {
						t.Errorf("%d is not 0 but returned error", tc.b)
					}
				} else if got != tc.want {
					t.Errorf("got: %d - want: %d", got, tc.want)
				}
			default:
				t.Errorf("invalid operator: %s", tc.operator)
			}
		})
	}
}

func TestPow(t *testing.T) {
	testCases := []struct {
		x    float64
		y    float64
		want float64
	}{
		{1, 10, 1},
		{100, 0, 1},
		{2, 3, 8},
		{9, 2, 81},
		{5, 5, 3125},
	}
	for _, tc := range testCases {
		if got := calculator.Pow(tc.x, tc.y); got != tc.want {
			t.Errorf("got: %f - want: %f", got, tc.want)
		}
	}
}

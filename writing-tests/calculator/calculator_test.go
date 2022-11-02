package calculator_test

import (
	"fmt"
	"testing"
)

func TestCalculator(t *testing.T) {
	testCases := []struct {
		a          int
		b          int
		operator   string
		wantResult int
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
		t.Run(fmt.Sprintf("%d %s %d", tc.a, tc.operator, tc.b), func(t *testing.T) {
			switch tc.operator {
			case "+":
			case "-":
			case "*":
			case "/":
			default:
			}
		})
	}
}

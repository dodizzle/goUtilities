package main

import (
	"testing"
)

func TestAbs(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{input: 5, expected: 5},
		{input: -5, expected: 5},
		{input: 0, expected: 0},
	}

	for _, test := range tests {
		result := positiveValue(test.input)
		if result != test.expected {
			t.Errorf("positiveValue(%d) = %d; expected %d", test.input, result, test.expected)
		}
	}
}

// write a test for the closestToZero function
func TestClosestToZero(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{input: []int{10, -2, 2, 3, 4, 5}, expected: 2},
		{input: []int{0, 1, 2, 3, 4, 5}, expected: 0},
		{input: []int{-1, -2, -3, -4, -5}, expected: -1},
		{input: nil, expected: -1},
		{input: []int{}, expected: 0},
	}

	for _, test := range tests {
		result := closestToZero(test.input)
		if result != test.expected {
			t.Errorf("closestToZero(%v) = %d; expected %d", test.input, result, test.expected)
		}
	}
}

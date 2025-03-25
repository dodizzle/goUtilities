package work_test

import (
	"testing"

	"github.com/dodizzle/goUtilities/interview1/work"
)

func TestPositiveValue(t *testing.T) {
	result := work.PositiveValue(-2)
	if result != 2 {
		t.Errorf("Expected 2, got %d", result)
	}
}

func TestClosestToZero(t *testing.T) {
	list := []int{6, 10, -2, 2, 3, 1, 4, 5}
	result := work.ClosestToZero(list)
	if result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}
}

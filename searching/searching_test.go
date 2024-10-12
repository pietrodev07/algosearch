package searching

import (
	"testing"
)

func TestLinearSearch(t *testing.T) {
	arr := []int{10, 20, 30, 40, 50}
	target := 30
	result := LinearSearch(arr, target)
	if result != 2 {
		t.Errorf("LinearSearch failed, expected %d, got %d", 2, result)
	}
}

func TestBinarySearch(t *testing.T) {
	arr := []int{10, 20, 30, 40, 50}
	target := 40
	result := BinarySearch(arr, target)
	if result != 3 {
		t.Errorf("BinarySearch failed, expected %d, got %d", 3, result)
	}
}

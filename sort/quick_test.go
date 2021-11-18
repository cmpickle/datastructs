package sort

import (
	"testing"
)

func TestSortBackwardsArray(t *testing.T) {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	expecting := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := Quick(arr, 0, len(arr)-1)
	if len(result) != len(expecting) {
		t.Errorf("Test sort backwards array failed, expected  length %d but got %d", len(expecting), len(result))
		return
	}
	for i, value := range result {
		if value != expecting[i] {
			t.Errorf("Test sort backwards array failed, expected %d at index %d but got %d", expecting[i], i, result[i])
			return
		}
	}
	t.Logf("Test sort backwards array success %v", expecting)
}

func TestSortPositiveAndNegativeArray(t *testing.T) {
	arr := []int{-9, 8, -17, 6, 5, -4, 3, 2, 0, 1, 0}
	expecting := []int{-17, -9, -4, 0, 0, 1, 2, 3, 5, 6, 8}
	result := Quick(arr, 0, len(arr)-1)
	if len(result) != len(expecting) {
		t.Errorf("Test sort backwards array failed, expected  length %d but got %d", len(expecting), len(result))
		return
	}
	for i, value := range result {
		if value != expecting[i] {
			t.Errorf("Test sort backwards array failed, expected %d at index %d but got %d", expecting[i], i, result[i])
			return
		}
	}
	t.Logf("Test sort backwards array success %v", expecting)
}

func TestSortLargeNumbersArray(t *testing.T) {
	arr := []int{1000000, -1000000, 9999999, -9999999}
	expecting := []int{-9999999, -1000000, 1000000, 9999999}
	result := Quick(arr, 0, len(arr)-1)
	if len(result) != len(expecting) {
		t.Errorf("Test sort backwards array failed, expected  length %d but got %d", len(expecting), len(result))
		return
	}
	for i, value := range result {
		if value != expecting[i] {
			t.Errorf("Test sort backwards array failed, expected %d at index %d but got %d", expecting[i], i, result[i])
			return
		}
	}
	t.Logf("Test sort backwards array success %v", expecting)
}

func TestSortSameNumberArray(t *testing.T) {
	arr := []int{1, 1, 1, 1}
	expecting := []int{1, 1, 1, 1}
	result := Quick(arr, 0, len(arr)-1)
	if len(result) != len(expecting) {
		t.Errorf("Test sort backwards array failed, expected  length %d but got %d", len(expecting), len(result))
		return
	}
	for i, value := range result {
		if value != expecting[i] {
			t.Errorf("Test sort backwards array failed, expected %d at index %d but got %d", expecting[i], i, result[i])
			return
		}
	}
	t.Logf("Test sort backwards array success %v", expecting)
}

func TestSortEmptyArray(t *testing.T) {
	arr := []int{}
	expecting := []int{}
	result := Quick(arr, 0, len(arr)-1)
	if len(result) != len(expecting) {
		t.Errorf("Test sort backwards array failed, expected  length %d but got %d", len(expecting), len(result))
		return
	}
	for i, value := range result {
		if value != expecting[i] {
			t.Errorf("Test sort backwards array failed, expected %d at index %d but got %d", expecting[i], i, result[i])
			return
		}
	}
	t.Logf("Test sort backwards array success %v", expecting)
}

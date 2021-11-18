package sort

import "testing"

func TestMergeSort(t *testing.T) {
	arr := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	expecting := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	Merge(arr, 0, len(arr)-1)
	result := arr
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

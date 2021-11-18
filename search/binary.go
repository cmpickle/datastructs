package search

func Binary(arr []int, target int) int {
	upper := len(arr) - 1
	lower := 0
	if target > arr[upper] || target < arr[lower] {
		return -1
	}
	for lower <= upper {
		idx := (upper + lower) / 2
		if arr[idx] == target {
			return idx
		}
		if target < arr[idx] {
			upper = idx - 1
		}
		if target > arr[idx] {
			lower = idx + 1
		}
	}
	return -1
}

// Shifted Binary Search (log n)

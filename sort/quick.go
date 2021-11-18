package sort

func Quick(arr []int, low int, high int) []int {
	if low <= high {
		pi := partition(arr, low, high)

		Quick(arr, low, pi-1)
		Quick(arr, pi+1, high)
	}

	return arr
}

func partition(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			tmp := arr[i]
			arr[i] = arr[j]
			arr[j] = tmp
		}
	}
	tmp := arr[i+1]
	arr[i+1] = arr[high]
	arr[high] = tmp

	return i + 1
}

package sort

func Merge(arr []int, l int, r int) {
	if l < r {
		m := l + (r-l)/2

		Merge(arr, l, m)
		Merge(arr, m+1, r)

		combine(arr, l, m, r)
	}
}

func combine(arr []int, l int, m int, r int) {
	// Sizes of two sub arrays
	t1 := m - l + 1
	t2 := r - m

	// Create temp arrays
	tmpL := make([]int, t1)
	tmpR := make([]int, t2)
	var i, j int

	// Copy data into temp arrays
	for i = 0; i < t1; i++ {
		tmpL[i] = arr[l+i]
	}
	for j = 0; j < t2; j++ {
		tmpR[j] = arr[m+1+j]
	}

	//Merge temp arrays
	i = 0
	j = 0
	k := l
	for i < t1 && j < t2 {
		if tmpL[i] < tmpR[j] {
			arr[k] = tmpL[i]
			i++
		} else {
			arr[k] = tmpR[j]
			j++
		}
		k++
	}

	// Copy remaining elements
	for i < t1 {
		arr[k] = tmpL[i]
		i++
		k++
	}
	for j < t2 {
		arr[k] = tmpR[j]
		j++
		k++
	}
}

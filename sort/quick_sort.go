package sort

// QuickSort QuickSort
func QuickSort(items IntSlice) IntSlice {
	size := len(items)

	if size < 2 {
		return items
	}

	pivot := items[0]
	less := make(IntSlice, 0, size)
	greater := make(IntSlice, 0, size)

	for i := 1; i < size; i++ {
		value := items[i]

		if value <= pivot {
			less = append(less, value)
		} else {
			greater = append(greater, value)
		}
	}

	QuickSort(less)
	QuickSort(greater)

	copy(items, less)
	lessSize := len(less)
	items[lessSize] = pivot
	copy(items[lessSize+1:], greater)

	return items
}

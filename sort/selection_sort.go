package sort

func findSmallestIndex(items IntSlice, begin int, end int) int {
	smallestIndex := begin
	smallest := items[begin]

	for i := begin + 1; i < end; i++ {
		if items[i] < smallest {
			smallestIndex = i
			smallest = items[i]
		}
	}

	return smallestIndex
}

// SelectionSort SelectionSort with slice
func SelectionSort(items IntSlice) IntSlice {
	length := len(items)

	// select smallest into first slot to last slot
	for i := 0; i < length; i++ {
		smallestIndex := findSmallestIndex(items, i, length)
		smallest := items[smallestIndex]

		// keep out the slot
		for j := smallestIndex; j > i; j-- {
			items[j] = items[j-1]
		}

		items[i] = smallest
	}

	return items
}

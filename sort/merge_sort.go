package sort

// MergeSort MergeSort
func MergeSort(items IntSlice) IntSlice {
	_mergeSort(items, 0, len(items))
	return items
}

func _mergeSort(items IntSlice, l, r int) {
	if r-l < 2 {
		return
	}

	m := (l + r) / 2

	_mergeSort(items, l, m)
	_mergeSort(items, m, r)

	_mergeHalfSorted(items, l, m, r)
}

func _mergeHalfSorted(items IntSlice, l, m, r int) {
	count := r - l
	tmpItems := make(IntSlice, count)

	i, j := l, m

	for c := 0; c < count; c++ {
		if i >= m {
			tmpItems[c] = items[j]
			j++
		} else if j >= r {
			tmpItems[c] = items[i]
			i++
		} else if items[i] <= items[j] {
			tmpItems[c] = items[i]
			i++
		} else {
			tmpItems[c] = items[j]
			j++
		}
	}

	copy(items[l:r], tmpItems)
}

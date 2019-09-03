package sort

// QuickSort Quick Sort
func QuickSort(items IntSlice) IntSlice {
	quickSortInPlace(items, 0, len(items)-1)
	return items
}

// QuickSort Quick Sort in place, include A[q] end A[r]
func quickSortInPlace(A IntSlice, p, r int) {
	if r <= p {
		return
	}

	q := quickSortPartition(A, p, r)
	quickSortInPlace(A, p, q-1)
	quickSortInPlace(A, q+1, r)
}

func quickSortPartition(A IntSlice, p, r int) int {
	// 最后一个位参考值
	x := A[r]
	i := p - 1

	// 从左到右，交换<x的值到前面位置
	for j := p; j < r; j++ {
		if A[j] <= x {
			i++
			A[i], A[j] = A[j], A[i]
		}
	}

	// 交换参考值
	A[i+1], A[r] = A[r], A[i+1]
	return i + 1
}

// QuickSort Quick sort
func quickSortMergeSlice(items IntSlice) IntSlice {
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

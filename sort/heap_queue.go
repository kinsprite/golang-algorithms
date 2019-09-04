package sort

// HeapMaximum HeapMaximum
func HeapMaximum(A IntSlice) int {
	return A[0]
}

// HeapExtractMax HeapExtractMax
func HeapExtractMax(A IntSlice) (int, IntSlice) {
	if len(A) < 1 {
		panic("Heap underflow")
	}

	max := A[0]
	newLen := len(A) - 1
	A[0] = A[newLen]
	newHeap := A[:newLen]
	MaxHeapify(newHeap, 0)
	return max, newHeap
}

// HeapIncreaseKey HeapIncreaseKey
func HeapIncreaseKey(A IntSlice, i, key int) IntSlice {
	if key < A[i] {
		panic("new key is smaller than current key")
	}

	A[i] = key

	for i >= 0 && A[parent(i)] < A[i] {
		A[parent(i)], A[i] = A[i], A[parent(i)]
		i = parent(i)
	}

	return A
}

// MaxHeapInsert MaxHeapInsert
func MaxHeapInsert(A IntSlice, key int) IntSlice {
	newA := append(A, key)
	HeapIncreaseKey(newA, len(newA)-1, key)
	return newA
}

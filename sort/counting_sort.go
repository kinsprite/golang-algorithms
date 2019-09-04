package sort

// CountingSort Counting Sort, 0 <= A[x] <= k
func CountingSort(A IntSlice, k int) IntSlice {
	C := make(IntSlice, k+1)

	// 各值的个数
	for _, a := range A {
		C[a]++
	}

	// 小于等于当前值的个数，即倒序插入的位置
	for i := 1; i <= k; i++ {
		C[i] += C[i-1]
	}

	B := make(IntSlice, len(A))

	for i := len(A) - 1; i >= 0; i-- {
		pos := C[A[i]] - 1
		B[pos] = A[i]
		C[A[i]]--
	}

	return B
}

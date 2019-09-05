package sort

import "math"

// BucketSort 0~999分成n个桶
func BucketSort(A IntSlice) IntSlice {
	const min = 0
	const max = 999

	n := len(A)
	step := float64(max-min+1) / float64(n)

	B := make([]IntSlice, n)

	// 把A[i]放到对应的桶中
	for i := 0; i < n; i++ {
		j := int(math.Floor(float64(A[i]-min) / step))
		B[j] = append(B[j], A[i])
	}

	// 对每个桶单独 Stable Sorting
	for i := 0; i < n; i++ {
		InsertionSort(B[i])
	}

	// 所有的桶按顺序充填回A
	c := 0

	for i := 0; i < n; i++ {
		len := len(B[i])

		for j := 0; j < len; j++ {
			A[c] = B[i][j]
			c++
		}
	}

	return A
}

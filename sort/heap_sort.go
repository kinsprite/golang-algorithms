package sort

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2 * (i + 1)
}

// MaxHeapify 假定 A[i] 的左右子树均为最大堆，即: 除A[i]之外一, 其余子节点均满足最大堆
// 下沉 A[i]，使得 A[i] 满足最大堆
func MaxHeapify(A IntSlice, i int) {
	l := left(i)
	r := right(i)
	largest := i

	if l < len(A) && A[l] > A[i] {
		largest = l
	}

	if r < len(A) && A[r] > A[largest] {
		largest = r
	}

	if largest != i {
		A[i], A[largest] = A[largest], A[i]
		MaxHeapify(A, largest)
	}
}

// BuildMaxHeap 从倒数第二层向上开始构建最大推. 才符合 MaxHeapify() 函数的条件
func BuildMaxHeap(A IntSlice) {
	for i := len(A)/2 - 1; i >= 0; i-- {
		MaxHeapify(A, i)
	}
}

// HeapSort 依次取堆顶元素
func HeapSort(A IntSlice) IntSlice {
	BuildMaxHeap(A)

	// 依次堆顶元素(最大值)与最后一个位置交换， 剩余堆满足  MaxHeapify() 函数的条件
	for i := len(A) - 1; i >= 1; i-- {
		A[i], A[0] = A[0], A[i]
		// 下沉交换之后的元素
		MaxHeapify(A[:i], 0)
	}

	return A
}

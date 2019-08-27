package sort

// BubbleSortInterface BubbleSort Interface
type BubbleSortInterface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

// Less IntSlice's less compare
func (items *IntSlice) Swap(i, j int) {
	(*items)[i], (*items)[j] = (*items)[j], (*items)[i]
}

// BubbleSort BubbleSort(items)
func BubbleSort(items BubbleSortInterface) BubbleSortInterface {
	length := items.Len()

	for i := length - 1; i > 0; i-- {
		for j := 1; j <= i; j++ {
			if items.Less(j, j-1) {
				items.Swap(j, j-1)
			}
		}
	}

	return items
}

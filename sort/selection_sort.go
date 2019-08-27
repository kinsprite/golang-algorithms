package sort

// SelectionSortInterface SelectionSort Interface
type SelectionSortInterface interface {
	Len() int
	Less(i, j int) bool
	Get(i int) interface{}
	Set(i int, value interface{})
	Clone() SelectionSortInterface
	Pop(i int) SelectionSortInterface
}

// IntSlice int slice
type IntSlice []int

// Len IntSlice's len
func (items *IntSlice) Len() int {
	return len(*items)
}

// Less IntSlice's less compare
func (items *IntSlice) Less(i, j int) bool {
	return (*items)[i] < (*items)[j]
}

// Get IntSlice's Get(i)
func (items *IntSlice) Get(i int) interface{} {
	return (*items)[i]
}

// Set IntSlice's Set(i, value)
func (items *IntSlice) Set(i int, value interface{}) {
	(*items)[i] = value.(int)
}

// Clone IntSlice's Clone()
func (items *IntSlice) Clone() SelectionSortInterface {
	newItems := make(IntSlice, items.Len())
	copy(newItems, *items)
	return &newItems
}

// Pop IntSlice's Pop(i)
func (items *IntSlice) Pop(i int) SelectionSortInterface {
	result := append((*items)[0:i], (*items)[i+1:len(*items)]...)
	return &result
}

func findSmallestIndex(items SelectionSortInterface) int {
	length := items.Len()
	smallestIndex := 0

	for i := 1; i < length; i++ {
		if items.Less(i, smallestIndex) {
			smallestIndex = i
		}
	}

	return smallestIndex
}

// SelectionSort SelectionSort with slice
func SelectionSort(items SelectionSortInterface) SelectionSortInterface {
	tmpArray := items.Clone()
	pushedIndex := 0

	for tmpArray.Len() > 0 {
		smallestIndex := findSmallestIndex(tmpArray)
		smallest := tmpArray.Get(smallestIndex)
		tmpArray = tmpArray.Pop(smallestIndex).(SelectionSortInterface)

		items.Set(pushedIndex, smallest)
		pushedIndex++
	}

	return items
}

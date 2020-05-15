package sort

// IntSlice int slice
type IntSlice []int

// Len IntSlice's len
func (items IntSlice) Len() int {
	return len(items)
}

// Less IntSlice's less compare
func (items IntSlice) Less(i, j int) bool {
	return items[i] < items[j]
}

// Less IntSlice's less compare
func (items IntSlice) Swap(i, j int) {
	items[i], items[j] = items[j], items[i]
}

// Get IntSlice's Get(i)
func (items IntSlice) Get(i int) interface{} {
	return items[i]
}

// Set IntSlice's Set(i, value)
func (items IntSlice) Set(i int, value interface{}) {
	items[i] = value.(int)
}

// Clone IntSlice's Clone()
func (items IntSlice) Clone() IntSlice {
	newItems := make(IntSlice, items.Len())
	copy(newItems, items)
	return newItems
}

// Pop IntSlice's Pop(i)
func (items IntSlice) Pop(i int) IntSlice {
	return append(items[0:i], items[i+1:]...)
}

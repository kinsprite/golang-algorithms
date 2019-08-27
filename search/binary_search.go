package search

// BinarySearchInterface BinarySearch interface
type BinarySearchInterface interface {
	Len() int
	// returns: <0, less than; =0, equal to; >0, greater than
	Comp(i int, value interface{}) int
}

// IntSlice int slice
type IntSlice []int

// Len IntSlice's len
func (items *IntSlice) Len() int {
	return len(*items)
}

// Comp IntSlice's compare
func (items *IntSlice) Comp(i int, value interface{}) int {
	return (*items)[i] - value.(int)
}

// BinarySearch binary search
func BinarySearch(items BinarySearchInterface, value interface{}) (bool, int) {
	low := 0
	high := items.Len() - 1

	for low <= high {
		mid := (low + high) / 2
		comp := items.Comp(mid, value)

		if comp == 0 {
			return true, mid
		}

		// Original:           low               mid             high
		// mid item >  value:  low       highNew mid             highOld
		// mid item < value:   lowOld            mid lowNew      high
		if comp > 0 {
			high = mid - 1
		} else if comp < 0 {
			low = mid + 1
		}
	}

	return false, 0
}

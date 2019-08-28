package sort

// InsertionSort InsertionSort
func InsertionSort(items IntSlice) IntSlice {
	length := len(items)

	for i := 1; i < length; i++ {
		sel := items[i]
		j := i - 1

		for ; j >= 0 && items[j] > sel; j-- {
			// 大于当前值的，都往后挪一个位置。 -- 找插入位置
			items[j+1] = items[j]
		}

		// 插入位置 j+1
		items[j+1] = sel

	}

	return items
}

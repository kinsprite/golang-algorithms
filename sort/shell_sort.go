package sort

// ShellSort ShellSort
func ShellSort(arr IntSlice) IntSlice {
	n := len(arr)

	for gap := n / 2; gap > 0; gap /= 2 {
		for i := gap; i < n; i++ {
			// add a[i] to the elements that have been gap sorted
			// save a[i] in temp and make a hole at position i
			temp := arr[i]

			// shift earlier gap-sorted elements up until the correct
			// location for a[i] is found
			j := i

			for ; j >= gap && arr[j-gap] > temp; j -= gap {
				arr[j] = arr[j-gap]
			}

			//  put temp (the original a[i]) in its correct location
			arr[j] = temp
		}
	}

	return arr
}

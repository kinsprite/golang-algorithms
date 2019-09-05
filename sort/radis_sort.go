package sort

import "math"

// RadixSort RadixSort
func RadixSort(A IntSlice, d int) IntSlice {
	slots := makeRadixSlots(len(A))

	// 从低位往高位开始排序， 确保低位先入容器
	for i := 0; i < d; i++ {
		counts := [10]int{}

		// 放到容器中
		for _, a := range A {
			r := a / int(math.Pow10(i)) % 10
			slots[r][counts[r]] = a
			counts[r]++
		}

		// 从容器中取出
		c := 0

		for j, count := range counts {
			for k := 0; k < count; k++ {
				A[c] = slots[j][k]
				c++
			}
		}
	}

	return A
}

func makeRadixSlots(count int) [10]IntSlice {
	slots := [10]IntSlice{}

	for i := range slots {
		slots[i] = make(IntSlice, count)
	}

	return slots
}

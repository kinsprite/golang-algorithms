package main

import "testing"
import search "algorithms/search"

func BenchmarkBinarySearchParallel(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			items := &search.IntSlice{1, 2, 3, 4, 5}
			value := 3
			search.BinarySearch(items, value)
		}
	})
}

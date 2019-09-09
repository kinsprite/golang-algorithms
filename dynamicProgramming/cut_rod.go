package dynamicProgramming

import (
	"math"
)

// RodPrice RodPrice
var RodPrice = []int{0, 1, 5, 8, 9, 10, 17, 17, 20, 24, 30}

// FoolishCutRod Cut Rod
func FoolishCutRod(price []int, n int) int {
	if n == 0 {
		return 0
	}

	q := math.MinInt32

	for i := 1; i <= n; i++ {
		rest := price[i] + FoolishCutRod(price, n-i)

		if rest > q {
			q = rest
		}
	}

	return q
}

// MemorialCutRod MemorialCutRod
func MemorialCutRod(price []int, n int) int {
	// init memorial
	r := make([]int, n+1)

	for i := range r {
		r[i] = math.MinInt32
	}

	return memorialCutRodAux(price, n, r)
}

func memorialCutRodAux(price []int, n int, r []int) int {
	if r[n] >= 0 {
		return r[n]
	}

	q := math.MinInt32

	if n == 0 {
		q = 0
	} else {
		for i := 1; i <= n; i++ {
			rest := price[i] + memorialCutRodAux(price, n-i, r)

			if rest > q {
				q = rest
			}
		}
	}

	r[n] = q
	return q
}

// BottomUpCutRod BottomUpCutRod
func BottomUpCutRod(price []int, n int) int {
	// init memorial
	r := make([]int, n+1)

	for i := range r {
		r[i] = math.MinInt32
	}

	r[0] = 0

	// price
	for j := 1; j <= n; j++ {
		q := math.MinInt32

		for i := 1; i <= j; i++ {
			rest := price[i] + r[j-i]

			if rest > q {
				q = rest
			}
		}

		r[j] = q
	}

	return r[n]
}

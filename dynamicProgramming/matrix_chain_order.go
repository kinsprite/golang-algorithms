package dynamicProgramming

import (
	"fmt"
	"math"
)

// MatrixChainOrder MatrixChainOrder, such as matrix A1[30x35] A2[35x15], p =[30, 35, 15]
func MatrixChainOrder(p []int) (m [][]int, s [][]int) {
	n := len(p) - 1

	// m[n+1, n+1]
	m = make([][]int, n+1)
	// s[n, n+1]
	s = make([][]int, n)

	for i := 0; i <= n; i++ {
		m[i] = make([]int, n+1)
	}

	for i := 0; i < n; i++ {
		s[i] = make([]int, n+1)
	}

	// init
	for i := 1; i <= n; i++ {
		m[i][i] = 0
	}

	for l := 2; l <= n; l++ { // l is the chain length
		for i := 1; i <= n-l+1; i++ {
			j := i + l - 1
			m[i][j] = math.MaxInt32

			for k := i; k <= j-1; k++ {
				q := m[i][k] + m[k+1][j] + p[i-1]*p[k]*p[j]

				if q < m[i][j] {
					m[i][j] = q
					s[i][j] = k
				}
			}
		}
	}

	return m, s
}

// PrintOptimalParens PrintOptimalParens, i,j >= 1
func PrintOptimalParens(s [][]int, i, j int) string {
	if i == j {
		return fmt.Sprint("A", i)
	}

	result := fmt.Sprint("(")
	result += PrintOptimalParens(s, i, s[i][j])
	result += PrintOptimalParens(s, s[i][j]+1, j)
	result += fmt.Sprint(")")
	return result
}

// DoMatrixChainOrder DoMatrixChainOrder
func DoMatrixChainOrder(p []int) string {
	_, s := MatrixChainOrder(p)
	return PrintOptimalParens(s, 1, len(p)-1)
}

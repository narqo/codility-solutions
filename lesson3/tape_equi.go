package lesson3

import (
	"math"
)

func TapeEquilibriumSolution(A []int) int {
	res := math.MaxInt16
	var sum int
	for _, c := range A {
		sum += c
	}

	head := A[0]
	for i := 1; i < len(A); i++ {
		tail := sum - head
		d := head - tail
		if d < 0 {
			d = -d
		}
		if d < res {
			res = d
		}
		head += A[i]
	}

	return res
}

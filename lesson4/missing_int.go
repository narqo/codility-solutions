package lesson4

import "sort"

// MissingIntegerSolution returns the minimal positive integer (greater than 0) that does not occur in A.
func MissingIntegerSolution(A []int) int {
	if len(A) == 0 {
		return 1
	}
	sort.Ints(A)
	p := A[0]
	if p > 1 {
		return 1
	}
	for i := 1; i < len(A); i++ {
		diff := A[i] - A[i-1]
		if diff <= 0 || diff == 1 {
			continue
		}
		return A[i-1] + 1
	}
	p = A[len(A)-1]
	if p > 0 {
		return p + 1
	}
	return 1
}

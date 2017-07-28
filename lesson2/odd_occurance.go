package lesson2

import "sort"

func OddOccuranceSolution(A []int) int {
	if len(A) == 1 {
		return A[0]
	}
	sort.Ints(A)
	for i := 0; i < len(A)-1; i += 2 {
		if A[i] != A[i+1] {
			return A[i]
		}
	}
	return A[len(A)-1]
}

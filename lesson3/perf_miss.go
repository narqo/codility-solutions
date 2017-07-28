package lesson3

import "sort"

func PermMissingElemSolution(A []int) (r int) {
	sort.Ints(A)
	r = len(A) + 1
	for i := 0; i < len(A); i++ {
		if A[i] != i+1 {
			r = i + 1
			break
		}
	}
	return
}

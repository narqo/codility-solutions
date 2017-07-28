package main

// Solution finds an index in an array such that its prefix sum equals its suffix sum.
func Solution(A []int) (p int) {
	if len(A) == 0 {
		return -1
	}
	var sum int
	for _, c := range A {
		sum += c
	}

	var left int
	for p = 1; p < len(A); p++ {
		left += A[p-1]
		if sum-left-A[p] == left {
			return p
		}
	}

	if sum-A[0] == 0 || sum-A[len(A)-1] == 0 {
		return 0
	}

	return -1
}

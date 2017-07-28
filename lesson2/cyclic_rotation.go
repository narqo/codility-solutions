package lesson2

func CyclicRotationSolution(A []int, K int) []int {
	for n := 0; n < K; n++ {
		for i := len(A) - 1; i > 0; i-- {
			if i > 0 {
				A[i], A[i-1] = A[i-1], A[i]
			} else {
				A[i], A[len(A)-1] = A[0], A[i]
			}
		}
	}
	return A
}

/*
Your test case:  [[3, 8, 9, 7, 6], 42]
Returned value: [7, 6, 3, 8, 9]

[7, 6, 1, 2, 3, 4, 5]

*/

/*
func CyclicRotationSolution(A []int, K int) []int {
	if K > len(A) {
		K %= len(A)
	}
	for i := len(A) - 1; i > 0; i-- {
		if i-K >= 0 {
			A[i], A[i-K] = A[i-K], A[i]
		} else {
			n := -i + K
			if n >= len(A) {
				n %= len(A)
			}
			if n > i {
				n %= i
			}
			A[i], A[n] = A[n], A[i]
		}
	}
	return A
}
*/

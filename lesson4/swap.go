package lesson4

/*
You are given an integer m (1 <= m <= 1 000 000) and two non-empty,
zero-indexed arrays A and B of n integers
	a0,a1,...,an−1 and b0,b1,...,bn−1
respectively (0 <= ai,bi <= m).

The goal is to check whether there is a swap operation which can
be performed on these arrays in such a way that the sum of elements
in array A equals the sum of elements in array B after the swap.
By swap operation we mean picking one element from array A and
one element from array B and exchanging them.
*/

func SwapSolution(A, B []int, m int) bool {
	aSum := SumInts(A)
	bSum := SumInts(B)
	diff := bSum - aSum
	if diff%2 != 0 {
		return false
	}
	diff /= 2
	aCounts := Counting(A, m)
	for _, c := range B {
		// d is a theoretical element of A, that could be swapped with c
		// to get an equal sums.
		d := diff - c
		if d > 0 && d < m && aCounts[d] > 0 {
			return true
		}
	}
	return false
}

func SumInts(A []int) (s int) {
	for _, c := range A {
		s += c
	}
	return
}

// Counting returns an array of couters.
// Each number in counted in the array by using an index that
// corresponds to the value of the given number.
func Counting(A []int, m int) []int {
	counts := make([]int, m)
	for _, c := range A {
		if c >= m {
			panic("out of range")
		}
		counts[c]++
	}
	return counts
}

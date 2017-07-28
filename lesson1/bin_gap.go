// Package lesson1 provides solution for "binary gap" challenge.
// A binary gap within a positive integer N is any maximal sequence
// of consecutive zeros that is surrounded by ones at both ends
// in the binary representation of N.
package lesson1

import "fmt"

// BinGapSolution returns the length the longest binary gap in N.
func BinGapSolution(N int) int {
	var (
		gapping bool
		gap     = 0
		g       = 0
	)
	for mask := 1; mask < N; mask <<= 1 {
		if N&mask != 0 {
			if gapping && g > gap {
				gap = g
			}
			gapping = true
			g = 0
			continue
		}
		if gapping {
			g++
		}
	}
	return gap
}

func BinGapSolution1(N int) int {
	s := fmt.Sprintf("%b", N)
	var (
		gapping bool
		gap     = 0
		g       = 0
	)
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c == '1' {
			if gapping && g > gap {
				gap = g
			}
			gapping = true
			g = 0
			continue
		}
		if gapping {
			g++
		}
	}
	return gap
}

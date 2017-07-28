package lesson3

import (
	"testing"
)

func TestTapeEquilibriumSolution(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		out  int
	}{
		{"example", []int{3, 1, 2, 4, 3}, 1},
		{"double", []int{1000, -1000}, 2000},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := TapeEquilibriumSolution(tc.in)
			if got != tc.out {
				t.Errorf("got %v, want %v", got, tc.out)
			}
		})
	}
}

package lesson4

import (
	"testing"
)

func TestMissingIntegerSolution(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		out  int
	}{
		{"example", []int{1, 3, 6, 4, 1, 2}, 5},
		{"single 1", []int{1}, 2},
		{"single 2", []int{2}, 1},
		{"double 1", []int{1, 2}, 3},
		{"double 2", []int{4, 5}, 1},
		{"simple", []int{4, 5, 6, 2}, 1},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := MissingIntegerSolution(tc.in)
			if got != tc.out {
				t.Errorf("got %v, want %v", got, tc.out)
			}
		})
	}
}

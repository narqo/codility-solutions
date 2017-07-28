package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		out  int
	}{
		{"example", []int{-1, 3, -4, 5, 1, -6, 2, 1}, 1},
		{"single1", []int{0}, 0},
		{"single2", []int{100}, 0},
		{"extreme large numbers", []int{2147483648}, 0},
		{"extreme negative numbers", []int{-2147483648}, 0},
		{"one large", []int{500, 1, -2, -1, 2}, 0},
		{"combinations of two 1", []int{-1, 0}, 0},
		{"combinations of two 2", []int{1, 0}, 0},
		{"combinations of three 1", []int{-1, -1, 1}, 0},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, want := Solution(tc.in), tc.out
			if got != want {
				t.Errorf("expected %v, got %v", want, got)
			}
		})
	}
}

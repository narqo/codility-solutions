package lesson4

import (
	"testing"
)

func TestSwapSolution(t *testing.T) {
	tests := []struct {
		name string
		inA  []int
		inB  []int
		m    int
		out  bool
	}{
		{
			name: "example",
			m:    10000,
			out:  true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := SwapSolution(tc.inA, tc.inB, tc.m)
			if got != tc.out {
				t.Errorf("got %v, want %v", got, tc.out)
			}
		})
	}
}

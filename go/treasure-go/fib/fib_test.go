package main

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	t.Parallel()
	cases := []struct {
		in, expected int
	}{
		{0, 0},
		{1, 1},
		{10, 55},
	}
	for i, tc := range cases {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			if got := fib(tc.in); got != tc.expected {
				t.Errorf("want %d, but got %d\n", tc.expected, got)
			}
		})
	}
}

package main

import "testing"

func TestSimpleArraySum(t *testing.T) {
	cases := []struct {
		size, result int
		array        []int
	}{
		{0, 0, []int{}},
		{1, 2, []int{2}},
		{5, 15, []int{1, 2, 3, 4, 5}},
	}
	for _, c := range cases {
		got := simpleArraySum(c.size, c.array)
		if got != c.result {
			t.Errorf("simpleArraySim(%d, %d) == %d, want %d", c.size, c.array, got, c.result)
		}
	}
}

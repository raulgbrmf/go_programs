package solvemefirst

import "testing"

func TestsolveMeFirst(t *testing.T) {
	cases := []struct {
		num1, num2, want uint32
	}{
		{1, 1, 2},
		{1, 0, 1},
		{0, 0, 0},
	}
	for _, c := range cases {
		got := solveMeFirst(c.num1, c.num2)
		if got != c.want {
			t.Errorf("solveMeFirst(%d, %d) == %d, want %d", c.num1, c.num2, got, c.want)
		}
  }
}

package main

import "testing"

func TestExtractDiagonals(t *testing.T){
  cases := []struct {
    size int
    matrix [][]int
    result float64
    }{
      {3, [][]int{{11, 2, 4}, {4, 5, 6}, {10, 8, -12}}, 15},
      {4, [][]int{{-1, 1, -7, -8}, {-10, -8, -5, -2}, {0, 9, 7, -1}, {4, 4, -2, 1}}, 1},
  }
  for _, c := range cases {
    got := extractDiagonals(c.size, c.matrix)
    if got != c.result {
      t.Errorf("simpleArraySim(%d, %d) == %d, want %d", c.size, c.matrix, got, c.result)
    }
  }
}

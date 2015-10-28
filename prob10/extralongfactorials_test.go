package main

import (
	"math/big"
	"testing"
)

func TestExtralongfactorials(t *testing.T) {
	cases := []struct {
		N      int64
		result *big.Int
	}{
		{0, big.NewInt(1)},
		{1, big.NewInt(1)},
		{2, big.NewInt(2)},
		{3, big.NewInt(6)},
		{20, big.NewInt(2432902008176640000)},
	}
	for _, c := range cases {
		got := fatorial(c.N)
		if got.Cmp(c.result) != 0 {
			t.Errorf("extralongfactorials(%d) == %d, want %d", c.N, got, c.result)
		}
	}
}

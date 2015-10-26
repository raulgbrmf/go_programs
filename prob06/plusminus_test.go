package main

import "testing"

import "bytes"

func TestPlusMinus(t *testing.T) {
	cases := []struct {
		size   int
		array  []int
		types  []float64
		result string
		// tenho que adicionar uma string de resultado
	}{
		{3, []int{0, 0, 0}, []float64{1, 0, 0}, "001.3333333333333333"}, //esse.333333 eh uma anomalia que nao entendi o porque
		{3, []int{1, 2, 3}, []float64{1, 0, 0}, "100.3333333333333333"},
		{3, []int{-1, -2, -3}, []float64{1, 0, 0}, "010.3333333333333333"},
	}
	for _, c := range cases {
		var b bytes.Buffer
		plusminus(c.size, c.array, c.types, &b)
		// fmt.Println("B")
		// fmt.Print(b.String())
		// fmt.Println("\n B fim")
		if b.String() != c.result { // e verificar se o buffer lido vai ser igual à string de resultado
			t.Errorf("plusminus(%d, %d, %f, %s) == %s, want %s", c.size, c.array, c.types, &b, b.String(), c.result)
		}
	}
}

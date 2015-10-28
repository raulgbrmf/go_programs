package main

import "testing"

func TestLibraryFine(t *testing.T) {
	cases := []struct {
		yearRecv, monthRecv, dayRecv, yearShould, monthShould, dayShould, hackos int
	}{
		{2015, 5, 30, 2015, 5, 2, 420},
		{2015, 6, 9, 2015, 6, 6, 45},
		{2004, 8, 31, 2004, 1, 20, 3500},
		{2015, 1, 1, 2014, 12, 31, 10000},
	}
	for _, c := range cases {
		got := libraryfine(c.yearRecv, c.monthRecv, c.dayRecv, c.yearShould, c.monthShould, c.dayShould)
		if got != c.hackos {
			t.Errorf("libraryfine(%d, %d, %d, %d, %d, %d) == %d, want %d", c.yearRecv, c.monthRecv, c.dayRecv, c.yearShould, c.monthShould, c.dayShould, got, c.hackos)
		}
	}
}

// Tá estranho, por que meu teste está retornando um valor diferente quando rodo a funcao

package main

import (
	"reflect"
	"testing"
)

func TestStairCase(t *testing.T) {
	cases := []struct {
		height int
		result []string
	}{
		{0, []string{}},
		{1, []string{"#", "\n"}},
		{2, []string{" ", "#", "\n", "#", "#", "\n"}},
	}
	for _, c := range cases {
		got := staircase(c.height)
		if !reflect.DeepEqual(got, c.result) {
			t.Errorf("simpleArraySim(%d) == %s, want %s", c.height, got, c.result)
		}
	}
}

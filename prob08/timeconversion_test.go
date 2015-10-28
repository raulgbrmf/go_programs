package main

import (
	"reflect"
	"testing"
)

func TestTimeconversion(t *testing.T) {
	cases := []struct {
		palavra string
		result  []string
	}{
		{"10:10:10PM", []string{"22", ":", "10", ":", "10"}},
		{"10:10:10AM", []string{"10", ":", "10", ":", "10"}},
		{"12:10:10AM", []string{"00", ":", "10", ":", "10"}},
		{"12:10:10PM", []string{"12", ":", "10", ":", "10"}},
	}
	for _, c := range cases {
		got := timeconversion(c.palavra)
		if !reflect.DeepEqual(got, c.result) {
			t.Errorf("timeconversion(%s) == %s, want %s", c.palavra, got, c.result)
		}
	}
}

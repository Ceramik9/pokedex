package main

import (
	"testing"
	"reflect"
)

func TestCleanInput(t *testing.T) {
	type test struct {
		input string
		want []string
	}

	tests := []test{
		{input: "Charmander Bulbasaur PIKACHU", want: []string{"charmander", "bulbasaur","pikachu"}},
		{input: "", want: []string{}},
		{input: " ", want: []string{}},
		{input: "Pikachu", want: []string{"pikachu"}},
		{input: "     foo bar ", want: []string{"foo", "bar"}},
	}

	for i, tc := range tests {
		got := cleanInput(tc.input)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("Test: %d failed. Expected: %v, got: %v", i+1, tc.want, got)
		}
	}
}

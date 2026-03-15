package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello   world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "charmander bulbasaur squirtle",
			expected: []string{"charmander", "bulbasaur", "squirtle"},
		},
		{
			input:    "  PiKaChU  ",
			expected: []string{"pikachu"},
		},
		{
			input:    "   ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		for i := range actual {
			if actual[i] != c.expected[i] {
				t.Errorf("expected %s, got %s", c.expected, actual)
			}
		}
	}
}

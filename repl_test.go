package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    " NEveR  goNna   GiVE  yOu UP",
			expected: []string{"never", "gonna", "give", "you", "up"},
		},
		{
			input:    "Bulbasaur Charmander Squirtle",
			expected: []string{"bulbasaur", "charmander", "squirtle"},
		},
		{
			input:    "   ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("expected: %v\nactual: %v\nno match",
				actual, c.expected)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("expected: %v\nactual: %v\nno match",
					actual, c.expected)
			}
		}
	}
}

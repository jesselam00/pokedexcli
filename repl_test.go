package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("cleanInput(%q) == %v, expected %v", c.input, actual, c.expected)
			t.Fail()
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%q)[%d] == %q, expected %q", c.input, i, word, expectedWord)
				t.Fail()
				break
			}
		}
	}
}

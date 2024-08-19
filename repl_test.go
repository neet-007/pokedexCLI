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
			input: "hello world",
			expected: []string{
				"hello",
				"world",
			},
		},
		{
			input: "HEllO wOrLd",
			expected: []string{
				"hello",
				"world",
			},
		},
	}

	for _, cs := range cases {
		acutal := cleanInput(cs.input)
		expected := cs.expected

		if len(acutal) != len(expected) {
			t.Errorf("failed test case length not matchin %s vs %s", acutal, expected)
			continue
		}

		for i := range acutal {
			if acutal[i] != expected[i] {
				t.Errorf("failed test case words not matchin %s vs %s", acutal[i], expected[i])
				continue
			}
		}
	}
}

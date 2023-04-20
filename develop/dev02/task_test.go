package main

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "a4bc2d5e",
			expected: "aaaabccddddde",
		},
		{
			input:    "abcd",
			expected: "abcd",
		},
		{
			input:    "45",
			expected: "",
		},
		{
			input:    "",
			expected: "",
		},
		{
			input:    "qwe\\4\\5",
			expected: "qwe45",
		},
		{
			input:    "qwe\\45",
			expected: "qwe44444",
		},
		{
			input:    "qwe\\\\5",
			expected: "qwe\\\\\\\\\\",
		},
		{
			input:    "a0b0c0",
			expected: "",
		},
		{
			input:    "a1b1c1",
			expected: "abc",
		},
		{
			input:    "a10",
			expected: "aaaaaaaaaa",
		},
		{
			input:    "\\\\3",
			expected: "\\\\\\",
		},
		{
			input:    "\\\\",
			expected: "\\",
		},
	}
	for i := range testCases {
		tc := testCases[i]
		if output, _ := Unpack(tc.input); output != tc.expected {
			t.Errorf("Unpack doesn't work correctly, %v not equal %v", output, tc.expected)
		}
	}

}

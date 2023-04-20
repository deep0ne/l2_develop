package main

import (
	"testing"
)

func TestCut(t *testing.T) {
	testCases := []struct {
		pattern   string
		fields    fieldFlags
		delimeter string
		separated bool
		expected  string
	}{
		{
			pattern:   "a,b,c,d,e",
			fields:    fieldFlags{1, 3},
			delimeter: ",",
			separated: false,
			expected:  "a,c",
		},
		{
			pattern:   "abcde",
			fields:    fieldFlags{1, 3},
			delimeter: ",",
			separated: true,
			expected:  "",
		},
		{
			pattern:   "abcde",
			fields:    fieldFlags{1, 3},
			delimeter: ",",
			separated: false,
			expected:  "abcde",
		},
		{
			pattern:   "hey:siri:tell:me:a:story",
			fields:    fieldFlags{2, 4, 6},
			delimeter: ":",
			separated: false,
			expected:  "siri:me:story",
		},
	}

	for i := range testCases {
		tc := testCases[i]
		if output, _ := CutForTests(tc.pattern, tc.fields, tc.delimeter, tc.separated); output != tc.expected {
			t.Errorf("Cut does not work as expected, %v not equal %v", output, tc.expected)
		}
	}

}

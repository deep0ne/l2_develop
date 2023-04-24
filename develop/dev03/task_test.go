package main

import (
	"bufio"
	"dev03/sort"
	"os"
	"testing"
)

func TestGrep(t *testing.T) {
	testCases := []struct {
		input    string
		expected []string
		params   sort.SortParams
	}{
		{
			input: `abc def ghi
ghi def abc
def abc ghi
def ghi abc`,
			expected: []string{"abc def ghi", "def abc ghi", "def ghi abc", "ghi def abc"},
			params:   sort.SortParams{},
		},
		{
			input: `1 2 3
4 5 6
7 8 9
10 11 12
13 14 15`,
			expected: []string{"1 2 3", "10 11 12", "13 14 15", "4 5 6", "7 8 9"},
			params:   sort.SortParams{},
		},
		{
			input: `1 2 3
4 5 6
7 8 9
10 11 12
13 14 15`,
			expected: []string{"1 2 3", "4 5 6", "7 8 9", "10 11 12", "13 14 15"},
			params: sort.SortParams{
				Numerical: true,
			},
		},
		{
			input: `1 2 3
4 5 6
7 8 9
10 11 12
13 14 15`,
			expected: []string{"13 14 15", "10 11 12", "7 8 9", "4 5 6", "1 2 3"},
			params: sort.SortParams{
				Numerical: true,
				Reverse:   true,
			},
		},
		{
			input: `1 5 3
4 0 6
1 2 4
10 1 5
666 3 0`,
			expected: []string{"4 0 6", "10 1 5", "1 2 4", "666 3 0", "1 5 3"},
			params: sort.SortParams{
				Numerical: true,
				Column:    2,
			},
		},
		{
			input: `1 1 1
1 2 3
1 1 1
3 2 1
4 0 5
6 4 2
2 4 6
`,
			expected: []string{"2 4 6", "4 0 5", "1 2 3", "6 4 2", "3 2 1", "1 1 1"},
			params: sort.SortParams{
				Numerical: true,
				Column:    3,
				Reverse:   true,
				Unique:    true,
				Sorted:    true,
			},
		},
		{
			input: `10M
2K
1G
100
20
`,
			expected: []string{"20", "100", "2K", "10M", "1G"},
			params: sort.SortParams{
				HumanNumeric: true,
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		input, _ := os.Create("input.txt")
		input.WriteString(tc.input)

		fileInput, _ := os.Open("input.txt")
		sort.Sort(fileInput, tc.params)

		fileOutput, _ := os.Open("sorted_input.txt")
		scanner := bufio.NewScanner(fileOutput)
		for _, expectedString := range tc.expected {
			scanner.Scan()
			text := scanner.Text()
			if expectedString != text {
				t.Errorf("Sort worked wrong. Expected %s. Got %s.", expectedString, text)
			}
		}
	}

}

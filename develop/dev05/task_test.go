package main

import (
	"bufio"
	"dev05/grep"
	"os"
	"regexp"
	"testing"
)

func TestGrep(t *testing.T) {
	testCases := []struct {
		options  grep.SearchOptions
		regex    *regexp.Regexp
		input    string
		expected []string
	}{
		{
			options: grep.SearchOptions{
				AfterFlag: 2,
			},
			regex: regexp.MustCompile(grep.FormRegex("apple", grep.SearchOptions{AfterFlag: 2})),
			input: `apple for far free appledog bro
is that applejuice on the shelf yours?`,
			expected: []string{"apple for far", "applejuice on the"},
		},
		{
			options: grep.SearchOptions{
				ContextFlag: 2,
				LineNumFlag: true,
				FixedFlag:   true,
			},
			regex: regexp.MustCompile(grep.FormRegex("string", grep.SearchOptions{
				ContextFlag: 2,
				LineNumFlag: true,
				FixedFlag:   true})),

			input: `this is string number one
then comes stringzzz with subpattern
then string with only one word before it`,

			expected: []string{"1: this is string number one"},
		},

		{
			options: grep.SearchOptions{
				AfterFlag:   2,
				BeforeFlag:  1,
				LineNumFlag: true,
				FixedFlag:   false,
			},
			regex: regexp.MustCompile(grep.FormRegex("string", grep.SearchOptions{
				AfterFlag:   2,
				BeforeFlag:  1,
				LineNumFlag: true,
				FixedFlag:   false,
			})),

			input: `this is string number one
then comes stringzzz with subpattern
then string with only one word before it`,

			expected: []string{"1: is string number one", "2: comes stringzzz with subpattern", "3: then string with only"},
		},
		{
			options: grep.SearchOptions{
				AfterFlag:   2,
				BeforeFlag:  1,
				LineNumFlag: true,
				FixedFlag:   false,
				CountFlag:   1,
			},
			regex: regexp.MustCompile(grep.FormRegex("string", grep.SearchOptions{
				AfterFlag:   2,
				BeforeFlag:  1,
				LineNumFlag: true,
				FixedFlag:   false,
				CountFlag:   1,
			})),

			input: `this is string number one
then comes stringzzz with subpattern
then string with only one word before it`,

			expected: []string{"1: is string number one"},
		},

		{
			options: grep.SearchOptions{
				AfterFlag:   2,
				BeforeFlag:  1,
				LineNumFlag: true,
				FixedFlag:   false,
				InvertFlag:  true,
			},
			regex: regexp.MustCompile(grep.FormRegex("string", grep.SearchOptions{
				AfterFlag:   2,
				BeforeFlag:  1,
				LineNumFlag: true,
				FixedFlag:   false,
				CountFlag:   1,
				InvertFlag:  true,
			})),

			input: `this is string number one
then comes stringzzz with subpattern
then string with only one word before it`,

			expected: []string{"1: this ", "2: then ", "3: one word before it "},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		input, _ := os.Create("input.txt")
		input.WriteString(tc.input)

		fileInput, _ := os.Open("input.txt")
		grep.Grep(tc.regex, fileInput, tc.options, true)

		fileOutput, _ := os.Open("output.txt")
		scanner := bufio.NewScanner(fileOutput)
		for _, expectedString := range tc.expected {
			scanner.Scan()
			text := scanner.Text()
			if expectedString != text {
				t.Errorf("Grep worked wrong. Expected %s. Got %s.", expectedString, text)
			}
		}
	}

}

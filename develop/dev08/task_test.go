package main

import (
	"log"
	"os"
	"strings"
	"testing"
)

func TestOS(t *testing.T) {
	testCases := []struct {
		command  string
		expected string
		checker  func() string
	}{
		{
			command: "cd ..",
			checker: func() string {
				dir, _ := os.Getwd()
				current := strings.Split(dir, "/")
				return current[len(current)-1]
			},
			expected: "develop",
		},
		{
			command: "cd dev08",
			checker: func() string {
				dir, _ := os.Getwd()
				current := strings.Split(dir, "/")
				return current[len(current)-1]
			},
			expected: "dev08",
		},
		{
			command: "cd ../..",
			checker: func() string {
				dir, _ := os.Getwd()
				log.Println(dir)
				current := strings.Split(dir, "/")
				return current[len(current)-1]
			},
			expected: "l2_develop",
		},
	}

	for i := range testCases {
		tc := testCases[i]
		args := strings.Split(tc.command, " ")
		err := Execute(args)
		if err != nil {
			t.Fatal(err)
		}
		if res := tc.checker(); res != tc.expected {
			t.Errorf("something went wrong, %v not equal %v", res, tc.expected)
		}
	}

}

package grep

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type SearchOptions struct {
	AfterFlag, BeforeFlag, ContextFlag, CountFlag  int
	IgnoreFlag, InvertFlag, FixedFlag, LineNumFlag bool
}

func FormRegex(pattern string, options SearchOptions) string {
	var regex string

	if options.IgnoreFlag {
		regex = regexp.QuoteMeta(strings.ToLower(pattern))
	} else {
		regex = regexp.QuoteMeta(pattern)
	}

	if options.FixedFlag {
		regex = "\\b" + regex + "\\b"
	} else {
		regex = "\\w*" + regex + "\\w*"
	}

	if options.BeforeFlag > 0 {
		regex = "(?:\\w+\\W+){" + strconv.Itoa(options.BeforeFlag) + "}" + regex
	}

	if options.AfterFlag > 0 {
		regex = regex + "(?:\\W+\\w+){" + strconv.Itoa(options.AfterFlag) + "}"
	}

	if options.ContextFlag > 0 {
		regex = "(?:\\w+\\W+){" + strconv.Itoa(options.ContextFlag) + "}" + regex + "(?:\\W+\\w+){" + strconv.Itoa(options.ContextFlag) + "}"
	}

	return regex
}

func Grep(regex *regexp.Regexp, file *os.File, options SearchOptions, testmode bool) error {
	var (
		iterate    int
		outputFile *os.File
	)
	if testmode {
		outputFile, _ = os.Create("output.txt")
		defer outputFile.Close()
	}

	scanner := bufio.NewScanner(file)
	findings := make([]string, 0, options.CountFlag)

	for scanner.Scan() {
		text := scanner.Text()
		if options.IgnoreFlag {
			text = strings.ToLower(text)
		}
		finds := regex.FindAllString(text, -1)

		if options.InvertFlag {
			notFound := make(map[string]struct{})
			for _, f := range finds {
				for _, word := range strings.Split(f, " ") {
					notFound[word] = struct{}{}
				}
			}
			notFoundString := ""
			for _, word := range strings.Split(text, " ") {
				if _, ok := notFound[word]; !ok {
					notFoundString += word + " "
				}
			}
			findings = append(findings, notFoundString)
		} else {
			findings = append(findings, strings.Join(finds, " "))
		}
	}

	if options.CountFlag > 0 {
		iterate = options.CountFlag
	} else {
		iterate = len(findings)
	}

	for line := 0; line < iterate; line++ {
		if testmode {
			if options.LineNumFlag {
				s := strconv.Itoa(line+1) + ": "
				outputFile.WriteString(s)
			}
			outputFile.WriteString(findings[line])
			outputFile.WriteString("\n")
		} else {
			if options.LineNumFlag {
				fmt.Println(line+1, ": ")
			}
			fmt.Println(findings[line])
		}
	}
	return nil
}

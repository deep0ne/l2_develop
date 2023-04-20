package main

import (
	"dev05/grep"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {

	var (
		AfterFlag, BeforeFlag, ContextFlag, CountFlag  int
		IgnoreFlag, InvertFlag, FixedFlag, LineNumFlag bool
	)

	flag.IntVar(&AfterFlag, "A", 0, "Usage: \"after\" - печатать N строк после совпадения")
	flag.IntVar(&BeforeFlag, "B", 0, "Usage: \"before\" - печатать N строк до совпадения")
	flag.IntVar(&ContextFlag, "C", 0, "Usage: - печатать N строк до и после совпадения")
	flag.IntVar(&CountFlag, "c", 0, "Usage: \"after\" - печатать N строк после совпадения")
	flag.BoolVar(&IgnoreFlag, "i", false, "Usage: \"ignore-case\" - игнорировать регистр")
	flag.BoolVar(&InvertFlag, "v", false, "Usage: \"invert\" - Исключает вместо совпадения")
	flag.BoolVar(&FixedFlag, "F", false, "Usage: \"fixed\" - искать только полное совпадение, а не паттерн")
	flag.BoolVar(&LineNumFlag, "n", false, "Usage: \"line num\" - печатать номер строки")

	flag.Parse()

	if ContextFlag != 0 && (AfterFlag != 0 || BeforeFlag != 0) {
		fmt.Fprintf(os.Stderr, "Флаг \"C\" может быть использован, только если флаги \"A\" и \"B\" не были введены.")
		os.Exit(1)
	}

	search := os.Args[len(os.Args)-1]
	searchOptions := grep.SearchOptions{
		AfterFlag:   AfterFlag,
		BeforeFlag:  BeforeFlag,
		ContextFlag: ContextFlag,
		CountFlag:   CountFlag,
		IgnoreFlag:  IgnoreFlag,
		InvertFlag:  InvertFlag,
		FixedFlag:   FixedFlag,
		LineNumFlag: LineNumFlag,
	}

	regex := grep.FormRegex(search, searchOptions)
	re := regexp.MustCompile(regex)
	fmt.Println(re)

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	err = grep.Grep(re, file, searchOptions, true)

}

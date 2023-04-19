package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type fieldFlags []int

func (i *fieldFlags) String() string {
	return "my string representation"
}

func (i *fieldFlags) Set(value string) error {
	values := strings.Split(value, ",")
	for _, val := range values {
		intVal, err := strconv.Atoi(val)
		if err != nil {
			return err
		}
		*i = append(*i, intVal)
	}
	return nil
}

func Cut(fields fieldFlags, delimeter string, separated bool) error {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			if err == io.EOF {
				break
			}
		}
		line := strings.Split(scanner.Text(), delimeter)
		if separated && len(line) == 1 {
			continue
		}

		for _, f := range fields {
			if f > len(line) || f < 1 {
				return errors.New("You chose wrong number for fields. See usage.")
			}
			fmt.Print(line[f-1], " ")
		}
		fmt.Println()
	}
	return nil
}

func main() {
	var (
		Fields        fieldFlags
		DelimeterFlag string
		SeparatedFlag bool
	)

	flag.Var(&Fields, "f", "Usage: choose columns. Number (int) of column must be >= 1 and <= total number of columns")
	flag.StringVar(&DelimeterFlag, "d", "\t", "Usage: choose a delimeter")
	flag.BoolVar(&SeparatedFlag, "s", false, "Usage: type \"-s=true/false\". When true, lines without delimeter won't be in the output.")
	flag.Parse()
	if Fields == nil {
		log.Fatal("You must provide number of fields.")
	}

	err := Cut(Fields, DelimeterFlag, SeparatedFlag)
	if err != nil {
		log.Fatal(err)
	}
}

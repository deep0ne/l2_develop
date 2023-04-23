/*

Реализовать утилиту аналог консольной команды cut (man cut). Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB) на колонки и выводить запрошенные.

Реализовать поддержку утилитой следующих ключей:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

*/

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

// реализация методов интерфейса для флагов
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
		matches := make([]string, 0, len(scanner.Text()))
		// выходим по CTRL+D
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
				if !separated {
					continue
				} else {
					return errors.New("You chose wrong number for fields. See usage.")
				}
			}
			matches = append(matches, line[f-1])
		}
		fmt.Println(strings.Join(matches, delimeter))
	}
	return nil
}

func CutForTests(pattern string, fields fieldFlags, delimeter string, separated bool) (string, error) {

	line := strings.Split(pattern, delimeter)
	if len(line) == 1 {
		if separated {
			return "", nil
		}
		return pattern, nil
	}
	words := make([]string, 0)
	for _, f := range fields {
		if f > len(line) || f < 1 {
			return "", errors.New("You chose wrong number for fields. See usage.")
		}
		words = append(words, line[f-1])
	}
	return strings.Join(words, delimeter), nil
}

func main() {
	var (
		Fields        fieldFlags
		DelimeterFlag string
		SeparatedFlag bool
	)

	flag.Var(&Fields, "f", "Usage: choose columns. Number (int) of column must be >= 1 and <= total number of columns")
	flag.StringVar(&DelimeterFlag, "d", "\t", "Usage: choose a delimeter")
	flag.BoolVar(&SeparatedFlag, "s", true, "Usage: type \"-s=true/false\". When true, lines without delimeter won't be in the output.")
	flag.Parse()
	if Fields == nil {
		log.Fatal("You must provide number of fields.")
	}

	err := Cut(Fields, DelimeterFlag, SeparatedFlag)
	if err != nil {
		log.Fatal(err)
	}
}

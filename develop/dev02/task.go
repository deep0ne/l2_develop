package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Unpack(word string) (string, error) {
	if len(word) > 0 && unicode.IsDigit(rune(word[0])) {
		return "", errors.New("Неверная строка")
	}

	var (
		sb   strings.Builder
		idx  int
		prev rune
	)

	for len(word) > 0 {
		prev = rune(word[idx])
		word = word[idx+1:]
		numIndex := strings.IndexFunc(word, func(r rune) bool { return !unicode.IsDigit(r) })
		if numIndex == -1 {
			numIndex = len(word)
		}
		if idx != numIndex-1 {
			sb.WriteRune(prev)
		} else {
			num, _ := strconv.Atoi(word[:numIndex])
			sb.WriteString(strings.Repeat(string(prev), num))
			word = word[numIndex:]
		}
	}
	return sb.String(), nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	result, err := Unpack(scanner.Text())
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}

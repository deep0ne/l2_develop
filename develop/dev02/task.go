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
		if prev == '\\' {
			if len(word) == 0 {
				return "", errors.New("Неверная строка")
			}
			next := rune(word[idx])
			word = word[idx+1:]
			if next == '\\' && len(word) != 0 {
				if unicode.IsDigit(rune(word[idx])) {
					repeat, _ := strconv.Atoi(string(word[idx]))
					sb.WriteString(strings.Repeat(string(next), repeat))
					word = word[idx+1:]
				} else {
					sb.WriteRune('\\')
				}
			} else {
				if idx >= len(word) {
					sb.WriteRune(next)
					break
				}
				if unicode.IsDigit(rune(word[idx])) {
					repeat, _ := strconv.Atoi(string(word[idx]))
					sb.WriteString(strings.Repeat(string(next), repeat))
					word = word[idx+1:]
				} else {
					sb.WriteRune(next)
				}
			}
		} else {
			numIndex := strings.IndexFunc(word, func(r rune) bool { return !unicode.IsDigit(r) })
			if numIndex == -1 {
				numIndex = len(word)
			}
			if numIndex == 0 {
				sb.WriteRune(prev)
			} else {
				num, _ := strconv.Atoi(word[:numIndex])
				sb.WriteString(strings.Repeat(string(prev), num))
				word = word[numIndex:]
			}
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

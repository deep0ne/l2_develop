package main

/*
Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:
"a4bc2d5e" => "aaaabccddddde"
"abcd" => "abcd"
"45" => "" (некорректная строка)
"" => ""

Дополнительно
Реализовать поддержку escape-последовательностей.
Например:
qwe\4\5 => qwe45 (*)
qwe\45 => qwe44444 (*)
qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка, функция должна возвращать ошибку. Написать unit-тесты.

*/

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
		if prev == '\\' { // обработка escape последовательностей
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
			numIndex := strings.IndexFunc(word, func(r rune) bool { return !unicode.IsDigit(r) }) // ищем индекс первого НЕ числа
			if numIndex == -1 {
				numIndex = len(word) // значит строка заканчивается числом
			}
			if numIndex == 0 {
				sb.WriteRune(prev)
			} else {
				num, _ := strconv.Atoi(word[:numIndex]) // всё, что было до буквы - число. конвертируем
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

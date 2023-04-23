package main

import (
	"bufio"
	"errors"
	"flag"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type SortParams struct {
	column    int
	numerical bool
	reverse   bool
	unique    bool
}

type Order struct {
	columnWord string
	index      int
}

func Sort(file *os.File, params SortParams) error {
	filename := "sorted_" + file.Name()
	sortedFile, err := os.Create(filename)
	if err != nil {
		return errors.New("Не удалось создать файл с отсортированными данными")
	}

	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if params.column != 0 {
		orderMap := make([]Order, 0)
		for idx, line := range lines {
			splitted := strings.Split(line, " ")
			if params.column > len(splitted) {
				return errors.New("В строке нет такого количества колонок")
			}
			order := Order{
				columnWord: splitted[params.column-1],
				index:      idx,
			}
			orderMap = append(orderMap, order)
		}
		sort.Slice(orderMap, func(i, j int) bool {
			if params.numerical {
				iNum, err := strconv.Atoi(orderMap[i].columnWord)
				if err != nil {
					log.Fatal("В данных для сортировки находится не число")
				}
				jNum, err := strconv.Atoi(orderMap[j].columnWord)
				if err != nil {
					log.Fatal("В данных для сортировки находится не число")
				}
				return iNum < jNum
			}
			return orderMap[i].columnWord < orderMap[j].columnWord
		})
		newLines := make([]string, 0)
		for _, order := range orderMap {
			newLines = append(newLines, lines[order.index])
		}
		lines = nil
		lines = append(lines, newLines...)
	} else {
		sort.Slice(lines, func(i, j int) bool {
			if params.numerical {
				// в таком случае всегда будем сортировать по первому числу
				iNum, err := strconv.Atoi(strings.Split(lines[i], " ")[0])
				if err != nil {
					log.Fatal("В данных для сортировки находится не число")
				}
				jNum, err := strconv.Atoi(strings.Split(lines[j], " ")[0])
				if err != nil {
					log.Fatal("В данных для сортировки находится не число")
				}
				return iNum < jNum
			}
			return lines[i] < lines[j]
		})
	}

	if params.reverse {
		for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
			lines[i], lines[j] = lines[j], lines[i]
		}
	}

	if params.unique {
		tmp := make([]string, 0)
		for i := 0; i < len(lines)-1; i++ {
			if lines[i] != lines[i+1] {
				tmp = append(tmp, lines[i])
			}
		}
		tmp = append(tmp, lines[len(lines)-1])
		lines = nil
		lines = append(lines, tmp...)
	}

	for _, line := range lines {
		sortedFile.WriteString(line)
		sortedFile.WriteString("\n")
	}
	return nil
}

func main() {

	var (
		column                     int
		numerical, reverse, unique bool
	)

	flag.IntVar(&column, "k", 0, "Usage: выбрать колонку для сортировки")
	flag.BoolVar(&numerical, "n", false, "Usage: сортировать по числовому значению")
	flag.BoolVar(&reverse, "r", false, "Usage: сортировать в обратном порядке")
	flag.BoolVar(&unique, "u", false, "Usage: не выводить повторяющиеся строки")

	flag.Parse()

	params := SortParams{
		column:    column,
		numerical: numerical,
		reverse:   reverse,
		unique:    unique,
	}

	file, err := os.Open("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	err = Sort(file, params)
	if err != nil {
		log.Fatal(err)
	}
}

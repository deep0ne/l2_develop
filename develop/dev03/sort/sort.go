package sort

import (
	"bufio"
	"dev03/utils"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

type SortParams struct {
	Column        int
	Numerical     bool
	Reverse       bool
	Unique        bool
	Sorted        bool
	TrailingSpace bool
	HumanNumeric  bool
	MonthSort     bool
}

// структура для удобства сортировки по полям
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
		if params.TrailingSpace {
			strings.TrimRight(line, " ")
		}
		lines = append(lines, line)
	}

	checker := make([]string, len(lines))
	copy(checker, lines)

	if params.Column != 0 {
		orderMap := make([]Order, 0)
		for idx, line := range lines {
			splitted := strings.Split(line, " ")
			if params.Column > len(splitted) {
				return errors.New("В строке нет такого количества колонок")
			}
			order := Order{
				columnWord: splitted[params.Column-1],
				index:      idx,
			}
			orderMap = append(orderMap, order)
		}
		// сортируем, сохраняя изначальные индексы для вывода
		sort.Slice(orderMap, func(i, j int) bool {
			if params.Numerical {
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
			if params.MonthSort {
				return utils.GetMonth(strings.ToLower(orderMap[i].columnWord)) < utils.GetMonth(strings.ToLower(orderMap[j].columnWord))
			}
			if params.HumanNumeric {
				return utils.GetMultiplier(orderMap[i].columnWord) < utils.GetMultiplier(orderMap[j].columnWord)
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
		if params.HumanNumeric {
			sort.Slice(lines, func(i, j int) bool {
				return utils.GetMultiplier(lines[i]) < utils.GetMultiplier(lines[j])
			})
		} else if params.MonthSort {
			sort.Slice(lines, func(i, j int) bool {
				return utils.GetMonth(strings.ToLower(lines[i])) < utils.GetMonth(strings.ToLower(lines[j]))
			})
		} else {
			sort.Slice(lines, func(i, j int) bool {
				if params.Numerical {
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
	}

	if params.Sorted {
		if reflect.DeepEqual(checker, lines) {
			fmt.Println("Данные уже отсортированы")
			return nil
		}
	}

	if params.Reverse {
		for i, j := 0, len(lines)-1; i < j; i, j = i+1, j-1 {
			lines[i], lines[j] = lines[j], lines[i]
		}
	}

	if params.Unique {
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

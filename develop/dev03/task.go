/*

Отсортировать строки в файле по аналогии с консольной утилитой sort (man sort — смотрим описание и основные параметры): на входе подается файл из несортированными строками, на выходе — файл с отсортированными.

Реализовать поддержку утилитой следующих ключей:

-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительно

Реализовать поддержку утилитой следующих ключей:

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учетом суффиксов
*/

package main

import (
	"dev03/sort"
	"flag"
	"log"
	"os"
)

func main() {

	var (
		column                                                                     int
		numerical, reverse, unique, sorted, trailingSpace, humanNumeric, monthSort bool
	)

	flag.IntVar(&column, "k", 0, "Usage: выбрать колонку для сортировки")
	flag.BoolVar(&numerical, "n", false, "Usage: сортировать по числовому значению")
	flag.BoolVar(&reverse, "r", false, "Usage: сортировать в обратном порядке")
	flag.BoolVar(&unique, "u", false, "Usage: не выводить повторяющиеся строки")
	flag.BoolVar(&sorted, "c", false, "Usage: проверяет, отсортированы ли данные")
	flag.BoolVar(&trailingSpace, "b", false, "Usage: игнорирует хвостовые пробелы")
	flag.BoolVar(&humanNumeric, "h", false, "Usage: сортировать по числовому значению с учетом суффиксов")
	flag.BoolVar(&monthSort, "M", false, "Usage: отсортировать месяцы")

	flag.Parse()

	if humanNumeric && numerical {
		log.Fatal("Для сортировки нужно выбрать либо numeric, либо human-numeric")
	}

	params := sort.SortParams{
		Column:        column,
		Numerical:     numerical,
		Reverse:       reverse,
		Unique:        unique,
		Sorted:        sorted,
		TrailingSpace: trailingSpace,
		HumanNumeric:  humanNumeric,
		MonthSort:     monthSort,
	}

	file, err := os.Open("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	err = sort.Sort(file, params)
	if err != nil {
		log.Fatal(err)
	}
}

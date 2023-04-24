package utils

import (
	"regexp"
	"strconv"
)

func GetMultiplier(s string) float64 {
	// Поиск последних нецифровых символов
	re := regexp.MustCompile(`(\d+)([KMGTPEZY]?B?)?$`)
	matches := re.FindStringSubmatch(s)

	// Если нет соответствия, то возвращаем ноль
	if len(matches) < 2 {
		return 0
	}

	// Определение множителя (KB, MB, GB, и т.д.)
	multiplier := 1.0
	switch matches[2] {
	case "KB", "K":
		multiplier = 1024.0
	case "MB", "M":
		multiplier = 1024.0 * 1024.0
	case "GB", "G":
		multiplier = 1024.0 * 1024.0 * 1024.0
	case "TB", "T":
		multiplier = 1024.0 * 1024.0 * 1024.0 * 1024.0
	case "PB", "P":
		multiplier = 1024.0 * 1024.0 * 1024.0 * 1024.0 * 1024.0
	case "EB", "E":
		multiplier = 1024.0 * 1024.0 * 1024.0 * 1024.0 * 1024.0 * 1024.0
	case "ZB", "Z":
		multiplier = 1024.0 * 1024.0 * 1024.0 * 1024.0 * 1024.0 * 1024.0 * 1024.0
	case "YB", "Y":
		multiplier = 1024.0 * 1024.0 * 1024.0 * 1024.0 * 1024.0 * 1024.0 * 1024.0 * 1024.0
	}

	// Преобразование числа из строки во float64 и умножение на множитель
	num, _ := strconv.ParseFloat(matches[1], 64)
	return num * multiplier
}

func GetMonth(s string) int {
	switch s {
	case "january":
		return 1
	case "february":
		return 2
	case "march":
		return 3
	case "april":
		return 4
	case "may":
		return 5
	case "june":
		return 6
	case "july":
		return 7
	case "august":
		return 8
	case "september":
		return 9
	case "october":
		return 10
	case "november":
		return 11
	case "december":
		return 12
	}
	return 0
}

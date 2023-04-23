/*
Создать программу печатающую точное время с использованием NTP -библиотеки. Инициализировать как go module.
Использовать библиотеку github.com/beevik/ntp. Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Требования:
Программа должна быть оформлена как go module
Программа должна корректно обрабатывать ошибки библиотеки: выводить их в STDERR и возвращать ненулевой код выхода в OS

*/

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

const NTPServer = "0.beevik-ntp.pool.ntp.org"

func GetNTPTime() time.Time {
	t, err := ntp.Time(NTPServer)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка получения времени", err)
		os.Exit(1)
	}
	return t
}

func main() {
	fmt.Println(GetNTPTime().Format("2006-01-02 15:04:05.000"))
}

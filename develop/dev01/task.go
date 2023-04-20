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

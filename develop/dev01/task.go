package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

const NTPServer = "0.beevik-ntp.pool.ntp.org"

func main() {
	t, err := ntp.Time(NTPServer)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка получения времени", err)
		os.Exit(1)
	}

	fmt.Println(t.Format("2006-01-02 15:04:05.000"))
}

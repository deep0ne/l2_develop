/*

Реализовать функцию, которая будет объединять один или более done-каналов в single-канал, если один из его составляющих каналов закроется.
Очевидным вариантом решения могло бы стать выражение при использованием select, которое бы реализовывало эту связь,
Однако иногда неизвестно общее число done-каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or-каналов, реализовывала бы весь функционал.

Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}

*/

package main

import (
	"fmt"
	"time"
)

var or func(channels ...<-chan interface{}) <-chan interface{} = func(channels ...<-chan interface{}) <-chan interface{} {

	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	output := make(chan any)
	donechan := make(chan struct{})

	fanIn := func(input <-chan interface{}) {

		for n := range input {
			output <- n
		}

		// сигнал каналу
		donechan <- struct{}{}
	}

	// Запускаем горутину для каждого канала
	for _, ch := range channels {
		go fanIn(ch)
	}

	// Ожидаем, пока все горутины завершатся
	go func() {
		<-donechan
		close(output)
	}()

	return output
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)

	fmt.Printf("fone after %v", time.Since(start))

}

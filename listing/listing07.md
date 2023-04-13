Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)
	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}
		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
1
2
3
4
5
6
7
8
0
0
0
...
```
После считывания чисел от 1 до 8 мы будем получать на выводе бесконечное количество нулей. Это происходит, потому что в функции `merge`, когда мы пишем значения в возвращаемый из функции канал, мы не проверяем, что канал закрыт. Когда канал закрыт, любое считывание из него вернёт значение по умолчанию, которое равняется нулю для `chan int`, поэтому чтобы бесконечно не выводить нули, нужно добавить проверку на закрытие канала и завершить цикл, когда оба канала закрыты.
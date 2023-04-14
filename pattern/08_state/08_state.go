/*
Паттерн Состояние относится к поведенческим паттернам.
Он позволяет объекту изменять своё поведение в зависимости от внутреннего состояния.
Мне кажется, этот паттерн можно реализовать на примере логгера, состоянием будет считаться его поток вывода
В данном примере состояние может быть stdout или stderr, но так же можно добавить состояние записи в файл, в внешний сервис и т.д.
*/

package main

import (
	"fmt"
	"os"
)

type State interface {
	Log(message string)
}

type Logger struct {
	state State
}

type STDOutState struct{}

func NewLogger() *Logger {
	return &Logger{&STDOutState{}}
}

func (l *Logger) SetState(state State) {
	l.state = state
}

func (l *Logger) Log(message string) {
	l.state.Log(message)
}

func (s *STDOutState) Log(message string) {
	fmt.Println("STDOUT LOG -->", message)
}

type STDErrState struct{}

func (s *STDErrState) Log(message string) {
	fmt.Fprintf(os.Stderr, "STDERR LOG --> ")
	fmt.Fprintln(os.Stderr, message)
}

func main() {
	logger := NewLogger()
	logger.Log("Hi, this is the first log")
	logger.SetState(&STDErrState{})
	logger.Log("Hi, this is the second log")
}

/*

Необходимо реализовать свой собственный UNIX-шелл-утилиту с поддержкой ряда простейших команд:

- cd <args> - смена директории (в качестве аргумента могут быть то-то и то)
- pwd - показать путь до текущего каталога
- echo <args> - вывод аргумента в STDOUT
- kill <args> - "убить" процесс, переданный в качесте аргумента (пример: такой-то пример)
- ps - выводит общую информацию по запущенным процессам в формате *такой-то формат*

Так же требуется поддерживать функционал fork/exec-команд

Дополнительно необходимо поддерживать конвейер на пайпах (linux pipes, пример cmd1 | cmd2 | .... | cmdN).

*/

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func Execute(arguments []string) error {
	command := arguments[0]
	switch command {
	case "cd":
		err := os.Chdir(arguments[1])
		if err != nil {
			fmt.Println("Cannot change directory. Change your input.")
			fmt.Println("Type \"pwd\" if you want to know your current directory")
		}

	case "pwd":
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println("Could not get current directory...")
		}
		fmt.Println(dir)

	case "echo":
		fmt.Println(strings.Join(arguments[1:], " "))

	case "ps":
		processes, err := exec.Command("ps", arguments[1:]...).Output()
		if err != nil {
			fmt.Println("There are no processes")
		}
		fmt.Println(string(processes))
	// убиваем процесс по его PID
	case "kill":
		var pid int
		for _, arg := range arguments[1:] { // ищем PID
			if arg[0] != '-' { // может быть статус для kill
				pidConverted, err := strconv.Atoi(arg)
				if err != nil {
					fmt.Println("Wrong PID. To find PID you can type \"ps -e\"")
				}
				pid = pidConverted
			}
		}

		process, err := os.FindProcess(pid)
		if err != nil {
			fmt.Println("Wrong PID. To find PID you can type \"ps -e\"")
		}
		err = process.Kill()
		if err != nil {
			fmt.Println("Could not kill process. Are you sure your PID is correct?")
		} else {
			fmt.Printf("Process %d is killed!\n", pid)
		}

	case "exec":
		if len(arguments[1:]) > 1 {
			fmt.Println("Too many arguments.")
			fmt.Println("Exec usage: exec firefox")
		}
		cmd := exec.Command(arguments[1])
		err := cmd.Run()
		if err != nil {
			fmt.Println("Wrong argument for exec.")

		}
	default:
		return errors.New("Wrong Command")
	}
	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		arguments := strings.Split(scanner.Text(), " ")
		if arguments[0] == "quit" {
			break
		}
		Execute(arguments)
	}
}

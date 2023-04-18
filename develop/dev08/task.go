package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		arguments := strings.Split(scanner.Text(), " ")
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

		case "kill":
			var pid int
			for _, arg := range arguments[1:] {
				if arg[0] != '-' {
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
		case "quit":
			return
		}
	}
}

/*


Реализовать простейший telnet-клиент.

Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Требования:
Программа должна подключаться к указанному хосту (ip или доменное имя + порт) по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s)
При нажатии Ctrl+D программа должна закрывать сокет и завершаться.
Если сокет закрывается со стороны сервера, программа должна также завершаться. При подключении к несуществующему сервер,
Программа должна завершаться через timeout

*/

package main

import (
	"bufio"
	"context"
	"div10/server"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	timeout := flag.Duration("timeout", 10*time.Second, "connection timeout")
	flag.Parse()

	if flag.NArg() != 2 {
		fmt.Println("Usage of telnet: go-telnet host port")
		os.Exit(1)
	}

	host := flag.Arg(0)
	port := flag.Arg(1)
	server.RunServer(":" + port) // создаём сервер

	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), *timeout) // подключаемся по tcp к созданному серверу
	if err != nil {
		fmt.Println("Error connecting to server")
		os.Exit(1)
	}

	defer conn.Close()

	fmt.Println("Now connected to:", conn.RemoteAddr())

	ctx, cancel := context.WithCancel(context.Background())

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-interrupt // для закрытия сокета по CTRL+D
		cancel()
	}()

	// пишем в сокет
	go func() {
		reader := bufio.NewReader(os.Stdin)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				word, err := reader.ReadString('\n')
				if err != nil {
					if err == io.EOF {
						cancel()
					}
				}
				_, err = fmt.Fprint(conn, word)
				if err != nil {
					fmt.Println(err)
					return
				}

			}
		}
	}()

	// читаем из сокета
	go func() {
		reader := bufio.NewReader(conn)
		for {
			select {
			case <-ctx.Done():
				return
			default:
				word, _ := reader.ReadString('\n')
				fmt.Printf("Server says %s", word)
			}
		}

	}()

	<-ctx.Done()
}

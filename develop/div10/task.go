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
	server.RunServer(":" + port)

	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), *timeout)
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
		<-interrupt
		cancel()
	}()

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

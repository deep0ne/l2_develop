package server

import "github.com/reiver/go-telnet"

func RunServer(port string) {
	handler := telnet.EchoHandler

	go func() {
		err := telnet.ListenAndServe(port, handler)
		if err != nil {
			panic(err)
		}
	}()
}

package main

import (
	"fmt"
	"net"
)

func main() {
	Listenner, _ := net.Listen("tcp", "localhost:8080")
	for {
		Connect, err := Listenner.Accept()
		if err != nil {
			continue
		}
		go WithClient(Connect)
	}
}

func WithClient(Connect net.Conn) {
	defer Connect.Close()

	ClientBuffer := make([]byte, 100)
	for {
		Connect.Write([]byte("Обработка заявки\n"))

		ReadBuffer, Err := Connect.Read(ClientBuffer)
		if Err != nil {
			fmt.Println(Err)
			break
		}

		Connect.Write(append([]byte("Закончили обработку"), ClientBuffer[:ReadBuffer]...))
	}
}

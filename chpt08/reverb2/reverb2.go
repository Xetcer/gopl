/*
go run reverb2.go
эхо-сервер может имитировать реверберацию обычного эха, сначала отвечая громко ("HELLO! "),
затем, после задержки, — умеренно ("Hello! "), а потом — совсем тихо ("hello! ").
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		//Аргументы функции, запускаемой с помощью go, вычисляются при выполнении самой инструкции go;
		// таким образом, input.Text() вычисляется в главной go- подпрограмме.
		// Обеспечиваем параллелизм в рамках одного соединения.
		go echo(c, input.Text(), 1*time.Second)
	}
	// Игнорируем потенциальные ошибки input.Err()
	c.Close()
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

// Netcatl - TCP-клиент только для чтения
// работает совместно с reverи1.go
// go run netcat2.go
/*
	отправляет входные данные на сервер и в то же время копирует ответ сервера на выход
	go run netcat2.go
	отправить строку типа Hello ?;
*/
package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go mustCopy(os.Stdout, conn)
	mustCopy(conn, os.Stdin)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

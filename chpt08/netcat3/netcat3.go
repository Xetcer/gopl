/*
 работает в паре с netcat2
 go run netcat3.go
 hello
 для завершения подпрограммы отправить ctrl+z +enter

go sub-programm: starting
hello
         HELLO
         hello
         hello
^Z
main programm waiting...
2025/01/09 13:22:48 done
go sub-programm: finished
main programm is done, exit programm

*/

package main

import (
	"fmt"
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
	done := make(chan struct{})
	go func() {
		fmt.Println("go sub-programm: starting")
		io.Copy(os.Stdout, conn) // игнорируем ошибки
		log.Println("\ngo sub-programm: done")
		done <- struct{}{} // сигнал главной подпрограмме
		fmt.Println("go sub-programm: finished")
	}()
	mustCopy(conn, os.Stdin)
	conn.Close()
	fmt.Println("main programm waiting...")
	<-done
	fmt.Println("main programm is done, exit programm")
	// ожидание завершения фоновой подпрограммы
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

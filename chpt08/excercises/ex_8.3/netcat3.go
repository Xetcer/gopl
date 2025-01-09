/*
 работает в паре с reverb1
 go run netcat3.go
 hello
 для завершения подпрограммы отправить ctrl+z +enter

В программе netcat3 значение интерфейса conn имеет конкретный тип *net .TCPConn, который представляет TCP-соединение.
ТСР-соединение состоит из двух половин, которые могут быть закрыты независимо с использованием методов CloseRead и CloseWrite.
Измените главную go-подпрограмму netcat3 так, чтобы она закрывала только записывающую половину соединения, так, чтобы программа
продолжала выводить последние эхо от сервера reverbl даже после того, как стандартный ввод будет закрыт.
(Сделать это для сервера reverb2 труднее; см. упражнение 8.4.)

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
	if TCPConn, ok := conn.(*net.TCPConn); ok {
		go func() {
			fmt.Println("go sub-programm: starting")
			io.Copy(os.Stdout, conn) // игнорируем ошибки
			log.Println("\ngo sub-programm: done")
			done <- struct{}{} // сигнал главной подпрограмме
			fmt.Println("go sub-programm: finished")
		}()
		mustCopy(conn, os.Stdin)
		TCPConn.CloseWrite()
		// conn.Close()
		fmt.Println("main programm waiting...")
		<-done
		fmt.Println("main programm is done, exit programm")
		TCPConn.CloseRead()
		// ожидание завершения фоновой подпрограммы
	} else {
		log.Fatal("can't convert conn to *net.TCPConn")
	}
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

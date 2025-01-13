/*
пользователям обмениваться текстовыми сообщениями друг с другом. В этой программе есть четыре вида go-подпрограмм.
Имеется по экземпляру go-подпрограмм main и broadcaster, а для каждого подключенного клиента имеется по одной go- подпрограмме handleConn и clientWriter.
Go-подпрограмма broadcaster является хорошей иллюстрацией использования инструкции select, так как она должна реагировать на три различных вида сообщений.
Работа главный go-подпрограммы, показанной ниже, состоит в прослушивании и приеме входящих сетевых подключений от клиентов.
Для каждого из них она создает новую go-подпрограмму handleConn,
go run chat.go
Работает с несколькими экземплярами netcat3 из этой главы, сообщения от каждого клиента отправляются каждому клиенту
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

type client chan<- string // канала исходящих сообщений
var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // Все входящие сообщения клиента
)

/*
broadcaster. Ее локальная переменная clients записывает текущее множество подключенных клиентов.
Единственная информация, записываемая о каждом клиенте, — это его канал для исходящих сообщений
*/
func broadcaster() {
	clients := make(map[client]bool) //все подключенные клиенты
	for {
		select {
		case msg := <-messages:
			// широковещательное входящее сообщение во все
			// каналы исходящих сообщений для клиентов
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

/*
Функция handleConn создает новый канал исходящих сообщений для своего клиента
и объявляет широковещателю о поступлении этого клиента по каналу entering.
Затем она считывает каждую строку текста от клиента, отправляя каждую строку
широковещателю по глобальному каналу входящих сообщений и предваряя каждое сообщение
указанием отправителя. Когда от клиента получена вся информация, handleConn объявляет
об убытии клиента по каналу leaving и закрывает подключение.
*/
func handleConn(conn net.Conn) {
	ch := make(chan string) // Исходящие сообщения клиентов
	go clientWriter(conn, ch)
	who := conn.RemoteAddr().String()
	ch <- "Вы " + who
	messages <- who + " подключился"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	leaving <- ch
	messages <- who + " отключился"
	conn.Close()
}

// clientWriter как только в канал поступает новое сообщение, оно будет отправлено,
// цикл будет выполняться до тех пор пока канал открыт
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}

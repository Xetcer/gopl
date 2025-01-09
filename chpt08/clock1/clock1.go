// Clockl является TCP-сервером, периодически выводящим время.
// go run clock1.go
package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	// Listen создает объект net.Listener, который прослушивает входящие соединения на сетевом порту, в данном случае это TCP-порт localhost: 8000
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		// Метод Accept прослушивателя блокируется до тех пор, пока не будет сделан входящий запрос на подключение, после чего возвращает объект net.Conn, представляющий соединение.
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // к примеру обрыв соединения
			continue
		}
		handleConn(conn)
	}
}

/*
Функция handleConn обрабатывает одно полное клиентское соединение. Она в цикле выводит клиенту текущее временя, time.Now().
Поскольку net.Conn соответствует интерфейсу io.Writer, мы можем осуществлять вывод непосредственно в него. Цикл завершается,
когда выполнение записи не удается, например потому, что клиент был отключен, и при этом handleConn закрывает свою сторону
соединения с помощью отложенного вызова Close и переходит в состояние ожидания очередного запроса на подключение.
*/
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return // например отключение клиента
		}
		time.Sleep(1 * time.Second)
	}
}

/*
go run reverb2.go
эхо-сервер может имитировать реверберацию обычного эха, сначала отвечая громко ("HELLO! "),
затем, после задержки, — умеренно ("Hello! "), а потом — совсем тихо ("hello! ").
go run reverb2.go
для проверки запустить netcat2 утилиту и не вводить ничего 10 секунд, в результате получим сообщение от сервера
в консоли netcat о том что мы не ввели ничего You don't enter message. Exit...
*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

const timeout = 10 * time.Second

func echo(c net.Conn, shout string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(conn net.Conn) {
	wg := &sync.WaitGroup{}
	defer func() {
		wg.Wait()
		conn.Close()
	}()
	input := bufio.NewScanner(conn)

	ch := make(chan string)

	go func() {
		for input.Scan() {
			ch <- input.Text()
		}
	}()

	for {
		select {
		case text := <-ch:
			wg.Add(1)
			go echo(conn, text, 1*time.Second, wg)
		case <-time.After(timeout):
			fmt.Fprintln(conn, "\rYou don't enter message. Exit...")
			return
		}
	}
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

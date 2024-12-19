/*
go run wait.go  https://golang.org
при успешном запуске ничего не увидишь)
*/
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

/*
WaitForServer пытается соединиться с сервером заданного URL.
Попытки предпринимаются в течение минуты с растущими интервалами.
Сообщает об ошибке, если все попытки неудачны,
*/
func WaitForServer(url string) error {
	const timeout = 1 * time.Minute
	deadLine := time.Now().Add(timeout)
	for tries := 0; time.Now().Before(deadLine); tries++ {
		_, err := http.Head(url)
		if err == nil {
			return nil // Успешное соединение
		}
		log.Printf("Сервер не отвечает (%s); повтор...", err)
		time.Sleep(time.Second << uint(tries)) // увеличение задержки
	}
	return fmt.Errorf("Сервер %s не отвечает; время %s ", url, timeout)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "usage: wait url\n")
		os.Exit(1)
	}
	url := os.Args[1]
	if err := WaitForServer(url); err != nil {
		fmt.Fprintf(os.Stderr, "Site is down: %v\n", err)
		os.Exit(1)
	}
}

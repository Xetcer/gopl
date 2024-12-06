/*
осуществляет выборку содержимого по каждому из указанных URL и выводит его как не интерпретированный текст
go run fetch.go http://www.gopl.io
*/
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix("http://", url) { // 1.8 добавляем префикс, если его нет проверка go run fetch.go www.gopl.io
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		// b, err := io.ReadAll(resp.Body)
		_, err = io.Copy(os.Stdout, resp.Body) // 1.7 копируем напрямую в выходной терминал тело, без выделения буфера
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: чтение %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Println("HTTP status: ", resp.Status) // 1.9 добавил вывод состояния HTTP
		// fmt.Printf("%s", b)
	}
}

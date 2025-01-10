package main

import (
	"fmt"
	"gopl/chpt05/links"
	"log"
	"os"
)

// tokens представляет собой подсчитывающий семафор, используемый
// для ограничения количества параллельных запросов величиной 20.
var tokens = make(chan struct{}, 20)

func crawl(url string) []string {
	fmt.Println(url)
	tokens <- struct{}{} // захват маркера
	list, err := links.Extract(url)
	<-tokens // освобождщение маркера
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	worklist := make(chan []string)  // Список URL могут быть дубли
	unseenLinks := make(chan string) // удаление дублей

	// Запуск с аргументами командной строки
	go func() { worklist <- os.Args[1:] }()

	// Создание 20 подпрограмм сканирования для выборки всех непросмотренных ссылок
	// Параллельное сканирование
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundlinks := crawl(link)
				go func() { worklist <- foundlinks }()
			}
		}()
	}
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}

}

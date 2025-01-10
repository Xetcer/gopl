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
	worklist := make(chan []string)
	var n int // количество ожидающих отправки в список

	// Запуск с аргументами командной строки
	n++
	go func() { worklist <- os.Args[1:] }()

	// Параллельное сканирование
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string) {
					worklist <- crawl(link)
				}(link)
			}
		}
	}
}

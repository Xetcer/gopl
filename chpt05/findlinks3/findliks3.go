package main

// go run findliks3.go https://golang.org

import (
	"fmt"
	"log"
	"os"

	"gopl/chpt05/links"
)

// breadthFirst вызывает f для каждого элемента в worklist.
// Все элементы, возвращаемые f, добавляются в worklist.
// f вызывается для каждого элемента не более одного раза,
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool) // Создаем карту, в которой отмечаем вызывалась ли f для этого элемента уже
	for len(worklist) > 0 {
		items := worklist            // сохраняем worklist в items
		worklist = nil               // зануляем worklist
		for _, item := range items { // для каждого элемента
			if !seen[item] { // если он еще не обрабатывался f
				seen[item] = true                       // ставим отметку об обработке в f
				worklist = append(worklist, f(item)...) // добавляем в worklist все элементы среза, возвращаемые f
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)                // выводит исходный url
	list, err := links.Extract(url) // исзвлекаем ссылки из url
	if err != nil {
		log.Print(err) // если есть ошибка, выводим ее в лог
	}
	return list // воззвращаем список ссылок, найденных в URL
}

func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, os.Args[1:])
}

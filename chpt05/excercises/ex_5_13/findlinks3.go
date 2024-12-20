package main

/*
Упражнение 5.13. Модифицируйте функцию crawl так, чтобы она делала локальные копии найденных ею страниц,
при необходимости создавая каталоги. Не делайте копии страниц, полученных из других доменов. Например, если
исходная страница поступает с адреса golang.org, сохраняйте все страницы оттуда,
но не сохраняйте страницы, например, с vimeo. com.

Запуск go run findlinks3.go http://gopl.io
*/

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

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

var baseHost string

func crawl(URL string) []string {
	_, err := url.Parse(URL)
	if err != nil {
		log.Print(err)
	}

	if baseHost == "" {
		baseHost = URL
	}

	if !strings.Contains(URL, baseHost) {
		return nil
	}

	err = downLoadPage(URL)
	if err != nil {
		fmt.Println("pageSave error:", err)
		log.Println(err)
	}

	fmt.Println(URL) // выводит исходный url
	list, err := links.Extract(URL)
	if err != nil {
		log.Println(err)
	}
	return list // воззвращаем список ссылок, найденныех
}

func downLoadPage(URL string) error {
	page, err := http.Get(URL)
	if err != nil {
		return err
	}

	// сорфмируем имя файла
	_, pageName := path.Split(URL)
	pageName += ".html"

	file, err := os.OpenFile(pageName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	pageBytes, err := io.ReadAll(page.Body)
	if err != nil {
		return err
	}
	_, err = file.Write(pageBytes)
	if err != nil {
		return err
	}
	return nil
}

// !+main
func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	breadthFirst(crawl, os.Args[1:])
}

//!-main

/*
Findlinksl выводит ссылки в HTML-документе,прочитанном со стандартного входа.
Запустим findlinks для начальной страницы Go, передав выход программы fetch (раздел 1.5) на вход findlinks.
Для запуска надо собрать командой go build два пакета fetch и findlinks

	./fetch https://golang.org | ./findlinks
*/

/*
Exercise 5.1
Измените программу findlinks так, чтобы она обходила связанный список n.FirstChild с помощью рекурсивных вызовов visit,
а не с помощью цикла.

Change the findlinks program to traverse the n.FirstChild linked list using recursive calls to visit instead of a loop.
*/
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

// visit добавляет в links все ссылки,
// найденные в п, и возвращает результат.
// Для спуска по дереву для узла п функция visit рекурсивно вызывает себя для каждого из дочерних узлов п,
// которые хранятся в связанном списке FirstChild.
func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	links = visit(links, n.FirstChild)
	links = visit(links, n.NextSibling)
	return links
}

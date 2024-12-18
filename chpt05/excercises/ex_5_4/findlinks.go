/*
Findlinksl выводит ссылки в HTML-документе,прочитанном со стандартного входа.
Запустим findlinks для начальной страницы Go, передав выход программы fetch (раздел 1.5) на вход findlinks.
Для запуска надо собрать командой go build два пакета fetch и findlinks

	./fetch https://golang.org | ./findlinks

Упражнение 5.4. Расширьте функцию visit так, чтобы она извлекала другие разновидности ссылок из
документа, такие как изображения, сценарии и листы стилей.
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
	if n.Type == html.ElementNode {
		switch n.Data {
		case "a", "img", "style", "script":
			for _, a := range n.Attr {
				links = append(links, a.Val)
			}

		}

		// for _, a := range n.Attr {
		// 	if a.Key == "href" || a.Key ==  || a.Key == "style" || a.Key == "script" {
		// 		links = append(links, a.Val)
		// 	}
		// }
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

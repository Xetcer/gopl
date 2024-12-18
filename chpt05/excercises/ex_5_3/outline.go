/*
программа использует рекурсию по дереву узлов HTML для вывода наброска структуры дерева;
когда программа встречает каждый элемент, она помещает дескриптор(имя ТЭГА сам ТЭГ ) элемента в стек, а затем выводит стек
На вход попадает вывод программы fetch.exe, которую надо предварительно построить (смотри глава 1)
запуск: ./fetch https://golang.org | ./outline
*/
/*
Упражнение 5.3.Напишите функцию для вывода содержимого всех текстовых узлов в дереве документа HTML.
Не входите в элементы <script> и <style>, поскольку их содержимое в веб-браузере не является видимым.
*/
package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	for _, text := range getTextFromHTML(nil, doc) {
		fmt.Println(text)
	}
}
func getTextFromHTML(texts []string, n *html.Node) []string {
	if n == nil {
		return texts
	}
	if n.Type == html.TextNode && n.Parent.Data != "script" && n.Parent.Data != "style" { // Если это текстовый узел и не скрипт или стиль
		if len(strings.TrimSpace(n.Data)) != 0 {
			for _, line := range strings.Split(n.Data, "\n") {
				if len(line) != 0 {
					texts = append(texts, strings.TrimSpace(line))
				}
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		texts = getTextFromHTML(texts, c)
	}
	return texts
}

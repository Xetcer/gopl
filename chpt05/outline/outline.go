/*
программа использует рекурсию по дереву узлов HTML для вывода наброска структуры дерева;
когда программа встречает каждый элемент, она помещает дескриптор(имя ТЭГА сам ТЭГ ) элемента в стек, а затем выводит стек
На вход попадает вывод программы fetch.exe, которую надо предварительно построить (смотри глава 1)
запуск: ./fetch https://golang.org | ./outline
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
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}
func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // Внесение html ТЭГ в стек
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}

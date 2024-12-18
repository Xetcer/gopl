/*
программа использует рекурсию по дереву узлов HTML для вывода наброска структуры дерева;
когда программа встречает каждый элемент, она помещает дескриптор(имя ТЭГА сам ТЭГ ) элемента в стек, а затем выводит стек
На вход попадает вывод программы fetch.exe, которую надо предварительно построить (смотри глава 1)
запуск: ./fetch https://golang.org | ./outline
*/
/*
Упражнение 5.2. Напишите функцию для заполнения отображения, ключами которого являются имена элементов (р, div, span и т.д.),
а значениями — количество элементов с таким именем в дереве HTML-документа.
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
	tagMap := outline(nil, doc)
	fmt.Println("Tag\t\tCount")
	for tag, count := range tagMap {
		fmt.Printf("%s\t\t%d\n", tag, count)
	}
}
func outline(stack []string, n *html.Node) (stackMap map[string]int) {
	stackMap = make(map[string]int)
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // Внесение html ТЭГ в стек
		stackMap[n.Data]++
		// fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		tempMap := outline(stack, c)
		for tag, count := range tempMap {
			stackMap[tag] += count
		}
	}
	return stackMap
}

/*
Упражнение 5.9. Напишите функцию expand(s string, f func(string) string) string,
которая заменяет каждую подстроку "$foo” в s текстом, который возвращается вызовом f ("foo").
*/
package main

import (
	"fmt"
	"strings"
)

func expand(s string, f func(string) string) string {
	newString := s
	if f == nil {
		return ""
	}
	words := strings.Fields(s)
	fmt.Println(words, len(words))
	for _, fooStr := range words {
		if strings.HasPrefix(fooStr, "$") {
			newString = strings.Replace(newString, fooStr, f(fooStr), 1)
		}
	}
	return newString
}

func testReplace(s string) string {
	if len(s) != 0 {
		return "test"
	}
	return ""
}

func main() {
	testStr := "Это $заменить строка, в которой $нужно $заменить элементы."
	fmt.Println(expand(testStr, testReplace))
}

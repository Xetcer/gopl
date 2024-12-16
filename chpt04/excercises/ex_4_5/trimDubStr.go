/*
	Упражнение 4.5. Напишите функцию, которая без выделения дополнительной памяти удаляет все смежные дубликаты в срезе [ ] string.
*/

package main

import "fmt"

func trimDubStr(str []string) []string {
	cStr := 0
	for _, s := range str {
		if str[cStr] == s {
			continue
		}
		cStr++
		str[cStr] = s
	}
	return str[:cStr+1]
}

func main() {
	str := []string{"a", "a", "a", "b", "b", "c", "c"}
	fmt.Println("Init string:", str)
	fmt.Println("After trim:", trimDubStr(str))
}

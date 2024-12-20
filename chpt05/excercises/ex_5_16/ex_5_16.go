package main

/*
Упражнение 5.16. Напишите вариативную версию функции strings. Join.
*/
import "fmt"

func join(sep string, elems ...string) string {
	strLen := len(elems)
	switch strLen {
	case 0:
		return ""
	case 1:
		return elems[0]
	default:
		newStr := elems[0]
		for i := 1; i < strLen; i++ {
			newStr += sep + elems[i]
		}
		return newStr
	}
}

func main() {
	fmt.Println(join("_"))                        // ""
	fmt.Println(join("_", "one"))                 // one
	fmt.Println(join("_", "one", "two"))          // one_two
	fmt.Println(join("_", "one", "two", "three")) // one_two_three
}

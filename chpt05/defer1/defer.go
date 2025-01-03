package main

/*
Когда программа сталкивается с аварийной ситуацией, все отложенные функции выполняются в порядке,
обратном их появлению в исходном тексте, начиная с функции на вершине стека и опускаясь до функции main
*/
// go run defer.go

import "fmt"

func main() {
	f(3)
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // сбой при х==0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

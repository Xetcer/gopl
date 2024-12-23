package main

import "fmt"

func main() {
	fmt.Println(returnRecover())
}

/*
returnRecover в случае паники возвращает значение из восстановления и сообщения о панике
*/
func returnRecover() (result int, err error) {
	defer func() {
		if p := recover(); p != nil {
			result = 42
			err = fmt.Errorf("recover msg: %v", p)
		}
	}()
	panic("It is a panic!")
}

package main

import (
	"log"
	"time"
)

/*
Инструкция defer может также использоваться для пары отладочных записей о входе в некоторую функцию и выходе из нее.
Показанная ниже функция BigSlowOperation немедленно вызывает функцию trace, которая выполняет запись о входе в
функцию и возвращает значение-функцию, которая при вызове выполняет запись о выходе из функции. Таким образом,
с помощью отложенного вызова возвращаемой функции мы можем выполнять запись о входе в функцию и выходе из нее в одной
инструкции и даже передавать между этими двумя действиями значения, например время начала работы функции. Но не забывайте
о завершающей паре скобок в инструкции defer, иначе “входное’4 действие будет выполнено на выходе из функции, а “выходное’'
не будет выполнено вовсе!

 go run trace.go
*/

func bigSlowOperation() {
	defer trace("bigCloswOperation")() // Не забываем о скобках!
	// имитация длительной работы
	time.Sleep(10 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("вход в %s", msg)
	return func() { log.Printf("выход из %s (%s)", msg, time.Since(start)) }
}

func main() {
	bigSlowOperation()
}

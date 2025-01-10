/*
Приведенная ниже программа осуществляет обратный отсчет для запуска ракеты.
Функция time. Tic к возвращает канал, по которому она периодически отправляет события,
действуя как метроном. Каждое событие представляет собой значение момента времени,
но оно не так интересно, как сам факт его доставки.

	go run countdown1.go
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("начинаю отсчет.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
	launch()
}

func launch() {
	fmt.Println("Lift off!")
}

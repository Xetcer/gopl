/*
Давайте сделаем нашу программу запуска выводящей обратный отсчет.
Инструкция select ниже приводит к тому, что на каждой итерации цикла
выполняется ожидание сигнала прерывания в течение секунды, но не дольше.
*/
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()

	fmt.Println("Начинаю отсчет. Нажмите <Enter> для отказа")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			// ничего не делаем
		case <-abort:
			fmt.Println("Запуск отменен Jopta!")
			return
		}
	}
	launch()
}

func launch() {
	fmt.Println("Lift off!")
}

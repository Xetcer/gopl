package main

var deposits = make(chan int) // Обновление вклада
var balances = make(chan int) // Получение баланса

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

// Это монитор который позволяет извне получить данные внутри данного модуля.
func teller() {
	var balance int //balance ограничен подпрограммой teller
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // запуск управляющей подпрограммы
}

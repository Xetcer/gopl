package main

type cashOP struct {
	amount   int       // количество наличных для снятия
	opResult chan bool // канал для возвращения результата
}

var deposits = make(chan int)    // Обновление вклада
var balances = make(chan int)    // Получение баланса
var withdraw = make(chan cashOP) // Снятие наличных

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	newCashOp := cashOP{amount: amount, opResult: make(chan bool)}
	withdraw <- newCashOp

	return <-newCashOp.opResult
}

// Это монитор который позволяет извне получить данные внутри данного модуля.
func teller() {
	var balance int //balance ограничен подпрограммой teller
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case cashOpertion := <-withdraw: // получаем
			if balance < cashOpertion.amount {
				cashOpertion.opResult <- false
			} else {
				balance -= cashOpertion.amount
				cashOpertion.opResult <- true
			}

		}
	}
}

func init() {
	go teller() // запуск управляющей подпрограммы
}

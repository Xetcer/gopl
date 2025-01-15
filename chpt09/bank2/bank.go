/*
Мы используем ту же самую идею и канал емкостью 1 для того, чтобы гарантировать,
что одновременно к совместно используемой переменной может обратиться только одна go-подпрограмма.
Семафор, которым ведет подсчет только до 1, называется бинарным семафором.

Такой шаблон взаимного исключения настолько полезен, что поддерживается непосредственно типом Mutex из пакета sync.
Его метод Lock захватывает маркер (вызывает блокировку), а метод Unlock его освобождает
*/
package bank

var (
	sema    = make(chan struct{}, 1) // Бинарый семафор для
	balance int                      // защиты balance
)

func Deposit(amount int) {
	sema <- struct{}{} // Захват маркера
	balance = balance + amount
	<-sema // Освобождение маркера
}

func Balance() int {
	sema <- struct{}{} // Захват маркера
	b := balance
	<-sema // освобождение маркера
	return b
}

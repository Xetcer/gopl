/*
каждый элемент отображения является указателем на структуру entry.
Каждый элемент entry содержит записанный результат вызова функции f,
как и прежде, но теперь он дополнительно содержит канал с именем ready.
Сразу после того, как поле result структуры entry оказывается записанным,
этот канал закрывается, широковещательно оповещая (раздел 8.9) все прочие
go-подпрограммы о том, что теперь можно безопасно читать результат из этого
элемента entry.
*/
package memo

import "sync"

// Func является типом функции с запоминанием
type Func func(key string) (interface{}, error)
type result struct {
	value interface{}
	err   error
}
type entry struct {
	res   result
	ready chan struct{} // Закрывается когда res готов
}

type Memo struct {
	f     Func
	mu    sync.Mutex //Защита cache
	cache map[string]*entry
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// Это первый запрос данного ключа.
		// Эта go-подпрограмма становится ответственной за
		// вычисление значения и оповещение о готовности.
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()
		e.res.value, e.res.err = memo.f(key)
		close(e.ready) // Широковещательное оповещение о готовности
	} else {
		// повторный запрос данного ключа.
		memo.mu.Unlock()
		<-e.ready // Ожидание готовности
	}
	return e.res.value, e.res.err
}

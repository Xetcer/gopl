package memo

import (
	"sync"
)

// Func является типом функции с запоминанием
type Func func(key string) (interface{}, error)
type result struct {
	value interface{}
	err   error
}

type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]result
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

/*
Теперь детектор гонки молчит даже при параллельном выполнении тестов.
К сожалению, это изменение Memo отменяет добытый ранее прирост производительности.
 Применяя блокировку на время, равное продолжительности каждого вызова f, Get сериализует
 все операции ввода-вывода, которые мы намеревались сделать параллельными.
*/
// Get безопасный с точки зрения параллельности метод
func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	memo.mu.Unlock()
	return res.value, res.err
}

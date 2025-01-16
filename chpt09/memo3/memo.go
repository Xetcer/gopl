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

func (memo *Memo) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	memo.mu.Unlock()
	if !ok {
		res.value, res.err = memo.f(key)
		// между этими двумя критическими разделами
		// несколько горутин могут вычислять f(key)
		// и обновлять карту кэша
		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}
	return res.value, res.err
}

/*
Производительность снова улучшается, но теперь мы замечаем, что некоторые URL выбираются дважды.
Это происходит, когда две или более go-подпрограмм вызывают Get для одного и того же URL в одно и то же время.
Обе выполняют поиск в кеше, не находят искомого значения, а затем вызывают медленную функцию f.
После этого обе go-подпрограммы обновляют отображение полученным результатом. В итоге один результат заменяется другим.
В идеале хотелось бы избежать такой лишней работы. Эта возможность иногда называется подавлением повторений.
*/

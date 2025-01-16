/*
альтернативный вариант относительно memo4,
в котором переменная отображения ограничена
управляющей go-подпрограммой, которой абоненты
Get должны отправлять сообщение.
*/
package memo

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

// request сообщение, требующее применение Func к Key
type request struct {
	key      string
	response chan<- result
}

type Memo struct{ requests chan request }

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

/*
	Get создает канал response, помещает его в запрос и отправляет запрос
	управляющей go-подпрограмме, после чего сразу же переходит
	к получению ответа из этого канала.
*/
func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

/*
Переменная cache ограничена управляющей go-подпрограммой (*Memo). server, показанной ниже.
 Она считывает запросы в цикле до тех пор, пока канал не будет закрыт с помощью метода Close.
  Для каждого запроса она обращается к кешу, создавая и вставляя новую запись entry,
  если таковая не была найдена.
*/
func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// This is the first request for this key.
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key) // call f(key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	// Вычисление функции.
	e.res.value, e.res.err = f(key)
	//Оповещение о готовности
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// Ожидание готовности
	<-e.ready
	// отправка результата клиенту
	response <- e.res
}

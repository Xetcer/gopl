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

// request - запрос, содержащий ключ и канал для отправки ответа.
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
	response := make(chan result)           // Создаем канал ответа
	memo.requests <- request{key, response} // закидываем его в канал запроса
	res := <-response                       // ожидаем получения ответа из канала, который будет в функции deliver
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
	cache := make(map[string]*entry) // создаем кэш
	for req := range memo.requests { // для каждого нового запроса поступившего в канал запросов
		e := cache[req.key] // смотрим есть такой запрос в кэше или нет
		if e == nil {       // если его нет
			// This is the first request for this key.
			e = &entry{ready: make(chan struct{})} // создаем экземпляр записи в кэш, где создаем канал который будет сообщать о готовности результата
			cache[req.key] = e                     // Добавляем запись в кэш
			go e.call(f, req.key)                  // call f(key)	// вызываем функцию вызова функции отправки запроса по ссылке
		}
		go e.deliver(req.response) // сообщаем о том что получили ответ на текущую ссылку
	}
}

func (e *entry) call(f Func, key string) {
	// Вычисление функции.
	e.res.value, e.res.err = f(key) // производим запрос по ссылке
	//Оповещение о готовности
	close(e.ready) // закрываем канал
}

func (e *entry) deliver(response chan<- result) {
	// Ожидание готовности
	<-e.ready // пока не закрыт канал, который говорит о том что запрос по ссылке завершен ждем окончания
	// отправка результата клиенту
	response <- e.res // передаем результат в канал ответа
}

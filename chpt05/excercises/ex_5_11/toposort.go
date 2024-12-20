/*
проблему вычисления последовательности курсов информатики, которая удовлетворяет требованиям каждого из них,
заключающимся в том, что определенный курс опирается на другие курсы, которые должны быть изучены до него.
Условия приведены в таблице prereqs ниже, которая представляет собой карту каждого курса на список курсов,
которые должны быть пройдены до данного курса.

Упражнение 5.11. Преподаватель линейной алгебры (linear algebra) считает, что до его курса следует прослушать курс матанализа (calculus).
Перепишите функцию topoSort так, чтобы она сообщала о наличии циклов.
*/

package main

import (
	"fmt"
	"log"
	"sort"
)

// для проверки отработки зацикливания раскоментить любую строку
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	//"linear algebra": {"calculus"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	//"computer organization": {"compilers"},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	sorted, err := topoSort(prereqs)
	if err != nil {
		log.Println(err)
	}
	for i, course := range sorted {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, error) {
	var order []string                      // Создаем срез в котором будет порядок курсов
	seen := make(map[string]bool)           // создаем карту, в которой отмечаем был ли данный курс уже добавлен в выходной срез
	var visitAll func(items []string) error // объявляем анонимную функцию, которая будет еще и рекурсивной

	visitAll = func(items []string) error {
		for _, item := range items { // бежим по всем ключам карты с курсами
			if !seen[item] { // если элемент еще не просматривался
				seen[item] = true                         // ставим отметку о просмотре
				if err := visitAll(m[item]); err != nil { // просматриваем все курсы которые необходимо пройти для изучения текущего
					return err
				}
				order = append(order, item) // если нет ошибок то добавим в выходной срез, найденные курсы.
			} else { // если элемент уже просматривался
				isCycle := true                // ставим флаг, о том что обнаружили цикл
				for _, course := range order { // бежим по всем курсам, которые уже добавили в выходной список
					if course == item { // если в выходном срезе есть зацикленный курс
						isCycle = false // сбрасывам флаг зацикливания
					}
				}
				if isCycle { // если же мы еще не добавили в выходной срез зацикленный курс, значит это ошибка
					return fmt.Errorf("has cycle: %s", item)
				}
			}
		}
		return nil
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	if err := visitAll(keys); err != nil {
		return nil, err
	}
	return order, nil
}

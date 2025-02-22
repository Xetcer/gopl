/*
проблему вычисления последовательности курсов информатики, которая удовлетворяет требованиям каждого из них,
заключающимся в том, что определенный курс опирается на другие курсы, которые должны быть изучены до него.
Условия приведены в таблице prereqs ниже, которая представляет собой отображение каждого курса на список курсов,
которые должны быть пройдены до данного курса.
*/
package main

import (
	"fmt"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	for key, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", key, course)
	}
}

func topoSort(m map[string][]string) map[int]string {
	order := make(map[int]string)
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order[len(order)] = item
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	// sort.Strings(keys)
	visitAll(keys)
	return order
}

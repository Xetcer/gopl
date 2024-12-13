package annagram

import (
	"fmt"
	"testing"
)

func TestAnnagram(t *testing.T) {
	tests := []struct {
		s1, s2 string
		want   bool
	}{
		{s1: "test1", s2: "test1", want: true},
		{s1: "test2", s2: "test", want: false},
		{s1: "test3", s2: "test4", want: false},
		{s1: "test4 test4 test4", s2: "test4 test4 test4", want: true},
		{s1: "test5", s2: "test5 test5 test5", want: true},
		{"Statue of Liberty", "Built to stay free", true},
		{"Statue Liberty", "Built to stay free", false},
		{"eat", "tea", true},
		{"listen", "silent", true},
		{"anagram", "nag a ram", true},
		{"Elvis", "lives", true},
		{"A gentleman", "elegant man", true},
		{"Clint Eastwood", "old west action", true},
		{"Tom Marvolo Riddle", "I am Lord Voldemort", true},
		{"dormitory", "dirty room", true},
		{"the eyes", "they see", true},
		{"slot machines", "cash lost in me", true},
		{"debit card", "bad credit", true},
		{"astronomer", "moon starer", true},
		{"tea", "coffee", false},
		{"Statue Liberty", "Built to stay free", false},
		{"hello", "world", false},
		{"мама мыла раму", "раму мыла мама", true},
		{"воз и ныне там", "там и ныне воз", true},
		{"я с миром", "мир со мной", false},
		{"была цель, жить правильно", "жить было правильно, цель аль", true},
		{"нам дали тепло", "тепло дали нам", true},
		{"отвертка лежит рядом с телом", "отвертка рядом лежит с телом", true},
		{"возможности", "положительный", false},
		{"честный человек", "мудрость", false},
		{"чувство ответственности", "ответственность чувство", false},
		{"тяжело в учении - легко в бою", "в учении легко - тяжело в бою", true},
		{"надо отдать должное", "отдать надо должное", true},
		{"今天是个好天气，啊好天气", "好天气个啊好天气是，天今", true},
	}

	for _, test := range tests {
		fmt.Println(len(test.s1), len(test.s2))
		got := anagram(test.s1, test.s2)
		if got != test.want {
			t.Errorf("annagram(%q, %q) result=%t, but %t expected", test.s1, test.s2, got, test.want)
		}
	}
}

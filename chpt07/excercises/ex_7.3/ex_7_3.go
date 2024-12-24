package main

import (
	"bytes"
	"strconv"
)

type tree struct {
	value       int
	left, right *tree
}

func (t *tree) String() string {
	var b bytes.Buffer

	b.WriteRune('[')
	t.appendToBuilder(&b)
	if b.Len() > 0 {
		b.Truncate(b.Len() - 1)
	}
	b.WriteRune(']')

	return b.String()
}

func (t *tree) appendToBuilder(b *bytes.Buffer) {
	if t == nil {
		return
	}
	t.left.appendToBuilder(b)
	b.WriteString(strconv.Itoa(t.value))
	b.WriteRune(' ')
	t.right.appendToBuilder(b)
}

// Sort сортирует значения на месте
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues добавляет элементы t к values в требуемом
// порядке и возвращает результирующий срез
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		//Эквивалентно возврату &tree{value: value}
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

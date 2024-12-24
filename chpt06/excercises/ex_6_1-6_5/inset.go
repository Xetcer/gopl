package main

import (
	"bytes"
	"fmt"
)

// PLATFORM 6.5 определяет разрядность платформы 32|64 в зависимости от размера uint на текущей платформе
const PLATFORM = 32 << (^uint(0) >> 64)

type IntSet struct {
	words []uint64
}

// Has указывает, содежит ли множество неотрицательное значениче х
func (s *IntSet) Has(x int) bool {
	word, bit := x/PLATFORM, uint(x%PLATFORM) // получаем слово и бит от числа X
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add добавляет неотрицательное значение x в множество
func (s *IntSet) Add(x int) {
	word, bit := x/PLATFORM, uint(x%PLATFORM)
	// fmt.Printf("value =%d word=%d bit=%d, slice len=%d\n", x, word, bit, len(s.words))
	for word >= len(s.words) { // Пока слово не равно длине среза words в s
		s.words = append(s.words, 0) // добавляем 0 к слову
		// fmt.Println("add new 0 to s.words:", s.words)
	}
	s.words[word] |= 1 << bit // s.words[word] ставим в 1 бит
	// fmt.Println("XOR 1<<bit:", s.words)
}

// AddAll 6.2 добавляет список элеметов в множество
func (s *IntSet) AddAll(vals ...int) {
	for _, value := range vals {
		s.Add(value)
	}
}

// Remove 6.1 Удаляет x из множества
func (s *IntSet) Remove(x int) {
	word, bit := x/PLATFORM, uint(x%PLATFORM)
	// fmt.Printf("value =%d word=%d bit=%d, slice len=%d\n", x, word, bit, len(s.words))
	if word < len(s.words) {
		// fmt.Println("remove 1 to s.words:", s.words)
		// mask := ^uint64(1 << bit) // Создаем инвертированную маску числа
		s.words[word] &^= 1 << bit // сравниваем с маской текущее число
	}
}

// Clear 6.1 Удаляет все элементы множества
func (s *IntSet) Clear() {
	s.words = s.words[:0]
	s.words = nil
}

// Copy 6.1 Возвращает копию множества
func (s *IntSet) Copy() *IntSet {
	newSet := &IntSet{}
	newSet.words = make([]uint64, len(s.words))
	copy(newSet.words, s.words)
	return newSet
}

// UnionWith делает множество s равным объединению множеств s и t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// IntersectWidth 6.3 делает множество равным пересечению множеств s и t
func (s *IntSet) IntersectWidth(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// DifferenceWidth 6.3 делает множество s равным разнице множеств s и t
// Разность множеств — это все элементы, которые содержатся в одном множестве (где вызывается метод),
// но не содержатся в другом (куда передается аргументом).
func (s *IntSet) DifferenceWidth(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// SymmetricDifference 6.3 Симметричная разность двух множеств содержит элементы, имеющиеся в одном из множеств, но не в обоих одновременно
func (s *IntSet) SymmetricDifference(t *IntSet) {
	// c := s.Copy()        // сохраняем s
	// s.UnionWith(t)       // объединяем s t
	// c.IntersectWidth(t)  // находим пересечения t и исходного s
	// s.DifferenceWidth(c) // исключаем из объединения - пересечения
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}

}

// Elems 6.4 выводим читабельный срез значений множества
func (s *IntSet) Elems() (result []int) {
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < PLATFORM; j++ {
			if word&(1<<uint(j)) != 0 {
				result = append(result, PLATFORM*i+j)
			}
		}
	}
	return result
}

// String возвращает множество как строку вида "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < PLATFORM; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", PLATFORM*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Len 6.1 Возвращает количество чисел лежащих в множестве.
func (s *IntSet) Len() int {
	len := 0
	for _, word := range s.words {
		for i := 0; i < PLATFORM; i++ {
			if word&(1<<i) != 0 {
				len++
			}
		}
	}
	return len
}

func main() {
	var x IntSet
	x.Add(5)
	fmt.Println(x.String())
	fmt.Println("Len=", x.Len())

	x.Add(15)
	fmt.Println(x.String())
	fmt.Println("Len=", x.Len())

	new := x.Copy()
	fmt.Println("New is:", new.String())

	x.AddAll(4, 8, 15, 16, 23, 42)
	fmt.Println("AddAll: ", x.String())

	// x.Remove(1)
	x.Remove(15)
	fmt.Println(x.String())
	fmt.Println("Len=", x.Len())

	// intersection
	var y IntSet
	y.AddAll(4, 8, 15, 16, 23, 42)
	x.IntersectWidth(&y)
	fmt.Println("Intersection:", x.String())

	// DifferenceTest
	x.Clear()
	fmt.Println("After clear Len=", x.Len())
	x.DifferenceWidth(&y)
	fmt.Println("Differense with y:", x.String())

	// SymmetricDifference test
	x.Clear()
	y.Clear()
	x.AddAll(4, 8, 15, 16)
	y.AddAll(15, 16, 23, 42)
	x.SymmetricDifference(&y)
	fmt.Println("Differense with y:", x.String())

	// Elems
	fmt.Println("Elems of x:", x.Elems())

	// Example_One()
	// fmt.Println()
	// Example_Two()
}

func Example_One() {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String(), " count:", x.Len()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String(), " count:", y.Len()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String(), " count:", x.Len()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func Example_Two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x, " count:", x.Len())         // "{1 9 42 144}"
	fmt.Println(x.String(), " count:", x.Len()) // "{1 9 42 144}"
	fmt.Println(x, " count:", x.Len())          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}

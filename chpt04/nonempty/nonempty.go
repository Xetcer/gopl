package main

import "fmt"

// nonempty возвращает срез, содержащий только непустые строки.
// Содержимое базового массива при работе функции изменяется,
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func main() {
	//!+main
	data := []string{"one", "", "three"}
	fmt.Println("First nonempty function:", data)
	fmt.Printf("%q\n", nonempty(data)) // `["one" "three"]`
	fmt.Printf("%q\n", data)           // `["one" "three" "three"]`
	data = []string{"one", "", "three"}
	fmt.Println("First nonempty function:", data)
	fmt.Printf("%q\n", nonempty2(data)) // `["one" "three"]`
	fmt.Printf("%q\n", data)            // `["one" "three" "three"]`
	//!-main
}

//!+alt
func nonempty2(strings []string) []string {
	out := strings[:0] // zero-length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

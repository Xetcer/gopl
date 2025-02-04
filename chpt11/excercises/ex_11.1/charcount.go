// Charcount вычисляет количество символов Unicode,
package charcount

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func GetCharCount(r io.Reader) [utf8.UTFMax + 1]int {
	counts := make(map[rune]int)    // количество символов в Unicode
	var utflen [utf8.UTFMax + 1]int // Количество длин кодировок UTF-8
	invalid := 0                    // Количество некорректных символов UTF-8

	in := bufio.NewReader(r)
	for {
		r, n, err := in.ReadRune() // Возврат руны, размера и ошибки
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		counts[r]++
		utflen[n]++
	}
	return utflen
}

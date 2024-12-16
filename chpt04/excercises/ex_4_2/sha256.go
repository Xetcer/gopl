/*
go run sha256.go -help

функция получает контрольную сумму введенной в терминале строки с указанным при запуске алгоритмом
go run sha256.go -hash=0  // SHA256
go run sha256.go -hash=1  // SHA384
go run sha256.go -hash=2  // SHA512
*/
package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func shaApply(shatype int, text string) string {
	switch shatype {
	default:
		c1 := sha256.Sum256([]byte(text))
		return fmt.Sprintf("SHA256 \"%s\": %x\n", text, c1)
	case 1:
		c1 := sha512.Sum384([]byte(text))
		return fmt.Sprintf("SHA384 \"%s\": %x\n", text, c1)
	case 2:
		c1 := sha512.Sum512([]byte(text))
		return fmt.Sprintf("SHA512 \"%s\": %x\n", text, c1)
	}
}

func main() {
	var hash int
	flag.IntVar(&hash, "hash", 0, "0-sha256, 1-sha384, 2-sha512")
	flag.Parse()

	// c1 := sha256.Sum256([]byte("x"))
	// c2 := sha256.Sum256([]byte("X"))
	// fmt.Printf("%X\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(shaApply(hash, input.Text()))
	}
}

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func comma(s string) string {
	var buffer bytes.Buffer

	groupLen := len(s) % 3
	// fmt.Printf("\"%s\" len=%d, first group len = %d\n", s, len(s), groupLen)
	if groupLen == 0 {
		groupLen = 3
	}
	buffer.WriteString(s[:groupLen])
	for i := groupLen; i < len(s); i += 3 {
		buffer.WriteString(",")
		buffer.WriteString(s[i : i+3])
	}
	return buffer.String()
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(comma(input.Text()))
	}
}

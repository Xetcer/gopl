/*
функция Sum256 из пакета crypto/ sha256 генерирует криптографический хеш, или дайджест, SHA256 сообщения,
хранящегося в произвольном байтовом срезе. Дайджест состоит из 256 битов, поэтому его типом является [32] byte.
Если два дайджеста совпадают, то очень вероятно, что соответствующие сообщения одинаковы; если же дайджесты различаются,
то различаются и сообщения. Приведенная далее программа выводит и сравнивает дайджесты SHA256 для "х" и "X":
*/
package main

import (
	"crypto/sha256"
	"fmt"
)

func compareBits(b1, b2 byte) int {
	difCount := 0
	// fmt.Printf("%b : %b\n", b1, b2)
	for i := 0; i < 8; i++ {
		if (b1 & byte(1<<i)) != (b2 & byte(1<<i)) {
			// fmt.Printf("bit %d: b1&1<<%[1]d=%d,  b2&1<<%[1]d=%d\n", i, b1&1<<i, b2&1<<i)
			difCount++
		}
	}
	return difCount
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%X\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	difCount := 0
	for i := 0; i < len(c1); i++ {
		difCount += compareBits(c1[i], c2[i])
	}
	fmt.Println("Differet bytes count:", difCount)

}

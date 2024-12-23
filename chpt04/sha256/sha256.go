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

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%X\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}

/*
Упражнение 9.2. Перепишите пример PopCount из раздела 2.6.2 так, чтобы он инициализировал
таблицу поиска с использованием sync .Once при первом к ней обращении. (В реальности стоимость
синхронизации для таких малых и высокооптими- зированных функций, как PopCount, является чрезмерно высокой.)
*/
package main

import (
	"fmt"
	"sync"
)

// pc[i] - количество единичных битов в i.
var pc [256]byte
var loadTableOnce sync.Once
var wg sync.WaitGroup

// func init() {
// 	for i := range pc {
// 		pc[i] = pc[i/2] + byte(i&1)
// 	}
// }

func loadTable() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
	loadTableOnce.Do(loadTable)
	return int(
		pc[byte(x>>(0*8))] +
			pc[byte(x>>(1*8))] +
			pc[byte(x>>(2*8))] +
			pc[byte(x>>(3*8))] +
			pc[byte(x>>(4*8))] +
			pc[byte(x>>(5*8))] +
			pc[byte(x>>(6*8))] +
			pc[byte(x>>(7*8))])
}

func PopCount_2_3(x uint64) int {
	var popCount = 0
	loadTableOnce.Do(loadTable)
	for i := 0; i < 8; i++ {
		popCount += int(pc[byte(x>>(i*8))])
	}
	return popCount
}

func PopCount_2_4(x uint64) int {
	var popCount = 0
	temp := x
	loadTableOnce.Do(loadTable)
	for i := 0; i < 64; i++ {
		popCount += int(temp & 1)
		temp = temp >> 1
	}
	return popCount
}

func PopCount_2_5(x uint64) int {
	var popCount = 0
	// fmt.Printf("x = %d(%b)\n", x, x)
	loadTableOnce.Do(loadTable)
	for x != 0 {
		x &= x - 1
		popCount++
		// fmt.Printf("x = %d(%b)\n", x, x)
	}
	return popCount
}

func main() {
	wg.Add(1)
	go func() {
		fmt.Println(PopCount(100))
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		fmt.Println(PopCount_2_3(100))
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		fmt.Println(PopCount_2_4(100))
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		fmt.Println(PopCount_2_5(100))
		wg.Done()
	}()
	wg.Wait()
}

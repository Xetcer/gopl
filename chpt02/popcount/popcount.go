package main

import "fmt"

// pc[i] - количество единичных битов в i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(x uint64) int {
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
	for i := 0; i < 8; i++ {
		popCount += int(pc[byte(x>>(i*8))])
	}
	return popCount
}

func PopCount_2_4(x uint64) int {
	var popCount = 0
	temp := x
	for i := 0; i < 64; i++ {
		popCount += int(temp & 1)
		temp = temp >> 1
	}
	return popCount
}

func PopCount_2_5(x uint64) int {
	var popCount = 0
	// fmt.Printf("x = %d(%b)\n", x, x)
	for x != 0 {
		x &= x - 1
		popCount++
		// fmt.Printf("x = %d(%b)\n", x, x)
	}
	return popCount
}

func main() {
	fmt.Println(PopCount(100))
	fmt.Println(PopCount_2_3(100))
	fmt.Println(PopCount_2_4(100))
	fmt.Println(PopCount_2_5(100))
}

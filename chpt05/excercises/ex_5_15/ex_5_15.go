package main

import (
	"fmt"
	"math"
)

func min(vars ...int) (int, error) {
	if len(vars) == 0 {
		return 0, fmt.Errorf("%s", "no values to compare")
	}
	min := math.MaxInt
	for _, val := range vars {
		if val < min {
			min = val
		}
	}
	return min, nil
}

func max(vars ...int) (int, error) {
	if len(vars) == 0 {
		return 0, fmt.Errorf("%s", "no values to compare")
	}
	max := math.MinInt
	for _, val := range vars {
		if val > max {
			max = val
		}
	}
	return max, nil
}

func main() {
	fmt.Println(min())        // error
	fmt.Println(max())        // error
	fmt.Println(min(1))       // 1
	fmt.Println(max(1))       // 1
	fmt.Println(min(1, 2, 3)) //1
	fmt.Println(max(1, 2, 3)) // 3
}

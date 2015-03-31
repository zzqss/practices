package main

import (
	"fmt"
)

// calculate hamming weight
// base on one fact x&x-1 make the left 1 move to right postion
func HammingWeight(input uint64) int {
	count := 0
	for ; input > 0; count++ {
		input &= input - 1
	}
	return count
}

func main() {
	fmt.Println(HammingWeight(uint64(11111)))
}

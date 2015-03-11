package bignumber

import (
	"fmt"
	"testing"
)

func TestPow(t *testing.T) {
	a := NewBigNumber(100)
	a.Pow(10)
	fmt.Println(a.number)
	if len(a.number) != 21 {
		panic("wrong pow")
	}
	for index, num := range a.number {
		if index == 20 {
			if num != byte(1) {
				panic("wrong pow")
			}
			continue
		}
		if num != byte(0) {
			panic("wrong pow")
		}
	}
}

func TestAdd(t *testing.T) {
	a := NewBigNumber(11)
	b := NewBigNumber(100)
	a.Add(b)
	fmt.Println(a.number)
	for _, num := range a.number {
		if num != byte(1) {
			panic("wrong add")
		}
	}
}

func TestSubtract(t *testing.T) {
	a := NewBigNumber(111)
	b := NewBigNumber(110)
	a.Subtract(b)
	fmt.Println(a.number)
	if len(a.number) != 1 || a.number[0] != byte(1) {
		panic("wrong subtract")
	}
}

func TestMultiply(t *testing.T) {
	a := NewBigNumber(100)
	b := NewBigNumber(100)
	a.Multiply(b)
	fmt.Println(a.number)
	if len(a.number) != 5 {
		panic("wrong multiply")
	}
	for index, num := range a.number {
		if index == 4 {
			if num != byte(1) {
				panic("wrong multiply")
			}
			continue
		}
		if num != byte(0) {
			panic("wrong multiply")
		}
	}
}

func TestDivide(t *testing.T) {
	a := NewBigNumber(110)
	b := NewBigNumber(10)
	a.Divide(b)
	fmt.Println(a.number)
	if len(a.number) != 2 {
		panic("wrong divide")
	}
	if a.number[0] != byte(1) {
		panic("wrong divide")
	}
	if a.number[1] != byte(1) {
		panic("wrong divide")
	}
}

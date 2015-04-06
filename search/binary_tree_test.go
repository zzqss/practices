package search

import (
	"testing"
)

func TestBinaryTree(t *testing.T) {
	slice := [...]int{1, 6, 8, 4, 9, 3, 5, 2, 7, 4, 5, 7}
	length := len(slice)
	topTree := NewBinaryTree(slice[0])
	for i := 1; i < length; i++ {
		topTree.Insert(slice[i])
	}
	topTree.Print()
}

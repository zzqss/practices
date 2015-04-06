package search

import (
	"fmt"
)

type BinaryTree struct {
	value int
	left  *BinaryTree
	right *BinaryTree
}

func NewBinaryTree(value int) *BinaryTree {
	return &BinaryTree{value, nil, nil}
}
func (this *BinaryTree) Value() int {
	return this.value
}

func (this *BinaryTree) Left() *BinaryTree {
	return this.left
}

func (this *BinaryTree) Right() *BinaryTree {
	return this.right
}
func (this *BinaryTree) Search(value int) *BinaryTree {
	if value == this.value {
		return this
	}
	var binaryTree *BinaryTree
	if value < this.value {
		binaryTree = this.left
	} else {
		binaryTree = this.right
	}
	if binaryTree == nil {
		return nil
	} else {
		return binaryTree.Search(value)
	}
}

// insert if not exist
func (this *BinaryTree) Insert(value int) *BinaryTree {
	if this.value == value {
		return this
	}
	var binaryTree **BinaryTree
	if value < this.value {
		binaryTree = &this.left
	} else {
		binaryTree = &this.right
	}
	if *binaryTree == nil {
		*binaryTree = NewBinaryTree(value)
		return *binaryTree
	} else {
		return (*binaryTree).Insert(value)
	}
}
func (this *BinaryTree) Print() {
	if this.left != nil {
		this.left.Print()
	}
	fmt.Println(this.value)
	if this.right != nil {
		this.right.Print()
	}
}

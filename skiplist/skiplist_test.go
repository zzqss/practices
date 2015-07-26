package main

import (
	"testing"
)

func TestSkipList(t *testing.T) {
	skipList := NewSkipList()
	skipList.Insert(NewIntKey(1), "a")
	skipList.Insert(NewIntKey(2), "b")
	skipList.Insert(NewIntKey(3), "c")
	skipList.Get(NewIntKey(1))
	skipList.Get(NewIntKey(2))
	skipList.Get(NewIntKey(3))
	skipList.Delete(NewIntKey(1))
	skipList.Delete(NewIntKey(2))
	skipList.Delete(NewIntKey(3))
}

package main

import (
	"testing"
)

func TestSkipList(t *testing.T) {
	skipList := NewSkipList()
	skipList.Insert(NewIntKey(1), "a")
	skipList.Get(NewIntKey(1))
	skipList.Delete(NewIntKey(1))
}

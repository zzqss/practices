package main

import (
	"log"
	"testing"
)

func TestSkipList(t *testing.T) {
	skipList := NewSkipList()
	skipList.Insert(NewIntKey(2), "a")
	log.Println(skipList.length)
	skipList.Insert(NewIntKey(1), "b")
	log.Println(skipList.length)
	skipList.Insert(NewIntKey(3), "c")
	log.Println(skipList.length)
	skipList.Get(NewIntKey(1))
	skipList.Get(NewIntKey(2))
	skipList.Get(NewIntKey(3))
	skipList.Delete(NewIntKey(1))
	log.Println(skipList.length)
	skipList.Delete(NewIntKey(3))
	log.Println(skipList.length)
	skipList.Delete(NewIntKey(2))
	log.Println(skipList.length)
}

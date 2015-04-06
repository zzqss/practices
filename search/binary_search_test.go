package search

import (
	"log"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	slice := []int{1, 3, 5, 7}
	log.Println(BinarySearch(slice, 0, len(slice)-1, 7))
}

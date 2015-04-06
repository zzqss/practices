package sort

import (
	"log"
	"testing"
)

func TestInsertSort(t *testing.T) {
	slice := []int{8, 2, 4, 5, 1, 9, 100, 99}
	OptimizeInsertSortAsc(slice)
	log.Println(slice)
}

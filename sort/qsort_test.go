package sort

import (
	"fmt"
	"testing"
)

func TestQSort(t *testing.T) {
	array := []int{1, 2, 4, 3, 2, 1}
	QSort(array, 0, len(array)-1)
	fmt.Println(array)
}

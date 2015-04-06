package search

// binary search
// slice must be in asc order
func BinarySearch(slice []int, start, end, find int) int {
	if start > end {
		return -1
	}
	middle := (start + end) / 2
	middleValue := slice[middle]
	if middleValue == find {
		return middle
	}
	if find > middleValue {
		return BinarySearch(slice, middle+1, end, find)
	} else {
		return BinarySearch(slice, start, middle-1, find)
	}
}

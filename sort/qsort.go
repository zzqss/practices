package sort

// quick sort
// find a middle
// split to little array
func QSort(array []int, start, end int) {
	if start > end {
		return
	}
	splitIndex := findSplitIndex(array, start, end)
	QSort(array, start, splitIndex-1)
	QSort(array, splitIndex+1, end)
}

// mv first element of array to middle of array,
// where left of it less than it
// right of it bigger than it
func findSplitIndex(array []int, start, end int) int {
	index := start
	value := array[index]
	for i := start + 1; i <= end; i++ {
		if array[i] < value {
			next := array[index+1]
			now := array[i]
			array[index] = now
			array[index+1] = value
			array[i] = next
			index += 1
		}
	}
	return index
}

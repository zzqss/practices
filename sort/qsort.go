package sort

// quick sort
// find a middle
// split to little array
func QSort(array []int, start, end int) {
	if start > end {
		return
	}
	splitIndex := optimizeFindSplitIndex1(array, start, end)
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

// how to optimize
// reduce one swap
// we know ,there is one place for array[start]
// so swap rest of array,to the right position
// record the right position
// swap number bigger than it to right
func optimizeFindSplitIndex1(array []int, start, end int) int {
	position := start
	for i := start + 1; i <= end; i++ {
		if array[i] < array[start] {
			position += 1
			temp := array[position]
			array[position] = array[i]
			array[i] = temp
		}
	}
	temp := array[start]
	array[start] = array[position]
	array[position] = temp
	return position
}

// more about find split index
// if most of array are equal,last optimize func is about O(n*n)
// so loop from both edge,
// if all the number are same ,we should find it middle
// first
func optimizeFindSplitIndex2(array []int, start, end int) {

}

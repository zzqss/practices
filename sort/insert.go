package sort

// what is basic insert sort
// the origin thought of insert sort
// assume the prev slice is  sorted
// loop from 0 to n
// from the begin to i-1
// if
func BasicInsertSortAsc(slice []int) {
	length := len(slice)
	for i := 1; i < length; i++ {
		for j := i; j > 0; j-- {
			if slice[j] < slice[j-1] {
				temp := slice[j]
				slice[j] = slice[j-1]
				slice[j-1] = temp
			}
		}
	}
}

// what do i optimize
// detail of swap
// the basic swap ,swap all the number before we find the right position
// just think about it
// we know it is there,just mv the right slice to right,then we put the number to the position
// it is fine
//[]int{8, 2, 4, 5, 1, 9, 100, 99}
func OptimizeInsertSortAsc(slice []int) {
	length := len(slice)
	for i := 1; i < length; i++ {
		now := slice[i]
		index := i
		for j := i - 1; j >= 0; j-- {
			if now < slice[j] {
				slice[j+1] = slice[j]
				index -= 1
			} else {
				break
			}
		}
		slice[index] = now
	}
}

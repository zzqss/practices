package main

import (
	"log"
)

/**
The Wagner–Fischer algorithm computes edit distance
based on the observation that if we reserve a matrix to hold the edit distances between all prefixes of the first string and all prefixes of the second,
then we can compute the values in the matrix by flood filling the matrix,
and thus find the distance between the two full strings as the last value computed
**/
func wagnerFischer(start, end string) int {
	startLength := len(start)
	endLength := len(end)
	if startLength == 0 {
		return endLength
	}
	if endLength == 0 {
		return startLength
	}
	//init the matrix
	table := make([][]int, startLength+1, startLength+1)
	for index, _ := range table {
		table[index] = make([]int, endLength+1, endLength+1)
	}
	// init the border of matrix
	for i := 0; i <= startLength; i++ {
		table[i][0] = i
	}
	for i := 0; i <= endLength; i++ {
		table[0][i] = i
	}
	// start to calculate the edit distance
	for i := 1; i <= startLength; i++ {
		for j := 1; j <= endLength; j++ {
			if start[i-1:i] == end[j-1:j] {
				table[i][j] = table[i-1][j-1]
				continue
			}
			normalCost := table[i-1][j-1] + 1
			minCost := normalCost
			deleteCost := table[i-1][j] + 1
			if deleteCost < minCost {
				minCost = deleteCost
			}
			insertCost := table[i][j-1] + 1
			if insertCost < minCost {
				minCost = insertCost
			}
			table[i][j] = minCost
		}
	}
	return table[startLength][endLength]
}

func main() {
	log.Println(wagnerFischer("asdadsa", "asdsadsa"))
}

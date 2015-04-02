package main

import (
	"log"
)

/**
what levenshtein would do
at a char judge what it should do
left char length
substitution is the best
but if delete works delete
if insert better for later one insert
if delete what result
if substitution what result
if insert what result
**/
// if insert better for latter one insert
func isInsert(start, end string) bool {
	if len(end) < 2 || len(start) < 1 {
		return false
	}
	if start[0:1] == end[1:2] {
		return true
	}
	return false
}

// when do we delete in two
// if delete this one will do ,delete
func IsDelete(start, end string) bool {
	if len(start) < 2 || len(end) < 1 {
		return false
	}
	if start[1:2] == end[0:1] {
		return true
	}
	return false
}

// calculate Levenshtein distince of tow string
// just think
func Levenshtein(start, end string) int {
	startLength := len(start)
	endLength := len(end)
	if startLength == 0 {
		return endLength
	}
	if endLength == 0 {
		return startLength
	}
	length := 0
	for i := 0; i < endLength; i++ {
		if i >= startLength {
			length += endLength - i
			start += end[i:endLength]
			break
		}
		if start[i:i+1] == end[i:i+1] {
			continue
		}
		length += 1
		if isInsert(start[i:startLength], end[i:endLength]) {
			start = start[0:i] + end[i:i+1] + start[i:startLength]
			startLength = len(start)
			continue
		}
		if IsDelete(start[i:startLength], end[i:endLength]) {
			start = start[0:i] + start[i+1:startLength]
			startLength = len(start)
			continue
		}
		start = start[0:i] + end[i:i+1] + start[i+1:startLength]
		startLength = len(start)
	}
	if startLength > endLength {
		length += startLength - endLength
	}
	return length
}

func main() {
	log.Println(Levenshtein("kitten", "sitting"))
}

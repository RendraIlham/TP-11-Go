package main

import "fmt"

func SeqSearch(T tabInt, N, X int) int {
	var idx int
	var k int
	
	idx = -1
	k = 0
	
	for idx == -1 && k < N {
		if T[k] == X {
			idx = k
		}
		k++
	}
	return idx
}
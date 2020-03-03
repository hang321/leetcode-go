package main

import (
	"fmt"
	"sort"
)

func main() {
	test1 := []int {-4,-1,0,3,10}
	test2 := []int {-7,-3,2,3,11}
	
	fmt.Println(sortedSquares(test1))
	fmt.Println(sortedSquares(test2))
}

func sortedSquares(A []int) []int {
	var sq []int
	for i := 0; i < len(A); i++ {
		sq = append(sq, A[i] * A[i])
	}
	
	sort.Ints(sq)
	
	return sq
}
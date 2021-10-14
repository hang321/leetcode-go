package main

import (
	"fmt"
)

func main() {

	test1 := []byte("hello")
	test2 := []byte("Hannah")

	reverseString344(test1)
	reverseString344(test2)
}

func reverseString344(bytes []byte) {
	n := len(bytes)
	if n == 0 {
		return
	}
	for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
		bytes[l], bytes[r] = bytes[r], bytes[l]
	}

	fmt.Println(string(bytes[:]))
}

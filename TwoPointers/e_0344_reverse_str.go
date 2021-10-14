package TwoPointers

import "fmt"

/*
Write a function that reverses a string. The input string is given as an array of characters s.
Example 1:
Input: ["h","e","l","l","o"]
Output: ["o","l","l","e","h"]
Example 2:
Input: ["H","a","n","n","a","h"]
Output: ["h","a","n","n","a","H"]
*/

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

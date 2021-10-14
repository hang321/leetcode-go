package main

import (
	"fmt"
)

func main() {

	numbers := []int{2, 7, 11, 15}
	target := 9

	fmt.Println(twoSum(numbers, target))
}

func twoSum(numbers []int, target int) []int {
	// using a map to store the complement value
	var aMap = make(map[int]int)

	for i := 0; i <= len(numbers); i++ {
		complement := target - numbers[i]
		if val, ok := aMap[complement]; ok {
			return []int{val + 1, i + 1}
		}
		aMap[numbers[i]] = i
	}
	return nil
}

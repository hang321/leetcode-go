package main

import (
	"fmt"
	"math"
)

func main() {

	test1 := []int{1, 1}
	test2 := []int{4, 3, 2, 1, 4}
	test3 := []int{1, 2, 1}

	fmt.Println(maxArea(test1))
	fmt.Println(maxArea(test2))
	fmt.Println(maxArea(test3))
}

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	max := 0

	for i < j {
		a, b := height[i], height[j]
		h := int(math.Min(float64(a), float64(b)))

		area := h * (j - i)
		if max < area {
			max = area
		}

		if a < b {
			i++
		} else {
			j--
		}
	}

	return max
}

package TwoPointers

import "math"

/*
Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai).
n vertical lines are drawn such that the two endpoints of the line i is at (i, ai) and (i, 0).
Find two lines, which, together with the x-axis forms a container, such that the container contains the most water.
Note: You may not slant the container
*/

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

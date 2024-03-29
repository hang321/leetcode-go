package TwoPointers

/*
Given an array of integers that is already sorted in ascending order, find two numbers such that they add up to a specific target number.

The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2.

Note:

Your returned answers (both index1 and index2) are not zero-based.
You may assume that each input would have exactly one solution and you may not use the same element twice.
Example:

Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.
*/

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

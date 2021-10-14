package TwoPointers

/**
Given a singly linked list, determine if it is a palindrome.

Example 1:

Input: 1->2
Output: false
Example 2:

Input: 1->2->2->1
Output: true
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	fast, slow := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = head.Next.Next
		slow = head.Next
	}
	if fast != nil {
		slow = slow.Next
	}
	reverse(fast)
	for slow != nil && fast.Val == slow.Val {
		fast = fast.Next
		slow = slow.Next
	}
	return slow == nil
}

/**
Use the fast and slow pointer to find the middle of the list.
Which means when fast pointer , reaches the end of the last,
slow pointer would reach the middle.
Then reverse the second half of the last and
compare each element from the first half of the list with the ones in second half.
*/

func reverse(head *ListNode) *ListNode {
	var prev *ListNode = nil
	for head != nil {
		next := head.Next
		head.Next = prev
		prev = head
		head = next
	}
	return prev
}

type ListNode struct {
	Val  int
	Next *ListNode
}

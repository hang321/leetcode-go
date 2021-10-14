package main

import (
	"fmt"
)

func main() {
	test1a := ListNode{Val: 2, Next: nil}
	test1 := ListNode{Val: 1, Next: &test1a}
	fmt.Println(isPalindrome(&test1))

	test2c := ListNode{Val: 1, Next: nil}
	test2b := ListNode{Val: 2, Next: &test2c}
	test2a := ListNode{Val: 2, Next: &test2b}
	test2 := ListNode{Val: 1, Next: &test2a}
	fmt.Println(isPalindrome(&test2))
}

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

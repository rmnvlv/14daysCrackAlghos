package day05

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/* Medium */

type ListNode struct {
	Val  int
	Next *ListNode
}

//Runtime: 0 ms, Memory Usage: 2.2 MB
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	fast, slow := head, head

	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	if fast == nil {
		return head.Next
	}

	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return head
}

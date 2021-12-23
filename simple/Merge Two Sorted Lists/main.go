package main

import (
	"fmt"
)

/*
You are given the heads of two sorted linked lists list1 and list2.

Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.

Return the head of the merged linked list.



Example 1:


Input: list1 = [1,2,4], list2 = [1,3,4]
Output: [1,1,2,3,4,4]
Example 2:

Input: list1 = [], list2 = []
Output: []
Example 3:

Input: list1 = [], list2 = [0]
Output: [0]


Constraints:

The number of nodes in both lists is in the range [0, 50].
-100 <= Node.val <= 100
Both list1 and list2 are sorted in non-decreasing order.
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l1, l2 = &ListNode{}, &ListNode{}
	l1 = &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 = &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	fmt.Println(mergeTwoLists(l1, l2))

	l1 = &ListNode{}
	l2 = &ListNode{}
	fmt.Println(mergeTwoLists(l1, l2))
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	res := &ListNode{}
	curr := res
	for list1 != nil || list2 != nil {
		if list1 == nil {
			curr.Next = list2
			break
		}
		if list2 == nil {
			curr.Next = list1
			break
		}

		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}
	return res.Next
}

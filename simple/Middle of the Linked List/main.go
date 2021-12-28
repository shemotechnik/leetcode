package main

import (
	"fmt"
	"math"
)

/*
Given the head of a singly linked list, return the middle node of the linked list.

If there are two middle nodes, return the second middle node.



Example 1:


Input: head = [1,2,3,4,5]
Output: [3,4,5]
Explanation: The middle node of the list is node 3.
Example 2:


Input: head = [1,2,3,4,5,6]
Output: [4,5,6]
Explanation: Since the list has two middle nodes with values 3 and 4, we return the second one.


Constraints:

The number of nodes in the list is in the range [1, 100].
1 <= Node.val <= 100
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l1 = &ListNode{}
	l1 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6}}}}}}
	fmt.Println(middleNode(l1))

	l1 = &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	fmt.Println(middleNode(l1))
}

func middleNode(head *ListNode) *ListNode {
	var sl = []int{}
	for head != nil {
		sl = append(sl, head.Val)
		head = head.Next
	}
	if len(sl) == 0 {
		return &ListNode{}
	}
	from := int(math.Round(float64(len(sl) / 2)))
	ret := SliceToListNode(head, sl[from:])
	return &ret
}

func SliceToListNode(head *ListNode, slice []int) ListNode {
	if len(slice) == 0 {
		return ListNode{}
	}

	if len(slice) > 1 {
		next := SliceToListNode(head, slice[1:])
		return ListNode{Val: slice[0], Next: &next}
	}
	return ListNode{Val: slice[0]}
}

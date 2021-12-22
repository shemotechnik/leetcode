package main

/*
You are given the head of a singly linked-list. The list can be represented as:

L0 → L1 → … → Ln - 1 → Ln
Reorder the list to be on the following form:

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
You may not modify the values in the list's nodes. Only nodes themselves may be changed.


Input: head = [1,2,3,4]
Output: [1,4,2,3]

Input: head = [1,2,3,4,5]
Output: [1,5,2,4,3]
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var lNode = ListNode{}
	lNode.Val = 1
	lNode.Next = &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}
	reorderList(&lNode)
}

func reorderList(head *ListNode) {
	var slice = []int{}
	var newSlice = []int{}
	ListNodeToSlice(head, &slice)
	l, r := 0, len(slice)-1
	for l <= r {
		newSlice = append(newSlice, slice[l])
		if l != r {
			newSlice = append(newSlice, slice[r])
		}
		l++
		r--
	}
	*head = SliceToListNode(head, newSlice)
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

func ListNodeToSlice(head *ListNode, slice *[]int) {
	*slice = append(*slice, head.Val)
	if head.Next == nil {
		return
	}
	ListNodeToSlice(head.Next, slice)
}

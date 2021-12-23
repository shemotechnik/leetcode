package main

import "fmt"

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order, and each of their nodes contains a single digit. Add the two numbers and return the sum as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.



Example 1:


Input: l1 = [2,4,3], l2 = [5,6,4]
Output: [7,0,8]
Explanation: 342 + 465 = 807.
Example 2:

Input: l1 = [0], l2 = [0]
Output: [0]
Example 3:

Input: l1 = [9,9,9,9,9,9,9], l2 = [9,9,9,9]
Output: [8,9,9,9,0,0,0,1]


Constraints:

The number of nodes in each linked list is in the range [1, 100].
0 <= Node.val <= 9
It is guaranteed that the list represents a number that does not have leading zeros.
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var l1, l2 = &ListNode{}, &ListNode{}
	l1 = &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 = &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	fmt.Println(addTwoNumbers(l1, l2))

	l1 = &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}}}}
	l2 = &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9}}}}
	fmt.Println(addTwoNumbers(l1, l2))

	l1 = &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 9}}}
	l2 = &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: &ListNode{Val: 9}}}}
	fmt.Println(addTwoNumbers(l1, l2))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var sl1 = []int{}
	var sl2 = []int{}

	ListNodeToSlice(l1, &sl1)
	ListNodeToSlice(l2, &sl2)

	slR := addSlices(sl1, sl2)
	lNode := SliceToListNode(&ListNode{}, slR)

	return &lNode
}

func addSlices(sl1, sl2 []int) []int {
	var res = []int{}
	additionalAdd := 0
	var summ = 0
	i := 0
	j := 0
	var sl1Var, sl2Var = 0, 0
	for i <= len(sl1)-1 || j <= len(sl2)-1 {
		if i > len(sl1)-1 {
			sl1Var = 0
		} else {
			sl1Var = sl1[i]
		}
		if j > len(sl2)-1 {
			sl2Var = 0
		} else {
			sl2Var = sl2[j]
		}

		summ = sl1Var + sl2Var + additionalAdd
		if additionalAdd == 1 {
			additionalAdd = 0
		}
		if summ >= 10 {
			res = append(res, summ-10)
			additionalAdd = 1
		} else {
			res = append(res, summ)
		}

		i++
		j++
	}
	if additionalAdd == 1 {
		res = append(res, 1)
	}

	return res
}

func ListNodeToSlice(head *ListNode, slice *[]int) {
	*slice = append(*slice, head.Val)
	if head.Next == nil {
		return
	}
	ListNodeToSlice(head.Next, slice)
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

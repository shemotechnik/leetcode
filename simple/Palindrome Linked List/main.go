package main

import "fmt"

/*
Given the head of a singly linked list, return true if it is a palindrome.



Example 1:


Input: head = [1,2,2,1]
Output: true
Example 2:


Input: head = [1,2]
Output: false


Constraints:

The number of nodes in the list is in the range [1, 105].
0 <= Node.val <= 9
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val int
	Next *ListNode
}

func main(){
	var lNode = ListNode{}
	lNode.Val=1
	lNode.Next=&ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}

	fmt.Println(isPalindrome(&lNode))
}

func isPalindrome(head *ListNode) bool {
	var slice = []int{}
	ListNodeToSlice(head, &slice)
	for i := 0; i < len(slice); i++ {
		j := len(slice)-1-i
		if slice[i] != slice[j] {
			return false
		}
	}
	return true
}

func ListNodeToSlice(head *ListNode, slice *[]int){
	*slice = append(*slice, head.Val)
	if head.Next == nil{
		return
	}
	ListNodeToSlice(head.Next, slice)
}


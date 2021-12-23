package main

import "fmt"

/*
Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.


Example 1:

Input: s = "()"
Output: true
Example 2:

Input: s = "()[]{}"
Output: true
Example 3:

Input: s = "(]"
Output: false


Constraints:

1 <= s.length <= 104
s consists of parentheses only '()[]{}'.
*/

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
	fmt.Println(isValid("({[]})"))
	fmt.Println(isValid("{{}[][[[]]]}"))
	fmt.Println(isValid("{{}[][[[]]]}{"))
}

func isValid(s string) bool {
	var opened = make(map[int]string)
	var position = 0
	for i := range s {
		if string(s[i]) == "(" || string(s[i]) == "{" || string(s[i]) == "[" {
			opened[position] = string(s[i])
			position++
		} else {
			switch string(s[i]) {
			case ")":
				if opened[position-1] != "(" {
					return false
				}
			case "]":
				if opened[position-1] != "[" {
					return false
				}
			case "}":
				if opened[position-1] != "{" {
					return false
				}
			}
			delete(opened, position-1)
			position--
		}
	}
	if len(opened) == 0 {
		return true
	}
	return false
}

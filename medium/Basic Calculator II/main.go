package main

import (
	"fmt"
	"unicode"
)

/*
Given a string s which represents an expression, evaluate this expression and return its value.

The integer division should truncate toward zero.

You may assume that the given expression is always valid. All intermediate results will be in the range of [-231, 231 - 1].

Note: You are not allowed to use any built-in function which evaluates strings as mathematical expressions, such as eval().



Example 1:

Input: s = "3+2*2"
Output: 7
Example 2:

Input: s = " 3/2 "
Output: 1
Example 3:

Input: s = " 3+5 / 2 "
Output: 5


Constraints:

1 <= s.length <= 3 * 105
s consists of integers and operators ('+', '-', '*', '/') separated by some number of spaces.
s represents a valid expression.
All the integers in the expression are non-negative integers in the range [0, 231 - 1].
The answer is guaranteed to fit in a 32-bit integer.
*/

func main() {
	fmt.Println(calculate("3+2*2"))
	fmt.Println(calculate("33/2+5/3"))
	fmt.Println(calculate("3+5 / 2"))
	fmt.Println(calculate(" 3/2 "))

	fmt.Println(calculate("1+1+1"))
	fmt.Println(calculate("0-2147483647"))
	fmt.Println(calculate("1-1+1"))
	fmt.Println(calculate("2-3+4"))
	fmt.Println(calculate("-222/1-11*2-4/1-1"))

	fmt.Println(calculate("6/2/3/10"))
	fmt.Println(calculate("100000000/1/2/3/4/5/6/7/8/9/10"))
	fmt.Println(calculate("1*2*3*4*5*6*7*8*9*10"))
	fmt.Println(calculate("1+2*5/3+6/4*2"))
}

func calculate(s string) int {
	var stack []int
	var curNum int
	var prevOp rune
	prevOp = '+'

	for i, v := range s {
		if unicode.IsDigit(v) {
			curNum = curNum*10 + int(v-'0')
		}
		if v == '+' || v == '-' || v == '*' || v == '/' || i == len(s)-1 {
			switch prevOp {
			case '+':
				stack = append(stack, curNum)
			case '-':
				stack = append(stack, -curNum)
			case '*':
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, top*curNum)

			case '/':
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, top/curNum)
			}
			curNum = 0
			prevOp = v
		}

	}

	var sum int
	for _, v := range stack {
		sum += v
	}

	return sum
}

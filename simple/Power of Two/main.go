package main

import (
	"fmt"
)

/*
Contest:
Given an integer n, return true if it is a power of two. Otherwise, return false.

An integer n is a power of two, if there exists an integer x such that n == 2x.


Example 1:


Input: n = 1
Output: true
Explanation: 20 = 1
Example 2:

Input: n = 16
Output: true
Explanation: 24 = 16
Example 3:

Input: n = 3
Output: false


Constraints:

-231 <= n <= 231 - 1
*/

func main() {
	fmt.Println(isPowerOfTwo(16))
	fmt.Println(isPowerOfTwo(1))
	fmt.Println(isPowerOfTwo(3))
	fmt.Println(isPowerOfTwo(6))
	fmt.Println(isPowerOfTwo(-16))
}

func isPowerOfTwo(n int) bool {
	if n == 0 {
		return false
	}
	for {
		if n == 1 {
			return true
		} else if n%2 != 0 {
			return false
		} else if n > -1 && n < 1 {
			return false
		} else {
			n = n / 2
		}
	}
}

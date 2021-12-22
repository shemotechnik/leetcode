package main

import (
	"bytes"
	"fmt"
)

/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".



Example 1:

Input: strs = ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: strs = ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.


Constraints:

1 <= strs.length <= 200
0 <= strs[i].length <= 200
strs[i] consists of only lower-case English letters.
*/

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println(longestCommonPrefix([]string{"ab", "a"}))
	fmt.Println(longestCommonPrefix([]string{"cir", "car"}))
}

func longestCommonPrefix(strs []string) string {
	var bf = bytes.Buffer{}
	if len(strs[0]) == 0 {
		return ""
	}
	word := strs[0]
	i := 0
	for i < len(word) {
		char := string(word[i])
		flag := true
		for k := range strs {
			if len(strs[k]) > i {
				if string(strs[k][i]) != char {
					flag = false
					break
				}
			} else {
				flag = false
				break
			}
		}
		if flag {
			bf.WriteString(char)
		} else {
			break
		}
		i++
	}
	return bf.String()
}

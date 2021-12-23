package main

import (
	"fmt"
)

/*
You are given a string s and an array of strings words of the same length. Return all starting indices of substring(s) in s that is a concatenation of each word in words exactly once, in any order, and without any intervening characters.

You can return the answer in any order.



Example 1:

Input: s = "barfoothefoobarman", words = ["foo","bar"]
Output: [0,9]
Explanation: Substrings starting at index 0 and 9 are "barfoo" and "foobar" respectively.
The output order does not matter, returning [9,0] is fine too.
Example 2:

Input: s = "wordgoodgoodgoodbestword", words = ["word","good","best","word"]
Output: []
Example 3:

Input: s = "barfoofoobarthefoobarman", words = ["bar","foo","the"]
Output: [6,9,12]

Constraints:

1 <= s.length <= 104
s consists of lower-case English letters.
1 <= words.length <= 5000
1 <= words[i].length <= 30
words[i] consists of lower-case English letters.
*/

func main() {
	fmt.Println(findSubstring("qwefoobarfoothefoobarfooman", []string{"foo", "bar", "foo"}))
	fmt.Println(findSubstring("barfoofoobarthefoobarman", []string{"bar", "foo", "the"}))
	fmt.Println(findSubstring("ababaab", []string{"ab", "ba", "ba"}))
}

func findSubstring(s string, words []string) []int {
	var ret = []int{}
	if len(s) == 0 || len(words) == 0 {
		return []int{}
	}

	for i := range s {
		flag := false
		for j := range words {
			if words[j][0] == s[i] {
				flag = true
				break
			}
		}
		if flag {
			if i+(len(words)*len(words[0])) <= len(s) {
				substr := s[i : i+(len(words)*len(words[0]))]
				small := len(words[0])
				m := make(map[string]int)
				for j := range words {
					subsubstr := substr[j*small : (j+1)*small]
					m[subsubstr]++
				}
				for j := range words {
					m[words[j]]--
				}
				ifAllNull := true
				for j := range m {
					if m[j] != 0 {
						ifAllNull = false
						break
					}
				}
				if ifAllNull {
					ret = append(ret, i)
				}
			}
		}
	}
	return ret
}
